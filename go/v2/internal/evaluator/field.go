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

package evaluator

import (
	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	"github.com/bufbuild/protovalidate/go/v2/internal/errors"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// field performs validation on a single message field, defined by its
// descriptor.
type field struct {
	// Value is the evaluator to apply to the field's value
	Value value
	// Descriptor is the FieldDescriptor targeted by this evaluator
	Descriptor protoreflect.FieldDescriptor
	// Required indicates that the field must have a set value.
	Required bool
	// Optional indicates that the evaluators should not be applied to this field
	// if the value is unset. Fields that contain messages, are prefixed with
	// `optional`, or are part of a oneof are considered optional. evaluators
	// will still be applied if the field is set as the zero value.
	Optional bool
}

func (f field) Evaluate(val protoreflect.Value, failFast bool) error {
	return f.EvaluateMessage(val.Message(), failFast)
}

func (f field) EvaluateMessage(msg protoreflect.Message, failFast bool) (err error) {
	if f.Required && !msg.Has(f.Descriptor) {
		return &errors.ValidationError{Violations: []*validate.Violation{{
			FieldPath:    string(f.Descriptor.Name()),
			ConstraintId: "required",
			Message:      "value is required",
		}}}
	}

	if (f.Optional || f.Value.IgnoreEmpty) && !msg.Has(f.Descriptor) {
		return nil
	}

	val := msg.Get(f.Descriptor)
	if err = f.Value.Evaluate(val, failFast); err != nil {
		errors.PrefixErrorPaths(err, string(f.Descriptor.Name()))
	}
	return err
}

func (f field) Tautology() bool {
	return !f.Required && f.Value.Tautology()
}

var _ MessageEvaluator = field{}
