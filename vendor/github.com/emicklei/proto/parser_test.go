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
	"strings"
	"testing"
)

func TestParseComment(t *testing.T) {
	proto := `
    // first
	// second

    /*
	ctyle
	multi
	line
    */

    // cpp style single line //

	message test{}
	`
	p := newParserOn(proto)
	pr, err := p.Parse()
	if err != nil {
		t.Fatal(err)
	}

	if got, want := len(collect(pr).Comments()), 3; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func newParserOn(def string) *Parser {
	p := NewParser(strings.NewReader(def))
	p.debug = true
	return p
}

func TestScanIgnoreWhitespace_Digits(t *testing.T) {
	p := newParserOn(" 1234 ")
	_, _, lit := p.next()
	if got, want := lit, "1234"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestScanIgnoreWhitespace_Minus(t *testing.T) {
	p := newParserOn(" -1234")
	_, _, lit := p.next()
	if got, want := lit, "-"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestNextIdentifier(t *testing.T) {
	ident := " aap.noot.mies "
	p := newParserOn(ident)
	_, tok, lit := p.nextIdentifier()
	if got, want := tok, tIDENT; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := lit, strings.TrimSpace(ident); got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestNextIdentifierWithKeyword(t *testing.T) {
	ident := " aap.rpc.mies.enum ="
	p := newParserOn(ident)
	_, tok, lit := p.nextIdentifier()
	if got, want := tok, tIDENT; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := lit, "aap.rpc.mies.enum"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	_, tok, _ = p.next()
	if got, want := tok, tEQUALS; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestNextIdentifierNoIdent(t *testing.T) {
	ident := "("
	p := newParserOn(ident)
	_, tok, lit := p.nextIdentifier()
	if got, want := tok, tLEFTPAREN; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := lit, "("; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

// https://github.com/google/protobuf/issues/4726
func TestProtobufIssue4726(t *testing.T) {
	src := `syntax = "proto3";

	service SomeService {
		rpc SomeMethod (Whatever) returns (Whatever) {
			option (google.api.http) = {
				delete : "/some/url"
				additional_bindings {
					delete: "/another/url"
				}
			};
		}
	}`
	p := newParserOn(src)
	_, err := p.Parse()
	if err != nil {
		t.Error(err)
	}
}

func TestProtoIssue92(t *testing.T) {
	src := `syntax = "proto3";

package test;

message Foo {
  .game.Resource one = 1 [deprecated = true];
  repeated .game.sub.Resource two = 2;
  map<string, .game.Resource> three = 3;
}`
	p := newParserOn(src)
	_, err := p.Parse()
	if err != nil {
		t.Error(err)
	}
}

func TestParseSingleQuotesStrings(t *testing.T) {
	p := newParserOn(` 'bohemian','' `)
	_, _, lit := p.next()
	if got, want := lit, "'bohemian'"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	_, tok, _ := p.next()
	if got, want := tok, tCOMMA; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	_, _, lit = p.next()
	if got, want := lit, "''"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}
