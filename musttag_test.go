package musttag_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/junk1tm/musttag"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, musttag.Analyzer)
}
