package musttag

import (
	"testing"

	"go-simpler.org/assert"
	. "go-simpler.org/assert/EF"
)

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
