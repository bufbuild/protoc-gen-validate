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
	"errors"
	"testing"

	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func TestMerge(t *testing.T) {
	t.Parallel()

	t.Run("no errors", func(t *testing.T) {
		t.Parallel()
		ok, err := Merge(nil, nil, true)
		assert.NoError(t, err)
		assert.True(t, ok)
		ok, err = Merge(nil, nil, false)
		assert.NoError(t, err)
		assert.True(t, ok)
	})

	t.Run("no dst", func(t *testing.T) {
		t.Parallel()

		t.Run("non-validation", func(t *testing.T) {
			t.Parallel()
			someErr := errors.New("some error")
			ok, err := Merge(nil, someErr, true)
			assert.Equal(t, someErr, err)
			assert.False(t, ok)
			ok, err = Merge(nil, someErr, false)
			assert.Equal(t, someErr, err)
			assert.False(t, ok)
		})

		t.Run("validation", func(t *testing.T) {
			t.Parallel()
			exErr := &ValidationError{Violations: []*validate.Violation{{ConstraintId: "foo"}}}
			ok, err := Merge(nil, exErr, true)
			var valErr *ValidationError
			require.ErrorAs(t, err, &valErr)
			assert.True(t, proto.Equal(exErr.ToProto(), valErr.ToProto()))
			assert.False(t, ok)
			ok, err = Merge(nil, exErr, false)
			require.ErrorAs(t, err, &valErr)
			assert.True(t, proto.Equal(exErr.ToProto(), valErr.ToProto()))
			assert.True(t, ok)
		})
	})

	t.Run("merge", func(t *testing.T) {
		t.Parallel()

		t.Run("non-validation dst", func(t *testing.T) {
			t.Parallel()
			dstErr := errors.New("some error")
			srcErr := &ValidationError{Violations: []*validate.Violation{{ConstraintId: "foo"}}}
			ok, err := Merge(dstErr, srcErr, true)
			assert.Equal(t, dstErr, err)
			assert.False(t, ok)
			ok, err = Merge(dstErr, srcErr, false)
			assert.Equal(t, dstErr, err)
			assert.False(t, ok)
		})

		t.Run("non-validation src", func(t *testing.T) {
			t.Parallel()
			dstErr := &ValidationError{Violations: []*validate.Violation{{ConstraintId: "foo"}}}
			srcErr := errors.New("some error")
			ok, err := Merge(dstErr, srcErr, true)
			assert.Equal(t, srcErr, err)
			assert.False(t, ok)
			ok, err = Merge(dstErr, srcErr, false)
			assert.Equal(t, srcErr, err)
			assert.False(t, ok)
		})

		t.Run("validation", func(t *testing.T) {
			t.Parallel()

			dstErr := &ValidationError{Violations: []*validate.Violation{{ConstraintId: "foo"}}}
			srcErr := &ValidationError{Violations: []*validate.Violation{{ConstraintId: "bar"}}}
			exErr := &ValidationError{Violations: []*validate.Violation{
				{ConstraintId: "foo"},
				{ConstraintId: "bar"},
			}}
			ok, err := Merge(dstErr, srcErr, true)
			var valErr *ValidationError
			require.ErrorAs(t, err, &valErr)
			assert.True(t, proto.Equal(exErr.ToProto(), valErr.ToProto()))
			assert.False(t, ok)
			dstErr = &ValidationError{Violations: []*validate.Violation{{ConstraintId: "foo"}}}
			ok, err = Merge(dstErr, srcErr, false)
			require.ErrorAs(t, err, &valErr)
			assert.True(t, proto.Equal(exErr.ToProto(), valErr.ToProto()))
			assert.True(t, ok)
		})
	})
}
