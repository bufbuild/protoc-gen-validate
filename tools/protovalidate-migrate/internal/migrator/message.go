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

package migrator

import (
	"github.com/bufbuild/protocompile/ast"
)

type MessageVisitor struct {
	emittedDisabled bool
	PrinterVisitor
}

func (v *MessageVisitor) VisitOptionNode(node *ast.OptionNode) error {
	name := node.Name.Parts[0].Value()
	isPGV := name == "(validate.disabled)" || name == "(validate.ignored)"
	isPV := name == "(buf.validate.message)"

	if v.ShouldPrintOriginal(isPGV, isPV) {
		if isPV && v.emittedDisabled {
			return nil
		} else if isPV {
			v.emittedDisabled = true
		}
		if err := v.PrintNodes(true, node); err != nil {
			return err
		}
	}

	if !v.emittedDisabled && v.ShouldPrintReplacement(isPGV) {
		v.emittedDisabled = true
		return v.PrintNodes(
			false,
			node.Keyword,
			v.replaceNode(node.Name, "(buf.validate.message).disabled"),
			node.Equals,
			node.Val,
			node.Semicolon,
		)
	}

	return nil
}

func (v *MessageVisitor) VisitMessageNode(node *ast.MessageNode) error {
	return ast.VisitChildren(node, &MessageVisitor{PrinterVisitor: v.PrinterVisitor})
}

func (v *MessageVisitor) VisitFieldNode(node *ast.FieldNode) error {
	return ast.VisitChildren(node, &FieldVisitor{PrinterVisitor: v.PrinterVisitor})
}

func (v *MessageVisitor) VisitMapFieldNode(node *ast.MapFieldNode) error {
	return ast.VisitChildren(node, &FieldVisitor{PrinterVisitor: v.PrinterVisitor})
}

func (v *MessageVisitor) VisitOneOfNode(node *ast.OneOfNode) error {
	return ast.VisitChildren(node, &OneOfVisitor{PrinterVisitor: v.PrinterVisitor})
}
