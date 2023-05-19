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

func sint32Suite() suites.Suite {
	return suites.Suite{
		"none/valid": {
			Message:  &cases.SInt32None{Val: 123},
			Expected: results.Success(true),
		},
		"const/valid": {
			Message:  &cases.SInt32Const{Val: 1},
			Expected: results.Success(true),
		},
		"const/invalid": {
			Message: &cases.SInt32Const{Val: 2},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "sint32.const"}),
		},
		"in/valid": {
			Message:  &cases.SInt32In{Val: 3},
			Expected: results.Success(true),
		},
		"in/invalid": {
			Message: &cases.SInt32In{Val: 5},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "sint32.in"}),
		},
		"not_in/valid": {
			Message:  &cases.SInt32NotIn{Val: 1},
			Expected: results.Success(true),
		},
		"not_in/invalid": {
			Message: &cases.SInt32NotIn{Val: 0},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "sint32.not_in"}),
		},
		"lt/valid/less": {
			Message:  &cases.SInt32LT{Val: -1},
			Expected: results.Success(true),
		},
		"lt/invalid/equal": {
			Message: &cases.SInt32LT{Val: 0},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "sint32.lt"}),
		},
		"lt/invalid/greater": {
			Message: &cases.SInt32LT{Val: 1},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "sint32.lt"}),
		},
		"lte/valid/less": {
			Message:  &cases.SInt32LTE{Val: 63},
			Expected: results.Success(true),
		},
		"lte/valid/equal": {
			Message:  &cases.SInt32LTE{Val: 64},
			Expected: results.Success(true),
		},
		"lte/invalid/greater": {
			Message: &cases.SInt32LTE{Val: 65},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "sint32.lte"}),
		},
		"gt/valid/greater": {
			Message:  &cases.SInt32GT{Val: 17},
			Expected: results.Success(true),
		},
		"gt/invalid/equal": {
			Message: &cases.SInt32GT{Val: 16},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "sint32.gt"}),
		},
		"gt/invalid/less": {
			Message: &cases.SInt32GT{Val: 15},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "sint32.gt"}),
		},
		"gte/valid/greater": {
			Message:  &cases.SInt32GTE{Val: 9},
			Expected: results.Success(true),
		},
		"gte/valid/equal": {
			Message:  &cases.SInt32GTE{Val: 8},
			Expected: results.Success(true),
		},
		"gte/invalid/less": {
			Message: &cases.SInt32GTE{Val: 7},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "sint32.gte"}),
		},
		"gt_lt/inclusive/valid/within": {
			Message:  &cases.SInt32GTLT{Val: 5},
			Expected: results.Success(true),
		},
		"gt_lt/inclusive/invalid/above": {
			Message: &cases.SInt32GTLT{Val: 11},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "sint32.gt_lt"}),
		},
		"gt_lt/inclusive/invalid/below": {
			Message: &cases.SInt32GTLT{Val: -1},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "sint32.gt_lt"}),
		},
		"gt_lt/inclusive/invalid/max": {
			Message: &cases.SInt32GTLT{Val: 10},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "sint32.gt_lt"}),
		},
		"gt_lt/inclusive/invalid/min": {
			Message: &cases.SInt32GTLT{Val: 0},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "sint32.gt_lt"}),
		},
		"gt_lt/exclusive/valid/above": {
			Message:  &cases.SInt32ExLTGT{Val: 11},
			Expected: results.Success(true),
		},
		"gt_lt/exclusive/valid/below": {
			Message:  &cases.SInt32ExLTGT{Val: -1},
			Expected: results.Success(true),
		},
		"gt_lt/exclusive/invalid/within": {
			Message: &cases.SInt32ExLTGT{Val: 5},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "sint32.gt_lt_exclusive"}),
		},
		"gt_lt/exclusive/invalid/max": {
			Message: &cases.SInt32ExLTGT{Val: 10},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "sint32.gt_lt_exclusive"}),
		},
		"gt_lt/exclusive/invalid/min": {
			Message: &cases.SInt32ExLTGT{Val: 0},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "sint32.gt_lt_exclusive"}),
		},
		"gte_lte/inclusive/valid/within": {
			Message:  &cases.SInt32GTELTE{Val: 200},
			Expected: results.Success(true),
		},
		"gte_lte/inclusive/valid/max": {
			Message:  &cases.SInt32GTELTE{Val: 256},
			Expected: results.Success(true),
		},
		"gte_lte/inclusive/valid/min": {
			Message:  &cases.SInt32GTELTE{Val: 128},
			Expected: results.Success(true),
		},
		"gte_lte/inclusive/invalid/above": {
			Message: &cases.SInt32GTELTE{Val: 300},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "sint32.gte_lte"}),
		},
		"gte_lte/inclusive/invalid/below": {
			Message: &cases.SInt32GTELTE{Val: 100},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "sint32.gte_lte"}),
		},
		"gte_lte/exclusive/valid/above": {
			Message:  &cases.SInt32ExGTELTE{Val: 300},
			Expected: results.Success(true),
		},
		"gte_lte/exclusive/valid/below": {
			Message:  &cases.SInt32ExGTELTE{Val: 100},
			Expected: results.Success(true),
		},
		"gte_lte/exclusive/valid/max": {
			Message:  &cases.SInt32ExGTELTE{Val: 256},
			Expected: results.Success(true),
		},
		"gte_lte/exclusive/valid/min": {
			Message:  &cases.SInt32ExGTELTE{Val: 128},
			Expected: results.Success(true),
		},
		"gte_lte/exclusive/invalid/within": {
			Message: &cases.SInt32ExGTELTE{Val: 200},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "sint32.gte_lte_exclusive"}),
		},
		"ignore_empty/valid/empty": {
			Message:  &cases.SInt32Ignore{Val: 0},
			Expected: results.Success(true),
		},
		"ignore_empty/valid/within": {
			Message:  &cases.SInt32Ignore{Val: 200},
			Expected: results.Success(true),
		},
		"ignore_empty/invalid/above": {
			Message: &cases.SInt32Ignore{Val: 300},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "sint32.gte_lte"}),
		},
		"compilation/wrong_type": {
			Message:  &cases.SInt32IncorrectType{Val: 123},
			Expected: results.CompilationError("double constraints on float field"),
		},
	}
}
