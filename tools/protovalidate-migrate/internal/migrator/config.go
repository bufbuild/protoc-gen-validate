// Copyright 2023 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package migrator

import (
	"fmt"
	"log"
	"os"

	flag "github.com/spf13/pflag"
)

type Config struct {
	Verbose   bool
	write     bool
	print     bool
	output    string
	PGVImport string
	PVImport  string
	RemovePGV bool
	ReplacePV bool
	Paths     []string
}

func ParseFlags() Config {
	cfg := Config{
		PGVImport: "validate/validate.proto",
		PVImport:  "buf/validate/validate.proto",
	}
	log.SetFlags(0)

	flag.Usage = func() {
		_, _ = fmt.Fprintf(os.Stderr,
			"Usage: %s [FLAGS ...] PROTO_PATHS ...\n\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}
	flag.BoolVarP(&cfg.Verbose, "verbose", "v", cfg.Verbose, "verbose logging")
	flag.BoolVarP(&cfg.write, "write", "w", cfg.write, "overwrite target files in-place; mutually exclusive with -o")
	flag.StringVarP(&cfg.output, "output", "o", cfg.output, "write output to the given path; mutually exclusive with -w")
	flag.StringVar(&cfg.PGVImport, "legacy-import", cfg.PGVImport, "protoc-gen-validate proto import path")
	flag.StringVar(&cfg.PVImport, "protovalidate-import", cfg.PVImport, "protovalidate proto import path")
	flag.BoolVar(&cfg.RemovePGV, "remove-legacy", cfg.RemovePGV, "remove protoc-gen-validate options")
	flag.BoolVar(&cfg.ReplacePV, "replace-protovalidate", cfg.ReplacePV,
		"replace protovalidate options to match protoc-gen-validate options (only if present)")
	flag.Parse()

	cfg.print = !cfg.write && cfg.output == ""
	cfg.Paths = flag.Args()

	if ct := len(cfg.Paths); ct == 0 {
		log.Fatal("error: no PROTO_PATH arguments provided")
	}

	if cfg.write && cfg.output != "" {
		log.Fatal("error: cannot use -w flag with -o flag")
	}

	return cfg
}
