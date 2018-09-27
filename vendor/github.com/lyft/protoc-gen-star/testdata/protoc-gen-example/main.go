package main

import (
	"github.com/lyft/protoc-gen-star"
	"github.com/lyft/protoc-gen-star/lang/go"
)

func main() {
	pgs.Init(
		pgs.DebugEnv("DEBUG"),
	).RegisterModule(
		ASTPrinter(),
		JSONify(),
	).RegisterPostProcessor(
		pgsgo.GoFmt(),
	).Render()
}
