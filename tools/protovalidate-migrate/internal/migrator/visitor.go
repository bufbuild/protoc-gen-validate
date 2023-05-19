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

func New(cfg Config, fileNode *ast.FileNode, out io.Writer) ast.Visitor {
	return &RootVisitor{
		printer: PrinterVisitor{
			Config: cfg,
			file:   fileNode,
			w:      out,
		},
	}
}

type RootVisitor struct {
	ast.NoOpVisitor
	printer PrinterVisitor
}

func (v *RootVisitor) VisitFileNode(node *ast.FileNode) error {
	for _, decl := range node.Decls {
		imp, ok := decl.(*ast.ImportNode)
		if !ok {
			continue
		}
		v.printer.HasPGV = v.printer.HasPGV || imp.Name.AsString() == v.printer.PGVImport
		v.printer.HasPV = v.printer.HasPV || imp.Name.AsString() == v.printer.PVImport
	}

	switch {
	case !v.printer.HasPGV:
		// nothing to remove or replace
		return v.printer.PrintNodes(true, node)
	case v.printer.HasPV && !v.printer.RemovePGV && !v.printer.ReplacePV:
		// not removing legacy and not replacing existing
		return v.printer.PrintNodes(true, node)
	default:
		// in all other cases we have work to do
		return ast.VisitChildren(node, &FileVisitor{PrinterVisitor: v.printer})
	}
}

type element interface {
	LeadingWhitespace() string
	RawText() string
}

type nodeInfo interface {
	element
	LeadingComments() ast.Comments
	TrailingComments() ast.Comments
}

type wrappedNode interface {
	ToNodeInfo() nodeInfo
}

type replacementNode struct {
	ast.Node
	info replacementNodeInfo
}

func (r replacementNode) ToNodeInfo() nodeInfo {
	return r.info
}

type replacementNodeInfo struct {
	ast.NodeInfo
	Replacement string
}

func (r replacementNodeInfo) RawText() string {
	return r.Replacement
}

type unwrappedMessageLiteral struct {
	*ast.MessageLiteralNode
}

func (u *unwrappedMessageLiteral) Children() []ast.Node {
	allChildren := u.MessageLiteralNode.Children()
	return allChildren[1 : len(allChildren)-1]
}
