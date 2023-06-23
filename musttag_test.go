package musttag

import (
	"go/token"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"go-simpler.org/assert"
	. "go-simpler.org/assert/dotimport"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()

	t.Run("tests", func(t *testing.T) {
		r := report
		defer func() { report = r }()

		report = func(pass *analysis.Pass, st *structType, fn Func, fnPos token.Position) {
			pass.Reportf(st.Pos, fn.shortName())
		}

		setupTestData(t, testdata, "tests")

		analyzer := New()
		analysistest.Run(t, testdata, analyzer, "tests")
	})

	t.Run("builtins", func(t *testing.T) {
		setupTestData(t, testdata, "builtins")

		analyzer := New(
			Func{Name: "example.com/custom.Marshal", Tag: "custom", ArgPos: 0},
			Func{Name: "example.com/custom.Unmarshal", Tag: "custom", ArgPos: 1},
		)
		analysistest.Run(t, testdata, analyzer, "builtins")
	})

	t.Run("bad Func.ArgPos", func(t *testing.T) {
		setupTestData(t, testdata, "tests")

		analyzer := New(
			Func{Name: "encoding/json.Marshal", Tag: "json", ArgPos: 10},
		)
		err := analysistest.Run(nopT{}, testdata, analyzer, "tests")[0].Err
		assert.Equal[E](t, err.Error(), `Func.ArgPos cannot be 10: encoding/json.Marshal accepts only 1 argument(s)`)
	})
}

func TestFlags(t *testing.T) {
	analyzer := New()
	analyzer.Flags.Usage = func() {}
	analyzer.Flags.SetOutput(io.Discard)

	t.Run("ok", func(t *testing.T) {
		err := analyzer.Flags.Parse([]string{"-fn=test.Test:test:0"})
		assert.NoErr[E](t, err)
	})

	t.Run("invalid format", func(t *testing.T) {
		err := analyzer.Flags.Parse([]string{"-fn=test.Test"})
		assert.Equal[E](t, err.Error(), `invalid value "test.Test" for flag -fn: invalid syntax`)
	})

	t.Run("non-number argpos", func(t *testing.T) {
		err := analyzer.Flags.Parse([]string{"-fn=test.Test:test:-"})
		assert.Equal[E](t, err.Error(), `invalid value "test.Test:test:-" for flag -fn: strconv.Atoi: parsing "-": invalid syntax`)
	})
}

type nopT struct{}

func (nopT) Errorf(string, ...any) {}

// NOTE: analysistest does not yet support modules;
// see https://github.com/golang/go/issues/37054 for details.
func setupTestData(t *testing.T, testDataDir, dir string) {
	t.Helper()

	err := os.Chdir(filepath.Join(testDataDir, "src", dir))
	if err != nil {
		t.Fatal(err)
	}

	output, err := exec.Command("go", "mod", "vendor").CombinedOutput()
	if err != nil {
		t.Log(string(output))
		t.Fatal(err)
	}
}
