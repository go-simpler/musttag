package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"golang.org/x/tools/go/analysis/singlechecker"

	"go-simpler.org/musttag"
)

var version = "dev" // injected at build time.

func main() {
	// override the builtin -V flag.
	flag.Var(versionFlag{}, "V", "print version and exit")
	singlechecker.Main(musttag.New(&musttag.Settings{}))
}

type versionFlag struct{}

func (versionFlag) String() string   { return "" }
func (versionFlag) IsBoolFlag() bool { return true }
func (versionFlag) Set(string) error {
	fmt.Printf("musttag version %s %s/%s\n", version, runtime.GOOS, runtime.GOARCH)
	os.Exit(0)
	return nil
}
