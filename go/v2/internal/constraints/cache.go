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
	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate/priv"
	"github.com/bufbuild/protovalidate/go/v2/internal/errors"
	"github.com/bufbuild/protovalidate/go/v2/internal/expression"
	"github.com/google/cel-go/cel"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// Cache is a build-through cache to computed standard constraints.
type Cache struct {
	cache map[protoreflect.FieldDescriptor]expression.ASTSet
}

// NewCache constructs a new build-through cache for the standard constraints.
func NewCache() Cache {
	return Cache{
		cache: map[protoreflect.FieldDescriptor]expression.ASTSet{},
	}
}

// Build creates the standard constraints for the given field. If forItems is
// true, the constraints for repeated list items is built instead of the
// constraints on the list itself.
func (c *Cache) Build(
	env *cel.Env,
	fieldDesc protoreflect.FieldDescriptor,
	fieldConstraints *validate.FieldConstraints,
	forItems bool,
) (set expression.ProgramSet, err error) {
	constraints, done, err := c.resolveConstraints(fieldDesc, fieldConstraints, forItems)
	if done {
		return nil, err
	}

	env, err = c.prepareEnvironment(env, fieldDesc, constraints, forItems)
	if err != nil {
		return nil, err
	}

	var asts expression.ASTSet
	constraints.Range(func(desc protoreflect.FieldDescriptor, val protoreflect.Value) bool {
		precomputedASTs, compileErr := c.loadOrCompileStandardConstraint(env, desc)
		if compileErr != nil {
			err = compileErr
			return false
		}
		asts = asts.Merge(precomputedASTs)
		return true
	})
	if err != nil {
		return nil, err
	}

	rulesGlobal := cel.Globals(&expression.Variable{Name: "rules", Val: constraints.Interface()})
	set, err = asts.ReduceResiduals(rulesGlobal)
	return set, err
}

// resolveConstraints extracts the standard constraints for the specified field. An
// error is returned if the wrong constraints are applied to a field (typically
// if there is a type-mismatch). The done result is true if an error is returned
// or if there are now standard constraints to apply to this field.
func (c *Cache) resolveConstraints(
	fieldDesc protoreflect.FieldDescriptor,
	fieldConstraints *validate.FieldConstraints,
	forItems bool,
) (rules protoreflect.Message, done bool, err error) {
	constraints := fieldConstraints.ProtoReflect()
	setOneof := constraints.WhichOneof(fieldConstraintsOneofDesc)
	if setOneof == nil {
		return nil, true, nil
	}
	expected, ok := c.getExpectedConstraintDescriptor(fieldDesc, forItems)
	if ok && setOneof.FullName() != expected.FullName() {
		return nil, true, errors.NewCompilationErrorf(
			"expected constraint %q, got %q on field %q",
			expected.FullName(),
			setOneof.FullName(),
			fieldDesc.FullName(),
		)
	}
	if !ok || !constraints.Has(setOneof) {
		return nil, true, nil
	}
	rules = constraints.Get(setOneof).Message()
	return rules, false, nil
}

// prepareEnvironment prepares the environment for compiling standard constraint
// expressions.
func (c *Cache) prepareEnvironment(
	env *cel.Env,
	fieldDesc protoreflect.FieldDescriptor,
	rules protoreflect.Message,
	forItems bool,
) (*cel.Env, error) {
	env, err := env.Extend(
		cel.Types(rules.Interface()),
		cel.Variable("this", c.getCELType(fieldDesc, forItems)),
		cel.Variable("rules",
			cel.ObjectType(string(rules.Descriptor().FullName()))),
	)
	if err != nil {
		return nil, errors.NewCompilationErrorf(
			"failed to extend base environment: %w", err)
	}
	return env, nil
}

// loadOrCompileStandardConstraint loads the precompiled ASTs for the
// specified constraint field from the Cache if present or precomputes them
// otherwise. The result may be empty if the constraint does not have associated
// CEL expressions.
func (c *Cache) loadOrCompileStandardConstraint(
	env *cel.Env,
	constraintFieldDesc protoreflect.FieldDescriptor,
) (set expression.ASTSet, err error) {
	if cachedConstraint, ok := c.cache[constraintFieldDesc]; ok {
		return cachedConstraint, nil
	}
	exprs, _ := proto.GetExtension(constraintFieldDesc.Options(), priv.E_Field).(*priv.FieldConstraints)
	set, err = expression.CompileASTs(exprs.GetCel(), env)
	if err != nil {
		return set, errors.NewCompilationErrorf(
			"failed to compile standard constraint %q: %w",
			constraintFieldDesc.FullName(), err)
	}
	c.cache[constraintFieldDesc] = set
	return set, nil
}

// getExpectedConstraintDescriptor produces the field descriptor from the
// validate.FieldConstraints 'type' oneof that matches the provided target
// field descriptor. If ok is false, the field does not expect any standard
// constraints.
func (c *Cache) getExpectedConstraintDescriptor(
	targetFieldDesc protoreflect.FieldDescriptor,
	forItems bool,
) (expected protoreflect.FieldDescriptor, ok bool) {
	switch {
	case targetFieldDesc.IsMap():
		return mapFieldConstraintsDesc, true
	case targetFieldDesc.IsList() && !forItems:
		return repeatedFieldConstraintsDesc, true
	case targetFieldDesc.Kind() == protoreflect.MessageKind:
		expected, ok = expectedWKTConstraints[targetFieldDesc.Message().FullName()]
		return expected, ok
	default:
		expected, ok = expectedStandardConstraints[targetFieldDesc.Kind()]
		return expected, ok
	}
}

// getCELType resolves the CEL value type for the provided FieldDescriptor. If
// forItems is true, the type for the repeated list items is returned instead of
// the list type itself.
func (c *Cache) getCELType(fieldDesc protoreflect.FieldDescriptor, forItems bool) *cel.Type {
	if !forItems {
		switch {
		case fieldDesc.IsMap():
			return cel.MapType(
				c.getCELType(fieldDesc.MapKey(), true),
				c.getCELType(fieldDesc.MapValue(), true),
			)
		case fieldDesc.IsList():
			return cel.ListType(c.getCELType(fieldDesc, true))
		}
	}

	if fieldDesc.Kind() == protoreflect.MessageKind {
		switch fqn := fieldDesc.Message().FullName(); fqn {
		case "google.protobuf.Any":
			return cel.AnyType
		case "google.protobuf.Duration":
			return cel.DurationType
		case "google.protobuf.Timestamp":
			return cel.TimestampType
		default:
			return cel.ObjectType(string(fqn))
		}
	}
	return ProtoKindToCELType(fieldDesc.Kind())
}
