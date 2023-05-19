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

func mapSuite() suites.Suite {
	return suites.Suite{
		"none/valid": {
			Message:  &cases.MapNone{Val: map[uint32]bool{123: true, 456: false}},
			Expected: results.Success(true),
		},
		"min_pairs/valid": {
			Message:  &cases.MapMin{Val: map[int32]float32{1: 2, 3: 4, 5: 6}},
			Expected: results.Success(true),
		},
		"min_pairs/equal/valid": {
			Message:  &cases.MapMin{Val: map[int32]float32{1: 2, 3: 4}},
			Expected: results.Success(true),
		},
		"min_pairs/invalid": {
			Message: &cases.MapMin{Val: map[int32]float32{1: 2}},
			Expected: results.Violations(
				&validate.Violation{
					FieldPath:    "val",
					ConstraintId: "map.min_pairs",
					Message:      "map must be at least 2 entries",
				},
			),
		},
		"max_pairs/valid": {
			Message:  &cases.MapMax{Val: map[int64]float64{1: 2, 3: 4}},
			Expected: results.Success(true),
		},
		"max_pairs/equal/valid": {
			Message:  &cases.MapMax{Val: map[int64]float64{1: 2, 3: 4, 5: 6}},
			Expected: results.Success(true),
		},
		"max_pairs/invalid": {
			Message: &cases.MapMax{Val: map[int64]float64{1: 2, 3: 4, 5: 6, 7: 8}},
			Expected: results.Violations(
				&validate.Violation{
					FieldPath:    "val",
					ConstraintId: "map.max_pairs",
					Message:      "map must be at most 3 entries",
				},
			),
		},
		"min/max/valid": {
			Message:  &cases.MapMinMax{Val: map[string]bool{"a": true, "b": false, "c": true}},
			Expected: results.Success(true),
		},
		"min/max/min/valid": {
			Message:  &cases.MapMinMax{Val: map[string]bool{"a": true, "b": false}},
			Expected: results.Success(true),
		},
		"min/max/max/valid": {
			Message:  &cases.MapMinMax{Val: map[string]bool{"a": true, "b": false, "c": true, "d": false}},
			Expected: results.Success(true),
		},
		"min/max/below/invalid": {
			Message: &cases.MapMinMax{Val: map[string]bool{}},
			Expected: results.Violations(
				&validate.Violation{
					FieldPath:    "val",
					ConstraintId: "map.min_pairs",
					Message:      "map must be at least 2 entries",
				},
			),
		},
		"min/max/above/invalid": {
			Message: &cases.MapMinMax{Val: map[string]bool{"a": true, "b": false, "c": true, "d": false, "e": true}},
			Expected: results.Violations(
				&validate.Violation{
					FieldPath:    "val",
					ConstraintId: "map.max_pairs",
					Message:      "map must be at most 4 entries",
				},
			),
		},
		"exact/valid": {
			Message:  &cases.MapExact{Val: map[uint64]string{1: "a", 2: "b", 3: "c"}},
			Expected: results.Success(true),
		},
		"exact/below/invalid": {
			Message: &cases.MapExact{Val: map[uint64]string{1: "a", 2: "b"}},
			Expected: results.Violations(
				&validate.Violation{
					FieldPath:    "val",
					ConstraintId: "map.min_pairs",
					Message:      "map must be at least 3 entries",
				},
			),
		},
		"exact/above/invalid": {
			Message: &cases.MapExact{Val: map[uint64]string{1: "a", 2: "b", 3: "c", 4: "d"}},
			Expected: results.Violations(
				&validate.Violation{
					FieldPath:    "val",
					ConstraintId: "map.max_pairs",
					Message:      "map must be at most 3 entries",
				},
			),
		},
		"keys/valid": {
			Message:  &cases.MapKeys{Val: map[int64]string{-1: "a", -2: "b"}},
			Expected: results.Success(true),
		},
		"keys/empty/valid": {
			Message:  &cases.MapKeys{Val: map[int64]string{}},
			Expected: results.Success(true),
		},
		"keys/pattern/valid": {
			Message:  &cases.MapKeysPattern{Val: map[string]string{"A": "a"}},
			Expected: results.Success(true),
		},
		"keys/invalid": {
			Message: &cases.MapKeys{Val: map[int64]string{1: "a"}},
			Expected: results.Violations(
				&validate.Violation{
					FieldPath:    "val[1]",
					ConstraintId: "sint64.lt",
					Message:      "must be less than 0",
				},
			),
		},
		"keys/pattern/invalid": {
			Message: &cases.MapKeysPattern{Val: map[string]string{"A": "a", "!@#$%^&*()": "b"}},
			Expected: results.Violations(
				&validate.Violation{
					FieldPath:    "val[\"!@#$%!^(MISSING)&*()\"]",
					ConstraintId: "string.pattern",
					Message:      "must match the pattern `(?i)^[a-z0-9]+$`",
				},
			),
		},
		"values/valid": {
			Message:  &cases.MapValues{Val: map[string]string{"a": "Alpha", "b": "Beta"}},
			Expected: results.Success(true),
		},
		"values/empty/valid": {
			Message:  &cases.MapValues{Val: map[string]string{}},
			Expected: results.Success(true),
		},
		"values/pattern/valid": {
			Message:  &cases.MapValuesPattern{Val: map[string]string{"a": "A"}},
			Expected: results.Success(true),
		},
		"values/invalid": {
			Message: &cases.MapValues{Val: map[string]string{"a": "A", "b": "B"}},
			Expected: results.Violations(
				&validate.Violation{
					FieldPath:    "val[\"a\"]",
					ConstraintId: "string.min_len",
					Message:      "must be at least 3 characters",
				},
				&validate.Violation{
					FieldPath:    "val[\"b\"]",
					ConstraintId: "string.min_len",
					Message:      "must be at least 3 characters",
				},
			),
		},
		"values/pattern/invalid": {
			Message: &cases.MapValuesPattern{Val: map[string]string{"a": "A", "b": "!@#$%^&*()"}},
			Expected: results.Violations(
				&validate.Violation{
					FieldPath:    "val[\"b\"]",
					ConstraintId: "string.pattern",
					Message:      "must match the pattern `(?i)^[a-z0-9]+$`",
				},
			),
		},
		"recursive/valid": {
			Message:  &cases.MapRecursive{Val: map[uint32]*cases.MapRecursive_Msg{1: {Val: "abc"}}},
			Expected: results.Success(true),
		},
		"recursive/invalid": {
			Message: &cases.MapRecursive{Val: map[uint32]*cases.MapRecursive_Msg{1: {}}},
			Expected: results.Violations(
				&validate.Violation{
					FieldPath:    "val[0x1].val",
					ConstraintId: "string.min_len",
					Message:      "must be at least 3 characters",
				},
			),
		},
		"exact/ignore_empty/valid": {
			Message:  &cases.MapExactIgnore{Val: nil},
			Expected: results.Success(true),
		},
		"multiple/valid": {
			Message:  &cases.MultipleMaps{First: map[uint32]string{1: "a", 2: "b"}, Second: map[int32]bool{-1: true, -2: false}},
			Expected: results.Success(true),
		},
	}
}
