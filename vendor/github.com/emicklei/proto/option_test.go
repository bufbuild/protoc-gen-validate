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

func TestOptionCases(t *testing.T) {
	for i, each := range []struct {
		proto     string
		name      string
		strLit    string
		nonStrLit string
	}{{
		`option (full).java_package = "com.example.foo";`,
		"(full).java_package",
		"com.example.foo",
		"",
	}, {
		`option Bool = true;`,
		"Bool",
		"",
		"true",
	}, {
		`option Float = -3.14E1;`,
		"Float",
		"",
		"-3.14E1",
	}, {
		`option (foo_options) = { opt1: 123 opt2: "baz" };`,
		"(foo_options)",
		"",
		"",
	}, {
		`option foo = []`,
		"foo",
		"",
		"",
	}, {
		`option optimize_for = SPEED;`,
		"optimize_for",
		"",
		"SPEED",
	}, {
		"option (my.enum.service.is.like).rpc = 1;",
		"(my.enum.service.is.like).rpc",
		"",
		"1",
	}, {
		`option (imported.oss.package).action = "literal-double-quotes";`,
		"(imported.oss.package).action",
		"literal-double-quotes",
		"",
	}, {
		`option (imported.oss.package).action = "key:\"literal-double-quotes-escaped\"";`,
		"(imported.oss.package).action",
		`key:\"literal-double-quotes-escaped\"`,
		"",
	}, {
		`option (imported.oss.package).action = 'literalsinglequotes';`,
		"(imported.oss.package).action",
		"literalsinglequotes",
		"",
	}, {
		`option (imported.oss.package).action = 'single-quotes.with/symbols';`,
		"(imported.oss.package).action",
		"single-quotes.with/symbols",
		"",
	}} {
		p := newParserOn(each.proto)
		pr, err := p.Parse()
		if err != nil {
			t.Fatal("testcase failed:", i, err)
		}
		if got, want := len(pr.Elements), 1; got != want {
			t.Fatalf("[%d] got [%v] want [%v]", i, got, want)
		}
		o := pr.Elements[0].(*Option)
		if got, want := o.Name, each.name; got != want {
			t.Errorf("[%d] got [%v] want [%v]", i, got, want)
		}
		if len(each.strLit) > 0 {
			if got, want := o.Constant.Source, each.strLit; got != want {
				t.Errorf("[%d] got [%v] want [%v]", i, got, want)
			}
		}
		if len(each.nonStrLit) > 0 {
			if got, want := o.Constant.Source, each.nonStrLit; got != want {
				t.Errorf("[%d] got [%v] want [%v]", i, got, want)
			}
		}
		if got, want := o.IsEmbedded, false; got != want {
			t.Errorf("[%d] got [%v] want [%v]", i, got, want)
		}
	}
}

func TestLiteralString(t *testing.T) {
	proto := `"string"`
	p := newParserOn(proto)
	l := new(Literal)
	if err := l.parse(p); err != nil {
		t.Fatal(err)
	}
	if got, want := l.IsString, true; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := l.Source, "string"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestOptionComments(t *testing.T) {
	proto := `
// comment
option Help = "me"; // inline`
	p := newParserOn(proto)
	pr, err := p.Parse()
	if err != nil {
		t.Fatal(err)
	}
	o := pr.Elements[0].(*Option)
	if got, want := o.IsEmbedded, false; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := o.Comment != nil, true; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
	if got, want := o.Comment.Lines[0], " comment"; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
	if got, want := o.InlineComment != nil, true; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
	if got, want := o.InlineComment.Lines[0], " inline"; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
	if got, want := o.Position.Line, 3; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
	if got, want := o.Comment.Position.Line, 2; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
	if got, want := o.InlineComment.Position.Line, 3; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
}

func TestAggregateSyntax(t *testing.T) {
	proto := `
// usage:
message Bar {
  // alternative aggregate syntax (uses TextFormat):
  int32 b = 2 [(foo_options) = {
    opt1: 123,
    opt2: "baz"
  }];
}
	`
	p := newParserOn(proto)
	pr, err := p.Parse()
	if err != nil {
		t.Fatal(err)
	}
	o := pr.Elements[0].(*Message)
	f := o.Elements[0].(*NormalField)
	if got, want := len(f.Options), 1; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
	ac := f.Options[0].Constant.Map
	if got, want := len(ac), 2; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
	if got, want := ac["opt1"].Source, "123"; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
	if got, want := o.Position.Line, 3; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
	if got, want := o.Comment.Position.String(), "<input>:2:1"; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
	if got, want := f.Position.String(), "<input>:5:3"; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
	if got, want := f.Options[0].Constant.Position.Line, 5; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
	// check for AggregatedConstants
	list := f.Options[0].AggregatedConstants
	if got, want := list[0].Source, "123"; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
	if got, want := list[1].Source, "baz"; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
}

func TestNonPrimitiveOptionComment(t *testing.T) {
	proto := `
// comment
option Help = { string_field: "value" }; // inline`
	p := newParserOn(proto)
	pr, err := p.Parse()
	if err != nil {
		t.Fatal(err)
	}
	o := pr.Elements[0].(*Option)
	if got, want := o.Comment != nil, true; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
	if got, want := o.Comment.Lines[0], " comment"; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
	if got, want := o.InlineComment != nil, true; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
	if got, want := o.InlineComment.Lines[0], " inline"; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
}

func TestFieldCustomOptions(t *testing.T) {
	proto := `foo.bar lots = 1 [foo={hello:1}, bar=2];`
	p := newParserOn(proto)
	f := newNormalField()
	err := f.parse(p)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := f.Type, "foo.bar"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := f.Name, "lots"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := len(f.Options), 2; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
	if got, want := f.Options[0].Name, "foo"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := f.Options[1].Name, "bar"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := f.Options[1].Constant.Source, "2"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	// check for AggregatedConstants
	if got, want := f.Options[0].AggregatedConstants[0].Name, "hello"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := f.Options[0].AggregatedConstants[0].PrintsColon, true; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := f.Options[0].AggregatedConstants[0].Source, "1"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestIgnoreIllegalEscapeCharsInAggregatedConstants(t *testing.T) {
	src := `syntax = "proto3";
	message Person {
	  string name  = 3 [(validate.rules).string = {
						  pattern:   "^[^\d\s]+( [^\d\s]+)*$",
						  max_bytes: 256,
					   }];
	}`
	p := newParserOn(src)
	d, err := p.Parse()
	if err != nil {
		t.Fatal(err)
	}
	f := d.Elements[1].(*Message).Elements[0].(*NormalField)
	if got, want := f.Options[0].Name, "(validate.rules).string"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := len(f.Options[0].Constant.Map), 2; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := f.Options[0].Constant.Map["pattern"].Source, "^[^\\d\\s]+( [^\\d\\s]+)*$"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := len(f.Options[0].Constant.OrderedMap), 2; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := f.Options[0].Constant.OrderedMap[0].Name, "pattern"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := f.Options[0].Constant.OrderedMap[0].Source, "^[^\\d\\s]+( [^\\d\\s]+)*$"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	// check for AggregatedConstants
	if got, want := f.Options[0].AggregatedConstants[0].Name, "pattern"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := f.Options[0].AggregatedConstants[0].PrintsColon, true; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := f.Options[0].AggregatedConstants[0].Source, "^[^\\d\\s]+( [^\\d\\s]+)*$"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestIgnoreIllegalEscapeCharsInConstant(t *testing.T) {
	src := `syntax = "proto2";
	message Person {
		optional string cpp_trigraph = 20 [default = "? \? ?? \?? \??? ??/ ?\?-"];
	}`
	p := newParserOn(src)
	d, err := p.Parse()
	if err != nil {
		t.Fatal(err)
	}
	f := d.Elements[1].(*Message).Elements[0].(*NormalField)
	if got, want := f.Options[0].Constant.Source, "? \\? ?? \\?? \\??? ??/ ?\\?-"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestFieldCustomOptionExtendedIdent(t *testing.T) {
	proto := `Type field = 1 [(validate.rules).enum.defined_only = true];`
	p := newParserOn(proto)
	f := newNormalField()
	err := f.parse(p)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := f.Options[0].Name, "(validate.rules).enum.defined_only"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

// issue #50
func TestNestedAggregateConstants(t *testing.T) {
	src := `syntax = "proto3";

	package baz;

	option (foo.bar) = {
	  woot: 100
	  foo {
		hello: 200
		hello2: 300
		bar {
			hello3: 400
		}
	  }
	};`
	p := newParserOn(src)
	proto, err := p.Parse()
	if err != nil {
		t.Error(err)
	}
	option := proto.Elements[2].(*Option)
	if got, want := option.Name, "(foo.bar)"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := len(option.Constant.Map), 2; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	m := option.Constant.Map
	if got, want := m["woot"].Source, "100"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := len(m["foo"].Map), 3; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	m = m["foo"].Map
	if got, want := len(m["bar"].Map), 1; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := m["bar"].Map["hello3"].Source, "400"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := option.Constant.OrderedMap[1].Name, "foo"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := option.Constant.OrderedMap[1].OrderedMap[2].Name, "bar"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := option.Constant.OrderedMap[1].OrderedMap[2].OrderedMap[0].Source, "400"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := len(option.AggregatedConstants), 4; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	for _, each := range option.AggregatedConstants {
		t.Logf("%#v=%v\n", each, each.SourceRepresentation())
	}
}

// Issue #59
func TestMultiLineOptionAggregateValue(t *testing.T) {
	src := `rpc ListTransferLogs(ListTransferLogsRequest)
	returns (ListTransferLogsResponse) {
		option (google.api.http) = {
		get: "/v1/{parent=projects/*/locations/*/transferConfigs/*/runs/*}/"
			"transferLogs"
		};
}`
	p := newParserOn(src)
	rpc := new(RPC)
	p.next()
	err := rpc.parse(p)
	if err != nil {
		t.Error(err)
	}
	get := rpc.Options[0].Constant.Map["get"]
	if got, want := get.Source, "/v1/{parent=projects/*/locations/*/transferConfigs/*/runs/*}/transferLogs"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

// issue #76
func TestOptionAggregateCanUseKeyword(t *testing.T) {
	src := `message User {
		string email = 3 [(validate.field) = {required: true}];
	}`
	p := newParserOn(src)
	_, err := p.Parse()
	if err != nil {
		t.Error(err)
	}
}

// issue #77
func TestOptionAggregateWithRepeatedValues(t *testing.T) {
	src := `message Envelope {
		int64 not_in = 15 [(validate.rules).int64 = {not_in: [40, 45]}];
		int64 in = 16 [(validate.rules).int64 = {in: [[1],[2]]}];
	}`
	p := newParserOn(src)
	def, err := p.Parse()
	if err != nil {
		t.Error(err)
	}
	field := def.Elements[0].(*Message).Elements[0].(*NormalField)
	notIn := field.Options[0].Constant.Map["not_in"]
	if got, want := len(notIn.Array), 2; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := notIn.Array[0].Source, "40"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := notIn.Array[1].Source, "45"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestInvalidOptionAggregateWithRepeatedValues(t *testing.T) {
	src := `message Bogus {
		int64 a = 1 [a = {not_in: [40 syntax]}];
	}`
	p := newParserOn(src)
	_, err := p.Parse()
	if err == nil {
		t.Error("expected syntax error")
	}
}

// issue #79
func TestUseOfSemicolonsInAggregatedConstants(t *testing.T) {
	src := `rpc Test(Void) returns (Void) {
				option (google.api.http) = {
					post: "/api/v1/test";
					body: "*"; // ignored comment
				};
			}`
	p := newParserOn(src)
	rpc := new(RPC)
	p.next()
	err := rpc.parse(p)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := len(rpc.Elements), 1; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	opt := rpc.Elements[0].(*Option)
	if got, want := len(opt.Constant.Map), 2; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
	// old access to map
	if got, want := opt.Constant.Map["body"].Source, "*"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	// new access to map
	body, ok := opt.Constant.OrderedMap.Get("body")
	if !ok {
		t.Fatal("expected body key")
	}
	if got, want := body.Source, "*"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	for _, each := range opt.Constant.OrderedMap {
		t.Log(each)
	}
}

func TestParseNestedSelectorInAggregatedConstant(t *testing.T) {
	src := `rpc Test(Void) returns (Void) {
		option (google.api.http) = {
			get: "/api/v1/test"
			additional_bindings.post: "/api/v1/test"
			additional_bindings.body: "*"
		};
	}`
	p := newParserOn(src)
	rpc := new(RPC)
	p.next()
	err := rpc.parse(p)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := rpc.Options[0].Constant.Map["get"].Source, "/api/v1/test"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := rpc.Options[0].Constant.Map["additional_bindings.post"].Source, "/api/v1/test"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := rpc.Options[0].AggregatedConstants[2].Name, "additional_bindings.body"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := rpc.Options[0].AggregatedConstants[2].Source, "*"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestParseMultilineStringConstant(t *testing.T) {
	src := `message Test {
		string description = 3 [
			(common.ui_field_desc) = "Description of the account"
									 " domain (e.g. Team,"
									 "Name User Account Directory)."
  		];
	}`
	p := newParserOn(src)
	m := new(Message)
	p.next()
	err := m.parse(p)
	if err != nil {
		t.Fatal(err)
	}
	s := m.Elements[0].(*NormalField).Options[0].Constant.Source
	if got, want := s, "Description of the account domain (e.g. Team,Name User Account Directory)."; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestOptionWithRepeatedMessageValues(t *testing.T) {
	src := `message Foo {
		int64 a = 1 [b = {repeated_message_field: [{hello: 1}, {hello: 2}]}];
	}`
	p := newParserOn(src)
	def, err := p.Parse()
	if err != nil {
		t.Errorf("expected no error but got %v", err)
	}
	opt := def.Elements[0].(*Message).Elements[0].(*NormalField).Options[0]
	hello, ok := opt.AggregatedConstants[0].Array[0].OrderedMap.Get("hello")
	if !ok {
		t.Fail()
	}
	if got, want := hello.SourceRepresentation(), "1"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestOptionWithRepeatedMessageValuesWithArray(t *testing.T) {
	src := `message Foo {
		int64 a = 1 [ (bar.repeated_field_dep_option) =
			{ hello: 1, repeated_dep: [
				{ hello: 1, repeated_bar: [1, 2] },
				{ hello: 3, repeated_bar: [3, 4] } ] } ];
	}`
	p := newParserOn(src)
	def, err := p.Parse()
	if err != nil {
		t.Errorf("expected no error but got %v", err)
	}
	opt := def.Elements[0].(*Message).Elements[0].(*NormalField).Options[0]
	hello, ok := opt.Constant.OrderedMap.Get("hello")
	if !ok {
		t.Fail()
	}
	if got, want := hello.SourceRepresentation(), "1"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	repeatedDep, ok := opt.Constant.OrderedMap.Get("repeated_dep")
	if !ok {
		t.Fail()
	}
	if got, want := len(repeatedDep.Array), 2; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	hello, ok = repeatedDep.Array[0].OrderedMap.Get("hello")
	if !ok {
		t.Fail()
	}
	if got, want := hello.SourceRepresentation(), "1"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	onetwo, ok := repeatedDep.Array[0].OrderedMap.Get("repeated_bar")
	if !ok {
		t.Fail()
	}
	if got, want := onetwo.Array[0].Source, "1"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := onetwo.Array[1].Source, "2"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

// https://github.com/emicklei/proto/issues/99
func TestFieldCustomOptionLeadingDot(t *testing.T) {
	proto := `string app_entity_id = 4 [(.common.v1.some_custom_option) = { opt1: true opt2: false }];`
	p := newParserOn(proto)
	f := newNormalField()
	err := f.parse(p)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := f.Type, "string"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	o := f.Options[0]
	if got, want := o.Name, "(.common.v1.some_custom_option)"; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
}

// https://github.com/emicklei/proto/issues/106
func TestEmptyArrayInOptionStructure(t *testing.T) {
	src := `
	option (grpc.gateway.protoc_gen_swagger.options.openapiv2_schema) = {
		json_schema : {
		  title : "Frob a request"
		  description : "blah blah blah"
		  required : [ ]
		  optional:["this"]
		}
	  };	
	`
	p := newParserOn(src)
	p.next()
	o := new(Option)
	if err := o.parse(p); err != nil {
		t.Fatal("testcase parse failed:", err)
	}
	s, ok := o.Constant.OrderedMap.Get("json_schema")
	if !ok {
		t.Fatal("expected json_schema literal")
	}
	// none
	a, ok := s.OrderedMap.Get("required")
	if !ok {
		t.Fatal("expected required literal")
	}
	if len(a.Array) != 0 {
		t.Fatal("expecting empty array")
	}
	// one
	a, ok = s.OrderedMap.Get("optional")
	if !ok {
		t.Fatal("expected required literal")
	}
	if len(a.Array) != 1 {
		t.Fatal("expecting one size array")
	}
	if got, want := a.Array[0].Source, "this"; got != want {
		t.Fatalf("got [%s] want [%s]", got, want)
	}
}

// https://github.com/emicklei/proto/issues/107
func TestQuoteNotDroppedInOption(t *testing.T) {
	src := `string name = 1 [ quote = '<="foo"' ];`
	f := newNormalField()
	if err := f.parse(newParserOn(src)); err != nil {
		t.Fatal(err)
	}
	sr := f.Options[0].Constant.SourceRepresentation()
	if got, want := sr, `'<="foo"'`; got != want {
		t.Errorf("got [%s] want [%s]", got, want)
	}
}

func TestWhatYouTypeIsWhatYouGetOptionValue(t *testing.T) {
	src := `string n = 1 [ quote = 'm"\"/"' ];`
	f := newNormalField()
	if err := f.parse(newParserOn(src)); err != nil {
		t.Fatal(err)
	}
	sr := f.Options[0].Constant.SourceRepresentation()
	if got, want := sr, `'m"\"/"'`; got != want {
		t.Errorf("got [%s] want [%s]", got, want)
	}
}

func TestLiteralNoQuoteRuneSet(t *testing.T) {
	l := Literal{
		Source:   "foo",
		IsString: true,
	}
	if got, want := l.SourceRepresentation(), "\"foo\""; got != want {
		t.Errorf("got [%s] want [%s]", got, want)
	}
}
