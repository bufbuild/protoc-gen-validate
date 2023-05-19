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
	"log"
	"os"

	"github.com/bufbuild/protovalidate/tools/protovalidate-conformance/internal/cases"
	"github.com/bufbuild/protovalidate/tools/protovalidate-conformance/internal/results"
	"github.com/bufbuild/protovalidate/tools/protovalidate-conformance/internal/suites"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func main() {
	cfg := parseFlags()

	resultSet := &results.Set{
		SuiteFilter: cfg.suiteFilterPattern,
		CaseFilter:  cfg.caseFilterPattern,
		Verbose:     cfg.verbose,
		Strict:      cfg.strict,
	}
	err := cases.GlobalSuites().Range(cfg.suiteFilter, func(suiteName string, suite suites.Suite) error {
		req, err := suite.ToRequestProto(cfg.caseFilter)
		if err != nil || len(req.Cases) == 0 {
			return err
		}
		resp, err := runCommand(&cfg, req)
		if err != nil {
			return err
		}
		res := suite.ProcessResults(
			suiteName,
			cfg.caseFilter,
			resp,
			cfg.verbose,
			cfg.strict,
		)
		resultSet.AddSuite(res, cfg.verbose)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	switch {
	case cfg.proto:
		err = resultSet.MarshalTo(os.Stdout, proto.Marshal)
	case cfg.json:
		err = resultSet.MarshalTo(os.Stdout, protojson.Marshal)
	default:
		resultSet.Print(os.Stderr)
	}

	if err != nil {
		log.Fatal(err)
	}

	os.Exit(int(resultSet.Failures))
}
