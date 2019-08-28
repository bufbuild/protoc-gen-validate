// Copyright (c) 2017 Ernest Micklei
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
	"testing"
	"text/scanner"
)

var startPosition = scanner.Position{Line: 1, Column: 1}

func TestCreateComment(t *testing.T) {
	c0 := newComment(startPosition, "")
	if got, want := len(c0.Lines), 1; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	c1 := newComment(startPosition, `hello
world`)
	if got, want := len(c1.Lines), 2; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := c1.Lines[0], "hello"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := c1.Lines[1], "world"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := c1.Cstyle, false; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestTakeLastComment(t *testing.T) {
	c0 := newComment(startPosition, "hi")
	c1 := newComment(startPosition, "there")
	_, l := takeLastCommentIfEndsOnLine([]Visitee{c0, c1}, 1)
	if got, want := len(l), 1; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
	if got, want := l[0], c0; got != want {
		t.Errorf("got [%v] want [%v]", c1, want)
	}
}

func TestParseCommentWithEmptyLinesIndentAndTripleSlash(t *testing.T) {
	proto := `
	// comment 1
	// comment 2
	//
	// comment 3
	/// comment 4`
	p := newParserOn(proto)
	def, err := p.Parse()
	if err != nil {
		t.Fatal(err)
	}
	//spew.Dump(def)
	if got, want := len(def.Elements), 1; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}

	if got, want := len(def.Elements[0].(*Comment).Lines), 5; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
	if got, want := def.Elements[0].(*Comment).Lines[4], " comment 4"; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
	if got, want := def.Elements[0].(*Comment).Position.Line, 2; got != want {
		t.Fatalf("got [%d] want [%d]", got, want)
	}
	if got, want := def.Elements[0].(*Comment).Cstyle, false; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
}

func TestParseCStyleComment(t *testing.T) {
	proto := `
/*comment 1
comment 2

comment 3
  comment 4
*/`
	p := newParserOn(proto)
	def, err := p.Parse()
	if err != nil {
		t.Fatal(err)
	}
	if got, want := len(def.Elements), 1; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}

	if got, want := len(def.Elements[0].(*Comment).Lines), 6; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
	if got, want := def.Elements[0].(*Comment).Lines[3], "comment 3"; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
	if got, want := def.Elements[0].(*Comment).Lines[4], "  comment 4"; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
	if got, want := def.Elements[0].(*Comment).Cstyle, true; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
}

func TestParseCStyleCommentWithIndent(t *testing.T) {
	proto := `
	/*comment 1
	comment 2

	comment 3
	  comment 4
	*/`
	p := newParserOn(proto)
	def, err := p.Parse()
	if err != nil {
		t.Fatal(err)
	}
	if got, want := len(def.Elements), 1; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}

	if got, want := len(def.Elements[0].(*Comment).Lines), 6; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
	if got, want := def.Elements[0].(*Comment).Lines[0], "comment 1"; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
	if got, want := def.Elements[0].(*Comment).Lines[3], "\tcomment 3"; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
	if got, want := def.Elements[0].(*Comment).Lines[4], "\t  comment 4"; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
	if got, want := def.Elements[0].(*Comment).Cstyle, true; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
}

func TestParseCStyleOneLineComment(t *testing.T) {
	proto := `/* comment 1 */`
	p := newParserOn(proto)
	def, err := p.Parse()
	if err != nil {
		t.Fatal(err)
	}
	if got, want := len(def.Elements), 1; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}

	if got, want := len(def.Elements[0].(*Comment).Lines), 1; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
	if got, want := def.Elements[0].(*Comment).Lines[0], " comment 1 "; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := def.Elements[0].(*Comment).Cstyle, true; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestParseCStyleInlineComment(t *testing.T) {
	proto := `message Foo {
		int64 hello = 1; /*
			comment 1
		*/
	}`
	p := newParserOn(proto)
	def := new(Proto)
	err := def.parse(p)
	if err != nil {
		t.Fatal(err)
	}
	m := def.Elements[0].(*Message)
	if len(m.Elements) != 1 {
		t.Fatal("expected one element", m.Elements)
	}
	f := m.Elements[0].(*NormalField)
	comment := f.InlineComment
	if comment == nil {
		t.Fatal("no inline comment")
	}
	if got, want := len(comment.Lines), 3; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
	if got, want := comment.Lines[0], ""; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := comment.Lines[1], "			comment 1"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := comment.Cstyle, true; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestParseCommentWithTripleSlash(t *testing.T) {
	proto := `
/// comment 1
`
	p := newParserOn(proto)
	def, err := p.Parse()
	if err != nil {
		t.Fatal(err)
	}
	if got, want := len(def.Elements), 1; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
	if got, want := def.Elements[0].(*Comment).ExtraSlash, true; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
	if got, want := def.Elements[0].(*Comment).Lines[0], " comment 1"; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
	if got, want := def.Elements[0].(*Comment).Position.Line, 2; got != want {
		t.Fatalf("got [%d] want [%d]", got, want)
	}
}

func TestCommentAssociation(t *testing.T) {
	src := `
	// foo1
	// foo2

	// bar
	
	syntax = "proto3";
	
	// baz
	
	// bat1
	// bat2
	package bat;
	
	// Oneway is the return type to use for an rpc method if
	// the method should be generated as oneway.
	message Oneway {
	  bool ack = 1;
	}`
	p := newParserOn(src)
	def, err := p.Parse()
	if err != nil {
		t.Fatal(err)
	}
	if got, want := len(def.Elements), 6; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
	pkg := def.Elements[4].(*Package)
	if got, want := pkg.Comment.Message(), " bat1"; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
	if got, want := len(pkg.Comment.Lines), 2; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
	if got, want := pkg.Comment.Lines[1], " bat2"; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
	if got, want := len(def.Elements[5].(*Message).Comment.Lines), 2; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
}
