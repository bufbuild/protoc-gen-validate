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
	"testing"

	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	"github.com/bufbuild/protovalidate/go/v2/internal/celext"
	"github.com/google/cel-go/cel"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestASTSet_Merge(t *testing.T) {
	t.Parallel()

	var set ASTSet
	other := ASTSet{
		env: &cel.Env{},
		asts: []compiledAST{
			{AST: &cel.Ast{}},
			{AST: &cel.Ast{}},
		},
	}
	merged := set.Merge(other)
	assert.Equal(t, other.env, merged.env)
	assert.Equal(t, other.asts, merged.asts)

	another := ASTSet{
		asts: []compiledAST{
			{AST: &cel.Ast{}},
			{AST: &cel.Ast{}},
			{AST: &cel.Ast{}},
		},
	}
	merged = other.Merge(another)
	assert.Equal(t, other.env, merged.env)
	assert.Equal(t, other.asts, merged.asts[0:2])
	assert.Equal(t, another.asts, merged.asts[2:])
}

func TestASTSet_ToProgramSet(t *testing.T) {
	t.Parallel()

	env, err := celext.DefaultEnv(false)
	require.NoError(t, err)

	expr := &validate.Constraint{Expression: "foo"}
	asts, err := CompileASTs([]*validate.Constraint{expr}, env,
		cel.Variable("foo", cel.BoolType))
	require.NoError(t, err)
	assert.Len(t, asts.asts, 1)
	set, err := asts.ToProgramSet()
	require.NoError(t, err)
	assert.Len(t, set, 1)
	assert.Equal(t, asts.asts[0].Source, set[0].Source)

	empty := ASTSet{}
	set, err = empty.ToProgramSet()
	assert.Empty(t, set)
	assert.NoError(t, err)
}

func TestASTSet_ReduceResiduals(t *testing.T) {
	t.Parallel()

	env, err := celext.DefaultEnv(false)
	require.NoError(t, err)

	expr := &validate.Constraint{Expression: "foo"}
	asts, err := CompileASTs([]*validate.Constraint{expr}, env,
		cel.Variable("foo", cel.BoolType))
	require.NoError(t, err)
	assert.Len(t, asts.asts, 1)
	set, err := asts.ReduceResiduals(cel.Globals(&Variable{Name: "foo", Val: true}))
	assert.NoError(t, err)
	assert.Empty(t, set)
}
