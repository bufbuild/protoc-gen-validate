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
	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate/conformance/cases"
	"github.com/bufbuild/protovalidate/tools/protovalidate-conformance/internal/results"
	"github.com/bufbuild/protovalidate/tools/protovalidate-conformance/internal/suites"
)

func nestedSuite() suites.Suite {
	return suites.Suite{
		"wkt/uuid/field/valid": {
			Message: &cases.WktLevelOne{
				Two: &cases.WktLevelOne_WktLevelTwo{
					Three: &cases.WktLevelOne_WktLevelTwo_WktLevelThree{
						Uuid: "f81d16ef-40e2-40c6-bebc-89aaf5292f9a",
					},
				},
			},
			Expected: results.Success(true),
		},
		"wkt/uuid/field/invalid": {
			Message: &cases.WktLevelOne{
				Two: &cases.WktLevelOne_WktLevelTwo{
					Three: &cases.WktLevelOne_WktLevelTwo_WktLevelThree{
						Uuid: "not-a-valid-uuid",
					},
				},
			},
			Expected: results.Violations(&validate.Violation{
				FieldPath:    "two.three.uuid",
				ConstraintId: "string.uuid",
			}),
		},
	}
}
