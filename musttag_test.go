package musttag

import (
	"go/ast"
	"go/token"
	"path/filepath"
	"strings"
	"testing"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/analysistest"
	"golang.org/x/tools/go/types/typeutil"
)

// TODO(junk1tm): do not depend on tests execution order
func TestAll(t *testing.T) {
	t.Run("examples", func(t *testing.T) {
		dir := filepath.Join(analysistest.TestData(), "examples")
		analysistest.Run(t, dir, Analyzer)
	})

	t.Run("tests", func(t *testing.T) {
		// for the tests we want to record reports from all functions.
		reportOnce = false
		reportf = func(pass *analysis.Pass, call *ast.CallExpr, pos token.Pos, tag string) {
			fn := typeutil.StaticCallee(pass.TypesInfo, call)
			pass.Reportf(pos, shortName(fn.FullName()))
		}
		dir := filepath.Join(analysistest.TestData(), "tests")
		analysistest.Run(t, dir, Analyzer)
	})
}

func shortName(name string) string {
	name = strings.ReplaceAll(name, "*", "")
	name = strings.ReplaceAll(name, "(", "")
	name = strings.ReplaceAll(name, ")", "")
	name = strings.TrimPrefix(name, "encoding/")
	return name
}
