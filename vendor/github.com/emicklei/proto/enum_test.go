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

func TestEnum(t *testing.T) {
	proto := `
// enum
enum EnumAllowingAlias {
  reserved 998, 1000 to 2000;
  reserved "HELLO", "WORLD";
  option allow_alias = true;
  UNKNOWN = 0;
  STARTED = 1;
  RUNNING = 2 [(custom_option) = "hello world"];
  NEG = -42;
  SOMETHING_FOO = 0 [
    (bar.enum_value_option) = true,
    (bar.enum_value_dep_option) = { hello: 1 }
  ];  
}`
	p := newParserOn(proto)
	pr, err := p.Parse()
	if err != nil {
		t.Fatal(err)
	}
	enums := collect(pr).Enums()
	if got, want := len(enums), 1; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := len(enums[0].Elements), 8; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := enums[0].Comment != nil, true; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
	if got, want := enums[0].Comment.Message(), " enum"; got != want {
		t.Errorf("got [%v] want [%v]", enums[0].Comment, want)
	}
	if got, want := enums[0].Position.Line, 3; got != want {
		t.Errorf("got [%d] want [%d]", got, want)
	}
	// enum reserved ids
	e1 := enums[0].Elements[0].(*Reserved)
	if got, want := len(e1.Ranges), 2; got != want {
		t.Errorf("got [%d] want [%d]", got, want)
	}
	e1rg0 := e1.Ranges[0]
	if got, want := e1rg0.From, 998; got != want {
		t.Errorf("got [%d] want [%d]", got, want)
	}
	if got, want := e1rg0.From, e1rg0.To; got != want {
		t.Errorf("got [%d] want [%d]", got, want)
	}
	e1rg1 := e1.Ranges[1]
	if got, want := e1rg1.From, 1000; got != want {
		t.Errorf("got [%d] want [%d]", got, want)
	}
	if got, want := e1rg1.To, 2000; got != want {
		t.Errorf("got [%d] want [%d]", got, want)
	}
	// enum reserved field names
	e2 := enums[0].Elements[1].(*Reserved)
	if got, want := len(e2.FieldNames), 2; got != want {
		t.Errorf("got [%d] want [%d]", got, want)
	}
	e2fn0 := e2.FieldNames[0]
	if got, want := e2fn0, "HELLO"; got != want {
		t.Errorf("got [%s] want [%s]", got, want)
	}
	e2fn1 := e2.FieldNames[1]
	if got, want := e2fn1, "WORLD"; got != want {
		t.Errorf("got [%s] want [%s]", got, want)
	}
	ef1 := enums[0].Elements[3].(*EnumField)
	if got, want := ef1.Integer, 0; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := ef1.Position.Line, 7; got != want {
		t.Errorf("got [%d] want [%d]", got, want)
	}
	ef3 := enums[0].Elements[5].(*EnumField)
	if got, want := ef3.Integer, 2; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	ef3opt := ef3.Elements[0].(*Option)
	if got, want := ef3opt.Name, "(custom_option)"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	// test for deprecated field
	if got, want := ef3opt, ef3.ValueOption; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := ef3opt.Constant.Source, "hello world"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := ef3.Position.Line, 9; got != want {
		t.Errorf("got [%d] want [%d]", got, want)
	}
	ef4 := enums[0].Elements[6].(*EnumField)
	if got, want := ef4.Integer, -42; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestEnumWithHex(t *testing.T) {
	src := `enum Flags {
		  FLAG1 = 0x11;
		}`
	p := newParserOn(src)
	enum := new(Enum)
	p.next()
	if err := enum.parse(p); err != nil {
		t.Fatal(err)
	}
	if got, want := len(enum.Elements), 1; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := enum.Elements[0].(*EnumField).Integer, 17; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}
