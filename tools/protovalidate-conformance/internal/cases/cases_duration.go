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
	"google.golang.org/protobuf/types/known/durationpb"
)

func durationSuite() suites.Suite {
	return suites.Suite{
		"none/valid": {
			Message: &cases.DurationNone{
				Val: &durationpb.Duration{
					Seconds: 123,
				},
			},
			Expected: results.Success(true),
		},
		"required/valid": {
			Message: &cases.DurationRequired{
				Val: &durationpb.Duration{},
			},
			Expected: results.Success(true),
		},
		"required/invalid": {
			Message: &cases.DurationRequired{
				Val: nil,
			},
			Expected: results.Violations(&validate.Violation{
				FieldPath:    "val",
				ConstraintId: "required",
				Message:      "value is required",
			}),
		},
		"const/valid": {
			Message: &cases.DurationConst{
				Val: &durationpb.Duration{
					Seconds: 3,
				},
			},
			Expected: results.Success(true),
		},
		"const/valid/empty": {
			Message: &cases.DurationConst{}, Expected: results.Success(true),
		},
		"const/invalid": {
			Message: &cases.DurationConst{
				Val: &durationpb.Duration{
					Nanos: 3,
				},
			},
			Expected: results.Violations(&validate.Violation{
				FieldPath:    "val",
				ConstraintId: "duration.const",
				Message:      "field must be exactly the specified value",
			}),
		},
		"in/valid": {
			Message: &cases.DurationIn{
				Val: &durationpb.Duration{
					Seconds: 1,
				},
			},
			Expected: results.Success(true),
		},
		"in/valid/empty": {
			Message: &cases.DurationIn{}, Expected: results.Success(true),
		},
		"in/invalid": {
			Message: &cases.DurationIn{
				Val: &durationpb.Duration{},
			},
			Expected: results.Violations(&validate.Violation{
				FieldPath:    "val",
				ConstraintId: "duration.in",
				Message:      "field must be in the specified set of values",
			}),
		},
		"not in/valid": {
			Message: &cases.DurationNotIn{
				Val: &durationpb.Duration{
					Nanos: 1,
				},
			},
			Expected: results.Success(true),
		},
		"not in/valid/empty": {
			Message: &cases.DurationNotIn{}, Expected: results.Success(true),
		},
		"not in/invalid": {
			Message: &cases.DurationNotIn{
				Val: &durationpb.Duration{},
			},
			Expected: results.Violations(&validate.Violation{
				FieldPath:    "val",
				ConstraintId: "duration.not_in",
				Message:      "field must not be in the specified set of values",
			}),
		},
		"lt/valid": {
			Message: &cases.DurationLT{
				Val: &durationpb.Duration{
					Nanos: -1,
				},
			},
			Expected: results.Success(true),
		},
		"lt/valid/empty": {
			Message: &cases.DurationLT{}, Expected: results.Success(true),
		},
		"lt/invalid/equal": {
			Message: &cases.DurationLT{
				Val: &durationpb.Duration{},
			},
			Expected: results.Violations(&validate.Violation{
				FieldPath:    "val",
				ConstraintId: "duration.lt",
				Message:      "field must be less than 0s",
			}),
		},
		"lt/invalid": {
			Message: &cases.DurationLT{
				Val: &durationpb.Duration{
					Seconds: 1,
				},
			},
			Expected: results.Violations(&validate.Violation{
				FieldPath:    "val",
				ConstraintId: "duration.lt",
				Message:      "field must be less than 0s",
			}),
		},
		"lte/valid": {
			Message: &cases.DurationLTE{
				Val: &durationpb.Duration{},
			},
			Expected: results.Success(true),
		},
		"lte/valid/empty": {
			Message: &cases.DurationLTE{}, Expected: results.Success(true),
		},
		"lte/valid/equal": {
			Message: &cases.DurationLTE{
				Val: &durationpb.Duration{
					Seconds: 1,
				},
			},
			Expected: results.Success(true),
		},
		"lte/invalid": {
			Message: &cases.DurationLTE{
				Val: &durationpb.Duration{
					Seconds: 1, Nanos: 1,
				},
			},
			Expected: results.Violations(&validate.Violation{
				FieldPath:    "val",
				ConstraintId: "duration.lte",
				Message:      "",
			}),
		},
		"gt/valid": {
			Message: &cases.DurationGT{
				Val: &durationpb.Duration{
					Seconds: 1,
				},
			},
			Expected: results.Success(true),
		},
		"gt/valid/empty": {
			Message: &cases.DurationGT{}, Expected: results.Success(true),
		},
		"gt/invalid/equal": {
			Message: &cases.DurationGT{
				Val: &durationpb.Duration{
					Nanos: 1000,
				},
			},
			Expected: results.Violations(&validate.Violation{
				FieldPath:    "val",
				ConstraintId: "duration.gt",
				Message:      "",
			}),
		},
		"gt/invalid": {
			Message: &cases.DurationGT{
				Val: &durationpb.Duration{},
			},
			Expected: results.Violations(&validate.Violation{
				FieldPath:    "val",
				ConstraintId: "duration.gt",
				Message:      "must be greater than 0.000001s",
			}),
		},
		"gte/valid": {
			Message: &cases.DurationGTE{
				Val: &durationpb.Duration{
					Seconds: 3,
				},
			},
			Expected: results.Success(true),
		},
		"gte/valid/empty": {
			Message: &cases.DurationGTE{}, Expected: results.Success(true),
		},
		"gte/valid/equal": {
			Message: &cases.DurationGTE{
				Val: &durationpb.Duration{
					Nanos: 1000000,
				},
			},
			Expected: results.Success(true),
		},
		"gte/invalid": {
			Message: &cases.DurationGTE{
				Val: &durationpb.Duration{Seconds: -1},
			},
			Expected: results.Violations(&validate.Violation{
				FieldPath:    "val",
				ConstraintId: "duration.gte",
				Message:      "must be greater than or equal to 0.001s",
			}),
		},
		"gt_lt/valid": {
			Message: &cases.DurationGTLT{
				Val: &durationpb.Duration{
					Nanos: 1000,
				},
			},
			Expected: results.Success(true),
		},
		"gt_lt/valid/empty": {
			Message: &cases.DurationGTLT{}, Expected: results.Success(true),
		},
		"gt_lt/invalid/above": {
			Message: &cases.DurationGTLT{
				Val: &durationpb.Duration{
					Seconds: 1000,
				},
			},
			Expected: results.Violations(&validate.Violation{
				FieldPath:    "val",
				ConstraintId: "duration.gt_lt",
				Message:      "must be greater than 0s and less than 1s",
			}),
		},
		"gt_lt/invalid/below": {
			Message: &cases.DurationGTLT{
				Val: &durationpb.Duration{
					Nanos: -1000,
				},
			}, Expected: results.Violations(&validate.Violation{
				FieldPath:    "val",
				ConstraintId: "duration.gt_lt",
				Message:      "must be greater than 0s and less than 1s",
			}),
		},
		"gt_lt/invalid/max": {
			Message: &cases.DurationGTLT{
				Val: &durationpb.Duration{
					Seconds: 1,
				},
			},
			Expected: results.Violations(&validate.Violation{
				FieldPath:    "val",
				ConstraintId: "duration.gt_lt",
				Message:      "must be greater than 0s and less than 1s",
			}),
		},
		"gt_lt/invalid/min": {
			Message: &cases.DurationGTLT{
				Val: &durationpb.Duration{},
			},
			Expected: results.Violations(&validate.Violation{
				FieldPath:    "val",
				ConstraintId: "duration.gt_lt",
				Message:      "must be greater than 0s and less than 1s",
			}),
		},
		"gt_lt/exclusive/valid/empty": {
			Message: &cases.DurationExLTGT{}, Expected: results.Success(true),
		},
		"gt_lt/exclusive/valid/above": {
			Message: &cases.DurationExLTGT{
				Val: &durationpb.Duration{
					Seconds: 2,
				},
			},
			Expected: results.Success(true),
		},
		"gt_lt/exclusive/valid/below": {
			Message: &cases.DurationExLTGT{
				Val: &durationpb.Duration{
					Nanos: -1,
				},
			},
			Expected: results.Success(true),
		},
		"gt_lt/exclusive/invalid": {
			Message: &cases.DurationExLTGT{
				Val: &durationpb.Duration{
					Nanos: 1000,
				},
			},
			Expected: results.Violations(&validate.Violation{
				FieldPath:    "val",
				ConstraintId: "duration.gt_lt_exclusive",
				Message:      "must be greater than 1s or less than 0s",
			}),
		},
		"gt_lt/exclusive/invalid/max": {
			Message: &cases.DurationExLTGT{
				Val: &durationpb.Duration{
					Seconds: 1,
				},
			},
			Expected: results.Violations(&validate.Violation{
				FieldPath:    "val",
				ConstraintId: "duration.gt_lt_exclusive",
				Message:      "must be greater than 1s or less than 0s",
			}),
		},
		"gt_lt/exclusive/invalid/min": {
			Message: &cases.DurationExLTGT{
				Val: &durationpb.Duration{},
			},
			Expected: results.Violations(&validate.Violation{
				FieldPath:    "val",
				ConstraintId: "duration.gt_lt_exclusive",
				Message:      "must be greater than 1s or less than 0s",
			}),
		},
		"gte_lte/valid": {
			Message: &cases.DurationGTELTE{
				Val: &durationpb.Duration{Seconds: 60, Nanos: 1},
			},
			Expected: results.Success(true),
		},
		"gte_lte/valid/empty": {
			Message:  &cases.DurationGTELTE{},
			Expected: results.Success(true),
		},
		"gte_lte/valid/max": {
			Message: &cases.DurationGTELTE{
				Val: &durationpb.Duration{
					Seconds: 3600,
				},
			},
			Expected: results.Success(true),
		},
		"gte_lte/valid/min": {
			Message: &cases.DurationGTELTE{
				Val: &durationpb.Duration{
					Seconds: 60,
				},
			},
			Expected: results.Success(true),
		},
		"gte_lte/invalid/above": {
			Message: &cases.DurationGTELTE{
				Val: &durationpb.Duration{Seconds: 3600, Nanos: 1},
			},
			Expected: results.Violations(&validate.Violation{
				FieldPath:    "val",
				ConstraintId: "duration.gte_lte",
				Message:      "must be greater than or equal to 60s and less than or equal to 3600s",
			}),
		},
		"gte_lte/invalid/below": {
			Message: &cases.DurationGTELTE{
				Val: &durationpb.Duration{
					Seconds: 59,
				},
			},
			Expected: results.Violations(&validate.Violation{
				FieldPath:    "val",
				ConstraintId: "duration.gte_lte",
				Message:      "must be greater than or equal to 60s and less than or equal to 3600s",
			}),
		},
		"Ex gte_lte/valid/empty": {
			Message: &cases.DurationExGTELTE{}, Expected: results.Success(true),
		},
		"gte_lte/exclusive/valid/above": {
			Message: &cases.DurationExGTELTE{
				Val: &durationpb.Duration{
					Seconds: 3601,
				},
			},
			Expected: results.Success(true),
		},
		"gte_lte/exclusive/valid/below": {
			Message:  &cases.DurationExGTELTE{Val: &durationpb.Duration{}},
			Expected: results.Success(true),
		},
		"gte_lte/exclusive/valid/max": {
			Message: &cases.DurationExGTELTE{
				Val: &durationpb.Duration{
					Seconds: 3600,
				},
			},
			Expected: results.Success(true),
		},
		"gte_lte/exclusive/valid/min": {
			Message: &cases.DurationExGTELTE{
				Val: &durationpb.Duration{
					Seconds: 60,
				},
			},
			Expected: results.Success(true),
		},
		"gte_lte/exclusive/invalid": {
			Message: &cases.DurationExGTELTE{
				Val: &durationpb.Duration{
					Seconds: 61,
				},
			},
			Expected: results.Violations(&validate.Violation{
				FieldPath:    "val",
				ConstraintId: "duration.gte_lte_exclusive",
				Message:      "must be greater than or equal to 3600s or less than or equal to 60s",
			}),
		},
		"fields_with_other_fields/invalid_other_field": {
			Message: &cases.DurationFieldWithOtherFields{DurationVal: nil, IntVal: 12},
			Expected: results.Violations(&validate.Violation{
				FieldPath:    "int_val",
				ConstraintId: "int32.gt",
				Message:      "must be greater than 16",
			}),
		},
	}
}
