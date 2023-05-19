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
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func kitchenSinkSuite() suites.Suite {
	return suites.Suite{
		"field/valid": {
			Message:  &cases.KitchenSinkMessage{Val: &cases.ComplexTestMsg{Const: "abcd", IntConst: 5, BoolConst: false, FloatVal: &wrapperspb.FloatValue{Value: 1}, DurVal: &durationpb.Duration{Seconds: 3}, TsVal: &timestamppb.Timestamp{Seconds: 17}, FloatConst: 7, DoubleIn: 123, EnumConst: cases.ComplexTestEnum_COMPLEX_TEST_ENUM_TWO, AnyVal: &anypb.Any{TypeUrl: "type.googleapis.com/google.protobuf.Duration"}, RepTsVal: []*timestamppb.Timestamp{{Seconds: 3}}, MapVal: map[int32]string{-1: "a", -2: "b"}, BytesVal: []byte("\x00\x99"), O: &cases.ComplexTestMsg_X{X: "foobar"}}},
			Expected: results.Success(true),
		},
		"unset/valid": {
			Message:  &cases.KitchenSinkMessage{},
			Expected: results.Success(true),
		},
		"field/invalid": {
			Message: &cases.KitchenSinkMessage{Val: &cases.ComplexTestMsg{}},
			Expected: results.Violations(
				&validate.Violation{
					FieldPath:    "val.bytes_val",
					ConstraintId: "bytes.const",
					Message:      "field must be exactly the specified value",
				},
				&validate.Violation{
					FieldPath:    "val.double_in",
					ConstraintId: "double.in",
					Message:      "must be one of [456.789000, 123.000000]",
				},
				&validate.Violation{
					FieldPath:    "val.enum_const",
					ConstraintId: "enum.const",
					Message:      "must be exactly 2",
				},
				&validate.Violation{
					FieldPath:    "val.int_const",
					ConstraintId: "int32.const",
					Message:      "must be exactly 5",
				},
				&validate.Violation{
					FieldPath:    "val.o",
					ConstraintId: "required",
					Message:      "exactly one of oneof fields is required",
				},
				&validate.Violation{
					FieldPath:    "val.dur_val",
					ConstraintId: "required",
					Message:      "value is required",
				},
				&validate.Violation{
					FieldPath:    "val.const",
					ConstraintId: "string.const",
					Message:      "must be exactly `abcd`",
				},
			),
		},
		"field/embedded/invalid": {
			Message: &cases.KitchenSinkMessage{Val: &cases.ComplexTestMsg{Another: &cases.ComplexTestMsg{}}},
			Expected: results.Violations(
				&validate.Violation{
					FieldPath:    "val.another.bytes_val",
					ConstraintId: "bytes.const",
					Message:      "field must be exactly the specified value",
				},
				&validate.Violation{
					FieldPath:    "val.bytes_val",
					ConstraintId: "bytes.const",
					Message:      "field must be exactly the specified value",
				},
				&validate.Violation{
					FieldPath:    "val.another.double_in",
					ConstraintId: "double.in",
					Message:      "must be one of [456.789000, 123.000000]",
				},
				&validate.Violation{
					FieldPath:    "val.double_in",
					ConstraintId: "double.in",
					Message:      "must be one of [456.789000, 123.000000]",
				},
				&validate.Violation{
					FieldPath:    "val.another.enum_const",
					ConstraintId: "enum.const",
					Message:      "must be exactly 2",
				},
				&validate.Violation{
					FieldPath:    "val.enum_const",
					ConstraintId: "enum.const",
					Message:      "must be exactly 2",
				},
				&validate.Violation{
					FieldPath:    "val.another.int_const",
					ConstraintId: "int32.const",
					Message:      "must be exactly 5",
				},
				&validate.Violation{
					FieldPath:    "val.int_const",
					ConstraintId: "int32.const",
					Message:      "must be exactly 5",
				},
				&validate.Violation{
					FieldPath:    "val.o",
					ConstraintId: "required",
					Message:      "exactly one of oneof fields is required",
				},
				&validate.Violation{
					FieldPath:    "val.another.o",
					ConstraintId: "required",
					Message:      "exactly one of oneof fields is required",
				},
				&validate.Violation{
					FieldPath:    "val.another.dur_val",
					ConstraintId: "required",
					Message:      "value is required",
				},
				&validate.Violation{
					FieldPath:    "val.dur_val",
					ConstraintId: "required",
					Message:      "value is required",
				},
				&validate.Violation{
					FieldPath:    "val.another.const",
					ConstraintId: "string.const",
					Message:      "must be exactly `abcd`",
				},
				&validate.Violation{
					FieldPath:    "val.const",
					ConstraintId: "string.const",
					Message:      "must be exactly `abcd`",
				},
			),
		},
		"field/transitive/invalid": {
			Message: &cases.KitchenSinkMessage{Val: &cases.ComplexTestMsg{Const: "abcd", BoolConst: true, Nested: &cases.ComplexTestMsg{}}},
			Expected: results.Violations(
				&validate.Violation{
					FieldPath:    "val.bool_const",
					ConstraintId: "bool.const",
					Message:      "must be exactly false",
				},
				&validate.Violation{
					FieldPath:    "val.bytes_val",
					ConstraintId: "bytes.const",
					Message:      "field must be exactly the specified value",
				},
				&validate.Violation{
					FieldPath:    "val.nested.bytes_val",
					ConstraintId: "bytes.const",
					Message:      "field must be exactly the specified value",
				},
				&validate.Violation{
					FieldPath:    "val.double_in",
					ConstraintId: "double.in",
					Message:      "must be one of [456.789000, 123.000000]",
				},
				&validate.Violation{
					FieldPath:    "val.nested.double_in",
					ConstraintId: "double.in",
					Message:      "must be one of [456.789000, 123.000000]",
				},
				&validate.Violation{
					FieldPath:    "val.enum_const",
					ConstraintId: "enum.const",
					Message:      "must be exactly 2",
				},
				&validate.Violation{
					FieldPath:    "val.nested.enum_const",
					ConstraintId: "enum.const",
					Message:      "must be exactly 2",
				},
				&validate.Violation{
					FieldPath:    "val.int_const",
					ConstraintId: "int32.const",
					Message:      "must be exactly 5",
				},
				&validate.Violation{
					FieldPath:    "val.nested.int_const",
					ConstraintId: "int32.const",
					Message:      "must be exactly 5",
				},
				&validate.Violation{
					FieldPath:    "val.o",
					ConstraintId: "required",
					Message:      "exactly one of oneof fields is required",
				},
				&validate.Violation{
					FieldPath:    "val.dur_val",
					ConstraintId: "required",
					Message:      "value is required",
				},
				&validate.Violation{
					FieldPath:    "val.nested.o",
					ConstraintId: "required",
					Message:      "exactly one of oneof fields is required",
				},
				&validate.Violation{
					FieldPath:    "val.nested.dur_val",
					ConstraintId: "required",
					Message:      "value is required",
				},
				&validate.Violation{
					FieldPath:    "val.nested.const",
					ConstraintId: "string.const",
					Message:      "must be exactly `abcd`",
				},
			),
		},
		"many/all-non-message-fields/invalid": {
			Message: &cases.KitchenSinkMessage{Val: &cases.ComplexTestMsg{BoolConst: true, FloatVal: &wrapperspb.FloatValue{}, TsVal: &timestamppb.Timestamp{}, FloatConst: 8, AnyVal: &anypb.Any{TypeUrl: "asdf"}, RepTsVal: []*timestamppb.Timestamp{{Nanos: 1}}}},
			Expected: results.Violations(
				&validate.Violation{
					FieldPath:    "val.any_val",
					ConstraintId: "any.in",
					Message:      "type URL must be in the allow list",
				},
				&validate.Violation{
					FieldPath:    "val.bool_const",
					ConstraintId: "bool.const",
					Message:      "must be exactly false",
				},
				&validate.Violation{
					FieldPath:    "val.bytes_val",
					ConstraintId: "bytes.const",
					Message:      "field must be exactly the specified value",
				},
				&validate.Violation{
					FieldPath:    "val.double_in",
					ConstraintId: "double.in",
					Message:      "must be one of [456.789000, 123.000000]",
				},
				&validate.Violation{
					FieldPath:    "val.enum_const",
					ConstraintId: "enum.const",
					Message:      "must be exactly 2",
				},
				&validate.Violation{
					FieldPath:    "val.float_val",
					ConstraintId: "float.gt",
					Message:      "must be greater than 0",
				},
				&validate.Violation{
					FieldPath:    "val.float_const",
					ConstraintId: "float.lt",
					Message:      "must be less than 8",
				},
				&validate.Violation{
					FieldPath:    "val.int_const",
					ConstraintId: "int32.const",
					Message:      "must be exactly 5",
				},
				&validate.Violation{
					FieldPath:    "val.o",
					ConstraintId: "required",
					Message:      "exactly one of oneof fields is required",
				},
				&validate.Violation{
					FieldPath:    "val.dur_val",
					ConstraintId: "required",
					Message:      "value is required",
				},
				&validate.Violation{
					FieldPath:    "val.const",
					ConstraintId: "string.const",
					Message:      "must be exactly `abcd`",
				},
				&validate.Violation{
					FieldPath:    "val.ts_val",
					ConstraintId: "timestamp.gt",
					Message:      "must be greater than 1970-01-01T00:00:07Z",
				},
				&validate.Violation{
					FieldPath:    "val.rep_ts_val[0]",
					ConstraintId: "timestamp.gte",
					Message:      "must be greater than or equal to 1970-01-01T00:00:00.001Z",
				},
			),
		},
	}
}
