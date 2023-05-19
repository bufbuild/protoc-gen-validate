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
	"io"

	"github.com/bufbuild/protocompile/ast"
)

type PrinterVisitor struct {
	Config
	HasPGV bool
	HasPV  bool

	file *ast.FileNode
	w    io.Writer
}

func (v PrinterVisitor) PrintNodes(printComments bool, nodes ...ast.Node) error {
	for _, node := range nodes {
		var nodeInfo nodeInfo
		if infoer, ok := node.(wrappedNode); ok {
			nodeInfo = infoer.ToNodeInfo()
		} else {
			nodeInfo = v.file.NodeInfo(node)
		}
		if printComments {
			err := v.PrintComments(nodeInfo.LeadingComments())
			if err != nil {
				return err
			}
		}
		err := v.PrintElement(nodeInfo)
		if err != nil {
			return err
		}
		if printComments {
			err = v.PrintComments(nodeInfo.TrailingComments())
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (v PrinterVisitor) PrintComments(comments ast.Comments) error {
	for i, n := 0, comments.Len(); i < n; i++ {
		cmt := comments.Index(i)
		if err := v.PrintElement(cmt); err != nil {
			return err
		}
	}
	return nil
}

func (v PrinterVisitor) PrintElement(n element) error {
	_, err := v.w.Write([]byte(n.LeadingWhitespace()))
	if err != nil {
		return err
	}
	_, err = v.w.Write([]byte(n.RawText()))
	return err
}

func (v PrinterVisitor) ShouldPrintOriginal(isPGV, isPV bool) bool {
	return (isPGV && !v.RemovePGV) ||
		(isPV && !v.ReplacePV) ||
		(!isPGV && !isPV)
}

func (v PrinterVisitor) ShouldPrintReplacement(isPGV bool) bool {
	return isPGV && (v.ReplacePV || !v.HasPV)
}

func (v PrinterVisitor) replaceNode(original ast.Node, replacement string) replacementNode {
	return replacementNode{
		Node: original,
		info: replacementNodeInfo{
			NodeInfo:    v.file.NodeInfo(original),
			Replacement: replacement,
		},
	}
}

func (v PrinterVisitor) VisitFileNode(node *ast.FileNode) error {
	return v.PrintNodes(true, node)
}

func (v PrinterVisitor) VisitSyntaxNode(node *ast.SyntaxNode) error {
	return v.PrintNodes(true, node)
}

func (v PrinterVisitor) VisitPackageNode(node *ast.PackageNode) error {
	return v.PrintNodes(true, node)
}

func (v PrinterVisitor) VisitImportNode(node *ast.ImportNode) error {
	return v.PrintNodes(true, node)
}

func (v PrinterVisitor) VisitOptionNode(node *ast.OptionNode) error {
	return v.PrintNodes(true, node)
}

func (v PrinterVisitor) VisitOptionNameNode(node *ast.OptionNameNode) error {
	return v.PrintNodes(true, node)
}

func (v PrinterVisitor) VisitFieldReferenceNode(node *ast.FieldReferenceNode) error {
	return v.PrintNodes(true, node)
}

func (v PrinterVisitor) VisitCompactOptionsNode(node *ast.CompactOptionsNode) error {
	return v.PrintNodes(true, node)
}

func (v PrinterVisitor) VisitMessageNode(node *ast.MessageNode) error {
	return v.PrintNodes(true, node)
}

func (v PrinterVisitor) VisitExtendNode(node *ast.ExtendNode) error {
	return v.PrintNodes(true, node)
}

func (v PrinterVisitor) VisitExtensionRangeNode(node *ast.ExtensionRangeNode) error {
	return v.PrintNodes(true, node)
}

func (v PrinterVisitor) VisitReservedNode(node *ast.ReservedNode) error {
	return v.PrintNodes(true, node)
}

func (v PrinterVisitor) VisitRangeNode(node *ast.RangeNode) error {
	return v.PrintNodes(true, node)
}

func (v PrinterVisitor) VisitFieldNode(node *ast.FieldNode) error {
	return v.PrintNodes(true, node)
}

func (v PrinterVisitor) VisitGroupNode(node *ast.GroupNode) error {
	return v.PrintNodes(true, node)
}

func (v PrinterVisitor) VisitMapFieldNode(node *ast.MapFieldNode) error {
	return v.PrintNodes(true, node)
}

func (v PrinterVisitor) VisitMapTypeNode(node *ast.MapTypeNode) error {
	return v.PrintNodes(true, node)
}

func (v PrinterVisitor) VisitOneOfNode(node *ast.OneOfNode) error {
	return v.PrintNodes(true, node)
}

func (v PrinterVisitor) VisitEnumNode(node *ast.EnumNode) error {
	return v.PrintNodes(true, node)
}

func (v PrinterVisitor) VisitEnumValueNode(node *ast.EnumValueNode) error {
	return v.PrintNodes(true, node)
}

func (v PrinterVisitor) VisitServiceNode(node *ast.ServiceNode) error {
	return v.PrintNodes(true, node)
}

func (v PrinterVisitor) VisitRPCNode(node *ast.RPCNode) error {
	return v.PrintNodes(true, node)
}

func (v PrinterVisitor) VisitRPCTypeNode(node *ast.RPCTypeNode) error {
	return v.PrintNodes(true, node)
}

func (v PrinterVisitor) VisitIdentNode(node *ast.IdentNode) error {
	return v.PrintNodes(true, node)
}

func (v PrinterVisitor) VisitCompoundIdentNode(node *ast.CompoundIdentNode) error {
	return v.PrintNodes(true, node)
}

func (v PrinterVisitor) VisitStringLiteralNode(node *ast.StringLiteralNode) error {
	return v.PrintNodes(true, node)
}

func (v PrinterVisitor) VisitCompoundStringLiteralNode(node *ast.CompoundStringLiteralNode) error {
	return v.PrintNodes(true, node)
}

func (v PrinterVisitor) VisitUintLiteralNode(node *ast.UintLiteralNode) error {
	return v.PrintNodes(true, node)
}

func (v PrinterVisitor) VisitPositiveUintLiteralNode(node *ast.PositiveUintLiteralNode) error {
	return v.PrintNodes(true, node)
}

func (v PrinterVisitor) VisitNegativeIntLiteralNode(node *ast.NegativeIntLiteralNode) error {
	return v.PrintNodes(true, node)
}

func (v PrinterVisitor) VisitFloatLiteralNode(node *ast.FloatLiteralNode) error {
	return v.PrintNodes(true, node)
}

func (v PrinterVisitor) VisitSpecialFloatLiteralNode(node *ast.SpecialFloatLiteralNode) error {
	return v.PrintNodes(true, node)
}

func (v PrinterVisitor) VisitSignedFloatLiteralNode(node *ast.SignedFloatLiteralNode) error {
	return v.PrintNodes(true, node)
}

func (v PrinterVisitor) VisitArrayLiteralNode(node *ast.ArrayLiteralNode) error {
	return v.PrintNodes(true, node)
}

func (v PrinterVisitor) VisitMessageLiteralNode(node *ast.MessageLiteralNode) error {
	return v.PrintNodes(true, node)
}

func (v PrinterVisitor) VisitMessageFieldNode(node *ast.MessageFieldNode) error {
	return v.PrintNodes(true, node)
}

func (v PrinterVisitor) VisitKeywordNode(node *ast.KeywordNode) error {
	return v.PrintNodes(true, node)
}

func (v PrinterVisitor) VisitRuneNode(node *ast.RuneNode) error {
	return v.PrintNodes(true, node)
}

func (v PrinterVisitor) VisitEmptyDeclNode(node *ast.EmptyDeclNode) error {
	return v.PrintNodes(true, node)
}
