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

package constraints

import (
	"testing"

	"github.com/google/cel-go/cel"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func TestExpectedWrapperConstraints(t *testing.T) {
	t.Parallel()

	tests := map[protoreflect.FullName]*string{
		"google.protobuf.BoolValue":   proto.String("buf.validate.FieldConstraints.bool"),
		"google.protobuf.BytesValue":  proto.String("buf.validate.FieldConstraints.bytes"),
		"google.protobuf.DoubleValue": proto.String("buf.validate.FieldConstraints.double"),
		"google.protobuf.FloatValue":  proto.String("buf.validate.FieldConstraints.float"),
		"google.protobuf.Int32Value":  proto.String("buf.validate.FieldConstraints.int32"),
		"google.protobuf.Int64Value":  proto.String("buf.validate.FieldConstraints.int64"),
		"google.protobuf.StringValue": proto.String("buf.validate.FieldConstraints.string"),
		"google.protobuf.UInt32Value": proto.String("buf.validate.FieldConstraints.uint32"),
		"google.protobuf.UInt64Value": proto.String("buf.validate.FieldConstraints.uint64"),
		"foo.bar":                     nil,
	}

	for name, cons := range tests {
		fqn, constraint := name, cons
		t.Run(string(fqn), func(t *testing.T) {
			t.Parallel()
			desc, ok := ExpectedWrapperConstraints(fqn)
			if constraint != nil {
				assert.Equal(t, *constraint, string(desc.FullName()))
				assert.True(t, ok)
			} else {
				assert.False(t, ok)
			}
		})
	}
}

func TestProtoKindToCELType(t *testing.T) {
	t.Parallel()

	tests := map[protoreflect.Kind]*cel.Type{
		protoreflect.FloatKind:    cel.DoubleType,
		protoreflect.DoubleKind:   cel.DoubleType,
		protoreflect.Int32Kind:    cel.IntType,
		protoreflect.Int64Kind:    cel.IntType,
		protoreflect.Uint32Kind:   cel.UintType,
		protoreflect.Uint64Kind:   cel.UintType,
		protoreflect.Sint32Kind:   cel.IntType,
		protoreflect.Sint64Kind:   cel.IntType,
		protoreflect.Fixed32Kind:  cel.UintType,
		protoreflect.Fixed64Kind:  cel.UintType,
		protoreflect.Sfixed32Kind: cel.IntType,
		protoreflect.Sfixed64Kind: cel.IntType,
		protoreflect.BoolKind:     cel.BoolType,
		protoreflect.StringKind:   cel.StringType,
		protoreflect.BytesKind:    cel.BytesType,
		protoreflect.EnumKind:     cel.IntType,
		protoreflect.MessageKind:  cel.DynType,
		protoreflect.GroupKind:    cel.DynType,
		protoreflect.Kind(0):      cel.DynType,
	}

	for k, ty := range tests {
		kind, typ := k, ty
		t.Run(kind.String(), func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, typ, ProtoKindToCELType(kind))
		})
	}
}
