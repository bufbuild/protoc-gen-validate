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

func fixed32Suite() suites.Suite {
	return suites.Suite{
		"none/valid": {
			Message:  &cases.Fixed32None{Val: 123},
			Expected: results.Success(true),
		},
		"const/valid": {
			Message:  &cases.Fixed32Const{Val: 1},
			Expected: results.Success(true),
		},
		"const/invalid": {
			Message: &cases.Fixed32Const{Val: 2},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "fixed32.const"}),
		},
		"in/valid": {
			Message:  &cases.Fixed32In{Val: 3},
			Expected: results.Success(true),
		},
		"in/invalid": {
			Message: &cases.Fixed32In{Val: 5},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "fixed32.in"}),
		},
		"not_in/valid": {
			Message:  &cases.Fixed32NotIn{Val: 1},
			Expected: results.Success(true),
		},
		"not_in/invalid": {
			Message: &cases.Fixed32NotIn{Val: 0},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "fixed32.not_in"}),
		},
		"lt/valid/less": {
			Message:  &cases.Fixed32LT{Val: 4},
			Expected: results.Success(true),
		},
		"lt/invalid/equal": {
			Message: &cases.Fixed32LT{Val: 5},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "fixed32.lt"}),
		},
		"lt/invalid/greater": {
			Message: &cases.Fixed32LT{Val: 6},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "fixed32.lt"}),
		},
		"lte/valid/less": {
			Message:  &cases.Fixed32LTE{Val: 63},
			Expected: results.Success(true),
		},
		"lte/valid/equal": {
			Message:  &cases.Fixed32LTE{Val: 64},
			Expected: results.Success(true),
		},
		"lte/invalid/greater": {
			Message: &cases.Fixed32LTE{Val: 65},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "fixed32.lte"}),
		},
		"gt/valid/greater": {
			Message:  &cases.Fixed32GT{Val: 17},
			Expected: results.Success(true),
		},
		"gt/invalid/equal": {
			Message: &cases.Fixed32GT{Val: 16},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "fixed32.gt"}),
		},
		"gt/invalid/less": {
			Message: &cases.Fixed32GT{Val: 15},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "fixed32.gt"}),
		},
		"gte/valid/greater": {
			Message:  &cases.Fixed32GTE{Val: 9},
			Expected: results.Success(true),
		},
		"gte/valid/equal": {
			Message:  &cases.Fixed32GTE{Val: 8},
			Expected: results.Success(true),
		},
		"gte/invalid/less": {
			Message: &cases.Fixed32GTE{Val: 7},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "fixed32.gte"}),
		},
		"gt_lt/inclusive/valid/within": {
			Message:  &cases.Fixed32GTLT{Val: 6},
			Expected: results.Success(true),
		},
		"gt_lt/inclusive/invalid/above": {
			Message: &cases.Fixed32GTLT{Val: 11},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "fixed32.gt_lt"}),
		},
		"gt_lt/inclusive/invalid/below": {
			Message: &cases.Fixed32GTLT{Val: 4},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "fixed32.gt_lt"}),
		},
		"gt_lt/inclusive/invalid/max": {
			Message: &cases.Fixed32GTLT{Val: 10},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "fixed32.gt_lt"}),
		},
		"gt_lt/inclusive/invalid/min": {
			Message: &cases.Fixed32GTLT{Val: 5},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "fixed32.gt_lt"}),
		},
		"gt_lt/exclusive/valid/above": {
			Message:  &cases.Fixed32ExLTGT{Val: 11},
			Expected: results.Success(true),
		},
		"gt_lt/exclusive/valid/below": {
			Message:  &cases.Fixed32ExLTGT{Val: 4},
			Expected: results.Success(true),
		},
		"gt_lt/exclusive/invalid/within": {
			Message: &cases.Fixed32ExLTGT{Val: 6},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "fixed32.gt_lt_exclusive"}),
		},
		"gt_lt/exclusive/invalid/max": {
			Message: &cases.Fixed32ExLTGT{Val: 10},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "fixed32.gt_lt_exclusive"}),
		},
		"gt_lt/exclusive/invalid/min": {
			Message: &cases.Fixed32ExLTGT{Val: 5},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "fixed32.gt_lt_exclusive"}),
		},
		"gte_lte/inclusive/valid/within": {
			Message:  &cases.Fixed32GTELTE{Val: 200},
			Expected: results.Success(true),
		},
		"gte_lte/inclusive/valid/max": {
			Message:  &cases.Fixed32GTELTE{Val: 256},
			Expected: results.Success(true),
		},
		"gte_lte/inclusive/valid/min": {
			Message:  &cases.Fixed32GTELTE{Val: 128},
			Expected: results.Success(true),
		},
		"gte_lte/inclusive/invalid/above": {
			Message: &cases.Fixed32GTELTE{Val: 300},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "fixed32.gte_lte"}),
		},
		"gte_lte/inclusive/invalid/below": {
			Message: &cases.Fixed32GTELTE{Val: 100},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "fixed32.gte_lte"}),
		},
		"gte_lte/exclusive/valid/above": {
			Message:  &cases.Fixed32ExGTELTE{Val: 300},
			Expected: results.Success(true),
		},
		"gte_lte/exclusive/valid/below": {
			Message:  &cases.Fixed32ExGTELTE{Val: 100},
			Expected: results.Success(true),
		},
		"gte_lte/exclusive/valid/max": {
			Message:  &cases.Fixed32ExGTELTE{Val: 256},
			Expected: results.Success(true),
		},
		"gte_lte/exclusive/valid/min": {
			Message:  &cases.Fixed32ExGTELTE{Val: 128},
			Expected: results.Success(true),
		},
		"gte_lte/exclusive/invalid/within": {
			Message: &cases.Fixed32ExGTELTE{Val: 200},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "fixed32.gte_lte_exclusive"}),
		},
		"ignore_empty/valid/empty": {
			Message:  &cases.Fixed32Ignore{Val: 0},
			Expected: results.Success(true),
		},
		"ignore_empty/valid/within": {
			Message:  &cases.Fixed32Ignore{Val: 200},
			Expected: results.Success(true),
		},
		"ignore_empty/invalid/above": {
			Message: &cases.Fixed32Ignore{Val: 300},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "fixed32.gte_lte"}),
		},
		"compilation/wrong_type": {
			Message:  &cases.Fixed32IncorrectType{Val: 123},
			Expected: results.CompilationError("double constraints on float field"),
		},
	}
}
