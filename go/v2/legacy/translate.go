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
	pv "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	pgv "github.com/envoyproxy/protoc-gen-validate/validate"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// translateMessageOptions converts legacy protoc-gen-validate constraints off
// the provided message descriptor into protovalidate.MessageConstraints.
func translateMessageOptions(desc protoreflect.MessageDescriptor) *pv.MessageConstraints {
	disabled, _ := proto.GetExtension(desc.Options(), pgv.E_Disabled).(bool)
	if disabled {
		return &pv.MessageConstraints{
			Disabled: proto.Bool(true),
		}
	}
	ignored, _ := proto.GetExtension(desc.Options(), pgv.E_Ignored).(bool)
	if ignored {
		return &pv.MessageConstraints{
			Disabled: proto.Bool(true),
		}
	}
	return nil
}

// translateOneofOptions converts legacy protoc-gen-validate constraints off the
// provided oneof descriptor into protovalidate.OneofConstraints.
func translateOneofOptions(desc protoreflect.OneofDescriptor) *pv.OneofConstraints {
	required, _ := proto.GetExtension(desc.Options(), pgv.E_Required).(bool)
	if required {
		return &pv.OneofConstraints{
			Required: proto.Bool(true),
		}
	}
	return nil
}

// translateFieldOptions converts legacy protoc-gen-validate constraints off the
// provided field descriptor into protovalidate.FieldConstraints.
func translateFieldOptions(desc protoreflect.FieldDescriptor) *pv.FieldConstraints {
	if !proto.HasExtension(desc.Options(), pgv.E_Rules) {
		return nil
	}
	rules := desc.Options().ProtoReflect().Get(pgv.E_Rules.TypeDescriptor()).Message()
	return translateRules(rules)
}

func translateRules(rules protoreflect.Message) *pv.FieldConstraints {
	desc := rules.Descriptor()
	constraints := &pv.FieldConstraints{}

	if msgRules := rules.Get(desc.Fields().ByName("message")).Message(); msgRules.IsValid() {
		msgDesc := msgRules.Descriptor()
		if msgRules.Get(msgDesc.Fields().ByName("required")).Bool() {
			constraints.Required = true
		} else if msgRules.Get(msgDesc.Fields().ByName("skip")).Bool() {
			constraints.Skipped = true
		}
	}

	rulesOneof := rules.WhichOneof(rules.Descriptor().Oneofs().ByName("type"))
	if rulesOneof == nil {
		return constraints
	}
	fldRules := rules.Get(rulesOneof).Message()

	constraintsRefl := constraints.ProtoReflect()
	constraintsOneof := constraintsRefl.Descriptor().Fields().ByName(rulesOneof.Name())
	if constraintsOneof == nil {
		return constraints
	}
	fldConstraints := constraintsRefl.Mutable(constraintsOneof).Message()

	fldRules.Range(func(pgvDesc protoreflect.FieldDescriptor, value protoreflect.Value) bool {
		return translateRule(pgvDesc, value, constraints, fldConstraints)
	})
	return constraints
}

func translateRule(
	pgvDesc protoreflect.FieldDescriptor,
	value protoreflect.Value,
	constraints *pv.FieldConstraints,
	typConstraints protoreflect.Message,
) bool {
	if pgvDesc.Name() == "required" && pgvDesc.Kind() == protoreflect.BoolKind && value.Bool() {
		// old `required` fields on the WKTs need to be lifted to the top level
		constraints.Required = true
		return true
	} else if pgvDesc.Name() == "ignore_empty" && pgvDesc.Kind() == protoreflect.BoolKind && value.Bool() {
		// old `ignore_empty` fields on the type rules need to be lifted to the top level
		constraints.IgnoreEmpty = true
		return true
	}

	pvDesc := typConstraints.Descriptor().Fields().ByName(pgvDesc.Name())
	if pvDesc == nil ||
		pgvDesc.Kind() != pvDesc.Kind() ||
		pgvDesc.IsList() != pvDesc.IsList() ||
		pgvDesc.IsMap() != pvDesc.IsMap() {
		// removed field or mismatched shapes
		return true
	}

	switch {
	case pgvDesc.IsMap():
		// noop, there should be no map fields in PGV
	case pgvDesc.IsList():
		// only repeated fields are either scalar types or WKTs
		typConstraints.Set(pvDesc, value)
	case pgvDesc.Kind() == protoreflect.MessageKind &&
		pgvDesc.Message().FullName() == "validate.FieldRules":
		// recurse to translate RepeatedRules.items & MapRules.{keys,values}
		cons := translateRules(value.Message())
		typConstraints.Set(pvDesc, protoreflect.ValueOfMessage(cons.ProtoReflect()))
	case pgvDesc.Kind() == protoreflect.EnumKind:
		// try to use the same name, otherwise use the same number
		if pgvValDesc := pgvDesc.Enum().Values().ByNumber(value.Enum()); pgvValDesc == nil {
			typConstraints.Set(pvDesc, value)
		} else if pvValDesc := pvDesc.Enum().Values().ByName(pgvValDesc.Name()); pvValDesc == nil {
			typConstraints.Set(pvDesc, value)
		} else {
			typConstraints.Set(pvDesc, protoreflect.ValueOfEnum(pvValDesc.Number()))
		}
	default:
		typConstraints.Set(pvDesc, value)
	}
	return true
}
