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

package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"time"

	flag "github.com/spf13/pflag"
)

type config struct {
	suiteFilterPattern string
	suiteFilter        *regexp.Regexp
	caseFilterPattern  string
	caseFilter         *regexp.Regexp
	suiteTimeout       time.Duration
	verbose            bool
	strict             bool
	proto              bool
	json               bool
	print              bool
	cmd                string
	args               []string
}

func parseFlags() config {
	cfg := config{
		suiteTimeout: 5 * time.Second,
	}
	log.SetFlags(0)

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [FLAGS ...] CMD [ARGS ...]\n\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.StringVar(&cfg.suiteFilterPattern, "suite", cfg.suiteFilterPattern, "regex to filter suites")
	flag.StringVar(&cfg.caseFilterPattern, "case", cfg.caseFilterPattern, "regex to filter cases")
	flag.DurationVar(&cfg.suiteTimeout, "timeout", cfg.suiteTimeout, "per-suite timeout")
	flag.BoolVarP(&cfg.verbose, "verbose", "v", cfg.verbose, "verbose output")
	flag.BoolVar(&cfg.strict, "strict", cfg.strict, "strict mode")
	flag.BoolVar(&cfg.json, "json", cfg.json, "return results as JSON to stdout")
	flag.BoolVar(&cfg.proto, "proto", cfg.proto, "return results as binary serialized proto to stdout")
	flag.Parse()

	cfg.print = !cfg.json && !cfg.proto

	args := flag.Args()
	if len(args) == 0 {
		log.Fatal("a command must be specified")
	}
	cfg.cmd = args[0]
	cfg.args = args[1:]

	if cfg.suiteFilterPattern != "" {
		filter, err := regexp.Compile(cfg.suiteFilterPattern)
		if err != nil {
			log.Fatal(err)
		}
		cfg.suiteFilter = filter
	}

	if cfg.caseFilterPattern != "" {
		filter, err := regexp.Compile(cfg.caseFilterPattern)
		if err != nil {
			log.Fatal(err)
		}
		cfg.caseFilter = filter
	}

	return cfg
}
