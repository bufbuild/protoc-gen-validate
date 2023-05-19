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
	"time"

	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate/conformance/cases"
	"github.com/bufbuild/protovalidate/tools/protovalidate-conformance/internal/results"
	"github.com/bufbuild/protovalidate/tools/protovalidate-conformance/internal/suites"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func timestampSuite() suites.Suite {
	return suites.Suite{
		"none/valid": {
			Message:  &cases.TimestampNone{Val: &timestamppb.Timestamp{Seconds: 123}},
			Expected: results.Success(true),
		},
		"required/valid": {
			Message:  &cases.TimestampRequired{Val: &timestamppb.Timestamp{}},
			Expected: results.Success(true),
		},
		"required/invalid": {
			Message:  &cases.TimestampRequired{Val: nil},
			Expected: results.Violations(&validate.Violation{FieldPath: "val", ConstraintId: "required"}),
		},
		"const/valid": {
			Message:  &cases.TimestampConst{Val: &timestamppb.Timestamp{Seconds: 3, Nanos: 0}},
			Expected: results.Success(true),
		},
		"const/valid/empty": {
			Message:  &cases.TimestampConst{},
			Expected: results.Success(true),
		},
		"const/invalid": {
			Message:  &cases.TimestampConst{Val: &timestamppb.Timestamp{Nanos: 3}},
			Expected: results.Violations(&validate.Violation{FieldPath: "val", ConstraintId: "timestamp.const"}),
		},
		"lt/valid": {
			Message:  &cases.TimestampLT{Val: &timestamppb.Timestamp{Seconds: -1}},
			Expected: results.Success(true),
		},
		"lt/valid/empty": {
			Message:  &cases.TimestampLT{},
			Expected: results.Success(true),
		},
		"lt/invalid/equal": {
			Message:  &cases.TimestampLT{Val: &timestamppb.Timestamp{}},
			Expected: results.Violations(&validate.Violation{FieldPath: "val", ConstraintId: "timestamp.lt"}),
		},
		"lt/invalid": {
			Message:  &cases.TimestampLT{Val: &timestamppb.Timestamp{Seconds: 1}},
			Expected: results.Violations(&validate.Violation{FieldPath: "val", ConstraintId: "timestamp.lt"}),
		},
		"lte/valid": {
			Message:  &cases.TimestampLTE{Val: &timestamppb.Timestamp{}},
			Expected: results.Success(true),
		},
		"lte/valid/empty": {
			Message:  &cases.TimestampLTE{},
			Expected: results.Success(true),
		},
		"lte/valid/equal": {
			Message:  &cases.TimestampLTE{Val: &timestamppb.Timestamp{Seconds: 1}},
			Expected: results.Success(true),
		},
		"lte/invalid": {
			Message:  &cases.TimestampLTE{Val: &timestamppb.Timestamp{Seconds: 1, Nanos: 1}},
			Expected: results.Violations(&validate.Violation{FieldPath: "val", ConstraintId: "timestamp.lte"}),
		},
		"gt/valid": {
			Message:  &cases.TimestampGT{Val: &timestamppb.Timestamp{Seconds: 1}},
			Expected: results.Success(true),
		},
		"gt/valid/empty": {
			Message:  &cases.TimestampGT{},
			Expected: results.Success(true),
		},
		"gt/invalid/equal": {
			Message:  &cases.TimestampGT{Val: &timestamppb.Timestamp{Nanos: 1000}},
			Expected: results.Violations(&validate.Violation{FieldPath: "val", ConstraintId: "timestamp.gt"}),
		},
		"gt/invalid": {
			Message:  &cases.TimestampGT{Val: &timestamppb.Timestamp{}},
			Expected: results.Violations(&validate.Violation{FieldPath: "val", ConstraintId: "timestamp.gt"}),
		},
		"gte/valid": {
			Message:  &cases.TimestampGTE{Val: &timestamppb.Timestamp{Seconds: 3}},
			Expected: results.Success(true),
		},
		"gte/valid/empty": {
			Message:  &cases.TimestampGTE{},
			Expected: results.Success(true),
		},
		"gte/valid/equal": {
			Message:  &cases.TimestampGTE{Val: &timestamppb.Timestamp{Nanos: 1000000}},
			Expected: results.Success(true),
		},
		"gte/invalid": {
			Message:  &cases.TimestampGTE{Val: &timestamppb.Timestamp{Seconds: -1}},
			Expected: results.Violations(&validate.Violation{FieldPath: "val", ConstraintId: "timestamp.gte"}),
		},
		"gt_lt/valid": {
			Message:  &cases.TimestampGTLT{Val: &timestamppb.Timestamp{Nanos: 1000}},
			Expected: results.Success(true),
		},
		"gt_lt/valid/empty": {
			Message:  &cases.TimestampGTLT{},
			Expected: results.Success(true),
		},
		"gt_lt/invalid/above": {
			Message:  &cases.TimestampGTLT{Val: &timestamppb.Timestamp{Seconds: 1000}},
			Expected: results.Violations(&validate.Violation{FieldPath: "val", ConstraintId: "timestamp.gt_lt"}),
		},
		"gt_lt/invalid/below": {
			Message:  &cases.TimestampGTLT{Val: &timestamppb.Timestamp{Seconds: -1000}},
			Expected: results.Violations(&validate.Violation{FieldPath: "val", ConstraintId: "timestamp.gt_lt"}),
		},
		"gt_lt/invalid/max": {
			Message:  &cases.TimestampGTLT{Val: &timestamppb.Timestamp{Seconds: 1}},
			Expected: results.Violations(&validate.Violation{FieldPath: "val", ConstraintId: "timestamp.gt_lt"}),
		},
		"gt_lt/invalid/min": {
			Message:  &cases.TimestampGTLT{Val: &timestamppb.Timestamp{}},
			Expected: results.Violations(&validate.Violation{FieldPath: "val", ConstraintId: "timestamp.gt_lt"}),
		},
		"exclusive/gt_lt/valid/empty": {
			Message:  &cases.TimestampExLTGT{},
			Expected: results.Success(true),
		},
		"exclusive/gt_lt/valid/above": {
			Message:  &cases.TimestampExLTGT{Val: &timestamppb.Timestamp{Seconds: 2}},
			Expected: results.Success(true),
		},
		"exclusive/gt_lt/valid/below": {
			Message:  &cases.TimestampExLTGT{Val: &timestamppb.Timestamp{Seconds: -1}},
			Expected: results.Success(true),
		},
		"exclusive/gt_lt/invalid": {
			Message:  &cases.TimestampExLTGT{Val: &timestamppb.Timestamp{Nanos: 1000}},
			Expected: results.Violations(&validate.Violation{FieldPath: "val", ConstraintId: "timestamp.gt_lt_exclusive"}),
		},
		"exclusive/gt_lt/invalid/max": {
			Message:  &cases.TimestampExLTGT{Val: &timestamppb.Timestamp{Seconds: 1}},
			Expected: results.Violations(&validate.Violation{FieldPath: "val", ConstraintId: "timestamp.gt_lt_exclusive"}),
		},
		"exclusive/gt_lt/invalid/min": {
			Message:  &cases.TimestampExLTGT{Val: &timestamppb.Timestamp{}},
			Expected: results.Violations(&validate.Violation{FieldPath: "val", ConstraintId: "timestamp.gt_lt_exclusive"}),
		},
		"gte_lte/valid": {
			Message:  &cases.TimestampGTELTE{Val: &timestamppb.Timestamp{Seconds: 60, Nanos: 1}},
			Expected: results.Success(true),
		},
		"gte_lte/valid/empty": {
			Message:  &cases.TimestampGTELTE{},
			Expected: results.Success(true),
		},
		"gte_lte/valid/max": {
			Message:  &cases.TimestampGTELTE{Val: &timestamppb.Timestamp{Seconds: 3600}},
			Expected: results.Success(true),
		},
		"gte_lte/valid/min": {
			Message:  &cases.TimestampGTELTE{Val: &timestamppb.Timestamp{Seconds: 60}},
			Expected: results.Success(true),
		},
		"gte_lte/invalid/above": {
			Message:  &cases.TimestampGTELTE{Val: &timestamppb.Timestamp{Seconds: 3600, Nanos: 1}},
			Expected: results.Violations(&validate.Violation{FieldPath: "val", ConstraintId: "timestamp.gte_lte"}),
		},
		"gte_lte/invalid/below": {
			Message:  &cases.TimestampGTELTE{Val: &timestamppb.Timestamp{Seconds: 59}},
			Expected: results.Violations(&validate.Violation{FieldPath: "val", ConstraintId: "timestamp.gte_lte"}),
		},
		"exclusive/gte_lte/valid/empty": {
			Message:  &cases.TimestampExGTELTE{},
			Expected: results.Success(true),
		},
		"exclusive/gte_lte/valid/above": {
			Message:  &cases.TimestampExGTELTE{Val: &timestamppb.Timestamp{Seconds: 3601}},
			Expected: results.Success(true),
		},
		"exclusive/gte_lte/valid/below": {
			Message:  &cases.TimestampExGTELTE{Val: &timestamppb.Timestamp{}},
			Expected: results.Success(true),
		},
		"exclusive/gte_lte/valid/max": {
			Message:  &cases.TimestampExGTELTE{Val: &timestamppb.Timestamp{Seconds: 3600}},
			Expected: results.Success(true),
		},
		"exclusive/gte_lte/valid/min": {
			Message:  &cases.TimestampExGTELTE{Val: &timestamppb.Timestamp{Seconds: 60}},
			Expected: results.Success(true),
		},
		"exclusive/gte_lte/invalid": {
			Message:  &cases.TimestampExGTELTE{Val: &timestamppb.Timestamp{Seconds: 61}},
			Expected: results.Violations(&validate.Violation{FieldPath: "val", ConstraintId: "timestamp.gte_lte_exclusive"}),
		},
		"lt_now/valid": {
			Message:  &cases.TimestampLTNow{Val: &timestamppb.Timestamp{}},
			Expected: results.Success(true),
		},
		"lt_now/valid/empty": {
			Message:  &cases.TimestampLTNow{},
			Expected: results.Success(true),
		},
		"lt_now/invalid": {
			Message:  &cases.TimestampLTNow{Val: &timestamppb.Timestamp{Seconds: time.Now().Unix() + 7200}},
			Expected: results.Violations(&validate.Violation{FieldPath: "val", ConstraintId: "timestamp.lt_now"}),
		},
		"gt_now/valid": {
			Message:  &cases.TimestampGTNow{Val: &timestamppb.Timestamp{Seconds: time.Now().Unix() + 7200}},
			Expected: results.Success(true),
		},
		"gt_now/valid/empty": {
			Message:  &cases.TimestampGTNow{},
			Expected: results.Success(true),
		},
		"gt_now/invalid": {
			Message:  &cases.TimestampGTNow{Val: &timestamppb.Timestamp{}},
			Expected: results.Violations(&validate.Violation{FieldPath: "val", ConstraintId: "timestamp.gt_now"}),
		},
		"within/valid": {
			Message:  &cases.TimestampWithin{Val: timestamppb.Now()},
			Expected: results.Success(true),
		},
		"within/valid/empty": {
			Message:  &cases.TimestampWithin{},
			Expected: results.Success(true),
		},
		"within/invalid/below": {
			Message:  &cases.TimestampWithin{Val: &timestamppb.Timestamp{}},
			Expected: results.Violations(&validate.Violation{FieldPath: "val", ConstraintId: "timestamp.within"}),
		},
		"within/invalid/above": {
			Message:  &cases.TimestampWithin{Val: &timestamppb.Timestamp{Seconds: time.Now().Unix() + 7200}},
			Expected: results.Violations(&validate.Violation{FieldPath: "val", ConstraintId: "timestamp.within"}),
		},
		"lt_now/within/valid": {
			Message:  &cases.TimestampLTNowWithin{Val: &timestamppb.Timestamp{Seconds: time.Now().Unix() - 1800}},
			Expected: results.Success(true),
		},
		"lt_now/within/valid/empty": {
			Message:  &cases.TimestampLTNowWithin{},
			Expected: results.Success(true),
		},
		"lt_now/within/invalid/lt": {
			Message: &cases.TimestampLTNowWithin{Val: &timestamppb.Timestamp{Seconds: time.Now().Unix() + 1800}},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "timestamp.lt_now"}),
		},
		"lt_now/within/invalid/within": {
			Message:  &cases.TimestampLTNowWithin{Val: &timestamppb.Timestamp{Seconds: time.Now().Unix() / 7200}},
			Expected: results.Violations(&validate.Violation{FieldPath: "val", ConstraintId: "timestamp.within"}),
		},
		"gt_now/within/valid": {
			Message:  &cases.TimestampGTNowWithin{Val: &timestamppb.Timestamp{Seconds: time.Now().Unix() + 1800}},
			Expected: results.Success(true),
		},
		"gt_now/within/valid/empty": {
			Message:  &cases.TimestampGTNowWithin{},
			Expected: results.Success(true),
		},
		"gt_now/within/invalid/gt": {
			Message: &cases.TimestampGTNowWithin{Val: &timestamppb.Timestamp{Seconds: time.Now().Unix() / 1800}},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "timestamp.gt_now"},
				&validate.Violation{FieldPath: "val", ConstraintId: "timestamp.within"},
			),
		},
		"gt_now/within/invalid/within": {
			Message:  &cases.TimestampGTNowWithin{Val: &timestamppb.Timestamp{Seconds: time.Now().Unix() + 7200}},
			Expected: results.Violations(&validate.Violation{FieldPath: "val", ConstraintId: "timestamp.within"}),
		},
	}
}
