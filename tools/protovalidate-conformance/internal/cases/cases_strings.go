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

func stringSuite() suites.Suite {
	return suites.Suite{
		"none/valid": {
			Message:  &cases.StringNone{Val: "foobar"},
			Expected: results.Success(true),
		},
		"const/valid": {
			Message:  &cases.StringConst{Val: "foo"},
			Expected: results.Success(true),
		},
		"const/invalid": {
			Message: &cases.StringConst{Val: "bar"},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "string.const"}),
		},
		"in/valid": {
			Message:  &cases.StringIn{Val: "baz"},
			Expected: results.Success(true),
		},
		"in/invalid": {
			Message: &cases.StringIn{Val: "foo"},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "string.in"}),
		},
		"not_in/valid": {
			Message:  &cases.StringNotIn{Val: "bar"},
			Expected: results.Success(true),
		},
		"not_in/invalid": {
			Message: &cases.StringNotIn{Val: "fizz"},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "string.not_in"}),
		},
		"len/valid/ascii": {
			Message:  &cases.StringLen{Val: "foo"},
			Expected: results.Success(true),
		},
		"len/valid/multibyte": {
			Message:  &cases.StringLen{Val: "ʙǻɹ"},
			Expected: results.Success(true),
		},
		"len/valid/emoji/simple": {
			Message:  &cases.StringLen{Val: "😅😄👾"},
			Expected: results.Success(true),
		},
		"len/invalid": {
			Message: &cases.StringLen{Val: "fizz"},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "string.len"}),
		},
		"len/invalid/emoji/composite": {
			Message: &cases.StringLen{Val: "👩🏽‍💻🧑🏾‍💻👨🏼‍💻"},
			Expected: results.Violations(
				&validate.Violation{
					FieldPath:    "val",
					ConstraintId: "string.len",
					Message:      "composite emoji are treated as separate characters",
				}),
		},
		"min_len/valid/equal": {
			Message:  &cases.StringMinLen{Val: "foo"},
			Expected: results.Success(true),
		},
		"min_len/valid/greater": {
			Message:  &cases.StringMinLen{Val: "foobar"},
			Expected: results.Success(true),
		},
		"min_len/invalid/less": {
			Message: &cases.StringMinLen{Val: "pb"},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "string.min_len"}),
		},
		"max_len/valid/equal": {
			Message:  &cases.StringMaxLen{Val: "proto"},
			Expected: results.Success(true),
		},
		"max_len/valid/less": {
			Message:  &cases.StringMaxLen{Val: "buf"},
			Expected: results.Success(true),
		},
		"max_len/invalid/greater": {
			Message: &cases.StringMaxLen{Val: "validate"},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "string.max_len"}),
		},
		"min_max_len/valid/within": {
			Message:  &cases.StringMinMaxLen{Val: "quux"},
			Expected: results.Success(true),
		},
		"min_max_len/valid/min": {
			Message:  &cases.StringMinMaxLen{Val: "foo"},
			Expected: results.Success(true),
		},
		"min_max_len/valid/max": {
			Message:  &cases.StringMinMaxLen{Val: "proto"},
			Expected: results.Success(true),
		},
		"min_max_len/invalid/less": {
			Message: &cases.StringMinMaxLen{Val: "pb"},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "string.min_len"}),
		},
		"min_max_len/invalid/greater": {
			Message: &cases.StringMinMaxLen{Val: "validate"},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "string.max_len"}),
		},
		"min_max_len/equal/valid": {
			Message:  &cases.StringEqualMinMaxLen{Val: "proto"},
			Expected: results.Success(true),
		},
		"min_max_len/equal/invalid": {
			Message: &cases.StringEqualMinMaxLen{Val: "validate"},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "string.max_len"}),
		},
		"len_bytes/valid/ascii": {
			Message:  &cases.StringLenBytes{Val: "fizz"},
			Expected: results.Success(true),
		},
		"len_bytes/valid/multibyte": {
			Message:  &cases.StringLenBytes{Val: "Ƥƥ"},
			Expected: results.Success(true),
		},
		"len_bytes/valid/emoji": {
			Message:  &cases.StringLenBytes{Val: "😄"},
			Expected: results.Success(true),
		},
		"len_bytes/invalid": {
			Message: &cases.StringLenBytes{Val: "foo"},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "string.len_bytes"}),
		},
		"min_bytes/valid/equal": {
			Message:  &cases.StringMinBytes{Val: "fizz"},
			Expected: results.Success(true),
		},
		"min_bytes/valid/greater": {
			Message:  &cases.StringMinBytes{Val: "proto"},
			Expected: results.Success(true),
		},
		"min_bytes/invalid/less": {
			Message: &cases.StringMinBytes{Val: "foo"},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "string.min_bytes"}),
		},
		"max_bytes/valid/equal": {
			Message:  &cases.StringMaxBytes{Val: "validate"},
			Expected: results.Success(true),
		},
		"max_bytes/valid/less": {
			Message:  &cases.StringMaxBytes{Val: "proto"},
			Expected: results.Success(true),
		},
		"max_bytes/invalid/greater": {
			Message: &cases.StringMaxBytes{Val: "validation"},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "string.max_bytes"}),
		},
		"min_max_bytes/valid/within": {
			Message:  &cases.StringMinMaxBytes{Val: "quux"},
			Expected: results.Success(true),
		},
		"min_max_bytes/valid/min": {
			Message:  &cases.StringMinMaxBytes{Val: "fizz"},
			Expected: results.Success(true),
		},
		"min_max_bytes/valid/max": {
			Message:  &cases.StringMinMaxBytes{Val: "validate"},
			Expected: results.Success(true),
		},
		"min_max_bytes/invalid/less": {
			Message: &cases.StringMinMaxBytes{Val: "pb"},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "string.min_bytes"}),
		},
		"min_max_bytes/invalid/greater": {
			Message: &cases.StringMinMaxBytes{Val: "validation"},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "string.max_bytes"}),
		},
		"min_max_bytes/equal/valid": {
			Message:  &cases.StringEqualMinMaxBytes{Val: "fizz"},
			Expected: results.Success(true),
		},
		"min_max_bytes/equal/invalid": {
			Message: &cases.StringEqualMinMaxBytes{Val: "foo"},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "string.min_bytes"}),
		},
		"pattern/valid": {
			Message:  &cases.StringPattern{Val: "Foo123"},
			Expected: results.Success(true),
		},
		"pattern/invalid": {
			Message: &cases.StringPattern{Val: "!#@$#$%"},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "string.pattern"}),
		},
		"pattern/escapes/valid": {
			Message:  &cases.StringPatternEscapes{Val: "* \\ x"},
			Expected: results.Success(true),
		},
		"pattern/escapes/invalid": {
			Message: &cases.StringPatternEscapes{Val: "invalid"},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "string.pattern"}),
		},
		"prefix/valid/exact": {
			Message:  &cases.StringPrefix{Val: "foo"},
			Expected: results.Success(true),
		},
		"prefix/valid/starts_with": {
			Message:  &cases.StringPrefix{Val: "foobar"},
			Expected: results.Success(true),
		},
		"prefix/invalid": {
			Message: &cases.StringPrefix{Val: "fizz"},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "string.prefix"}),
		},
		"contains/valid/exact": {
			Message:  &cases.StringContains{Val: "bar"},
			Expected: results.Success(true),
		},
		"contains/valid/prefix": {
			Message:  &cases.StringContains{Val: "barfoo"},
			Expected: results.Success(true),
		},
		"contains/valid/suffix": {
			Message:  &cases.StringContains{Val: "foobar"},
			Expected: results.Success(true),
		},
		"contains/valid/within": {
			Message:  &cases.StringContains{Val: "foobarbaz"},
			Expected: results.Success(true),
		},
		"contains/invalid": {
			Message: &cases.StringContains{Val: "fizzbuzz"},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "string.contains"}),
		},
		"not_contains/valid": {
			Message:  &cases.StringNotContains{Val: "fizzbuzz"},
			Expected: results.Success(true),
		},
		"not_contains/invalid": {
			Message: &cases.StringNotContains{Val: "foobarbaz"},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "string.not_contains"}),
		},
		"suffix/valid/exact": {
			Message:  &cases.StringSuffix{Val: "baz"},
			Expected: results.Success(true),
		},
		"suffix/valid/ends_with": {
			Message:  &cases.StringSuffix{Val: "foobarbaz"},
			Expected: results.Success(true),
		},
		"suffix/invalid": {
			Message: &cases.StringSuffix{Val: "bazbarfoo"},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "string.suffix"}),
		},
		"email/valid/simple": {
			Message:  &cases.StringEmail{Val: "foo@bar.com"},
			Expected: results.Success(true),
		},
		"email/invalid/malformed": {
			Message: &cases.StringEmail{Val: "foobar"},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "string.email"}),
		},
		"email/invalid/local_segment_long": {
			Message: &cases.StringEmail{Val: "x0123456789012345678901234567890123456789012345678901234567890123456789@example.com"},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "string.email"}),
		},
		"email/invalid/host_segment_long": {
			Message: &cases.StringEmail{Val: "foo@x0123456789012345678901234567890123456789012345678901234567890123456789.com"},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "string.email"}),
		},
		"email/invalid/too_long": {
			Message: &cases.StringEmail{
				Val: "x123456789.x123456789.x123456789.x123456789.x123456789@x123456789.x123456789.x123456789.x123456789.x123456789.x123456789.x123456789.x123456789.x123456789.x123456789.x123456789.x123456789.x123456789.x123456789.x123456789.x123456789.x123456789.x123456789.x123456789",
			},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "string.email"}),
		},
		"email/invalid/bad_hostname": {
			Message: &cases.StringEmail{Val: "foo@-bar.com"},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "string.email"}),
		},
		"hostname/valid/lowercase": {
			Message:  &cases.StringHostname{Val: "example.com"},
			Expected: results.Success(true),
		},
		"hostname/valid/uppercase": {
			Message:  &cases.StringHostname{Val: "ASD.example.com"},
			Expected: results.Success(true),
		},
		"hostname/valid/hyphens": {
			Message:  &cases.StringHostname{Val: "foo-bar.com"},
			Expected: results.Success(true),
		},
		"hostname/valid/trailing_dot": {
			Message:  &cases.StringHostname{Val: "foo.bar."},
			Expected: results.Success(true),
		},
		"hostname/invalid/malformed": {
			Message: &cases.StringHostname{Val: "@!#$%^&*&^%$#"},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "string.hostname"}),
		},
		"hostname/invalid/underscore": {
			Message: &cases.StringHostname{Val: "foo_bar.com"},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "string.hostname"}),
		},
		"hostname/invalid/long": {
			Message: &cases.StringHostname{Val: "x0123456789012345678901234567890123456789012345678901234567890123456789.com"},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "string.hostname"}),
		},
		"hostname/invalid/trailing_hyphen": {
			Message: &cases.StringHostname{Val: "foo-bar-.com"},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "string.hostname"}),
		},
		"hostname/invalid/leading_hyphen": {
			Message: &cases.StringHostname{Val: "-foo-bar.com"},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "string.hostname"}),
		},
		"hostname/invalid/empty": {
			Message: &cases.StringHostname{Val: "foo..bar.com"},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "string.hostname"}),
		},
		"hostname/invalid/IDNs": {
			Message: &cases.StringHostname{Val: "你好.com"},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "string.hostname"}),
		},
		"ip/valid/v4": {
			Message:  &cases.StringIP{Val: "192.168.0.1"},
			Expected: results.Success(true),
		},
		"ip/valid/v6": {
			Message:  &cases.StringIP{Val: "3e::99"},
			Expected: results.Success(true),
		},
		"ip/invalid": {
			Message: &cases.StringIP{Val: "foobar"},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "string.ip"}),
		},
		"ipv4/valid": {
			Message:  &cases.StringIPv4{Val: "192.168.0.1"},
			Expected: results.Success(true),
		},
		"ipv4/invalid/malformed": {
			Message: &cases.StringIPv4{Val: "foobar"},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "string.ipv4"}),
		},
		"ipv4/invalid/erroneous": {
			Message: &cases.StringIPv4{Val: "256.0.0.0"},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "string.ipv4"}),
		},
		"ipv4/invalid/v6": {
			Message: &cases.StringIPv4{Val: "3e::99"},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "string.ipv4"}),
		},
		"ipv6/valid/expanded": {
			Message:  &cases.StringIPv6{Val: "2001:0db8:85a3:0000:0000:8a2e:0370:7334"},
			Expected: results.Success(true),
		},
		"ipv6/valid/collapsed": {
			Message:  &cases.StringIPv6{Val: "2001:0db8:85a3::8a2e:0370:7334"},
			Expected: results.Success(true),
		},
		"ipv6/invalid/malformed": {
			Message: &cases.StringIPv6{Val: "foobar"},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "string.ipv6"}),
		},
		"ipv6/invalid/erroneous": {
			Message: &cases.StringIPv6{Val: "ff::fff::0b"},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "string.ipv6"}),
		},
		"ipv6/invalid/v4": {
			Message: &cases.StringIPv6{Val: "192.168.0.1"},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "string.ipv6"}),
		},
		"uri/valid": {
			Message:  &cases.StringURI{Val: "https://example.com/foo/bar?baz=quux"},
			Expected: results.Success(true),
		},
		"uri/invalid/malformed": {
			Message: &cases.StringURI{Val: "!@#$%^&*"},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "string.uri"}),
		},
		"uri/invalid/relative": {
			Message: &cases.StringURI{Val: "/foo/bar?baz=quux"},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "string.uri"}),
		},
		"uri_ref/valid/absolute": {
			Message:  &cases.StringURIRef{Val: "https://example.com/foo/bar?baz=quux"},
			Expected: results.Success(true),
		},
		"uri_ref/valid/relative": {
			Message:  &cases.StringURIRef{Val: "/foo/bar?baz=quux"},
			Expected: results.Success(true),
		},
		"uri_ref/invalid": {
			Message: &cases.StringURIRef{Val: "!@#$%^&*"},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "string.uri_ref"}),
		},
		"address/valid/hostname": {
			Message:  &cases.StringAddress{Val: "foo.bar.com"},
			Expected: results.Success(true),
		},
		"address/valid/ipv4": {
			Message:  &cases.StringAddress{Val: "129.168.0.1"},
			Expected: results.Success(true),
		},
		"address/valid/ipv6": {
			Message:  &cases.StringAddress{Val: "3e::99"},
			Expected: results.Success(true),
		},
		"address/invalid/hostname": {
			Message: &cases.StringAddress{Val: "-foo.bar"},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "string.address"}),
		},
		"address/invalid/ipv6": {
			Message: &cases.StringAddress{Val: "ff::fff::0b"},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "string.address"}),
		},
		"uuid/valid/nil": {
			Message:  &cases.StringUUID{Val: "00000000-0000-0000-0000-000000000000"},
			Expected: results.Success(true),
		},
		"uuid/valid/v1/lowercase": {
			Message:  &cases.StringUUID{Val: "b45c0c80-8880-11e9-a5b1-000000000000"},
			Expected: results.Success(true),
		},
		"uuid/valid/v1/uppercase": {
			Message:  &cases.StringUUID{Val: "B45C0C80-8880-11E9-A5B1-000000000000"},
			Expected: results.Success(true),
		},
		"uuid/valid/v2/lowercase": {
			Message:  &cases.StringUUID{Val: "b45c0c80-8880-21e9-a5b1-000000000000"},
			Expected: results.Success(true),
		},
		"uuid/valid/v2/uppercase": {
			Message:  &cases.StringUUID{Val: "B45C0C80-8880-21E9-A5B1-000000000000"},
			Expected: results.Success(true),
		},
		"uuid/valid/v3/lowercase": {
			Message:  &cases.StringUUID{Val: "a3bb189e-8bf9-3888-9912-ace4e6543002"},
			Expected: results.Success(true),
		},
		"uuid/valid/v3/uppercase": {
			Message:  &cases.StringUUID{Val: "A3BB189E-8BF9-3888-9912-ACE4E6543002"},
			Expected: results.Success(true),
		},
		"uuid/valid/v4/lowercase": {
			Message:  &cases.StringUUID{Val: "8b208305-00e8-4460-a440-5e0dcd83bb0a"},
			Expected: results.Success(true),
		},
		"uuid/valid/v4/uppercase": {
			Message:  &cases.StringUUID{Val: "8B208305-00E8-4460-A440-5E0DCD83BB0A"},
			Expected: results.Success(true),
		},
		"uuid/valid/v5/lowercase": {
			Message:  &cases.StringUUID{Val: "a6edc906-2f9f-5fb2-a373-efac406f0ef2"},
			Expected: results.Success(true),
		},
		"uuid/valid/v5/uppercase": {
			Message:  &cases.StringUUID{Val: "A6EDC906-2F9F-5FB2-A373-EFAC406F0EF2"},
			Expected: results.Success(true),
		},
		"uuid/invalid/malformed": {
			Message: &cases.StringUUID{Val: "foobar"},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "string.uuid"}),
		},
		"uuid/invalid/erroneous": {
			Message: &cases.StringUUID{Val: "ffffffff-ffff-ffff-ffff-fffffffffffff"},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "string.uuid"}),
		},
		"well_known_regex/header_name/strict/valid/header": {
			Message:  &cases.StringHttpHeaderName{Val: "clustername"},
			Expected: results.Success(true),
		},
		"well_known_regex/header_name/strict/valid/pseudo_header": {
			Message:  &cases.StringHttpHeaderName{Val: ":authority"},
			Expected: results.Success(true),
		},
		"well_known_regex/header_name/strict/valid/numbers": {
			Message:  &cases.StringHttpHeaderName{Val: "abc-123"},
			Expected: results.Success(true),
		},
		"well_known_regex/header_name/strict/valid/special_token": {
			Message:  &cases.StringHttpHeaderName{Val: "!+#&.%"},
			Expected: results.Success(true),
		},
		"well_known_regex/header_name/strict/valid/period": {
			Message:  &cases.StringHttpHeaderName{Val: "FOO.BAR"},
			Expected: results.Success(true),
		},
		"well_known_regex/header_name/strict/invalid/empty": {
			Message: &cases.StringHttpHeaderName{Val: ""},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "string.well_known_regex.header_name"}),
		},
		"well_known_regex/header_name/strict/invalid/solo_colon": {
			Message: &cases.StringHttpHeaderName{Val: ":"},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "string.well_known_regex.header_name"}),
		},
		"well_known_regex/header_name/strict/invalid/trailing_colon": {
			Message: &cases.StringHttpHeaderName{Val: ":foo:"},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "string.well_known_regex.header_name"}),
		},
		"well_known_regex/header_name/strict/invalid/space": {
			Message: &cases.StringHttpHeaderName{Val: "foo bar"},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "string.well_known_regex.header_name"}),
		},
		"well_known_regex/header_name/strict/invalid/cr": {
			Message: &cases.StringHttpHeaderName{Val: "foo\rbar"},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "string.well_known_regex.header_name"}),
		},
		"well_known_regex/header_name/strict/invalid/lf": {
			Message: &cases.StringHttpHeaderName{Val: "foo\nbar"},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "string.well_known_regex.header_name"}),
		},
		"well_known_regex/header_name/strict/invalid/tab": {
			Message: &cases.StringHttpHeaderName{Val: "foo\tbar"},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "string.well_known_regex.header_name"}),
		},
		"well_known_regex/header_name/strict/invalid/nul": {
			Message: &cases.StringHttpHeaderName{Val: "foo\000bar"},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "string.well_known_regex.header_name"}),
		},
		"well_known_regex/header_name/strict/invalid/slash": {
			Message: &cases.StringHttpHeaderName{Val: "foo/bar"},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "string.well_known_regex.header_name"}),
		},
		"well_known_regex/header_name/loose/valid/slash": {
			Message:  &cases.StringHttpHeaderNameLoose{Val: "FOO/BAR"},
			Expected: results.Success(true),
		},
		"well_known_regex/header_name/loose/valid/tab": {
			Message:  &cases.StringHttpHeaderNameLoose{Val: "FOO\tBAR"},
			Expected: results.Success(true),
		},
		"well_known_regex/header_name/loose/valid/space": {
			Message:  &cases.StringHttpHeaderNameLoose{Val: "FOO BAR"},
			Expected: results.Success(true),
		},
		"well_known_regex/header_name/loose/invalid/empty": {
			Message: &cases.StringHttpHeaderNameLoose{Val: ""},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "string.well_known_regex.header_name"}),
		},
		"well_known_regex/header_name/loose/invalid/cr": {
			Message: &cases.StringHttpHeaderNameLoose{Val: "foo\rbar"},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "string.well_known_regex.header_name"}),
		},
		"well_known_regex/header_name/loose/invalid/lf": {
			Message: &cases.StringHttpHeaderNameLoose{Val: "foo\nbar"},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "string.well_known_regex.header_name"}),
		},
		"well_known_regex/header_value/strict/valid/empty": {
			Message:  &cases.StringHttpHeaderValue{Val: ""},
			Expected: results.Success(true),
		},
		"well_known_regex/header_value/strict/valid/periods": {
			Message:  &cases.StringHttpHeaderValue{Val: "foo.bar.baz"},
			Expected: results.Success(true),
		},
		"well_known_regex/header_value/strict/valid/uppercase": {
			Message:  &cases.StringHttpHeaderValue{Val: "/TEST/SOME/PATH"},
			Expected: results.Success(true),
		},
		"well_known_regex/header_value/strict/valid/spaces": {
			Message:  &cases.StringHttpHeaderValue{Val: "cluster name"},
			Expected: results.Success(true),
		},
		"well_known_regex/header_value/strict/valid/tab": {
			Message:  &cases.StringHttpHeaderValue{Val: "cluster\tname"},
			Expected: results.Success(true),
		},
		"well_known_regex/header_value/strict/valid/special_token": {
			Message:  &cases.StringHttpHeaderValue{Val: "!#%&./+"},
			Expected: results.Success(true),
		},
		"well_known_regex/header_value/strict/invalid/nul": {
			Message: &cases.StringHttpHeaderValue{Val: "foo\000bar"},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "string.well_known_regex.header_value"}),
		},
		"well_known_regex/header_value/strict/invalid/del": {
			Message: &cases.StringHttpHeaderValue{Val: "foo\007bar"},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "string.well_known_regex.header_value"}),
		},
		"well_known_regex/header_value/strict/invalid/cr": {
			Message: &cases.StringHttpHeaderValue{Val: "foo\rbar"},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "string.well_known_regex.header_value"}),
		},
		"well_known_regex/header_value/strict/invalid/lf": {
			Message: &cases.StringHttpHeaderValue{Val: "foo\nbar"},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "string.well_known_regex.header_value"}),
		},
		"well_known_regex/header_name/loose/valid/del": {
			Message:  &cases.StringHttpHeaderNameLoose{Val: "FOO\007BAR"},
			Expected: results.Success(true),
		},
		"well_known_regex/header_value/loose/invalid/nul": {
			Message: &cases.StringHttpHeaderValueLoose{Val: "\000"},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "string.well_known_regex.header_value"}),
		},
		"well_known_regex/header_value/loose/invalid/cr": {
			Message: &cases.StringHttpHeaderValueLoose{Val: "foo\rbar"},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "string.well_known_regex.header_value"}),
		},
		"well_known_regex/header_value/loose/invalid/lf": {
			Message: &cases.StringHttpHeaderValueLoose{Val: "foo\nbar"},
			Expected: results.Violations(
				&validate.Violation{FieldPath: "val", ConstraintId: "string.well_known_regex.header_value"}),
		},
	}
}
