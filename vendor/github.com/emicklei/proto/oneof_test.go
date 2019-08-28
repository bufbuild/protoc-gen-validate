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

func TestOneof(t *testing.T) {
	proto := `oneof foo {
	// just a name
    string	 name = 4;
    SubMessage sub_message = 9 [options=none];
}`
	p := newParserOn(proto)
	p.next() // consume first token
	o := new(Oneof)
	err := o.parse(p)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := o.Name, "foo"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := len(o.Elements), 2; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
	first := o.Elements[0].(*OneOfField)
	if got, want := first.Comment.Message(), " just a name"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := first.Position.Line, 3; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	second := o.Elements[1].(*OneOfField)
	if got, want := second.Name, "sub_message"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := second.Type, "SubMessage"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := second.Sequence, 9; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := second.Position.Line, 4; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestFieldOneofImported(t *testing.T) {
	fieldType := "foo.bar"
	proto := `message Value {
		oneof value {
			` + fieldType + ` value = 1;
		}
	}`
	p := newParserOn(proto)
	def := new(Proto)
	err := def.parse(p)
	if err != nil {
		t.Fatal(err)
	}
	m := def.Elements[0].(*Message)
	if len(m.Elements) != 1 {
		t.Errorf("expected one element but got %d", len(m.Elements))
	}
	o := m.Elements[0].(*Oneof)
	if len(o.Elements) != 1 {
		t.Errorf("expected one element but got %d", len(o.Elements))
	}
	f := o.Elements[0].(*OneOfField)
	if got, want := f.Type, fieldType; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := f.Name, "value"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestFieldOneofDotImported(t *testing.T) {
	fieldType := ".foo.bar"
	proto := `message Value {
		oneof value {
			` + fieldType + ` value = 1;
		}
	}`
	p := newParserOn(proto)
	def := new(Proto)
	err := def.parse(p)
	if err != nil {
		t.Fatal(err)
	}
	m := def.Elements[0].(*Message)
	if len(m.Elements) != 1 {
		t.Errorf("expected one element but got %d", len(m.Elements))
	}
	o := m.Elements[0].(*Oneof)
	if len(o.Elements) != 1 {
		t.Errorf("expected one element but got %d", len(o.Elements))
	}
	f := o.Elements[0].(*OneOfField)
	if got, want := f.Type, fieldType; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := f.Name, "value"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestOneOfWithOption(t *testing.T) {
	src := `oneof AnOneof {
		option (oneof_opt1) = -99;
		int32 oneof_field = 2;
	  }`
	p := newParserOn(src)
	p.next()
	oneof := new(Oneof)
	err := oneof.parse(p)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := oneof.Elements[0].(*Option).Name, "(oneof_opt1)"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}
