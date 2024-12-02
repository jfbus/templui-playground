package main

import (
	"fmt"
	"os"

	"github.com/jfbus/templui-playground/templplaygroundgenerator/generator"
)

func main() {
	g := generator.NewGenerator()
	g.Parse()
	if err := g.Generate(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err) // nolint
		os.Exit(1)
	}
}
