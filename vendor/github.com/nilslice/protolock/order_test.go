package protolock

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const ignoreArg = ""

func TestOrder(t *testing.T) {
	cfg, err := NewConfig(".", ".", ignoreArg, false)
	assert.NoError(t, err)

	// verify that the re-production of the same Protolock encoded as json
	// is equivalent to any previously encoded version of the same Protolock
	f, err := os.Open(cfg.LockFilePath())
	assert.NoError(t, err)

	current, err := FromReader(f)
	assert.NoError(t, err)

	r, err := Commit(*cfg)
	assert.NoError(t, err)
	assert.NotNil(t, r)

	updated, err := FromReader(r)
	assert.NoError(t, err)

	assert.Equal(t, current, updated)

	a, err := json.Marshal(current)
	assert.NoError(t, err)
	b, err := json.Marshal(updated)
	assert.NoError(t, err)

	assert.Equal(t, a, b)
}
