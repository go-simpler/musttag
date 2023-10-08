// Package musttag implements the musttag analyzer.
package musttag

import (
	"flag"
	"fmt"
	"go/ast"
	"go/token"
	"go/types"
	"path"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
	"golang.org/x/tools/go/types/typeutil"
)

// Func describes a function call to look for, e.g. json.Marshal.
type Func struct {
	Name   string // Name is the full name of the function, including the package.
	Tag    string // Tag is the struct tag whose presence should be ensured.
	ArgPos int    // ArgPos is the position of the argument to check.

	// a list of interface names (including the package);
	// if at least one is implemented by the argument, no check is performed.
	ifaceWhitelist []string
}

func (fn Func) shortName() string {
	name := strings.NewReplacer("*", "", "(", "", ")", "").Replace(fn.Name)
	return path.Base(name)
}

// New creates a new musttag analyzer.
// To report a custom function provide its description via Func,
// it will be added to the builtin ones.
func New(funcs ...Func) *analysis.Analyzer {
	var flagFuncs []Func
	return &analysis.Analyzer{
		Name:     "musttag",
		Doc:      "enforce field tags in (un)marshaled structs",
		Flags:    flags(&flagFuncs),
		Requires: []*analysis.Analyzer{inspect.Analyzer},
		Run: func(pass *analysis.Pass) (any, error) {
			l := len(builtins) + len(funcs) + len(flagFuncs)
			f := make(map[string]Func, l)

			toMap := func(slice []Func) {
				for _, fn := range slice {
					f[fn.Name] = fn
				}
			}
			toMap(builtins)
			toMap(funcs)
			toMap(flagFuncs)

			mainModule, err := getMainModule()
			if err != nil {
				return nil, err
			}

			return run(pass, mainModule, f)
		},
	}
}

// flags creates a flag set for the analyzer.
// The funcs slice will be filled with custom functions passed via CLI flags.
func flags(funcs *[]Func) flag.FlagSet {
	fs := flag.NewFlagSet("musttag", flag.ContinueOnError)
	fs.Func("fn", "report custom function (name:tag:argpos)", func(s string) error {
		parts := strings.Split(s, ":")
		if len(parts) != 3 || parts[0] == "" || parts[1] == "" {
			return strconv.ErrSyntax
		}
		pos, err := strconv.Atoi(parts[2])
		if err != nil {
			return err
		}
		*funcs = append(*funcs, Func{
			Name:   parts[0],
			Tag:    parts[1],
			ArgPos: pos,
		})
		return nil
	})
	return *fs
}

// for tests only.
var report = func(pass *analysis.Pass, st *structType, fn Func, fnPos token.Position) {
	const format = "`%s` should be annotated with the `%s` tag as it is passed to `%s` at %s"
	pass.Reportf(st.Pos, format, st.Name, fn.Tag, fn.shortName(), fnPos)
}

var trimVendor = regexp.MustCompile(`([^*/(]+/vendor/)`)

// run starts the analysis.
func run(pass *analysis.Pass, mainModule string, funcs map[string]Func) (any, error) {
	var err error

	walk := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	filter := []ast.Node{(*ast.CallExpr)(nil)}

	walk.Preorder(filter, func(n ast.Node) {
		if err != nil {
			return // there is already an error.
		}

		call, ok := n.(*ast.CallExpr)
		if !ok {
			return // not a function call.
		}

		callee := typeutil.StaticCallee(pass.TypesInfo, call)
		if callee == nil {
			return // not a static call.
		}

		name := trimVendor.ReplaceAllString(callee.FullName(), "")
		fn, ok := funcs[name]
		if !ok {
			return // the function is not supported.
		}

		if len(call.Args) <= fn.ArgPos {
			err = fmt.Errorf("Func.ArgPos cannot be %d: %s accepts only %d argument(s)", fn.ArgPos, fn.Name, len(call.Args))
			return
		}

		arg := call.Args[fn.ArgPos]
		if unary, ok := arg.(*ast.UnaryExpr); ok {
			arg = unary.X // e.g. json.Marshal(&foo)
		}

		initialPos := token.NoPos
		switch arg := arg.(type) {
		case *ast.Ident: // e.g. json.Marshal(foo)
			if arg.Obj == nil {
				return // e.g. json.Marshal(nil)
			}
			initialPos = arg.Obj.Pos()
		case *ast.CompositeLit: // e.g. json.Marshal(struct{}{})
			initialPos = arg.Pos()
		}

		typ := pass.TypesInfo.TypeOf(arg)
		if typ == nil {
			return // no type info found.
		}

		if implementsInterface(typ, fn.ifaceWhitelist, pass.Pkg.Imports()) {
			return // the type implements a Marshaler interface, nothing to check; see issue #64.
		}

		checker := checker{
			mainModule: mainModule,
			seenTypes:  make(map[string]struct{}),
		}

		st, ok := checker.parseStructType(typ, initialPos)
		if !ok {
			return // not a struct argument.
		}

		result, ok := checker.checkStructType(st, fn.Tag)
		if ok {
			return // nothing to report.
		}

		p := pass.Fset.Position(call.Pos())
		report(pass, result, fn, p)
	})

	return nil, err
}

// structType is an extension for types.Struct.
// The content of the fields depends on whether the type is named or not.
type structType struct {
	*types.Struct
	Name string    // for types.Named: the type's name; for anonymous: a placeholder string.
	Pos  token.Pos // for types.Named: the type's position; for anonymous: the corresponding identifier's position.
}

// checker parses and checks struct types.
type checker struct {
	mainModule string
	seenTypes  map[string]struct{} // prevent panic on recursive types; see issue #16.
}

// parseStructType parses the given types.Type, returning the underlying struct type.
func (c *checker) parseStructType(t types.Type, pos token.Pos) (*structType, bool) {
	for {
		// unwrap pointers (if any) first.
		ptr, ok := t.(*types.Pointer)
		if !ok {
			break
		}
		t = ptr.Elem()
	}

	switch t := t.(type) {
	case *types.Named: // a struct of the named type.
		pkg := t.Obj().Pkg() // may be nil; see issue #38.
		if pkg == nil {
			return nil, false
		}

		if !strings.HasPrefix(pkg.Path(), c.mainModule) {
			return nil, false
		}

		s, ok := t.Underlying().(*types.Struct)
		if !ok {
			return nil, false
		}

		return &structType{
			Struct: s,
			Pos:    t.Obj().Pos(),
			Name:   t.Obj().Name(),
		}, true

	case *types.Struct: // an anonymous struct.
		return &structType{
			Struct: t,
			Pos:    pos,
			Name:   "anonymous struct",
		}, true
	}

	return nil, false
}

// checkStructType recursively checks whether the given struct type is annotated with the tag.
// The result is the type of the first nested struct which fields are not properly annotated.
func (c *checker) checkStructType(st *structType, tag string) (*structType, bool) {
	c.seenTypes[st.String()] = struct{}{}

	for i := 0; i < st.NumFields(); i++ {
		field := st.Field(i)
		if !field.Exported() {
			continue
		}

		if _, ok := reflect.StructTag(st.Tag(i)).Lookup(tag); !ok {
			// tag is not required for embedded types; see issue #12.
			if !field.Embedded() {
				return st, false
			}
		}

		nested, ok := c.parseStructType(field.Type(), st.Pos) // TODO: or field.Pos()?
		if !ok {
			continue
		}
		if _, ok := c.seenTypes[nested.String()]; ok {
			continue
		}
		if result, ok := c.checkStructType(nested, tag); !ok {
			return result, false
		}
	}

	return nil, true
}

func implementsInterface(typ types.Type, ifaces []string, imports []*types.Package) bool {
	findScope := func(pkgName string) (*types.Scope, bool) {
		// fast path: check direct imports (e.g. looking for "encoding/json.Marshaler").
		for _, direct := range imports {
			if pkgName == trimVendor.ReplaceAllString(direct.Path(), "") {
				return direct.Scope(), true
			}
		}
		// slow path: check indirect imports (e.g. looking for "encoding.TextMarshaler").
		for _, direct := range imports {
			for _, indirect := range direct.Imports() {
				if pkgName == trimVendor.ReplaceAllString(indirect.Path(), "") {
					return indirect.Scope(), true
				}
			}
		}
		return nil, false
	}

	for _, ifacePath := range ifaces {
		// "encoding/json.Marshaler" -> "encoding/json" + "Marshaler"
		idx := strings.LastIndex(ifacePath, ".")
		if idx == -1 {
			continue
		}
		pkgName, ifaceName := ifacePath[:idx], ifacePath[idx+1:]

		scope, ok := findScope(pkgName)
		if !ok {
			continue
		}
		obj := scope.Lookup(ifaceName)
		if obj == nil {
			continue
		}
		iface, ok := obj.Type().Underlying().(*types.Interface)
		if !ok {
			continue
		}
		if types.Implements(typ, iface) {
			return true
		}
	}

	return false
}
