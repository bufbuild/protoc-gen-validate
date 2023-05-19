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
	"sync"

	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	"github.com/bufbuild/protovalidate/go/v2/internal/constraints"
	"github.com/bufbuild/protovalidate/go/v2/internal/errors"
	"github.com/bufbuild/protovalidate/go/v2/internal/expression"
	"github.com/google/cel-go/cel"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/dynamicpb"
)

// Builder is a build-through cache of message evaluators keyed off the provided
// descriptor.
type Builder struct {
	// TODO: (TCN-1708) based on benchmarks, about 50% of CPU time is spent obtaining a read
	//  lock on this mutex. Ideally, this can be reworked to be thread-safe while
	//  minimizing the need to obtain a lock.
	mtx         sync.RWMutex
	cache       map[protoreflect.MessageDescriptor]*message
	env         *cel.Env
	constraints constraints.Cache
	resolver    StandardConstraintResolver
	Load        func(desc protoreflect.MessageDescriptor) MessageEvaluator
}

// NewBuilder initializes a new Builder.
func NewBuilder(
	env *cel.Env,
	disableLazy bool,
	res StandardConstraintResolver,
	seedDesc ...protoreflect.MessageDescriptor,
) *Builder {
	bldr := &Builder{
		cache:       map[protoreflect.MessageDescriptor]*message{},
		env:         env,
		constraints: constraints.NewCache(),
		resolver:    res,
	}

	if disableLazy {
		bldr.Load = bldr.load
	} else {
		bldr.Load = bldr.loadOrBuild
	}

	for _, desc := range seedDesc {
		bldr.build(desc)
	}

	return bldr
}

// load returns a pre-cached MessageEvaluator for the given descriptor or, if
// the descriptor is unknown, returns an evaluator that always resolves to a
// errors.CompilationError.
func (bldr *Builder) load(desc protoreflect.MessageDescriptor) MessageEvaluator {
	if eval, ok := bldr.cache[desc]; ok {
		return eval
	}
	return unknownMessage{desc: desc}
}

// loadOrBuild either returns a memoized MessageEvaluator for the given
// descriptor, or lazily constructs a new one. This method is thread-safe via
// locking.
func (bldr *Builder) loadOrBuild(desc protoreflect.MessageDescriptor) MessageEvaluator {
	bldr.mtx.RLock()
	if eval, ok := bldr.cache[desc]; ok {
		bldr.mtx.RUnlock()
		return eval
	}
	bldr.mtx.RUnlock()

	bldr.mtx.Lock()
	defer bldr.mtx.Unlock()
	return bldr.build(desc)
}

func (bldr *Builder) build(desc protoreflect.MessageDescriptor) *message {
	if eval, ok := bldr.cache[desc]; ok {
		return eval
	}
	msgEval := &message{}
	bldr.cache[desc] = msgEval
	bldr.buildMessage(desc, msgEval)
	return msgEval
}

func (bldr *Builder) buildMessage(desc protoreflect.MessageDescriptor, msgEval *message) {
	msgConstraints := bldr.resolver.ResolveMessageConstraints(desc)
	if msgConstraints.GetDisabled() {
		return
	}

	steps := []func(
		desc protoreflect.MessageDescriptor,
		msgConstraints *validate.MessageConstraints,
		msg *message,
	){
		bldr.processMessageExpressions,
		bldr.processOneofConstraints,
		bldr.processFields,
	}

	for _, step := range steps {
		if step(desc, msgConstraints, msgEval); msgEval.Err != nil {
			break
		}
	}
}

func (bldr *Builder) processMessageExpressions(
	desc protoreflect.MessageDescriptor,
	msgConstraints *validate.MessageConstraints,
	msgEval *message,
) {
	compiledExprs, err := expression.Compile(
		msgConstraints.GetCel(),
		bldr.env,
		cel.Types(dynamicpb.NewMessage(desc)),
		cel.Variable("this", cel.ObjectType(string(desc.FullName()))),
	)
	if err != nil {
		msgEval.Err = err
		return
	}

	msgEval.Append(celPrograms(compiledExprs))
}

func (bldr *Builder) processOneofConstraints(
	desc protoreflect.MessageDescriptor,
	_ *validate.MessageConstraints,
	msgEval *message,
) {
	oneofs := desc.Oneofs()
	for i := 0; i < oneofs.Len(); i++ {
		oneofDesc := oneofs.Get(i)
		oneofConstraints := bldr.resolver.ResolveOneofConstraints(oneofDesc)
		oneofEval := oneof{
			Descriptor: oneofDesc,
			Required:   oneofConstraints.GetRequired(),
		}
		msgEval.Append(oneofEval)
	}
}

func (bldr *Builder) processFields(
	desc protoreflect.MessageDescriptor,
	_ *validate.MessageConstraints,
	msgEval *message,
) {
	fields := desc.Fields()
	for i := 0; i < fields.Len(); i++ {
		fdesc := fields.Get(i)
		fieldConstraints := bldr.resolver.ResolveFieldConstraints(fdesc)
		fldEval, err := bldr.buildField(fdesc, fieldConstraints)
		if err != nil {
			msgEval.Err = err
			return
		}
		msgEval.Append(fldEval)
	}
}

func (bldr *Builder) buildField(
	fieldDescriptor protoreflect.FieldDescriptor,
	fieldConstraints *validate.FieldConstraints,
) (field, error) {
	fld := field{
		Descriptor: fieldDescriptor,
		Required:   fieldConstraints.GetRequired(),
		Optional:   fieldDescriptor.HasPresence(),
	}
	err := bldr.buildValue(fieldDescriptor, fieldConstraints, false, &fld.Value)
	return fld, err
}

func (bldr *Builder) buildValue(
	fdesc protoreflect.FieldDescriptor,
	constraints *validate.FieldConstraints,
	forItems bool,
	valEval *value,
) (err error) {
	valEval.IgnoreEmpty = constraints.GetIgnoreEmpty()
	steps := []func(
		fdesc protoreflect.FieldDescriptor,
		fieldConstraints *validate.FieldConstraints,
		forItems bool,
		valEval *value,
	) error{
		bldr.processZeroValue,
		bldr.processFieldExpressions,
		bldr.processEmbeddedMessage,
		bldr.processWrapperConstraints,
		bldr.processStandardConstraints,
		bldr.processAnyConstraints,
		bldr.processEnumConstraints,
		bldr.processMapConstraints,
		bldr.processRepeatedConstraints,
	}

	for _, step := range steps {
		if err = step(fdesc, constraints, forItems, valEval); err != nil {
			return err
		}
	}
	return nil
}

func (bldr *Builder) processZeroValue(
	fdesc protoreflect.FieldDescriptor,
	_ *validate.FieldConstraints,
	forItems bool,
	val *value,
) error {
	val.Zero = fdesc.Default()
	if forItems && fdesc.IsList() {
		msg := dynamicpb.NewMessage(fdesc.ContainingMessage())
		val.Zero = msg.Get(fdesc).List().NewElement()
	}
	return nil
}

func (bldr *Builder) processFieldExpressions(
	fieldDesc protoreflect.FieldDescriptor,
	fieldConstraints *validate.FieldConstraints,
	_ bool,
	eval *value,
) error {
	exprs := fieldConstraints.GetCel()
	if len(exprs) == 0 {
		return nil
	}
	var opts []cel.EnvOption
	if fieldDesc.Kind() == protoreflect.MessageKind {
		opts = []cel.EnvOption{
			cel.TypeDescs(fieldDesc.ParentFile()),
			cel.Variable("this", cel.ObjectType(string(fieldDesc.Message().FullName()))),
		}
	} else {
		opts = []cel.EnvOption{
			cel.Variable("this", constraints.ProtoKindToCELType(fieldDesc.Kind())),
		}
	}
	compiledExpressions, err := expression.Compile(exprs, bldr.env, opts...)
	if err != nil {
		return err
	}
	if len(compiledExpressions) > 0 {
		eval.Constraints = append(eval.Constraints, celPrograms(compiledExpressions))
	}
	return nil
}

func (bldr *Builder) processEmbeddedMessage(
	fdesc protoreflect.FieldDescriptor,
	rules *validate.FieldConstraints,
	forItems bool,
	valEval *value,
) error {
	if fdesc.Kind() != protoreflect.MessageKind ||
		rules.GetSkipped() ||
		fdesc.IsMap() || (fdesc.IsList() && !forItems) {
		return nil
	}

	embedEval := bldr.build(fdesc.Message())
	if err := embedEval.Err; err != nil {
		return errors.NewCompilationErrorf(
			"failed to compile embedded type %s for %s: %w",
			fdesc.Message().FullName(), fdesc.FullName(), err)
	}
	valEval.Append(embedEval)

	return nil
}

func (bldr *Builder) processWrapperConstraints(
	fdesc protoreflect.FieldDescriptor,
	rules *validate.FieldConstraints,
	forItems bool,
	valEval *value,
) error {
	if fdesc.Kind() != protoreflect.MessageKind ||
		rules.GetSkipped() ||
		fdesc.IsMap() || (fdesc.IsList() && !forItems) {
		return nil
	}

	expectedWrapperDescriptor, ok := constraints.ExpectedWrapperConstraints(fdesc.Message().FullName())
	if !ok || !rules.ProtoReflect().Has(expectedWrapperDescriptor) {
		return nil
	}
	var unwrapped value
	err := bldr.buildValue(fdesc.Message().Fields().ByName("value"), rules, true, &unwrapped)
	if err != nil {
		return err
	}
	valEval.Append(unwrapped.Constraints)
	return nil
}

func (bldr *Builder) processStandardConstraints(
	fdesc protoreflect.FieldDescriptor,
	constraints *validate.FieldConstraints,
	forItems bool,
	valEval *value,
) error {
	stdConstraints, err := bldr.constraints.Build(
		bldr.env,
		fdesc,
		constraints,
		forItems,
	)
	if err != nil {
		return err
	}
	valEval.Append(celPrograms(stdConstraints))
	return nil
}

func (bldr *Builder) processAnyConstraints(
	fdesc protoreflect.FieldDescriptor,
	fieldConstraints *validate.FieldConstraints,
	forItems bool,
	valEval *value,
) error {
	if (fdesc.IsList() && !forItems) ||
		fdesc.Kind() != protoreflect.MessageKind ||
		fdesc.Message().FullName() != "google.protobuf.Any" {
		return nil
	}

	typeURLDesc := fdesc.Message().Fields().ByName("type_url")
	anyEval := anyPB{
		TypeURLDescriptor: typeURLDesc,
		In:                stringsToSet(fieldConstraints.GetAny().GetIn()),
		NotIn:             stringsToSet(fieldConstraints.GetAny().GetNotIn()),
	}
	valEval.Append(anyEval)
	return nil
}

func (bldr *Builder) processEnumConstraints(
	fdesc protoreflect.FieldDescriptor,
	fieldConstraints *validate.FieldConstraints,
	_ bool,
	valEval *value,
) error {
	if fdesc.Kind() != protoreflect.EnumKind {
		return nil
	}
	if fieldConstraints.GetEnum().GetDefinedOnly() {
		valEval.Append(definedEnum{ValueDescriptors: fdesc.Enum().Values()})
	}
	return nil
}

func (bldr *Builder) processMapConstraints(
	fieldDesc protoreflect.FieldDescriptor,
	constraints *validate.FieldConstraints,
	_ bool,
	valEval *value,
) error {
	if !fieldDesc.IsMap() {
		return nil
	}

	var mapEval kvPairs

	err := bldr.buildValue(
		fieldDesc.MapKey(),
		constraints.GetMap().GetKeys(),
		true,
		&mapEval.KeyConstraints)
	if err != nil {
		return errors.NewCompilationErrorf(
			"failed to compile key constraints for map %s: %w",
			fieldDesc.FullName(), err)
	}

	err = bldr.buildValue(
		fieldDesc.MapValue(),
		constraints.GetMap().GetValues(),
		true,
		&mapEval.ValueConstraints)
	if err != nil {
		return errors.NewCompilationErrorf(
			"failed to compile value constraints for map %s: %w",
			fieldDesc.FullName(), err)
	}

	valEval.Append(mapEval)
	return nil
}

func (bldr *Builder) processRepeatedConstraints(
	fdesc protoreflect.FieldDescriptor,
	fieldConstraints *validate.FieldConstraints,
	forItems bool,
	valEval *value,
) error {
	if !fdesc.IsList() || forItems {
		return nil
	}

	var listEval listItems
	err := bldr.buildValue(fdesc, fieldConstraints.GetRepeated().GetItems(), true, &listEval.ItemConstraints)
	if err != nil {
		return errors.NewCompilationErrorf(
			"failed to compile items constraints for repeated %v: %w", fdesc.FullName(), err)
	}

	valEval.Append(listEval)
	return nil
}
