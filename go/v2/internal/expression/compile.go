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
	"github.com/bufbuild/protovalidate/go/v2/internal/errors"
	"github.com/google/cel-go/cel"
)

// Expression is the read-only interface of either validate.Constraint or
// private.Constraint which can be the source of a CEL expression.
type Expression interface {
	GetId() string
	GetMessage() string
	GetExpression() string
}

// Compile produces a ProgramSet from the provided expressions in the given
// environment. If the generated cel.Program require cel.ProgramOption params,
// use CompileASTs instead with a subsequent call to ASTSet.ToProgramSet.
func Compile[T Expression](
	expressions []T,
	env *cel.Env,
	envOpts ...cel.EnvOption,
) (set ProgramSet, err error) {
	if len(expressions) == 0 {
		return nil, nil
	}

	if len(envOpts) > 0 {
		env, err = env.Extend(envOpts...)
		if err != nil {
			return nil, errors.NewCompilationErrorf(
				"failed to extend environment: %w", err)
		}
	}

	set = make(ProgramSet, len(expressions))
	for i, expr := range expressions {
		set[i].Source = expr

		ast, err := compileAST(env, expr)
		if err != nil {
			return nil, err
		}

		set[i], err = ast.toProgram(env)
		if err != nil {
			return nil, err
		}
	}
	return set, nil
}

// CompileASTs parses and type checks a set of expressions, producing a resulting
// ASTSet. The value can then be converted to a ProgramSet via
// ASTSet.ToProgramSet or ASTSet.ReduceResiduals. Use Compile instead if no
// cel.ProgramOption args need to be provided or residuals do not need to be
// computed.
func CompileASTs[T Expression](
	expressions []T,
	env *cel.Env,
	envOpts ...cel.EnvOption,
) (set ASTSet, err error) {
	set.env = env
	if len(expressions) == 0 {
		return set, nil
	}

	if len(envOpts) > 0 {
		set.env, err = env.Extend(envOpts...)
		if err != nil {
			return set, errors.NewCompilationErrorf(
				"failed to extend environment: %w", err)
		}
	}

	set.asts = make([]compiledAST, len(expressions))
	for i, expr := range expressions {
		set.asts[i], err = compileAST(set.env, expr)
		if err != nil {
			return set, err
		}
	}

	return set, nil
}

func compileAST(env *cel.Env, expr Expression) (out compiledAST, err error) {
	ast, issues := env.Compile(expr.GetExpression())
	if err := issues.Err(); err != nil {
		return out, errors.NewCompilationErrorf(
			"failed to compile expression %s: %w", expr.GetId(), err)
	}

	outType := ast.OutputType()
	if !(outType.IsAssignableType(cel.BoolType) || outType.IsAssignableType(cel.StringType)) {
		return out, errors.NewCompilationErrorf(
			"expression %s outputs %s, wanted either bool or string",
			expr.GetId(), outType.String())
	}

	return compiledAST{
		AST:    ast,
		Source: expr,
	}, nil
}
