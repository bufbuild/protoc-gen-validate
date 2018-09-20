package pgs

import (
	"bytes"
	"math/rand"
	"os"
	"strconv"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestDebugMode(t *testing.T) {
	t.Parallel()

	g := &Generator{}
	assert.False(t, g.debug)

	DebugMode()(g)
	assert.True(t, g.debug)
}

func TestDebugEnv(t *testing.T) {
	t.Parallel()

	g := &Generator{}
	assert.False(t, g.debug)

	e := strconv.Itoa(rand.Int())

	DebugEnv(e)(g)
	assert.False(t, g.debug)

	assert.NoError(t, os.Setenv(e, "1"))
	DebugEnv(e)(g)
	assert.True(t, g.debug)
}

func TestFileSystem(t *testing.T) {
	t.Parallel()

	p := dummyPersister(InitMockDebugger())
	g := &Generator{persister: p}

	fs := afero.NewMemMapFs()
	FileSystem(fs)(g)

	assert.Equal(t, fs, p.fs)
}

func TestProtocInput(t *testing.T) {
	t.Parallel()

	g := &Generator{}
	assert.Nil(t, g.in)

	b := &bytes.Buffer{}
	ProtocInput(b)(g)
	assert.Equal(t, b, g.in)
}

func TestProtocOutput(t *testing.T) {
	t.Parallel()

	g := &Generator{}
	assert.Nil(t, g.out)

	b := &bytes.Buffer{}
	ProtocOutput(b)(g)
	assert.Equal(t, b, g.out)
}
