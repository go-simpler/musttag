package musttag

import (
	"go/token"
	"io"
	"os"
	"path/filepath"
	"testing"

	"go-simpler.org/assert"
	. "go-simpler.org/assert/dotimport"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	// NOTE: analysistest does not yet support modules;
	// see https://github.com/golang/go/issues/37054 for details.
	// To be able to run tests with external dependencies,
	// we first need to write a GOPATH-like tree of stubs.
	prepareTestFiles(t)
	testPackages = []string{"tests", "builtins"}

	testdata := analysistest.TestData()

	t.Run("tests", func(t *testing.T) {
		r := report
		defer func() { report = r }()

		report = func(pass *analysis.Pass, st *structType, fn Func, fnPos token.Position) {
			pass.Reportf(st.Pos, fn.shortName())
		}

		analyzer := New()
		analysistest.Run(t, testdata, analyzer, "tests")
	})

	t.Run("builtins", func(t *testing.T) {
		analyzer := New(
			Func{Name: "example.com/custom.Marshal", Tag: "custom", ArgPos: 0},
			Func{Name: "example.com/custom.Unmarshal", Tag: "custom", ArgPos: 1},
		)
		analysistest.Run(t, testdata, analyzer, "builtins")
	})

	t.Run("bad Func.ArgPos", func(t *testing.T) {
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

func prepareTestFiles(t *testing.T) {
	testdata := analysistest.TestData()

	t.Cleanup(func() {
		err := os.RemoveAll(filepath.Join(testdata, "src"))
		assert.NoErr[F](t, err)
	})

	hardlink := func(dir, file string) {
		target := filepath.Join(testdata, "src", dir, file)

		err := os.MkdirAll(filepath.Dir(target), 0o777)
		assert.NoErr[F](t, err)

		err = os.Link(filepath.Join(testdata, file), target)
		assert.NoErr[F](t, err)
	}

	hardlink("tests", "tests.go")
	hardlink("builtins", "builtins.go")

	for file, data := range stubs {
		target := filepath.Join(testdata, "src", file)

		err := os.MkdirAll(filepath.Dir(target), 0o777)
		assert.NoErr[F](t, err)

		err = os.WriteFile(target, []byte(data), 0o666)
		assert.NoErr[F](t, err)
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
