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

func doubleSuite() suites.Suite {
	return suites.Suite{
		"none/valid": {
			Message:  &cases.DoubleNone{Val: -1.23456},
			Expected: results.Success(true),
		},
		"const/valid": {
			Message:  &cases.DoubleConst{Val: 1.23},
			Expected: results.Success(true),
		},
		"const/invalid": {
			Message: &cases.DoubleConst{Val: 4.56},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "double.const"}),
		},
		"in/valid": {
			Message:  &cases.DoubleIn{Val: 7.89},
			Expected: results.Success(true),
		},
		"in/invalid": {
			Message: &cases.DoubleIn{Val: 10.11},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "double.in"}),
		},
		"not_in/valid": {
			Message:  &cases.DoubleNotIn{Val: 1},
			Expected: results.Success(true),
		},
		"not_in/invalid": {
			Message: &cases.DoubleNotIn{Val: 0},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "double.not_in"}),
		},
		"lt/valid/less": {
			Message:  &cases.DoubleLT{Val: -1},
			Expected: results.Success(true),
		},
		"lt/invalid/equal": {
			Message: &cases.DoubleLT{Val: 0},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "double.lt"}),
		},
		"lt/invalid/greater": {
			Message: &cases.DoubleLT{Val: 1},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "double.lt"}),
		},
		"lte/valid/less": {
			Message:  &cases.DoubleLTE{Val: 63},
			Expected: results.Success(true),
		},
		"lte/valid/equal": {
			Message:  &cases.DoubleLTE{Val: 64},
			Expected: results.Success(true),
		},
		"lte/invalid/greater": {
			Message: &cases.DoubleLTE{Val: 65},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "double.lte"}),
		},
		"gt/valid/greater": {
			Message:  &cases.DoubleGT{Val: 17},
			Expected: results.Success(true),
		},
		"gt/invalid/equal": {
			Message: &cases.DoubleGT{Val: 16},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "double.gt"}),
		},
		"gt/invalid/less": {
			Message: &cases.DoubleGT{Val: 15},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "double.gt"}),
		},
		"gte/valid/greater": {
			Message:  &cases.DoubleGTE{Val: 9},
			Expected: results.Success(true),
		},
		"gte/valid/equal": {
			Message:  &cases.DoubleGTE{Val: 8},
			Expected: results.Success(true),
		},
		"gte/invalid/less": {
			Message: &cases.DoubleGTE{Val: 7},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "double.gte"}),
		},
		"gt_lt/inclusive/valid/within": {
			Message:  &cases.DoubleGTLT{Val: 5},
			Expected: results.Success(true),
		},
		"gt_lt/inclusive/invalid/above": {
			Message: &cases.DoubleGTLT{Val: 11},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "double.gt_lt"}),
		},
		"gt_lt/inclusive/invalid/below": {
			Message: &cases.DoubleGTLT{Val: -1},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "double.gt_lt"}),
		},
		"gt_lt/inclusive/invalid/max": {
			Message: &cases.DoubleGTLT{Val: 10},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "double.gt_lt"}),
		},
		"gt_lt/inclusive/invalid/min": {
			Message: &cases.DoubleGTLT{Val: 0},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "double.gt_lt"}),
		},
		"gt_lt/exclusive/valid/above": {
			Message:  &cases.DoubleExLTGT{Val: 11},
			Expected: results.Success(true),
		},
		"gt_lt/exclusive/valid/below": {
			Message:  &cases.DoubleExLTGT{Val: -1},
			Expected: results.Success(true),
		},
		"gt_lt/exclusive/invalid/within": {
			Message: &cases.DoubleExLTGT{Val: 5},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "double.gt_lt_exclusive"}),
		},
		"gt_lt/exclusive/invalid/max": {
			Message: &cases.DoubleExLTGT{Val: 10},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "double.gt_lt_exclusive"}),
		},
		"gt_lt/exclusive/invalid/min": {
			Message: &cases.DoubleExLTGT{Val: 0},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "double.gt_lt_exclusive"}),
		},
		"gte_lte/inclusive/valid/within": {
			Message:  &cases.DoubleGTELTE{Val: 200},
			Expected: results.Success(true),
		},
		"gte_lte/inclusive/valid/max": {
			Message:  &cases.DoubleGTELTE{Val: 256},
			Expected: results.Success(true),
		},
		"gte_lte/inclusive/valid/min": {
			Message:  &cases.DoubleGTELTE{Val: 128},
			Expected: results.Success(true),
		},
		"gte_lte/inclusive/invalid/above": {
			Message: &cases.DoubleGTELTE{Val: 300},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "double.gte_lte"}),
		},
		"gte_lte/inclusive/invalid/below": {
			Message: &cases.DoubleGTELTE{Val: 100},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "double.gte_lte"}),
		},
		"gte_lte/exclusive/valid/above": {
			Message:  &cases.DoubleExGTELTE{Val: 300},
			Expected: results.Success(true),
		},
		"gte_lte/exclusive/valid/below": {
			Message:  &cases.DoubleExGTELTE{Val: 100},
			Expected: results.Success(true),
		},
		"gte_lte/exclusive/valid/max": {
			Message:  &cases.DoubleExGTELTE{Val: 256},
			Expected: results.Success(true),
		},
		"gte_lte/exclusive/valid/min": {
			Message:  &cases.DoubleExGTELTE{Val: 128},
			Expected: results.Success(true),
		},
		"gte_lte/exclusive/invalid/within": {
			Message: &cases.DoubleExGTELTE{Val: 200},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "double.gte_lte_exclusive"}),
		},
		"ignore_empty/valid/empty": {
			Message:  &cases.DoubleIgnore{Val: 0},
			Expected: results.Success(true),
		},
		"ignore_empty/valid/within": {
			Message:  &cases.DoubleIgnore{Val: 200},
			Expected: results.Success(true),
		},
		"ignore_empty/invalid/above": {
			Message: &cases.DoubleIgnore{Val: 300},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "double.gte_lte"}),
		},
		"compilation/wrong_type": {
			Message:  &cases.DoubleIncorrectType{Val: 123},
			Expected: results.CompilationError("double constraints on float field"),
		},
	}
}
