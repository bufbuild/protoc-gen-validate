package protolock

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const simpleProto = `syntax = "proto3";
package test;

message Channel {
  int64 id = 1;
  string name = 2;
  string description = 3;
}

message NextRequest {}
message PreviousRequest {}

service ChannelChanger {
	rpc Next(stream NextRequest) returns (Channel);
	rpc Previous(PreviousRequest) returns (stream Channel);
}
`

const noUsingReservedFieldsProto = `syntax = "proto3";
package test;

message Channel {
  reserved 4, 8 to 11;
  reserved "foo", "bar";  
  int64 id = 1;
  string name = 2;
  string description = 3;
}

message Request {
  reserved 2;
  reserved "field2";
  .example.snth.Field field1 = 1;
}

message NextRequest {
  reserved 3;
  reserved "a_map";
}

message PreviousRequest {
  reserved 4;
  reserved "no_use";
  oneof test_oneof {
    int64 id = 1;
    bool is_active = 2;
  }
}

enum WithAllowAlias {
  reserved "DONTUSE";
  reserved 2;
  option allow_alias = true;
  UNKNOWN = 0;
  STARTED = 1;
  RUNNING = 1;
}

enum NoWithAllowAlias {
  reserved "DONTUSE2";
  reserved 2;
  UNKNOWN2 = 0;
  STARTED2 = 1;
}

message IHaveAnEnum {
  int32 id = 1;

  enum IAmTheEnum {
    reserved "NONE";
    reserved 101;
    ALL = 0;
    SOME = 100;
  }
}

service ChannelChanger {
	rpc Next(stream NextRequest) returns (Channel);
	rpc Previous(PreviousRequest) returns (stream Channel);
}
`

const usingReservedFieldsProto = `syntax = "proto3";
package test;

message Channel {
  int64 id = 1;
  string name = 2;
  string description = 3;
  string foo = 4;
  bool bar = 5;

  message A {
    int32 id = 1;
    string name = 2;
  }
}

message Request {
  .example.snth.Field field1 = 1;
  .example.snth.Field field2 = 2;
}

message NextRequest {
  string name = 1;
  map<string, int32> a_map = 3;
}

message PreviousRequest {
  oneof test_oneof {
    int64 id = 1;
    bool is_active = 2;
    string no_use = 3;
    float thing = 4;
  }
}

enum WithAllowAlias {
  option allow_alias = true;
  UNKNOWN = 0;
  STARTED = 1;
  RUNNING = 1;
  STOPPED = 2;
  DONTUSE = 3;
}

enum NoWithAllowAlias {
  UNKNOWN2 = 0;
  STARTED2 = 1;
  DONTUSE2 = 1;
  STOPPED2 = 2;
}

message IHaveAnEnum {
  int32 id = 1;

  enum IAmTheEnum {
    ALL = 0;
    SOME = 100;
    NONE = 1;
    FEW = 101;
  }
}

service ChannelChanger {
  rpc Next(stream NextRequest) returns (Channel);
  rpc Previous(PreviousRequest) returns (stream Channel);
}
`

const noRemoveReservedFieldsProto = `syntax = "proto3";
package test;

message Channel {
  reserved 44, 101, 103 to 110;
  reserved "no_more", "goodbye";
  int64 id = 1;
  string name = 2;
  string description = 3;
  string foo = 4;
  bool bar = 5;
}

message NextRequest {
  reserved 3;
  reserved "a_map";
}

message PreviousRequest {
  reserved 4;
  reserved "no_use";
  oneof test_oneof {
    int64 id = 1;
    bool is_active = 2;
  }

  enum NestedEnum {
    reserved 11;
    reserved "NOPE";

    LOCATION = 1;
  }

  NestedEnum value = 3;
}

enum AnotherEnum {
  reserved 2;
  reserved "DONTUSEIT";

  option allow_alias = true;

  USE = 2;
  OK = 3;
  FINE = 3;
}

service ChannelChanger {
  rpc Next(stream NextRequest) returns (Channel);
  rpc Previous(PreviousRequest) returns (stream Channel);
}
`

const removeReservedFieldsProto = `syntax = "proto3";
package test;

message Channel {
  reserved 101, 103 to 107;
  reserved "no_more";
  int64 id = 1;
  string name = 2;
  string description = 3;
  string foo = 4;
  bool bar = 5;
}

message NextRequest {
  map<string, int32> a_map = 3;  
}

message PreviousRequest {
  oneof test_oneof {
    int64 id = 1;
    bool is_active = 2;
  }

  enum NestedEnum {
    LOCATION = 1;
  }

  NestedEnum value = 3;
}

enum AnotherEnum {
  option allow_alias = true;

  USE = 2;
  OK = 3;
  FINE = 3;
}

service ChannelChanger {
  rpc Next(stream NextRequest) returns (Channel);
  rpc Previous(PreviousRequest) returns (stream Channel);
}
`

const noChangeFieldIDsProto = `syntax = "proto3";
package test;

message Channel {
  int64 id = 1;
  string name = 2;
  string description = 3;
  string foo = 4;
  bool bar = 5;
}

message NextRequest {
  map<string, int64> a_map = 1;
}

message PreviousRequest {
  reserved 4;
  reserved "no_use";
  oneof test_oneof {
    int64 id = 1;
    bool is_active = 2;
  }

  enum NestedEnum {
    option allow_alias = true;

    ONE = 1;
    UNO = 1;
    TWO = 2;
    THREE = 3;
  }

  NestedEnum value = 3;
}

enum AnotherEnum {
  ABC = 1;
  DEF = 2;
}

service ChannelChanger {
  rpc Next(stream NextRequest) returns (Channel);
  rpc Previous(PreviousRequest) returns (stream Channel);
}
`

const changeFieldIDsProto = `syntax = "proto3";
package test;

message Channel {
  int64 id = 1;
  string name = 2;
  string description = 3;
  string foo = 4443;
  bool bar = 59;
}

message NextRequest {
  map<string, int64> a_map = 2;
}

message PreviousRequest {
  reserved 4;
  reserved "no_use";
  oneof test_oneof {
    int64 id = 11;
    bool is_active = 32;
  }

  enum NestedEnum {
    option allow_alias = true;

    ONE = 1;
    UNO = 7;
    TWO = 2;
    THREE = 3;
  }

  NestedEnum value = 3;
}

enum AnotherEnum {
  ABC = 1;
  DEF = 99;
}

service ChannelChanger {
  rpc Next(stream NextRequest) returns (Channel);
  rpc Previous(PreviousRequest) returns (stream Channel);
}
`

const noChangingFieldTypesProto = `syntax = "proto3";
package test;

message Channel {
  int64 id = 1;
  string name = 2;
  string description = 3;
  string foo = 4;
  bool bar = 5;
}

message Request {
  .example.snth.Field field2 = 1;
}


message NextRequest {
  string name = 1;
  map<string, int32> a_map = 3;
}

message PreviousRequest {
  oneof test_oneof {
    int64 id = 1;
    bool is_active = 2;
  }
}

service ChannelChanger {
  rpc Next(stream NextRequest) returns (Channel);
  rpc Previous(PreviousRequest) returns (stream Channel);
}
`

const changingFieldTypesProto = `syntax = "proto3";
package test;

message Channel {
  int32 id = 1;
  bool name = 2;
  string description = 3;
  string foo = 4;
  repeated bool bar = 5;
}

message Request {
  .example.notSnth.NotField field2 = 1;
}

message NextRequest {
  string name = 1;
  map<int64, bool> a_map = 3;
}

message PreviousRequest {
  oneof test_oneof {
    int32 id = 1;
    bool is_active = 2;
  }
}

service ChannelChanger {
  rpc Next(stream NextRequest) returns (Channel);
  rpc Previous(PreviousRequest) returns (stream Channel);
}
`

const noChangingFieldNamesProto = `syntax = "proto3";
package test;

message Channel {
  int64 id = 1;
  string name = 2;
  string description = 3;
  string foo = 4;
  bool bar = 5;
}

message NextRequest {
  map<string, bool> a_map = 1;
}

message PreviousRequest {
  oneof test_oneof {
    string name = 4;
    bool is_active = 9;
  }

  enum NestedEnum {
    option allow_alias = true;

    ONE = 1;
    TWO = 2;
    DOS = 2;
  }
}

enum AnotherEnum {
  ABC = 1;
  DEF = 2;
}

service ChannelChanger {
  rpc Next(stream NextRequest) returns (Channel);
  rpc Previous(PreviousRequest) returns (stream Channel);
}
`

const changingFieldNamesProto = `syntax = "proto3";
package test;

message Channel {
  reserved "name", "foo";
  int64 channel_id = 1;
  string name_2 = 2;
  string description_3 = 3;
  string foo_baz = 4;
  bool bar = 5;
}

message NextRequest {
  map<string, bool> b_map = 1;
}

message PreviousRequest {
  oneof test_oneof {
    string name_2 = 4;
    bool is_active = 9;
  }

  enum NestedEnum {
    option allow_alias = true;

    UNO = 1;
    TWO = 2;
    DOS = 2;
  }
}

enum AnotherEnum {
  ABC = 1;
  GHI = 2;
}

service ChannelChanger {
  rpc Next(stream NextRequest) returns (Channel);
  rpc Previous(PreviousRequest) returns (stream Channel);
}
`

const noRemovingServicesRPCsProto = `syntax = "proto3";
package test;

message Channel {
  int64 id = 1;
  string name = 2;
  string description = 3;
  string foo = 4;
  bool bar = 5;
}

message NextRequest {}
message PreviousRequest {}

service ChannelChanger {
  rpc Next(stream NextRequest) returns (Channel);
  rpc Previous(PreviousRequest) returns (stream Channel);
}
`

const removingServicesRPCsProto = `syntax = "proto3";
package test;

message Channel {
  int64 id = 1;
  string name = 2;
  string description = 3;
  string foo = 4;
  bool bar = 5;
}

message NextRequest {}
message PreviousRequest {}

service ChannelChanger {
}
`

const noChangingRPCSignatureProto = `syntax = "proto3";
package test;

message Channel {
  int64 id = 1;
  string name = 2;
  string description = 3;
  string foo = 4;
  bool bar = 5;
}

message NextRequest {}
message PreviousRequest {}

service ChannelChanger {
  rpc Next(stream NextRequest) returns (Channel);
  rpc Previous(PreviousRequest) returns (stream Channel);
}
`

const changingRPCSignatureProto = `syntax = "proto3";
package test;

message Channel {
  int64 id = 1;
  string name = 2;
  string description = 3;
  string foo = 4;
  bool bar = 5;
}

message NextRequest {}
message PreviousRequest {}

service ChannelChanger {
  rpc Next(NextRequest) returns (ChannelDifferent);
  rpc Previous(stream PreviousRequest) returns (stream Channel);
}
`

const noRemovingFieldsWithoutReserveProto = `syntax = "proto3";
package test;

message Channel {
  int64 id = 1;
  string name = 2;
  string description = 3;
  string foo = 4;
  bool bar = 5;
}

message NextRequest {
  map<int32, bool> a_map = 1; 
}

message PreviousRequest {
  oneof test_oneof {
    int64 id = 1;
    bool is_active = 2;
  }

  enum NestedEnum {
    option allow_alias = true;

    ONE = 1;
    UNO = 1;
    TWO = 2;
  }

  NestEnum value = 4;
}

enum AnotherEnum {
  ABC = 1;
  DEF = 2;
}

service ChannelChanger {
  rpc Next(stream NextRequest) returns (Channel);
  rpc Previous(PreviousRequest) returns (stream Channel);
}
`

const removingFieldsWithoutReserveProto = `syntax = "proto3";
package test;

message Channel {
  reserved 5;
  int64 id = 1;
  string name_new = 2;
  string description = 3;
  string foo = 4;
}

message NextRequest {
  reserved 1;
}

message PreviousRequest {
  reserved 1;

  enum NestedEnum {
    reserved 1;
    reserved "ONE";
    option allow_alias = true;

    TWO = 2;
  }

  NestEnum value = 4;
}

enum AnotherEnum {
  DEF = 2;
}

service ChannelChanger {
  rpc Next(stream NextRequest) returns (Channel);
  rpc Previous(PreviousRequest) returns (stream Channel);
}
`

const noConflictSameNameNestedMessages = `syntax = "proto3";
package main;

message A {
    message I {
        int32 index = 1;
    }

    string id = 1;
    I i = 2;
}

message B {
    message I {
        reserved 2;
        int32 index = 1;
    }

    string id = 1;
    I i = 2;
}
`

const shouldConflictNestedMessage = `syntax = "proto3";
package main;

message A {
    message I {
        int32 index = 1;
    }

    string id = 1;
    I i = 2;
}

message B {
    message I {
        int32 index = 1;
        string name = 2;
    }

    string id = 1;
    I i = 2;
}
`

func TestParseOnReader(t *testing.T) {
	r := strings.NewReader(simpleProto)
	_, err := Parse("simpleProto", r)
	assert.NoError(t, err)
}

func TestChangingRPCSignature(t *testing.T) {
	SetDebug(true)
	curLock := parseTestProto(t, noChangingRPCSignatureProto)
	updLock := parseTestProto(t, changingRPCSignatureProto)

	warnings, ok := NoChangingRPCSignature(curLock, updLock)
	assert.False(t, ok)
	assert.Len(t, warnings, 3)

	warnings, ok = NoChangingRPCSignature(updLock, updLock)
	assert.True(t, ok)
	assert.Len(t, warnings, 0)
}

func TestRemovingServiceRPCs(t *testing.T) {
	SetDebug(true)
	curLock := parseTestProto(t, noRemovingServicesRPCsProto)
	updLock := parseTestProto(t, removingServicesRPCsProto)

	warnings, ok := NoRemovingRPCs(curLock, updLock)
	assert.False(t, ok)
	assert.Len(t, warnings, 2)

	warnings, ok = NoRemovingRPCs(updLock, updLock)
	assert.True(t, ok)
	assert.Len(t, warnings, 0)
}

func TestChangingFieldNames(t *testing.T) {
	SetDebug(true)
	curLock := parseTestProto(t, noChangingFieldNamesProto)
	updLock := parseTestProto(t, changingFieldNamesProto)

	warnings, ok := NoChangingFieldNames(curLock, updLock)
	assert.False(t, ok)
	assert.Len(t, warnings, 8)

	warnings, ok = NoChangingFieldNames(updLock, updLock)
	assert.True(t, ok)
	assert.Len(t, warnings, 0)
}

func TestChangingFieldTypes(t *testing.T) {
	SetDebug(true)
	curLock := parseTestProto(t, noChangingFieldTypesProto)
	updLock := parseTestProto(t, changingFieldTypesProto)

	warnings, ok := NoChangingFieldTypes(curLock, updLock)
	assert.False(t, ok)
	assert.Len(t, warnings, 7)

	warnings, ok = NoChangingFieldTypes(updLock, updLock)
	assert.True(t, ok)
	assert.Len(t, warnings, 0)
}

func TestUsingReservedFields(t *testing.T) {
	SetDebug(true)
	curLock := parseTestProto(t, noUsingReservedFieldsProto)
	updLock := parseTestProto(t, usingReservedFieldsProto)

	warnings, ok := NoUsingReservedFields(curLock, updLock)
	assert.False(t, ok)
	assert.Len(t, warnings, 15)

	warnings, ok = NoUsingReservedFields(updLock, updLock)
	assert.True(t, ok)
	assert.Len(t, warnings, 0)
}

func TestRemovingReservedFields(t *testing.T) {
	SetDebug(true)
	curLock := parseTestProto(t, noRemoveReservedFieldsProto)
	updLock := parseTestProto(t, removeReservedFieldsProto)

	warnings, ok := NoRemovingReservedFields(curLock, updLock)
	assert.False(t, ok)
	assert.Len(t, warnings, 13)

	warnings, ok = NoRemovingReservedFields(updLock, updLock)
	assert.True(t, ok)
	assert.Len(t, warnings, 0)
}

func TestChangingFieldIDs(t *testing.T) {
	SetDebug(true)
	curLock := parseTestProto(t, noChangeFieldIDsProto)
	updLock := parseTestProto(t, changeFieldIDsProto)

	warnings, ok := NoChangingFieldIDs(curLock, updLock)
	assert.False(t, ok)
	assert.Len(t, warnings, 7)

	warnings, ok = NoChangingFieldIDs(updLock, updLock)
	assert.True(t, ok)
	assert.Len(t, warnings, 0)
}

func TestRemovingFieldsWithoutReserve(t *testing.T) {
	SetDebug(true)
	curLock := parseTestProto(t, noRemovingFieldsWithoutReserveProto)
	updLock := parseTestProto(t, removingFieldsWithoutReserveProto)

	warnings, ok := NoRemovingFieldsWithoutReserve(curLock, updLock)
	assert.False(t, ok)
	assert.Len(t, warnings, 9)

	warnings, ok = NoRemovingFieldsWithoutReserve(updLock, updLock)
	assert.True(t, ok)
	assert.Len(t, warnings, 0)
}

func TestNoConflictSameNameNestedMessages(t *testing.T) {
	SetDebug(true)
	curLock := parseTestProto(t, noConflictSameNameNestedMessages)

	warnings, ok := NoUsingReservedFields(curLock, curLock)
	assert.True(t, ok)
	assert.Len(t, warnings, 0)
}

func TestShouldConflictReusingFieldsNestedMessages(t *testing.T) {
	SetDebug(true)
	curLock := parseTestProto(t, noConflictSameNameNestedMessages)
	updLock := parseTestProto(t, shouldConflictNestedMessage)

	warnings, ok := NoUsingReservedFields(curLock, updLock)
	assert.False(t, ok)
	assert.Len(t, warnings, 1)
}

func parseTestProto(t *testing.T, proto string) Protolock {
	r := strings.NewReader(proto)
	entry, err := Parse("proto", r)
	assert.NoError(t, err)
	return Protolock{
		Definitions: []Definition{
			{
				Filepath: Protopath("memory/io.Reader"),
				Def:      entry,
			},
		},
	}
}
