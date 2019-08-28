package protolock

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

var osp = filepath.Join("testdata", "test.proto")

func TestOSPathToProtoPath(t *testing.T) {
	path := Protopath(osp)
	p := ProtoPath(path)
	assert.Equal(t, "testdata:/:test.proto", string(p))
	assert.Equal(t, Protopath("testdata:/:test.proto"), p)
}

func TestProtoPathToOSPath(t *testing.T) {
	path := Protopath("testdata:/:test.proto")
	p := OSPath(path)
	assert.Equal(t, Protopath(osp), p)
	assert.Equal(t, osp, string(p))
}
