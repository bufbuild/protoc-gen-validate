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
)

func TestField(t *testing.T) {
	proto := `repeated foo.bar lots =1 [option1=a, option2=b, option3="happy"];`
	p := newParserOn(proto)
	f := newNormalField()
	err := f.parse(p)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := f.Repeated, true; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := f.Type, "foo.bar"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := f.Name, "lots"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := len(f.Options), 3; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
	if got, want := f.Options[0].Name, "option1"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := f.Options[0].Constant.Source, "a"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := f.Options[1].Name, "option2"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := f.Options[1].Constant.Source, "b"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := f.Options[2].Constant.Source, "happy"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestFieldSimple(t *testing.T) {
	proto := `string optional_string_piece = 24 [ctype=STRING_PIECE];`
	p := newParserOn(proto)
	f := newNormalField()
	err := f.parse(p)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := f.Type, "string"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := f.Name, "optional_string_piece"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := f.Sequence, 24; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := len(f.Options), 1; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
	if got, want := f.Options[0].Name, "ctype"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := f.Options[0].Constant.Source, "STRING_PIECE"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestFieldSyntaxErrors(t *testing.T) {
	for i, each := range []string{
		`repeatet foo.bar lots = 1;`,
		`string lots === 1;`,
	} {
		f := newNormalField()
		if f.parse(newParserOn(each)) == nil {
			t.Errorf("uncaught syntax error in test case %d, %#v", i, f)
		}
	}
}

func TestMapField(t *testing.T) {
	proto := ` <string, Project> projects = 3;`
	p := newParserOn(proto)
	f := newMapField()
	err := f.parse(p)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := f.KeyType, "string"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := f.Type, "Project"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := f.Name, "projects"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := f.Sequence, 3; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestMapFieldWithDotTypes(t *testing.T) {
	proto := ` <.Some.How, .Such.Project> projects = 3;`
	p := newParserOn(proto)
	f := newMapField()
	err := f.parse(p)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := f.KeyType, ".Some.How"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := f.Type, ".Such.Project"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestOptionalWithOption(t *testing.T) {
	proto := `optional int32 default_int32    = 61 [default =  41    ];`
	p := newParserOn(proto)
	f := newNormalField()
	err := f.parse(p)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := f.Sequence, 61; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	o := f.Options[0]
	if got, want := o.Name, "default"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := o.Constant.Source, "41"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestFieldInlineComment(t *testing.T) {
	proto := `message Hello {
		// comment
		bool foo = 1; // inline comment
	  }`
	p := newParserOn(proto)
	def := new(Proto)
	err := def.parse(p)
	if err != nil {
		t.Fatal(err)
	}
	m := def.Elements[0].(*Message)
	if len(m.Elements) != 1 {
		t.Error("expected one element", m.Elements)
	}
	f := m.Elements[0].(*NormalField)
	if f.InlineComment == nil {
		t.Error("expected inline comment")
	}
}

func TestFieldTypeStartsWithDot(t *testing.T) {
	proto := `.game.Resource foo = 1;`
	p := newParserOn(proto)
	f := newNormalField()
	err := f.parse(p)
	if err != nil {
		t.Fatal(err)
	}
	foot := f.Field.Type
	if got, want := foot, ".game.Resource"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestMultiLineFieldType(t *testing.T) {
	src := `google.ads.googleads.v1.enums.ConversionAdjustmentTypeEnum
	.ConversionAdjustmentType adjustment_type = 5;`
	p := newParserOn(src)
	f := newNormalField()
	err := f.parse(p)
	if err != nil {
		t.Fatal(err)
	}
}
