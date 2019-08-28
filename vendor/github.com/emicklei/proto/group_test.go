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

func TestGroup(t *testing.T) {
	oto := `message M {
		// group
        optional group OptionalGroup = 16 {
			// field
            optional int32 a = 17;
        }
    }`
	p := newParserOn(oto)
	p.next() // consume first token
	m := new(Message)
	err := m.parse(p)
	if err != nil {
		t.Error(err)
	}
	if got, want := len(m.Elements), 1; got != want {
		t.Logf("%#v", m.Elements)
		t.Fatalf("got [%v] want [%v]", got, want)
	}
	g := m.Elements[0].(*Group)
	if got, want := len(g.Elements), 1; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
	if got, want := g.Position.Line, 3; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
	if got, want := g.Comment != nil, true; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	f := g.Elements[0].(*NormalField)
	if got, want := f.Name, "a"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := f.Optional, true; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}
