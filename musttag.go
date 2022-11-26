package musttag

import (
	"go/ast"
	"go/token"
	"go/types"
	"strings"

	"golang.org/x/tools/go/analysis"
	inspectpass "golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
	"golang.org/x/tools/go/types/typeutil"

	// we need these dependencies for the tests in testdata to compile.
	// `go mod tidy` will remove them from go.mod if we don't import them here.
	_ "github.com/BurntSushi/toml"
	_ "gopkg.in/yaml.v3"
)

var Analyzer = &analysis.Analyzer{
	Name:     "musttag",
	Doc:      "check if struct fields used in Marshal/Unmarshal are annotated with the relevant tag",
	Requires: []*analysis.Analyzer{inspectpass.Analyzer},
	Run:      run,
}

// for tests only.
var (
	// should the same struct be reported only once for the same tag?
	reportOnce = true

	// reportf is a wrapper for pass.Reportf (as a variable, so it could be mocked in tests).
	reportf = func(pass *analysis.Pass, call *ast.CallExpr, pos token.Pos, tag string) {
		pass.Reportf(pos, "exported fields should be annotated with the %q tag", tag)
	}
)

// run starts the analysis.
func run(pass *analysis.Pass) (any, error) {
	inspect := pass.ResultOf[inspectpass.Analyzer].(*inspector.Inspector)

	filter := []ast.Node{
		(*ast.CallExpr)(nil),
	}

	type report struct {
		pos token.Pos
		tag string
	}
	reported := make(map[report]struct{})

	inspect.Preorder(filter, func(n ast.Node) {
		call := n.(*ast.CallExpr)

		tag, expr, ok := tagAndExpr(pass, call)
		if !ok {
			return
		}

		s, pos, ok := structAndPos(pass, expr)
		if !ok {
			return
		}

		if ok := checkStruct(s, tag, &pos); ok {
			return
		}

		r := report{pos, tag}
		if _, ok := reported[r]; ok && reportOnce {
			return
		}

		reportf(pass, call, pos, tag)
		reported[r] = struct{}{}
	})

	return nil, nil
}

// tagAndExpr analyses the given function call and returns the struct tag to
// look for and the expression that likely contains the struct to check.
func tagAndExpr(pass *analysis.Pass, call *ast.CallExpr) (string, ast.Expr, bool) {
	const (
		jsonTag = "json"
		xmlTag  = "xml"
		yamlTag = "yaml"
		tomlTag = "toml"
	)

	fn := typeutil.StaticCallee(pass.TypesInfo, call)
	if fn == nil {
		return "", nil, false
	}

	switch fn.FullName() {
	case "encoding/json.Marshal",
		"encoding/json.MarshalIndent",
		"(*encoding/json.Encoder).Encode",
		"(*encoding/json.Decoder).Decode":
		return jsonTag, call.Args[0], true
	case "encoding/json.Unmarshal":
		return jsonTag, call.Args[1], true

	case "encoding/xml.Marshal",
		"encoding/xml.MarshalIndent",
		"(*encoding/xml.Encoder).Encode",
		"(*encoding/xml.Decoder).Decode",
		"(*encoding/xml.Encoder).EncodeElement",
		"(*encoding/xml.Decoder).DecodeElement":
		return xmlTag, call.Args[0], true
	case "encoding/xml.Unmarshal":
		return xmlTag, call.Args[1], true

	case "gopkg.in/yaml.v3.Marshal",
		"(*gopkg.in/yaml.v3.Encoder).Encode",
		"(*gopkg.in/yaml.v3.Decoder).Decode":
		return yamlTag, call.Args[0], true
	case "gopkg.in/yaml.v3.Unmarshal":
		return yamlTag, call.Args[1], true

	case "(*github.com/BurntSushi/toml.Encoder).Encode",
		"(*github.com/BurntSushi/toml.Decoder).Decode":
		return tomlTag, call.Args[0], true
	case "github.com/BurntSushi/toml.Unmarshal",
		"github.com/BurntSushi/toml.Decode",
		"github.com/BurntSushi/toml.DecodeFile":
		return tomlTag, call.Args[1], true
	case "github.com/BurntSushi/toml.DecodeFS":
		return tomlTag, call.Args[2], true

	default:
		return "", nil, false
	}
}

// structAndPos analyses the given expression and returns the struct to check
// and the position to report if needed.
func structAndPos(pass *analysis.Pass, expr ast.Expr) (*types.Struct, token.Pos, bool) {
	t := pass.TypesInfo.TypeOf(expr)
	if ptr, ok := t.(*types.Pointer); ok {
		t = ptr.Elem()
	}

	switch t := t.(type) {
	case *types.Named: // named type
		s, ok := t.Underlying().(*types.Struct)
		if ok {
			return s, t.Obj().Pos(), true
		}

	case *types.Struct: // anonymous struct
		if unary, ok := expr.(*ast.UnaryExpr); ok {
			expr = unary.X // &x
		}
		//nolint:gocritic // commentedOutCode: these are examples
		switch arg := expr.(type) {
		case *ast.Ident: // var x struct{}; json.Marshal(x)
			return t, arg.Obj.Pos(), true
		case *ast.CompositeLit: // json.Marshal(struct{}{})
			return t, arg.Pos(), true
		}
	}

	return nil, 0, false
}

// checkStruct checks that exported fields of the given struct are annotated
// with the tag and updates the position to report in case a nested struct of a
// named type is found.
func checkStruct(s *types.Struct, tag string, pos *token.Pos) (ok bool) {
	for i := 0; i < s.NumFields(); i++ {
		if !s.Field(i).Exported() {
			continue
		}

		tagged := false
		for _, t := range strings.Split(s.Tag(i), " ") {
			// from the [reflect.StructTag] docs:
			// By convention, tag strings are a concatenation
			// of optionally space-separated key:"value" pairs.
			if strings.HasPrefix(t, tag+":") {
				tagged = true
			}
		}
		if !tagged {
			return false
		}

		// check if the field is a nested struct.
		t := s.Field(i).Type()
		if ptr, ok := t.(*types.Pointer); ok {
			t = ptr.Elem()
		}
		nested, ok := t.Underlying().(*types.Struct)
		if !ok {
			continue
		}
		if ok := checkStruct(nested, tag, pos); ok {
			continue
		}
		// update the position to point to the named type.
		if named, ok := t.(*types.Named); ok {
			*pos = named.Obj().Pos()
		}
		return false
	}

	return true
}
