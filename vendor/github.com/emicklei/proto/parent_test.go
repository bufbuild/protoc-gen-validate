// Copyright (c) 2018 Ernest Micklei
//
// MIT License
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
package proto

import (
	"fmt"
	"testing"
)

type parentChecker struct {
	errors []error
}

func checkParent(v Visitee, t *testing.T) {
	pc := new(parentChecker)
	v.Accept(pc)
	if len(pc.errors) == 0 {
		return
	}
	for _, each := range pc.errors {
		t.Error(each)
	}
}

func (pc *parentChecker) checkAll(list []Visitee, parent Visitee) {
	for _, each := range list {
		if _, ok := each.(*Comment); ok {
			continue
		}
		if got, want := getParent(each), parent; got != want {
			pc.errors = append(pc.errors, fmt.Errorf("%T has wrong parent set, got %v want %v", each, got, want))
		}
		each.Accept(pc)
	}
}
func (pc *parentChecker) check(astType, astName string, parent Visitee) {
	if parent == nil {
		pc.errors = append(pc.errors, fmt.Errorf("%s %s has no parent set", astType, astName))
	}
}

func (pc *parentChecker) VisitProto(p *Proto) {
	pc.checkAll(p.Elements, p)
}
func (pc *parentChecker) VisitMessage(m *Message) {
	pc.check("Message", m.Name, m.Parent)
	pc.checkAll(m.Elements, m)
}
func (pc *parentChecker) VisitService(v *Service) {
	pc.check("Service", v.Name, v.Parent)
	pc.checkAll(v.Elements, v)
}
func (pc *parentChecker) VisitSyntax(s *Syntax) {
	pc.check("Syntax", s.Value, s.Parent)
}
func (pc *parentChecker) VisitPackage(p *Package) {
	pc.check("Package", p.Name, p.Parent)
}
func (pc *parentChecker) VisitOption(o *Option) {
	pc.check("Option", o.Name, o.Parent)
}
func (pc *parentChecker) VisitImport(i *Import) {
	pc.check("Import", i.Filename, i.Parent)
}
func (pc *parentChecker) VisitNormalField(i *NormalField) {
	pc.check("NormalField", i.Name, i.Parent)
}
func (pc *parentChecker) VisitEnumField(i *EnumField) {
	pc.check("EnumField", i.Name, i.Parent)
}
func (pc *parentChecker) VisitEnum(e *Enum) {
	pc.check("Enum", e.Name, e.Parent)
	pc.checkAll(e.Elements, e)
}
func (pc *parentChecker) VisitComment(e *Comment) {}
func (pc *parentChecker) VisitOneof(o *Oneof) {
	pc.check("Oneof", o.Name, o.Parent)
	pc.checkAll(o.Elements, o)
}
func (pc *parentChecker) VisitOneofField(o *OneOfField) {
	pc.check("OneOfField", o.Name, o.Parent)
}
func (pc *parentChecker) VisitReserved(r *Reserved) {
	pc.check("Reserved", "", r.Parent)
}
func (pc *parentChecker) VisitRPC(r *RPC) {
	pc.check("RPC", r.Name, r.Parent)
	//pc.checkAll(r.Options, r)
	for _, each := range r.Options {
		pc.check("Option", each.Name, r)
	}
}
func (pc *parentChecker) VisitMapField(f *MapField) {
	pc.check("MapField", f.Name, f.Parent)
}

// proto2
func (pc *parentChecker) VisitGroup(g *Group) {
	pc.check("Group", g.Name, g.Parent)
	pc.checkAll(g.Elements, g)
}
func (pc *parentChecker) VisitExtensions(e *Extensions) {
	pc.check("Extensions", "", e.Parent)
}
