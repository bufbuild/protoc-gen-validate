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

// anyPB is a specialized evaluator for applying validate.AnyRules to an
// anypb.Any message. This is handled outside CEL which attempts to
// hydrate anyPB's within an expression, breaking evaluation if the type is
// unknown at runtime.
type anyPB struct {
	// TypeURLDescriptor is the descriptor for the TypeURL field
	TypeURLDescriptor protoreflect.FieldDescriptor
	// In specifies which type URLs the value may possess
	In map[string]struct{}
	// NotIn specifies which type URLs the value may not possess
	NotIn map[string]struct{}
}

func (a anyPB) Evaluate(val protoreflect.Value, failFast bool) error {
	typeURL := val.Message().Get(a.TypeURLDescriptor).String()

	err := &errors.ValidationError{Violations: []*validate.Violation{}}
	if len(a.In) > 0 {
		if _, ok := a.In[typeURL]; !ok {
			err.Violations = append(err.Violations, &validate.Violation{
				ConstraintId: "any.in",
				Message:      "type URL must be in the allow list",
			})
			if failFast {
				return err
			}
		}
	}

	if len(a.NotIn) > 0 {
		if _, ok := a.NotIn[typeURL]; ok {
			err.Violations = append(err.Violations, &validate.Violation{
				ConstraintId: "any.not_in",
				Message:      "type URL must not be in the block list",
			})
		}
	}

	if len(err.Violations) > 0 {
		return err
	}
	return nil
}

func (a anyPB) Tautology() bool {
	return len(a.In) == 0 && len(a.NotIn) == 0
}

func stringsToSet(ss []string) map[string]struct{} {
	if len(ss) == 0 {
		return nil
	}
	set := make(map[string]struct{}, len(ss))
	for _, s := range ss {
		set[s] = struct{}{}
	}
	return set
}

var _ evaluator = anyPB{}
