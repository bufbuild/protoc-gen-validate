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

func TestValidator_Validate(t *testing.T) {
	t.Parallel()

	t.Run("HasMsgExprs", func(t *testing.T) {
		t.Parallel()
		val, err := New()
		require.NoError(t, err)

		tests := []struct {
			msg   *pb.HasMsgExprs
			exErr bool
		}{
			{
				&pb.HasMsgExprs{X: 2, Y: 43},
				false,
			},
			{
				&pb.HasMsgExprs{X: 9, Y: 8},
				true,
			},
		}

		for _, test := range tests {
			err := val.Validate(test.msg)
			if test.exErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		}
	})
}

func TestRecursive(t *testing.T) {
	t.Parallel()
	val, err := New()
	require.NoError(t, err)

	selfRec := &pb.SelfRecursive{X: 123, Turtle: &pb.SelfRecursive{X: 456}}
	err = val.Validate(selfRec)
	assert.NoError(t, err)

	loopRec := &pb.LoopRecursiveA{B: &pb.LoopRecursiveB{}}
	err = val.Validate(loopRec)
	assert.NoError(t, err)
}

func TestValidator_ValidateOneof(t *testing.T) {
	t.Parallel()
	val, err := New()
	require.NoError(t, err)
	oneofMessage := &pb.MsgHasOneof{O: &pb.MsgHasOneof_X{X: "foo"}}
	err = val.Validate(oneofMessage)
	assert.NoError(t, err)

	oneofMessage = &pb.MsgHasOneof{O: &pb.MsgHasOneof_Y{Y: 42}}
	err = val.Validate(oneofMessage)
	assert.NoError(t, err)

	oneofMessage = &pb.MsgHasOneof{O: &pb.MsgHasOneof_Msg{Msg: &pb.HasMsgExprs{X: 4, Y: 50}}}
	err = val.Validate(oneofMessage)
	assert.NoError(t, err)

	oneofMessage = &pb.MsgHasOneof{}
	err = val.Validate(oneofMessage)
	assert.Error(t, err)
}

func TestValidator_ValidateRepeatedFoo(t *testing.T) {
	t.Parallel()
	val, err := New()
	require.NoError(t, err)
	repeatMessage := &pb.MsgHasRepeated{
		X: []float32{1, 2, 3},
		Y: []string{"foo", "bar"},
		Z: []*pb.HasMsgExprs{
			{
				X: 4,
				Y: 55,
			}, {
				X: 4,
				Y: 60,
			},
		},
	}
	err = val.Validate(repeatMessage)
	require.NoError(t, err)
}

func TestValidator_ValidateMapFoo(t *testing.T) {
	t.Parallel()
	val, err := New()
	require.NoError(t, err)
	mapMessage := &pb.MsgHasMap{
		Int32Map:   map[int32]int32{-1: 1, 2: 2},
		StringMap:  map[string]string{"foo": "foo", "bar": "bar", "baz": "baz"},
		MessageMap: map[int64]*pb.LoopRecursiveA{0: nil},
	}
	err = val.Validate(mapMessage)
	require.Error(t, err)
	t.Log(err)
}
