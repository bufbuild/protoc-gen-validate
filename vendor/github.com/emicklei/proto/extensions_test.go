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

import "testing"

func TestExtensions(t *testing.T) {
	proto := `message M {
		// extensions
		extensions 4, 20 to max; // max
	}`
	p := newParserOn(proto)
	p.next() // consume message
	m := new(Message)
	err := m.parse(p)
	if err != nil {
		t.Fatal(err)
	}
	if len(m.Elements) != 1 {
		t.Fatal("1 elements expected, got", len(m.Elements), m.Elements)
	}
	f := m.Elements[0].(*Extensions)
	if got, want := len(f.Ranges), 2; got != want {
		t.Fatalf("got [%d] want [%d]", got, want)
	}
	if got, want := f.Position.Line, 3; got != want {
		t.Fatalf("got [%d] want [%d]", got, want)
	}
	if got, want := f.Ranges[1].SourceRepresentation(), "20 to max"; got != want {
		t.Errorf("got [%s] want [%s]", got, want)
	}
	if f.Comment == nil {
		t.Fatal("comment expected")
	}
	if got, want := f.InlineComment.Message(), " max"; got != want {
		t.Errorf("got [%s] want [%s]", got, want)
	}
}
