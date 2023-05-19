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
	"math"

	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate/conformance/cases"
	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate/conformance/cases/other_package"
	"github.com/bufbuild/protovalidate/tools/protovalidate-conformance/internal/results"
	"github.com/bufbuild/protovalidate/tools/protovalidate-conformance/internal/suites"
)

func enumSuite() suites.Suite {
	return suites.Suite{
		"none/valid": {
			Message:  &cases.EnumNone{Val: cases.TestEnum_TEST_ENUM_ONE},
			Expected: results.Success(true),
		},
		"const/valid": {
			Message:  &cases.EnumConst{Val: cases.TestEnum_TEST_ENUM_TWO},
			Expected: results.Success(true),
		},
		"const/invalid": {
			Message: &cases.EnumConst{Val: cases.TestEnum_TEST_ENUM_ONE},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "enum.const"}),
		},
		"alias/const/valid": {
			Message:  &cases.EnumAliasConst{Val: cases.TestEnumAlias_TEST_ENUM_ALIAS_B},
			Expected: results.Success(true),
		},
		"alias/const/invalid": {
			Message: &cases.EnumAliasConst{Val: cases.TestEnumAlias_TEST_ENUM_ALIAS_GAMMA},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "enum.const"}),
		},
		"defined_only/valid/unspecified": {
			Message:  &cases.EnumDefined{Val: cases.TestEnum_TEST_ENUM_UNSPECIFIED},
			Expected: results.Success(true),
		},
		"defined_only/valid/specified": {
			Message:  &cases.EnumDefined{Val: cases.TestEnum_TEST_ENUM_ONE},
			Expected: results.Success(true),
		},
		"defined_only/invalid/unknown": {
			Message: &cases.EnumDefined{Val: math.MaxInt32},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "enum.defined_only"}),
		},
		"alias/defined_only/valid/unspecified": {
			Message:  &cases.EnumAliasDefined{Val: cases.TestEnumAlias_TEST_ENUM_ALIAS_UNSPECIFIED},
			Expected: results.Success(true),
		},
		"alias/defined_only/valid/specified": {
			Message:  &cases.EnumAliasDefined{Val: cases.TestEnumAlias_TEST_ENUM_ALIAS_C},
			Expected: results.Success(true),
		},
		"alias/defined_only/invalid/unknown": {
			Message: &cases.EnumAliasDefined{Val: math.MaxInt32},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "enum.defined_only"}),
		},
		"in/valid": {
			Message:  &cases.EnumIn{Val: cases.TestEnum_TEST_ENUM_TWO},
			Expected: results.Success(true),
		},
		"in/invalid": {
			Message: &cases.EnumIn{Val: cases.TestEnum_TEST_ENUM_ONE},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "enum.in"}),
		},
		"alias/in/valid": {
			Message:  &cases.EnumAliasIn{Val: cases.TestEnumAlias_TEST_ENUM_ALIAS_BETA},
			Expected: results.Success(true),
		},
		"alias/in/invalid": {
			Message: &cases.EnumAliasIn{Val: cases.TestEnumAlias_TEST_ENUM_ALIAS_A},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "enum.in"}),
		},
		"not_in/valid": {
			Message:  &cases.EnumNotIn{Val: cases.TestEnum_TEST_ENUM_UNSPECIFIED},
			Expected: results.Success(true),
		},
		"not_in/valid/unspecified": {
			Message:  &cases.EnumNotIn{Val: math.MaxInt32},
			Expected: results.Success(true),
		},
		"not_in/invalid": {
			Message: &cases.EnumNotIn{Val: cases.TestEnum_TEST_ENUM_ONE},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "enum.not_in"}),
		},
		"alias/not_in/valid": {
			Message:  &cases.EnumAliasNotIn{Val: cases.TestEnumAlias_TEST_ENUM_ALIAS_BETA},
			Expected: results.Success(true),
		},
		"alias/not_in/invalid": {
			Message: &cases.EnumAliasNotIn{Val: cases.TestEnumAlias_TEST_ENUM_ALIAS_A},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "enum.not_in"}),
		},
		"external/defined_only/valid": {
			Message:  &cases.EnumExternal{Val: other_package.Embed_ENUMERATED_VALUE},
			Expected: results.Success(true),
		},
		"external/defined_only/invalid": {
			Message: &cases.EnumExternal{Val: math.MaxInt32},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "enum.defined_only"}),
		},
	}
}
