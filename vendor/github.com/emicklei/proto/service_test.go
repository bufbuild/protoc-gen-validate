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

func TestService(t *testing.T) {
	proto := `service AccountService {
		// comment
		rpc CreateAccount (CreateAccount) returns (ServiceFault); // inline comment
		rpc GetAccounts   (stream Int64)  returns (Account) {} // inline comment2
		rpc Health(google.protobuf.Empty) returns (google.protobuf.Empty) {} // inline comment3
	}`
	pr, err := newParserOn(proto).Parse()
	if err != nil {
		t.Fatal(err)
	}
	srv := collect(pr).Services()[0]
	if got, want := len(srv.Elements), 3; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
	if got, want := srv.Position.String(), "<input>:1:1"; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
	rpc1 := srv.Elements[0].(*RPC)
	if got, want := rpc1.Name, "CreateAccount"; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
	if got, want := rpc1.Doc().Message(), " comment"; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
	if got, want := rpc1.InlineComment.Message(), " inline comment"; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
	if got, want := rpc1.Position.Line, 3; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
	rpc2 := srv.Elements[1].(*RPC)
	if got, want := rpc2.Name, "GetAccounts"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	rpc3 := srv.Elements[2].(*RPC)
	if got, want := rpc3.Name, "Health"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if rpc2.InlineComment == nil {
		t.Fatal("missing inline comment 2")
	}
	if got, want := rpc2.InlineComment.Message(), " inline comment2"; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
	if rpc3.InlineComment == nil {
		t.Fatal("missing inline comment 3")
	}
	if got, want := rpc3.InlineComment.Message(), " inline comment3"; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
}

func TestRPCWithOptionAggregateSyntax(t *testing.T) {
	proto := `service AccountService {
		// CreateAccount
		rpc CreateAccount (CreateAccount) returns (ServiceFault){
			// test_ident
			option (test_ident) = {
				test: "test"
				test2:"test2"
			}; // inline test_ident
		}
	}`
	pr, err := newParserOn(proto).Parse()
	if err != nil {
		t.Fatal(err)
	}
	srv := collect(pr).Services()[0]
	if got, want := len(srv.Elements), 1; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
	rpc1 := srv.Elements[0].(*RPC)
	if got, want := len(rpc1.Elements), 2; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	com := rpc1.Elements[0].(*Comment)
	if got, want := com.Message(), " test_ident"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	opt := rpc1.Elements[1].(*Option)
	if got, want := opt.Name, "(test_ident)"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := opt.InlineComment != nil, true; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
	if got, want := opt.InlineComment.Message(), " inline test_ident"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := len(opt.Constant.Map), 2; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
	// test deprecated field Options in RPC
	if got, want := len(rpc1.Options), 1; got != want {
		t.Errorf("got len Options %v want %v", got, want)
	}
}

func TestServiceWithOption(t *testing.T) {
	src := `service AnyService {
		option secure = true;
	  }`
	p := newParserOn(src)
	p.next()
	svc := new(Service)
	err := svc.parse(p)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := svc.Elements[0].(*Option).Name, "secure"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestRPCWithOneLineCommentInOptionBlock(t *testing.T) {
	proto := `service AccountService {
		rpc CreateAccount (CreateAccount) returns (ServiceFault) {
			// test comment
		}
	}`
	_, err := newParserOn(proto).Parse()
	if err != nil {
		t.Fatal(err)
	}
}

func TestRPCWithMultiLineCommentInOptionBlock(t *testing.T) {
	proto := `service AccountService {
		rpc CreateAccount (CreateAccount) returns (ServiceFault) {
			// test comment
			// test comment
		}
	}`
	def, err := newParserOn(proto).Parse()
	if err != nil {
		t.Fatal(err)
	}
	s := def.Elements[0].(*Service)
	r := s.Elements[0].(*RPC)
	if got, want := len(r.Elements), 1; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	c := r.Elements[0].(*Comment)
	if got, want := len(c.Lines), 2; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestRPCWithTypeThatHasLeadingDot(t *testing.T) {
	src := `service Dummy {
		rpc DeleteProgram (ProgramIdentifier) returns (.google.protobuf.Empty) {}
	}`
	_, err := newParserOn(src).Parse()
	if err != nil {
		t.Fatal(err)
	}
}
