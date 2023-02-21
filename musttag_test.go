package musttag

import (
	"go/token"
	"io"
	"os"
	"path/filepath"
	"testing"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	// NOTE(junk1tm): analysistest does not yet support modules
	// (see https://github.com/golang/go/issues/37054 for details).
	// So, to be able to run tests with external dependencies,
	// we first need to write a GOPATH-like tree of stubs.
	prepareTestFiles(t)
	testPackages = []string{"tests", "builtins"}

	analyzer := New(
		Func{Name: "example.com/custom.Marshal", Tag: "custom", ArgPos: 0},
		Func{Name: "example.com/custom.Unmarshal", Tag: "custom", ArgPos: 1},
	)

	t.Run("builtins", func(t *testing.T) {
		testdata := analysistest.TestData()
		analysistest.Run(t, testdata, analyzer, "builtins")
	})

	t.Run("tests", func(t *testing.T) {
		reportf = func(pass *analysis.Pass, st *structType, fn Func, fnPos token.Position) {
			pass.Reportf(st.Pos, fn.shortName())
		}
		testdata := analysistest.TestData()
		analysistest.Run(t, testdata, analyzer, "tests")
	})
}

func TestFlags(t *testing.T) {
	analyzer := New()
	analyzer.Flags.SetOutput(io.Discard) // TODO(junk1tm): does not work, the usage is still printed.

	t.Run("ok", func(t *testing.T) {
		err := analyzer.Flags.Parse([]string{"-fn=test.Test:test:0"})
		if err != nil {
			t.Errorf("got %v; want no error", err)
		}
	})

	t.Run("invalid format", func(t *testing.T) {
		const want = `invalid value "test.Test" for flag -fn: invalid syntax`
		err := analyzer.Flags.Parse([]string{"-fn=test.Test"})
		if got := err.Error(); got != want {
			t.Errorf("got %q; want %q", got, want)
		}
	})

	t.Run("non-number argpos", func(t *testing.T) {
		const want = `invalid value "test.Test:test:-" for flag -fn: strconv.Atoi: parsing "-": invalid syntax`
		err := analyzer.Flags.Parse([]string{"-fn=test.Test:test:-"})
		if got := err.Error(); got != want {
			t.Errorf("got %q; want %q", got, want)
		}
	})
}

func prepareTestFiles(t *testing.T) {
	testdata := analysistest.TestData()

	t.Cleanup(func() {
		_ = os.RemoveAll(filepath.Join(testdata, "src"))
	})

	hardlink := func(dir, file string) {
		target := filepath.Join(testdata, "src", dir, file)
		if err := os.MkdirAll(filepath.Dir(target), 0o777); err != nil {
			t.Fatal(err)
		}
		if err := os.Link(filepath.Join(testdata, file), target); err != nil {
			t.Fatal(err)
		}
	}
	hardlink("tests", "tests.go")
	hardlink("builtins", "builtins.go")

	for file, data := range stubs {
		target := filepath.Join(testdata, "src", file)
		if err := os.MkdirAll(filepath.Dir(target), 0o777); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(target, []byte(data), 0o666); err != nil {
			t.Fatal(err)
		}
	}
}

var stubs = map[string]string{
	"gopkg.in/yaml.v3/yaml.go": `package yaml
import "io"
func Marshal(_ any) ([]byte, error)   { return nil, nil }
func Unmarshal(_ []byte, _ any) error { return nil }
type Encoder struct{}
func NewEncoder(_ io.Writer) *Encoder { return nil }
func (*Encoder) Encode(_ any) error   { return nil }
type Decoder struct{}
func NewDecoder(_ io.Reader) *Decoder { return nil }
func (*Decoder) Decode(_ any) error   { return nil }`,

	"github.com/BurntSushi/toml/toml.go": `package toml
import "io"
import "io/fs"
func Unmarshal(_ []byte, _ any) error { return nil }
type MetaData struct{}
func Decode(_ string, _ any) (MetaData, error)            { return MetaData{}, nil }
func DecodeFS(_ fs.FS, _ string, _ any) (MetaData, error) { return MetaData{}, nil }
func DecodeFile(_ string, _ any) (MetaData, error)        { return MetaData{}, nil }
type Encoder struct{}
func NewEncoder(_ io.Writer) *Encoder { return nil }
func (*Encoder) Encode(_ any) error   { return nil }
type Decoder struct{}
func NewDecoder(_ io.Reader) *Decoder { return nil }
func (*Decoder) Decode(_ any) error   { return nil }`,

	"github.com/mitchellh/mapstructure/mapstructure.go": `package mapstructure
type Metadata struct{}
func Decode(_, _ any) error                          { return nil }
func DecodeMetadata(_, _ any, _ *Metadata) error     { return nil }
func WeakDecode(_, _ any) error                      { return nil }
func WeakDecodeMetadata(_, _ any, _ *Metadata) error { return nil }`,

	"example.com/custom/custom.go": `package custom
func Marshal(_ any) ([]byte, error)   { return nil, nil }
func Unmarshal(_ []byte, _ any) error { return nil }`,
}
