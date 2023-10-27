package musttag

import (
	"testing"

	"go-simpler.org/assert"
	. "go-simpler.org/assert/EF"
)

func Test_getMainModule(t *testing.T) {
	tests := map[string]struct {
		want, output string
	}{
		"single module": {
			want: "module1",
			output: `
{"Path": "module1", "Main": true, "Dir": "/path/to/module1"}`,
		},
		"multiple modules": {
			want: "module1",
			output: `
{"Path": "module1", "Main": true, "Dir": "/path/to/module1"}
{"Path": "module2", "Main": true, "Dir": "/path/to/module2"}`,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			gwd, co := getwd, commandOutput
			defer func() { getwd, commandOutput = gwd, co }()

			getwd = func() (string, error) {
				return "/path/to/module1/pkg", nil
			}
			commandOutput = func(string, ...string) (string, error) {
				return test.output, nil
			}

			got, err := getMainModule()
			assert.NoErr[F](t, err)
			assert.Equal[E](t, got, test.want)
		})
	}
}

func Test_cutVendor(t *testing.T) {
	tests := []struct {
		path, want string
	}{
		{"foo/bar.A", "foo/bar.A"},
		{"vendor/foo/bar.A", "foo/bar.A"},
		{"test/vendor/foo/bar.A", "foo/bar.A"},
		{"(test/vendor/foo/bar.A).B", "(foo/bar.A).B"},
		{"(*test/vendor/foo/bar.A).B", "(*foo/bar.A).B"},
	}

	for _, test := range tests {
		got := cutVendor(test.path)
		assert.Equal[E](t, got, test.want)
	}
}
