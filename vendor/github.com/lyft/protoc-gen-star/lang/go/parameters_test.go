package pgsgo

import (
	"testing"

	pgs "github.com/lyft/protoc-gen-star"
	"github.com/stretchr/testify/assert"
)

func TestParameters_Plugins(t *testing.T) {
	t.Parallel()

	p := pgs.Parameters{}
	plugins, all := Plugins(p)
	assert.Empty(t, plugins)
	assert.False(t, all)

	p[pluginsKey] = "foo+bar"
	plugins, all = Plugins(p)
	assert.Equal(t, []string{"foo", "bar"}, plugins)
	assert.False(t, all)

	p[pluginsKey] = ""
	plugins, all = Plugins(p)
	assert.Empty(t, plugins)
	assert.True(t, all)
}

func TestParameters_HasPlugin(t *testing.T) {
	t.Parallel()

	p := pgs.Parameters{}
	assert.False(t, HasPlugin(p, "foo"))

	p[pluginsKey] = "foo"
	assert.True(t, HasPlugin(p, "foo"))

	p[pluginsKey] = ""
	assert.True(t, HasPlugin(p, "foo"))

	p[pluginsKey] = "bar"
	assert.False(t, HasPlugin(p, "foo"))
}

func TestParameters_AddPlugin(t *testing.T) {
	t.Parallel()

	p := pgs.Parameters{}
	AddPlugin(p, "foo", "bar")
	assert.Equal(t, "foo+bar", p[pluginsKey])

	AddPlugin(p, "baz")
	assert.Equal(t, "foo+bar+baz", p[pluginsKey])

	AddPlugin(p)
	assert.Equal(t, "foo+bar+baz", p[pluginsKey])

	p[pluginsKey] = ""
	AddPlugin(p, "fizz", "buzz")
	assert.Equal(t, "", p[pluginsKey])
}

func TestParameters_EnableAllPlugins(t *testing.T) {
	t.Parallel()

	p := pgs.Parameters{pluginsKey: "foo"}
	_, all := Plugins(p)
	assert.False(t, all)

	EnableAllPlugins(p)
	_, all = Plugins(p)
	assert.True(t, all)
}

func TestParameters_ImportPrefix(t *testing.T) {
	t.Parallel()

	p := pgs.Parameters{}
	assert.Empty(t, ImportPrefix(p))
	SetImportPrefix(p, "foo")
	assert.Equal(t, "foo", ImportPrefix(p))
}

func TestParameters_ImportPath(t *testing.T) {
	t.Parallel()

	p := pgs.Parameters{}
	assert.Empty(t, ImportPath(p))
	SetImportPath(p, "foo")
	assert.Equal(t, "foo", ImportPath(p))
}

func TestParameters_ImportMap(t *testing.T) {
	t.Parallel()

	p := pgs.Parameters{
		"Mfoo.proto":       "bar",
		"Mfizz/buzz.proto": "baz",
	}

	AddImportMapping(p, "quux.proto", "shme")

	tests := []struct {
		proto, path string
		exists      bool
	}{
		{"quux.proto", "shme", true},
		{"foo.proto", "bar", true},
		{"fizz/buzz.proto", "baz", true},
		{"abcde.proto", "", false},
	}

	for _, test := range tests {
		t.Run(test.proto, func(t *testing.T) {
			path, ok := MappedImport(p, test.proto)
			if test.exists {
				assert.True(t, ok)
				assert.Equal(t, test.path, path)
			} else {
				assert.False(t, ok)
			}
		})
	}
}

func TestParameters_Paths(t *testing.T) {
	t.Parallel()

	p := pgs.Parameters{}

	assert.Equal(t, ImportPathRelative, Paths(p))
	SetPaths(p, SourceRelative)
	assert.Equal(t, SourceRelative, Paths(p))
}
