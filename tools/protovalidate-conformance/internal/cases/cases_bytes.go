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

func bytesSuite() suites.Suite {
	return suites.Suite{
		"none/valid": {
			Message:  &cases.BytesNone{Val: []byte("quux")},
			Expected: results.Success(true),
		},
		"const/valid": {
			Message:  &cases.BytesConst{Val: []byte("foo")},
			Expected: results.Success(true),
		},
		"const/invalid": {
			Message: &cases.BytesConst{Val: []byte("bar")},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "bytes.const"}),
		},
		"in/valid": {
			Message:  &cases.BytesIn{Val: []byte("bar")},
			Expected: results.Success(true),
		},
		"in/invalid": {
			Message: &cases.BytesIn{Val: []byte("quux")},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "bytes.in"}),
		},
		"not_in/valid": {
			Message:  &cases.BytesNotIn{Val: []byte("quux")},
			Expected: results.Success(true),
		},
		"not_in/invalid": {
			Message: &cases.BytesNotIn{Val: []byte("fizz")},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "bytes.not_in"}),
		},
		"len/valid": {
			Message:  &cases.BytesLen{Val: []byte("baz")},
			Expected: results.Success(true),
		},
		"len/invalid/less": {
			Message: &cases.BytesLen{Val: []byte("go")},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "bytes.len"}),
		},
		"len/invalid/greater": {
			Message: &cases.BytesLen{Val: []byte("fizz")},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "bytes.len"}),
		},
		"min_len/valid/equal": {
			Message:  &cases.BytesMinLen{Val: []byte("baz")},
			Expected: results.Success(true),
		},
		"min_len/valid/greater": {
			Message:  &cases.BytesMinLen{Val: []byte("buzz")},
			Expected: results.Success(true),
		},
		"min_len/invalid": {
			Message: &cases.BytesMinLen{Val: []byte("go")},
			Expected: results.Violations(&validate.Violation{
				FieldPath:    "val",
				ConstraintId: "bytes.min_len",
				Message:      "value must be at least 3 bytes long",
			}),
		},
		"max_len/valid": {
			Message:  &cases.BytesMaxLen{Val: []byte("foo")},
			Expected: results.Success(true),
		},
		"max_len/valid/max": {
			Message:  &cases.BytesMaxLen{Val: []byte("proto")},
			Expected: results.Success(true),
		},
		"max_len/invalid": {
			Message: &cases.BytesMaxLen{Val: []byte("1234567890")},
			Expected: results.Violations(&validate.Violation{
				FieldPath:    "val",
				ConstraintId: "bytes.max_len",
				Message:      "value must be at most 5 bytes long",
			}),
		},
		"min/max_len/valid-01": {
			Message:  &cases.BytesMinMaxLen{Val: []byte("quux")},
			Expected: results.Success(true),
		},
		"min/max_len/valid-02": {
			Message:  &cases.BytesMinMaxLen{Val: []byte("foo")},
			Expected: results.Success(true),
		},
		"min/max_len/valid-03": {
			Message:  &cases.BytesMinMaxLen{Val: []byte("proto")},
			Expected: results.Success(true),
		},
		"min/max_len/invalid/below": {
			Message: &cases.BytesMinMaxLen{Val: []byte("go")},
			Expected: results.Violations(&validate.Violation{
				FieldPath:    "val",
				ConstraintId: "bytes.min_len",
				Message:      "value must be at least 3 bytes long",
			}),
		},
		"min/max_len/invalid/above": {
			Message: &cases.BytesMinMaxLen{Val: []byte("validate")},
			Expected: results.Violations(&validate.Violation{
				FieldPath:    "val",
				ConstraintId: "bytes.max_len",
				Message:      "value must be at most 5 bytes",
			}),
		},
		"equal/min_len/max_len/valid": {
			Message:  &cases.BytesEqualMinMaxLen{Val: []byte("proto")},
			Expected: results.Success(true),
		},
		"equal/min_len/max_len/invalid": {
			Message: &cases.BytesEqualMinMaxLen{Val: []byte("validate")},
			Expected: results.Violations(&validate.Violation{
				FieldPath:    "val",
				ConstraintId: "bytes.max_len",
				Message:      "value must be at most 5 bytes",
			}),
		},
		"pattern/valid": {
			Message:  &cases.BytesPattern{Val: []byte("Foo123")},
			Expected: results.Success(true),
		},
		"pattern/invalid": {
			Message: &cases.BytesPattern{Val: []byte("你好你好")},
			Expected: results.Violations(&validate.Violation{
				FieldPath:    "val",
				ConstraintId: "bytes.pattern",
				Message:      "value must match the pattern \"你好你好\"",
			}),
		},
		"pattern/invalid/empty": {
			Message: &cases.BytesPattern{Val: []byte("")},
			Expected: results.Violations(&validate.Violation{
				FieldPath:    "val",
				ConstraintId: "bytes.pattern",
				Message:      "value must match the pattern \"\"",
			}),
		},
		"pattern/invalid/not_utf8": {
			Message:  &cases.BytesPattern{Val: []byte{0x99}},
			Expected: results.RuntimeError("value must be valid UTF-8 to apply regexp"),
		},
		"prefix/valid": {
			Message:  &cases.BytesPrefix{Val: []byte{0x99, 0x9f, 0x08}},
			Expected: results.Success(true),
		},
		"prefix/valid/only": {
			Message:  &cases.BytesPrefix{Val: []byte{0x99}},
			Expected: results.Success(true),
		},
		"prefix/invalid": {
			Message: &cases.BytesPrefix{Val: []byte("bar")},
			Expected: results.Violations(&validate.Violation{
				FieldPath:    "val",
				ConstraintId: "bytes.prefix",
				Message:      "value must have the prefix \"foo\"",
			}),
		},
		"contains/valid": {
			Message:  &cases.BytesContains{Val: []byte("candy bars")},
			Expected: results.Success(true),
		},
		"contains/valid/only": {
			Message:  &cases.BytesContains{Val: []byte("bar")},
			Expected: results.Success(true),
		},
		"contains/invalid": {
			Message: &cases.BytesContains{Val: []byte("candy bazs")},
			Expected: results.Violations(&validate.Violation{
				FieldPath:    "val",
				ConstraintId: "bytes.contains",
				Message:      "value must contain \"bar\"",
			}),
		},
		"suffix/valid": {
			Message:  &cases.BytesSuffix{Val: []byte{0x62, 0x75, 0x7A, 0x7A}},
			Expected: results.Success(true),
		},
		"suffix/valid/only": {
			Message:  &cases.BytesSuffix{Val: []byte("\x62\x75\x7A\x7A")},
			Expected: results.Success(true),
		},
		"suffix/invalid": {
			Message: &cases.BytesSuffix{Val: []byte("foobar")},
			Expected: results.Violations(&validate.Violation{
				FieldPath:    "val",
				ConstraintId: "bytes.suffix",
				Message:      "value must have the suffix \"baz\"",
			}),
		},
		"suffix/case_sensitive/invalid": {
			Message: &cases.BytesSuffix{Val: []byte("FooBaz")},
			Expected: results.Violations(&validate.Violation{
				FieldPath:    "val",
				ConstraintId: "bytes.suffix",
				Message:      "value must have the suffix \"baz\"",
			}),
		},
		"IP/valid/v4": {
			Message:  &cases.BytesIP{Val: []byte{0xC0, 0xA8, 0x00, 0x01}},
			Expected: results.Success(true),
		},
		"IP/valid/v6": {
			Message:  &cases.BytesIP{Val: []byte("\x20\x01\x0D\xB8\x85\xA3\x00\x00\x00\x00\x8A\x2E\x03\x70\x73\x34")},
			Expected: results.Success(true),
		},
		"IP/invalid": {
			Message: &cases.BytesIP{Val: []byte("foobar")},
			Expected: results.Violations(&validate.Violation{
				FieldPath:    "val",
				ConstraintId: "bytes.ip",
				Message:      "value must be a valid IP address",
			}),
		},
		"IPv4/valid": {
			Message:  &cases.BytesIPv4{Val: []byte{0xC0, 0xA8, 0x00, 0x01}},
			Expected: results.Success(true),
		},
		"IPv4/invalid": {
			Message: &cases.BytesIPv4{Val: []byte("foobar")},
			Expected: results.Violations(&validate.Violation{
				FieldPath:    "val",
				ConstraintId: "bytes.ipv4",
				Message:      "value must be a valid IPv4 address",
			}),
		},
		"IPv4/invalid/v6": {
			Message: &cases.BytesIPv4{Val: []byte("\x20\x01\x0D\xB8\x85\xA3\x00\x00\x00\x00\x8A\x2E\x03\x70\x73\x34")},
			Expected: results.Violations(&validate.Violation{
				FieldPath:    "val",
				ConstraintId: "bytes.ipv4",
				Message:      "value must be a valid IPv4 address",
			}),
		},
		"IPv6/valid": {
			Message:  &cases.BytesIPv6{Val: []byte("\x20\x01\x0D\xB8\x85\xA3\x00\x00\x00\x00\x8A\x2E\x03\x70\x73\x34")},
			Expected: results.Success(true),
		},
		"IPv6/invalid": {
			Message: &cases.BytesIPv6{Val: []byte("fooar")},
			Expected: results.Violations(&validate.Violation{
				FieldPath:    "val",
				ConstraintId: "bytes.ipv6",
				Message:      "value must be a valid IPv6 address",
			}),
		},
		"IPv6/invalid/v4": {
			Message: &cases.BytesIPv6{Val: []byte{0xC0, 0xA8, 0x00, 0x01}},
			Expected: results.Violations(&validate.Violation{
				FieldPath:    "val",
				ConstraintId: "bytes.ipv6",
			}),
		},
		"IPv6/valid/ignore_empty": {
			Message:  &cases.BytesIPv6Ignore{Val: nil},
			Expected: results.Success(true),
		},
	}
}
