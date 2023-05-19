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

package protovalidate

import (
	"testing"

	pb "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/tests/example/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkValidator(b *testing.B) {
	successMsg := &pb.HasMsgExprs{X: 2, Y: 43}
	failureMsg := &pb.HasMsgExprs{X: 9, Y: 2}

	b.Run("ColdStart", func(b *testing.B) {
		b.ReportAllocs()
		b.RunParallel(func(p *testing.PB) {
			for p.Next() {
				val, err := New()
				require.NoError(b, err)
				err = val.Validate(successMsg)
				assert.NoError(b, err)
			}
		})
	})

	b.Run("Lazy/Valid", func(b *testing.B) {
		b.ReportAllocs()
		val, err := New()
		require.NoError(b, err)
		b.ResetTimer()
		b.RunParallel(func(p *testing.PB) {
			for p.Next() {
				err := val.Validate(successMsg)
				assert.NoError(b, err)
			}
		})
	})

	b.Run("Lazy/Invalid", func(b *testing.B) {
		b.ReportAllocs()
		val, err := New()
		require.NoError(b, err)
		b.ResetTimer()
		b.RunParallel(func(p *testing.PB) {
			for p.Next() {
				err := val.Validate(failureMsg)
				assert.Error(b, err)
			}
		})
	})

	b.Run("Lazy/FailFast", func(b *testing.B) {
		b.ReportAllocs()
		val, err := New(WithFailFast(true))
		require.NoError(b, err)
		b.ResetTimer()
		b.RunParallel(func(p *testing.PB) {
			for p.Next() {
				err := val.Validate(failureMsg)
				assert.Error(b, err)
			}
		})
	})

	b.Run("PreWarmed/Valid", func(b *testing.B) {
		b.ReportAllocs()
		val, err := New(
			WithMessages(successMsg),
			WithDisableLazy(true),
		)
		require.NoError(b, err)
		b.ResetTimer()
		b.RunParallel(func(p *testing.PB) {
			for p.Next() {
				err := val.Validate(successMsg)
				assert.NoError(b, err)
			}
		})
	})

	b.Run("PreWarmed/Invalid", func(b *testing.B) {
		b.ReportAllocs()
		val, err := New(
			WithMessages(failureMsg),
			WithDisableLazy(true),
		)
		require.NoError(b, err)
		b.ResetTimer()
		b.RunParallel(func(p *testing.PB) {
			for p.Next() {
				err := val.Validate(failureMsg)
				assert.Error(b, err)
			}
		})
	})

	b.Run("PreWarmed/FailFast", func(b *testing.B) {
		b.ReportAllocs()
		val, err := New(
			WithFailFast(true),
			WithMessages(failureMsg),
			WithDisableLazy(true),
		)
		require.NoError(b, err)
		b.ResetTimer()
		b.RunParallel(func(p *testing.PB) {
			for p.Next() {
				err := val.Validate(failureMsg)
				assert.Error(b, err)
			}
		})
	})
}
