package musttag

import (
	"testing"

	"go-simpler.org/assert"
	. "go-simpler.org/assert/dotimport"
)

func Test_getMainModule(t *testing.T) {
	test := func(name, want, output string) {
		t.Helper()
		t.Run(name, func(t *testing.T) {
			t.Helper()

			gwd := getwd
			co := commandOutput
			defer func() {
				getwd = gwd
				commandOutput = co
			}()

			getwd = func() (string, error) {
				return "/path/to/module1/pkg", nil
			}
			commandOutput = func(name string, args ...string) (string, error) {
				return output, nil
			}

			got, err := getMainModule()
			assert.NoErr[F](t, err)
			assert.Equal[E](t, got, want)
		})
	}

	test("single module", "module1", `
{
	"Path": "module1",
	"Main": true,
	"Dir": "/path/to/module1"
}`)

	test("multiple modules", "module1", `
{
	"Path": "module1",
	"Main": true,
	"Dir": "/path/to/module1"
}
{
	"Path": "module2",
	"Main": true,
	"Dir": "/path/to/module2"
}`)
}
