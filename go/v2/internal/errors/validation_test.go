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

package errors

import (
	"testing"

	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	"github.com/stretchr/testify/assert"
)

func TestPrefixFieldPaths(t *testing.T) {
	t.Parallel()

	tests := []struct {
		fieldPath string
		format    string
		args      []any
		expected  string
	}{
		{
			"",
			"%s",
			[]any{"foo"},
			"foo",
		},
		{
			"bar",
			"%s",
			[]any{"foo"},
			"foo.bar",
		},
		{
			"bar",
			"[%d]",
			[]any{3},
			"[3].bar",
		},
		{
			"[3].bar",
			"%s",
			[]any{"foo"},
			"foo[3].bar",
		},
	}

	for _, tc := range tests {
		test := tc
		t.Run(test.expected, func(t *testing.T) {
			t.Parallel()
			err := &ValidationError{Violations: []*validate.Violation{
				{FieldPath: test.fieldPath},
				{FieldPath: test.fieldPath},
			}}
			PrefixFieldPaths(err, test.format, test.args...)
			for _, v := range err.Violations {
				assert.Equal(t, test.expected, v.FieldPath)
			}
		})
	}
}
