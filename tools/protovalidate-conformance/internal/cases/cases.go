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

package cases

import (
	"github.com/bufbuild/protovalidate/tools/protovalidate-conformance/internal/suites"
)

// GlobalSuites is used by the conformance harness, containing all tests to be
// run against implementation targets. The source proto messages used for the
// tests can be found in /protos/testing/buf/validate/conformance/cases.
func GlobalSuites() suites.Suites {
	return suites.Suites{
		"custom_constraints":                              customSuite(),
		"kitchen_sink":                                    kitchenSinkSuite(),
		"standard_constraints/bool":                       boolSuite(),
		"standard_constraints/bytes":                      bytesSuite(),
		"standard_constraints/double":                     doubleSuite(),
		"standard_constraints/enum":                       enumSuite(),
		"standard_constraints/fixed32":                    fixed32Suite(),
		"standard_constraints/fixed64":                    fixed64Suite(),
		"standard_constraints/float":                      floatSuite(),
		"standard_constraints/int32":                      int32Suite(),
		"standard_constraints/int64":                      int64Suite(),
		"standard_constraints/map":                        mapSuite(),
		"standard_constraints/message":                    messageSuite(),
		"standard_constraints/nested":                     nestedSuite(),
		"standard_constraints/oneof":                      oneofSuite(),
		"standard_constraints/repeated":                   repeatedSuite(),
		"standard_constraints/sfixed32":                   sfixed32Suite(),
		"standard_constraints/sfixed64":                   sfixed64Suite(),
		"standard_constraints/sint32":                     sint32Suite(),
		"standard_constraints/sint64":                     sint64Suite(),
		"standard_constraints/string":                     stringSuite(),
		"standard_constraints/uint32":                     uint32Suite(),
		"standard_constraints/uint64":                     uint64Suite(),
		"standard_constraints/well_known_types/any":       anySuite(),
		"standard_constraints/well_known_types/duration":  durationSuite(),
		"standard_constraints/well_known_types/timestamp": timestampSuite(),
		"standard_constraints/well_known_types/wrapper":   wrapperSuite(),
	}
}
