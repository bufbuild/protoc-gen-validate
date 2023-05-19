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
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func wrapperSuite() suites.Suite {
	return suites.Suite{
		"none/valid": {
			Message:  &cases.WrapperNone{Val: &wrapperspb.Int32Value{Value: 123}},
			Expected: results.Success(true),
		},
		"none/empty/valid": {
			Message:  &cases.WrapperNone{Val: nil},
			Expected: results.Success(true),
		},
		"float/valid": {
			Message:  &cases.WrapperFloat{Val: &wrapperspb.FloatValue{Value: 1}},
			Expected: results.Success(true),
		},
		"float/empty/valid": {
			Message:  &cases.WrapperFloat{Val: nil},
			Expected: results.Success(true),
		},
		"float/invalid": {
			Message: &cases.WrapperFloat{Val: &wrapperspb.FloatValue{Value: 0}},
			Expected: results.Violations(&validate.Violation{
				FieldPath:    "val",
				ConstraintId: "float.gt",
				Message:      "must be greater than 0",
			}),
		},
		"double/valid": {
			Message:  &cases.WrapperDouble{Val: &wrapperspb.DoubleValue{Value: 1}},
			Expected: results.Success(true),
		},
		"double/empty/valid": {
			Message:  &cases.WrapperDouble{Val: nil},
			Expected: results.Success(true),
		},
		"double/invalid": {
			Message: &cases.WrapperDouble{Val: &wrapperspb.DoubleValue{Value: 0}},
			Expected: results.Violations(&validate.Violation{
				FieldPath:    "val",
				ConstraintId: "double.gt",
				Message:      "must be greater than 0",
			}),
		},
		"int64/valid": {
			Message:  &cases.WrapperInt64{Val: &wrapperspb.Int64Value{Value: 1}},
			Expected: results.Success(true),
		},
		"int64/empty/valid": {
			Message:  &cases.WrapperInt64{Val: nil},
			Expected: results.Success(true),
		},
		"int64/invalid": {
			Message: &cases.WrapperInt64{Val: &wrapperspb.Int64Value{Value: 0}},
			Expected: results.Violations(&validate.Violation{
				FieldPath:    "val",
				ConstraintId: "int64.gt",
				Message:      "must be greater than 0",
			}),
		},
		"int32/valid": {
			Message:  &cases.WrapperInt32{Val: &wrapperspb.Int32Value{Value: 1}},
			Expected: results.Success(true),
		},
		"int32/empty/valid": {
			Message:  &cases.WrapperInt32{Val: nil},
			Expected: results.Success(true),
		},
		"int32/invalid": {
			Message: &cases.WrapperInt32{Val: &wrapperspb.Int32Value{Value: 0}},
			Expected: results.Violations(&validate.Violation{
				FieldPath:    "val",
				ConstraintId: "int32.gt",
				Message:      "must be greater than 0",
			}),
		},
		"uint64/valid": {
			Message:  &cases.WrapperUInt64{Val: &wrapperspb.UInt64Value{Value: 1}},
			Expected: results.Success(true),
		},
		"uint64/empty/valid": {
			Message:  &cases.WrapperUInt64{Val: nil},
			Expected: results.Success(true),
		},
		"uint64/invalid": {
			Message: &cases.WrapperUInt64{Val: &wrapperspb.UInt64Value{Value: 0}},
			Expected: results.Violations(&validate.Violation{
				FieldPath:    "val",
				ConstraintId: "uint64.gt",
				Message:      "must be greater than 0",
			}),
		},
		"uint32/valid": {
			Message:  &cases.WrapperUInt32{Val: &wrapperspb.UInt32Value{Value: 1}},
			Expected: results.Success(true),
		},
		"uint32/empty/valid": {
			Message:  &cases.WrapperUInt32{Val: nil},
			Expected: results.Success(true),
		},
		"uint32/invalid": {
			Message: &cases.WrapperUInt32{Val: &wrapperspb.UInt32Value{Value: 0}},
			Expected: results.Violations(&validate.Violation{
				FieldPath:    "val",
				ConstraintId: "uint32.gt",
				Message:      "must be greater than 0",
			}),
		},
		"bool/valid": {
			Message:  &cases.WrapperBool{Val: &wrapperspb.BoolValue{Value: true}},
			Expected: results.Success(true),
		},
		"bool/empty/valid": {
			Message:  &cases.WrapperBool{Val: nil},
			Expected: results.Success(true),
		},
		"bool/invalid": {
			Message: &cases.WrapperBool{Val: &wrapperspb.BoolValue{Value: false}},
			Expected: results.Violations(&validate.Violation{
				FieldPath:    "val",
				ConstraintId: "bool.const",
				Message:      "must be exactly true",
			}),
		},
		"string/valid": {
			Message:  &cases.WrapperString{Val: &wrapperspb.StringValue{Value: "foobar"}},
			Expected: results.Success(true),
		},
		"string/empty/valid": {
			Message:  &cases.WrapperString{Val: nil},
			Expected: results.Success(true),
		},
		"string/invalid": {
			Message: &cases.WrapperString{Val: &wrapperspb.StringValue{Value: "fizzbuzz"}},
			Expected: results.Violations(&validate.Violation{
				FieldPath:    "val",
				ConstraintId: "string.suffix",
				Message:      "must have the suffix `bar`",
			}),
		},
		"bytes/valid": {
			Message:  &cases.WrapperBytes{Val: &wrapperspb.BytesValue{Value: []byte("foo")}},
			Expected: results.Success(true),
		},
		"bytes/empty/valid": {
			Message:  &cases.WrapperBytes{Val: nil},
			Expected: results.Success(true),
		},
		"bytes/invalid": {
			Message: &cases.WrapperBytes{Val: &wrapperspb.BytesValue{Value: []byte("x")}},
			Expected: results.Violations(&validate.Violation{
				FieldPath:    "val",
				ConstraintId: "bytes.min_len",
				Message:      "must contain at least 3 bytes",
			}),
		},
		"required/string/valid": {
			Message:  &cases.WrapperRequiredString{Val: &wrapperspb.StringValue{Value: "bar"}},
			Expected: results.Success(true),
		},
		"required/string/invalid": {
			Message: &cases.WrapperRequiredString{Val: &wrapperspb.StringValue{Value: "foo"}},
			Expected: results.Violations(&validate.Violation{
				FieldPath:    "val",
				ConstraintId: "string.const",
				Message:      "must be exactly `bar`",
			}),
		},
		"required/string/empty/invalid": {
			Message: &cases.WrapperRequiredString{},
			Expected: results.Violations(&validate.Violation{
				FieldPath:    "val",
				ConstraintId: "required",
				Message:      "value is required",
			}),
		},
		"required/empty/string/valid": {
			Message:  &cases.WrapperRequiredEmptyString{Val: &wrapperspb.StringValue{Value: ""}},
			Expected: results.Success(true),
		},
		"required/empty/string/invalid": {
			Message: &cases.WrapperRequiredEmptyString{Val: &wrapperspb.StringValue{Value: "foo"}},
			Expected: results.Violations(&validate.Violation{
				FieldPath:    "val",
				ConstraintId: "string.const",
				Message:      "must be exactly ``",
			}),
		},
		"required/empty/string/empty/invalid": {
			Message:  &cases.WrapperRequiredEmptyString{},
			Expected: results.Violations(&validate.Violation{FieldPath: "val", ConstraintId: "required", Message: "value is required"}),
		},
		"optional/string/uuid/valid": {
			Message:  &cases.WrapperOptionalUuidString{Val: &wrapperspb.StringValue{Value: "8b72987b-024a-43b3-b4cf-647a1f925c5d"}},
			Expected: results.Success(true),
		},
		"optional/string/uuid/empty/valid": {
			Message:  &cases.WrapperOptionalUuidString{},
			Expected: results.Success(true),
		},
		"optional/string/uuid/invalid": {
			Message: &cases.WrapperOptionalUuidString{Val: &wrapperspb.StringValue{Value: "foo"}},
			Expected: results.Violations(&validate.Violation{
				FieldPath:    "val",
				ConstraintId: "string.uuid",
				Message:      "must be a valid UUID as defined by RFC 4122",
			}),
		},
		"required/float/valid": {
			Message:  &cases.WrapperRequiredFloat{Val: &wrapperspb.FloatValue{Value: 1}},
			Expected: results.Success(true),
		},
		"required/float/invalid": {
			Message: &cases.WrapperRequiredFloat{Val: &wrapperspb.FloatValue{Value: -5}},
			Expected: results.Violations(&validate.Violation{
				FieldPath:    "val",
				ConstraintId: "float.gt",
				Message:      "must be greater than 0",
			}),
		},
		"required/float/empty/invalid": {
			Message: &cases.WrapperRequiredFloat{},
			Expected: results.Violations(&validate.Violation{
				FieldPath:    "val",
				ConstraintId: "required",
				Message:      "value is required",
			}),
		},
	}
}
