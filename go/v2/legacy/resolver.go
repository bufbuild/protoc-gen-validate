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
	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	pv "github.com/bufbuild/protovalidate/go/v2"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// Mode determines the behavior of how the validator uses the legacy constraints
// relative to existing standard constraints.
type Mode uint8

const (
	// ModeMerge merges any existing standard constraints into the translated
	// legacy constraints. This mode uses proto.Merge to combine the results. This
	// is the default behavior.
	ModeMerge Mode = iota
	// ModeReplace replaces all existing standard constraints on the message,
	// oneof, or field with the translated legacy constraints.
	ModeReplace
	// ModeIfNotPresent uses the translated legacy constraints only if no standard
	// constraints are present on the message, oneof, or field.
	ModeIfNotPresent
)

// WithLegacySupport provides a protovalidate.ValidatorOption that transparently
// upgrades legacy constraints defined for protoc-gen-validate to be compatible
// with protovalidate. The specified mode determines the behavior of how the
// validator uses the legacy constraints.
func WithLegacySupport(mode Mode) pv.ValidatorOption {
	interceptor := migrateLegacyConstraintsInterceptor(mode)
	return pv.WithStandardConstraintInterceptor(interceptor)
}

func migrateLegacyConstraintsInterceptor(mode Mode) pv.StandardConstraintInterceptor {
	return func(res pv.StandardConstraintResolver) pv.StandardConstraintResolver {
		return legacyConstraintResolver{
			StandardConstraintResolver: res,
			mode:                       mode,
		}
	}
}

type legacyConstraintResolver struct {
	pv.StandardConstraintResolver
	mode Mode
}

func (l legacyConstraintResolver) ResolveMessageConstraints(
	desc protoreflect.MessageDescriptor,
) *validate.MessageConstraints {
	return resolveConstraints(
		translateMessageOptions,
		l.StandardConstraintResolver.ResolveMessageConstraints,
		l.mode,
		desc,
	)
}

func (l legacyConstraintResolver) ResolveOneofConstraints(
	desc protoreflect.OneofDescriptor,
) *validate.OneofConstraints {
	return resolveConstraints(
		translateOneofOptions,
		l.StandardConstraintResolver.ResolveOneofConstraints,
		l.mode,
		desc,
	)
}

func (l legacyConstraintResolver) ResolveFieldConstraints(
	desc protoreflect.FieldDescriptor,
) *validate.FieldConstraints {
	return resolveConstraints(
		translateFieldOptions,
		l.StandardConstraintResolver.ResolveFieldConstraints,
		l.mode,
		desc,
	)
}

func resolveConstraints[
	D protoreflect.Descriptor,
	C proto.Message,
](
	translate func(D) C,
	source func(D) C,
	mode Mode,
	desc D,
) C {
	switch mode {
	case ModeReplace:
		return translate(desc)
	case ModeIfNotPresent:
		constraints := source(desc)
		if constraints.ProtoReflect().IsValid() {
			return constraints
		}
		return translate(desc)
	case ModeMerge:
		fallthrough
	default:
		constraints := translate(desc)
		existing := source(desc)
		if !constraints.ProtoReflect().IsValid() {
			return existing
		}
		proto.Merge(constraints, source(desc))
		return constraints
	}
}

var _ pv.StandardConstraintResolver = legacyConstraintResolver{}
