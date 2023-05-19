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
	"testing"

	examplev1 "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/tests/example/v1"
	protovalidate "github.com/bufbuild/protovalidate/go/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func TestWithLegacySupport(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		mode  Mode
		msg   proto.Message
		exErr bool
	}{
		{
			name:  "MessageRules/Merge/Valid",
			mode:  ModeMerge,
			msg:   &examplev1.LegacyMixedMessage{X: 123},
			exErr: false,
		},
		{
			name:  "MessageRules/Merge/Invalid",
			mode:  ModeMerge,
			msg:   &examplev1.LegacyMixedMessage{},
			exErr: true,
		},
		{
			name:  "MessageRules/Replace/Valid",
			mode:  ModeReplace,
			msg:   &examplev1.LegacyMixedMessage{},
			exErr: false,
		},
		{
			name:  "MessageRules/IfNotPresent/Invalid",
			mode:  ModeIfNotPresent,
			msg:   &examplev1.LegacyMixedMessage{},
			exErr: true,
		},
		{
			name:  "OneofRules/Merge/Valid",
			mode:  ModeMerge,
			msg:   &examplev1.LegacyMixedOneof{},
			exErr: false,
		},
		{
			name:  "OneofRules/Replace/Invalid",
			mode:  ModeReplace,
			msg:   &examplev1.LegacyMixedOneof{},
			exErr: true,
		},
		{
			name:  "OneofRules/Replace/Valid",
			mode:  ModeReplace,
			msg:   &examplev1.LegacyMixedOneof{O: &examplev1.LegacyMixedOneof_X{X: 123}},
			exErr: false,
		},
		{
			name:  "OneofRules/IfNotPresent/Valid",
			mode:  ModeIfNotPresent,
			msg:   &examplev1.LegacyMixedOneof{},
			exErr: false,
		},
		{
			name:  "FieldRules/Merged/Valid",
			mode:  ModeMerge,
			msg:   &examplev1.LegacyMixedFields{X: 1},
			exErr: false,
		},
		{
			name:  "FieldRules/Merged/Invalid",
			mode:  ModeMerge,
			msg:   &examplev1.LegacyMixedFields{X: 123},
			exErr: true,
		},
		{
			name:  "FieldRules/Replace/Valid",
			mode:  ModeReplace,
			msg:   &examplev1.LegacyMixedFields{X: 123},
			exErr: false,
		},
		{
			name:  "FieldRules/Replace/Invalid",
			mode:  ModeReplace,
			msg:   &examplev1.LegacyMixedFields{X: -1},
			exErr: true,
		},
		{
			name:  "FieldRules/IfNotPresent/Valid",
			mode:  ModeIfNotPresent,
			msg:   &examplev1.LegacyMixedFields{X: -1},
			exErr: false,
		},
		{
			name:  "FieldRules/IfNotPresent/Invalid",
			mode:  ModeIfNotPresent,
			msg:   &examplev1.LegacyMixedFields{X: 123},
			exErr: true,
		},
	}

	for _, tc := range tests {
		test := tc
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			val, err := protovalidate.New(
				WithLegacySupport(test.mode),
				protovalidate.WithMessages(test.msg),
			)
			require.NoError(t, err)
			err = val.Validate(test.msg)
			if test.exErr {
				valErr := &protovalidate.ValidationError{}
				assert.ErrorAs(t, err, &valErr)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
