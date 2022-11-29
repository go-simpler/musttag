package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/junk1tm/musttag"
)

func main() {
	singlechecker.Main(musttag.New())
}
