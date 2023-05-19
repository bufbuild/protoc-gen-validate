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
	"fmt"

	"github.com/bufbuild/protovalidate/go/v2/internal/errors"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// listItems performs validation on the elements of a repeated field.
type listItems struct {
	// ItemConstraints are checked on every item of the list
	ItemConstraints value
}

func (r listItems) Evaluate(val protoreflect.Value, failFast bool) error {
	list := val.List()
	var ok bool
	var err error
	for i := 0; i < list.Len(); i++ {
		itemErr := r.ItemConstraints.Evaluate(list.Get(i), failFast)
		if itemErr != nil {
			errors.PrefixErrorPaths(itemErr, fmt.Sprintf("[%d]", i))
		}
		if ok, err = errors.Merge(err, itemErr, failFast); !ok {
			return err
		}
	}
	return err
}

func (r listItems) Tautology() bool {
	return r.ItemConstraints.Tautology()
}

var _ evaluator = listItems{}
