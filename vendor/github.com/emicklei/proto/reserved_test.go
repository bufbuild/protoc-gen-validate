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

func TestReservedRanges(t *testing.T) {
	r := new(Reserved)
	p := newParserOn(`reserved 2, 15, 9 to 11;`)
	_, tok, _ := p.next()
	if tRESERVED != tok {
		t.Fail()
	}
	err := r.parse(p)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := r.Ranges[0].SourceRepresentation(), "2"; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
	if got, want := r.Ranges[2].SourceRepresentation(), "9 to 11"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestReservedFieldNames(t *testing.T) {
	r := new(Reserved)
	p := newParserOn(`reserved "foo", "bar";`)
	_, _, _ = p.next()
	err := r.parse(p)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := len(r.FieldNames), 2; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
	if got, want := r.FieldNames[0], "foo"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := r.FieldNames[1], "bar"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}
