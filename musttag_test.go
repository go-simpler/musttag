package musttag

import (
	"go/ast"
	"go/token"
	"strings"
	"testing"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/analysistest"
	"golang.org/x/tools/go/types/typeutil"
)

func Test(t *testing.T) {
	// for the tests we want to record reports from all functions.
	reportOnce = false
	reportf = func(pass *analysis.Pass, call *ast.CallExpr, pos token.Pos, tag string) {
		fn := typeutil.StaticCallee(pass.TypesInfo, call)
		name := fn.FullName()
		pass.Reportf(pos, shortName(name))
	}
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, Analyzer)
}

func shortName(name string) string {
	name = strings.ReplaceAll(name, "*", "")
	name = strings.ReplaceAll(name, "(", "")
	name = strings.ReplaceAll(name, ")", "")
	name = strings.TrimPrefix(name, "encoding/")
	return name
}
