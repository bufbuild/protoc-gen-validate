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

package expression

import (
	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	"github.com/bufbuild/protovalidate/go/v2/internal/errors"
	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/interpreter"
	"google.golang.org/protobuf/reflect/protoreflect"
)

//nolint:gochecknoglobals // amortized, eliminates allocations for all CEL programs
var varPool = &VariablePool{New: func() any { return &Variable{} }}

// ProgramSet is a list of compiledProgram expressions that are evaluated
// together with the same input value. All expressions in a ProgramSet may refer
// to a `this` variable.
type ProgramSet []compiledProgram

// Eval applies the contained expressions to the provided `this` val, returning
// either *errors.ValidationError if the input is invalid or errors.RuntimeError
// if there is a type or range error. If failFast is true, execution stops at
// the first failed expression.
func (s ProgramSet) Eval(val any, failFast bool) error {
	binding := s.bindThis(val)
	defer varPool.Put(binding)

	var violations []*validate.Violation
	for _, expr := range s {
		violation, err := expr.eval(binding)
		if err != nil {
			return err
		}
		if violation != nil {
			violations = append(violations, violation)
			if failFast {
				break
			}
		}
	}

	if len(violations) > 0 {
		return &errors.ValidationError{Violations: violations}
	}

	return nil
}

func (s ProgramSet) bindThis(val any) *Variable {
	binding := varPool.Get()
	binding.Name = "this"

	switch value := val.(type) {
	case protoreflect.Message:
		binding.Val = value.Interface()
	case protoreflect.Map:
		// TODO: expensive to create this copy, but getting this into a ref.Val with
		//  traits.Mapper is not terribly feasible from this type.
		m := make(map[any]any, value.Len())
		value.Range(func(key protoreflect.MapKey, value protoreflect.Value) bool {
			m[key.Interface()] = value.Interface()
			return true
		})
		binding.Val = m
	default:
		binding.Val = value
	}

	return binding
}

// compiledProgram is a parsed and type-checked cel.Program along with the
// source Expression.
type compiledProgram struct {
	Program cel.Program
	Source  Expression
}

//nolint:nilnil // non-existence of violations is intentional
func (expr compiledProgram) eval(bindings interpreter.Activation) (*validate.Violation, error) {
	value, _, err := expr.Program.Eval(bindings)
	if err != nil {
		return nil, errors.NewRuntimeErrorf(
			"error evaluating %s: %w", expr.Source.GetId(), err)
	}
	switch val := value.Value().(type) {
	case string:
		if val == "" {
			return nil, nil
		}
		return &validate.Violation{
			ConstraintId: expr.Source.GetId(),
			Message:      val,
		}, nil
	case bool:
		if val {
			return nil, nil
		}
		return &validate.Violation{
			ConstraintId: expr.Source.GetId(),
			Message:      expr.Source.GetMessage(),
		}, nil
	default:
		return nil, errors.NewRuntimeErrorf(
			"resolved to an unexpected type %T", val)
	}
}
