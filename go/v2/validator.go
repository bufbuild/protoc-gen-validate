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
	"fmt"

	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	"github.com/bufbuild/protovalidate/go/v2/internal/celext"
	"github.com/bufbuild/protovalidate/go/v2/internal/errors"
	"github.com/bufbuild/protovalidate/go/v2/internal/evaluator"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type (
	// A ValidationError is returned if one or more constraints on a message are
	// violated. This error type can be converted into a validate.Violations
	// message via ToProto.
	//
	//    err = validator.Validate(msg)
	//    var valErr *ValidationError
	//    if ok := errors.As(err, &valErr); ok {
	//      pb := valErr.ToProto()
	//      // ...
	//    }
	ValidationError = errors.ValidationError

	// A CompilationError is returned if a CEL expression cannot be compiled &
	// type-checked or if invalid standard constraints are applied to a field.
	CompilationError = errors.CompilationError

	// A RuntimeError is returned if a valid CEL expression evaluation is
	// terminated, typically due to an unknown or mismatched type.
	RuntimeError = errors.RuntimeError
)

// Validator performs validation on any proto.Message values. The Validator is
// safe for concurrent use.
type Validator struct {
	builder  *evaluator.Builder
	failFast bool
}

// New creates a Validator with the given options. An error may occur in setting
// up the CEL execution environment if the configuration is invalid. See the
// individual ValidatorOption for how they impact the fallibility of New.
func New(options ...ValidatorOption) (*Validator, error) {
	cfg := config{resolver: evaluator.DefaultResolver{}}
	for _, opt := range options {
		opt(&cfg)
	}

	env, err := celext.DefaultEnv(cfg.useUTC)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to construct CEL environment: %w", err)
	}

	bldr := evaluator.NewBuilder(
		env,
		cfg.disableLazy,
		cfg.resolver,
		cfg.desc...,
	)

	return &Validator{
		failFast: cfg.failFast,
		builder:  bldr,
	}, nil
}

// Validate checks that message satisfies its constraints. Constraints are
// defined within the Protobuf file as options from the buf.validate package.
// An error is returned if the constraints are violated (ValidationError), the
// evaluation logic for the message cannot be built (CompilationError), or
// there is a type error when attempting to evaluate a CEL expression
// associated with the message (RuntimeError).
func (v *Validator) Validate(msg proto.Message) error {
	if msg == nil {
		return nil
	}
	refl := msg.ProtoReflect()
	eval := v.builder.Load(refl.Descriptor())
	return eval.EvaluateMessage(refl, v.failFast)
}

type config struct {
	failFast    bool
	useUTC      bool
	disableLazy bool
	desc        []protoreflect.MessageDescriptor
	resolver    StandardConstraintResolver
}

// A ValidatorOption modifies the default configuration of a Validator. See the
// individual options for their defaults and affects on the fallibility of
// configuring a Validator.
type ValidatorOption func(*config)

// WithUTC specifies whether timestamp operations should use UTC or the OS's
// local timezone for timestamp related values. By default, the local timezone
// is used.
func WithUTC(useUTC bool) ValidatorOption {
	return func(c *config) {
		c.useUTC = useUTC
	}
}

// WithFailFast specifies whether validation should fail on the first constraint
// violation encountered or if all violations should be accumulated. By default,
// all violations are accumulated.
func WithFailFast(failFast bool) ValidatorOption {
	return func(cfg *config) {
		cfg.failFast = failFast
	}
}

// WithMessages allows warming up the Validator with messages that are
// expected to be validated. Messages included transitively (i.e., fields with
// message values) are automatically handled.
func WithMessages(messages ...proto.Message) ValidatorOption {
	desc := make([]protoreflect.MessageDescriptor, len(messages))
	for i, msg := range messages {
		desc[i] = msg.ProtoReflect().Descriptor()
	}
	return WithDescriptors(desc...)
}

// WithDescriptors allows warming up the Validator with message
// descriptors that are expected to be validated. Messages included transitively
// (i.e., fields with message values) are automatically handled.
func WithDescriptors(descriptors ...protoreflect.MessageDescriptor) ValidatorOption {
	return func(cfg *config) {
		cfg.desc = append(cfg.desc, descriptors...)
	}
}

// WithDisableLazy prevents the Validator from lazily building validation logic
// for a message it has not encountered before. Disabling lazy logic
// additionally eliminates any internal locking as the validator becomes
// read-only.
//
// Note: All expected messages must be provided by WithMessages or
// WithDescriptors during initialization.
func WithDisableLazy(disable bool) ValidatorOption {
	return func(cfg *config) {
		cfg.disableLazy = disable
	}
}

// StandardConstraintResolver is responsible for resolving the standard
// constraints from the provided protoreflect.Descriptor. The default resolver
// can be intercepted and modified using WithStandardConstraintInterceptor.
type StandardConstraintResolver interface {
	ResolveMessageConstraints(desc protoreflect.MessageDescriptor) *validate.MessageConstraints
	ResolveOneofConstraints(desc protoreflect.OneofDescriptor) *validate.OneofConstraints
	ResolveFieldConstraints(desc protoreflect.FieldDescriptor) *validate.FieldConstraints
}

// StandardConstraintInterceptor can be provided to
// WithStandardConstraintInterceptor to allow modifying a
// StandardConstraintResolver.
type StandardConstraintInterceptor func(res StandardConstraintResolver) StandardConstraintResolver

// WithStandardConstraintInterceptor allows intercepting the
// StandardConstraintResolver used by the Validator to modify or replace it.
func WithStandardConstraintInterceptor(interceptor StandardConstraintInterceptor) ValidatorOption {
	return func(c *config) {
		c.resolver = interceptor(c.resolver)
	}
}
