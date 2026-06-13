package main

import (
	"flag"
	"fmt"
	"os"

	pgs "github.com/lyft/protoc-gen-star/v2"
	pgsgo "github.com/lyft/protoc-gen-star/v2/lang/go"
	"google.golang.org/protobuf/types/pluginpb"

	"github.com/envoyproxy/protoc-gen-validate/module"
	"github.com/envoyproxy/protoc-gen-validate/version"
)

func main() {
	versionFlag := flag.Bool("version", false, "print version information")
	flag.Parse()

	if *versionFlag {
		fmt.Println(version.String())
		os.Exit(0)
	}

	optional := uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
	pgs.
		Init(pgs.DebugEnv("DEBUG_PGV"), pgs.SupportedFeatures(&optional)).
		RegisterModule(module.ValidatorForLanguage("go")).
		RegisterPostProcessor(pgsgo.GoFmt()).
		Render()
}
