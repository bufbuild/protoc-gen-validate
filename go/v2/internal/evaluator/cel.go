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
	"github.com/bufbuild/protovalidate/go/v2/internal/expression"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// celPrograms is an evaluator that executes an expression.ProgramSet.
type celPrograms expression.ProgramSet

func (c celPrograms) Evaluate(val protoreflect.Value, failFast bool) error {
	return expression.ProgramSet(c).Eval(val.Interface(), failFast)
}

func (c celPrograms) EvaluateMessage(msg protoreflect.Message, failFast bool) error {
	return expression.ProgramSet(c).Eval(msg.Interface(), failFast)
}

func (c celPrograms) Tautology() bool {
	return len(c) == 0
}

var (
	_ evaluator        = (celPrograms)(nil)
	_ MessageEvaluator = (celPrograms)(nil)
)
