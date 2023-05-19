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

package legacy

import (
	"fmt"
	"testing"

	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	examplev1 "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/tests/example/v1"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func TestTranslateMessageOptions(t *testing.T) {
	t.Parallel()

	tests := []struct {
		msg proto.Message
		ex  *validate.MessageConstraints
	}{
		{
			msg: &examplev1.LegacyNone{},
			ex:  nil,
		},
		{
			msg: &examplev1.LegacyDisabled{},
			ex:  &validate.MessageConstraints{Disabled: proto.Bool(true)},
		},
		{
			msg: &examplev1.LegacyIgnored{},
			ex:  &validate.MessageConstraints{Disabled: proto.Bool(true)},
		},
	}

	for _, tc := range tests {
		test := tc
		t.Run(fmt.Sprintf("%T", test.msg), func(t *testing.T) {
			t.Parallel()
			desc := test.msg.ProtoReflect().Descriptor()
			out := translateMessageOptions(desc)
			assert.True(t, proto.Equal(test.ex, out))
		})
	}
}

func TestTranslateOneofOptions(t *testing.T) {
	t.Parallel()

	tests := []struct {
		msg   proto.Message
		oneof protoreflect.Name
		ex    *validate.OneofConstraints
	}{
		{
			msg:   &examplev1.LegacyNone{},
			oneof: "o",
			ex:    nil,
		},
		{
			msg:   &examplev1.LegacyOneofRequired{},
			oneof: "o",
			ex:    &validate.OneofConstraints{Required: proto.Bool(true)},
		},
	}

	for _, tc := range tests {
		test := tc
		t.Run(fmt.Sprintf("%T", test.msg), func(t *testing.T) {
			t.Parallel()
			desc := test.msg.ProtoReflect().Descriptor().Oneofs().ByName(test.oneof)
			out := translateOneofOptions(desc)
			assert.True(t, proto.Equal(test.ex, out))
		})
	}
}

func TestTranslateFieldOptions(t *testing.T) {
	t.Parallel()

	tests := []struct {
		msg   proto.Message
		field protoreflect.Name
		ex    *validate.FieldConstraints
	}{
		{
			msg:   &examplev1.LegacyNone{},
			field: "x",
			ex:    nil,
		},
		{
			msg:   &examplev1.LegacySimple{},
			field: "x",
			ex: &validate.FieldConstraints{Type: &validate.FieldConstraints_Int32{
				Int32: &validate.Int32Rules{Gt: proto.Int32(0)},
			}},
		},
		{
			msg:   &examplev1.LegacySkipped{},
			field: "x",
			ex: &validate.FieldConstraints{
				Skipped: true,
			},
		},
		{
			msg:   &examplev1.LegacyMessageRequired{},
			field: "x",
			ex: &validate.FieldConstraints{
				Required: true,
			},
		},
		{
			msg:   &examplev1.LegacyIn{},
			field: "x",
			ex: &validate.FieldConstraints{Type: &validate.FieldConstraints_Int32{
				Int32: &validate.Int32Rules{In: []int32{1, 2, 3}},
			}},
		},
		{
			msg:   &examplev1.LegacyRepeated{},
			field: "x",
			ex: &validate.FieldConstraints{Type: &validate.FieldConstraints_Repeated{
				Repeated: &validate.RepeatedRules{Items: &validate.FieldConstraints{
					Type: &validate.FieldConstraints_Int32{
						Int32: &validate.Int32Rules{Gt: proto.Int32(0)},
					},
				}},
			}},
		},
		{
			msg:   &examplev1.LegacyMap{},
			field: "x",
			ex: &validate.FieldConstraints{Type: &validate.FieldConstraints_Map{
				Map: &validate.MapRules{
					Keys: &validate.FieldConstraints{Type: &validate.FieldConstraints_String_{
						String_: &validate.StringRules{MinLen: proto.Uint64(3)},
					}},
					Values: &validate.FieldConstraints{Type: &validate.FieldConstraints_Int32{
						Int32: &validate.Int32Rules{Gt: proto.Int32(0)},
					}},
				},
			}},
		},
		{
			msg:   &examplev1.LegacyEnum{},
			field: "x",
			ex: &validate.FieldConstraints{Type: &validate.FieldConstraints_String_{
				String_: &validate.StringRules{WellKnown: &validate.StringRules_WellKnownRegex{
					WellKnownRegex: validate.KnownRegex_KNOWN_REGEX_HTTP_HEADER_NAME,
				}},
			}},
		},
		{
			msg:   &examplev1.LegacyWKTRequired{},
			field: "any",
			ex: &validate.FieldConstraints{
				Required: true,
				Type:     &validate.FieldConstraints_Any{},
			},
		},
		{
			msg:   &examplev1.LegacyWKTRequired{},
			field: "ts",
			ex: &validate.FieldConstraints{
				Required: true,
				Type:     &validate.FieldConstraints_Timestamp{},
			},
		},
		{
			msg:   &examplev1.LegacyWKTRequired{},
			field: "dur",
			ex: &validate.FieldConstraints{
				Required: true,
				Type:     &validate.FieldConstraints_Duration{},
			},
		},
		{
			msg:   &examplev1.LegacyIgnoreEmpty{},
			field: "x",
			ex: &validate.FieldConstraints{
				IgnoreEmpty: true,
				Type: &validate.FieldConstraints_Int32{Int32: &validate.Int32Rules{
					Gt: proto.Int32(123),
				}},
			},
		},
	}

	for _, tc := range tests {
		test := tc
		t.Run(fmt.Sprintf("%T", test.msg), func(t *testing.T) {
			t.Parallel()
			desc := test.msg.ProtoReflect().Descriptor().Fields().ByName(test.field)
			out := translateFieldOptions(desc)
			assert.True(t, proto.Equal(test.ex, out))
		})
	}
}
