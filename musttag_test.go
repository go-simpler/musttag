package musttag

import (
	"go/token"
	"path"
	"strings"
	"testing"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAll(t *testing.T) {
	// TODO(junk1tm): do not depend on tests execution order

	// NOTE(junk1tm): analysistest isn't aware of the main package's modules
	// (see https://github.com/golang/go/issues/37054), so to run tests with
	// external dependencies we have to be creative. Using vendor with symlinks
	// would work but the paths would contain the `vendor/` prefix, and that's
	// not what we want because we match full paths. The solution is to write
	// stubs of the dependencies (we don't need the actual code, only the
	// functions signatures to match) and to put them exactly at
	// testdata/src/path/to/pkg (GOPATH?), otherwise it won't work.

	analyzer := New(
		Func{Name: "example.com/custom.Marshal", Tag: "custom", ArgPos: 0},
		Func{Name: "example.com/custom.Unmarshal", Tag: "custom", ArgPos: 1},
	)

	t.Run("examples", func(t *testing.T) {
		testdata := analysistest.TestData()
		analysistest.Run(t, testdata, analyzer, "examples")
	})

	t.Run("tests", func(t *testing.T) {
		// for the tests we want to record reports from all functions.
		reportOnce = false
		reportf = func(pass *analysis.Pass, pos token.Pos, fn Func) {
			pass.Reportf(pos, shortName(fn.Name))
		}
		testdata := analysistest.TestData()
		analysistest.Run(t, testdata, analyzer, "tests")
	})
}

func shortName(name string) string {
	name = strings.ReplaceAll(name, "*", "")
	name = strings.ReplaceAll(name, "(", "")
	name = strings.ReplaceAll(name, ")", "")
	return path.Base(name)
}
