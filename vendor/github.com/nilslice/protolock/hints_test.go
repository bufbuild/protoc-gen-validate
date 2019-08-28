package protolock

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const hintSkip = `syntax = "proto3";
package dataset;

// @protolock:skip
message Channel {
  reserved 6, 8 to 11;
  int64 id = 1;
  string name = 2;
  string description = 3;
  string foo = 4;
  int32 age = 5;
}

message NextRequest {
	// @protolock:skip
	enum DontTrack {
		NOTHING = 1;
	}
}
// this text before our hint shouldn't matter +(#*)//.~  @protolock:skip
message PreviousRequest {}

// @protolock:skip
// @protolock:no-impl <- not a real hint, should pick up skip for ChannelChanger
// @protolock:internal <- real internal hint, for testing
service ChannelChanger {
  rpc Next(stream NextRequest) returns (Channel);
  rpc Previous(PreviousRequest) returns (stream Channel);
}

// @protolock:skip
message Volume {
	float32 level = 1;
}

// @protolock:skip
enum ShouldSkipEnum {
	ZERO = 0;
}

enum ShouldTrack {
	OK = 1;
}

service VolumeChanger {
	rpc Increase(stream IncreaseRequest) returns (Volume);
	rpc Decrease(DecreaseRequest) returns (Volume);
  }
`

func TestHints(t *testing.T) {
	SetDebug(true)
	lock := parseTestProto(t, hintSkip)

	for _, def := range lock.Definitions {
		t.Run("skip:messages", func(t *testing.T) {
			assert.Len(t, def.Def.Messages, 1)
			assert.Equal(t, def.Def.Messages[0].Name, "NextRequest")
		})
		t.Run("skip:services", func(t *testing.T) {
			assert.Len(t, def.Def.Services, 1)
			assert.Equal(t, def.Def.Services[0].Name, "VolumeChanger")
		})
		t.Run("skip:enums", func(t *testing.T) {
			assert.Len(t, def.Def.Enums, 1)
			assert.Equal(t, def.Def.Enums[0].Name, "ShouldTrack")
		})
	}
}
