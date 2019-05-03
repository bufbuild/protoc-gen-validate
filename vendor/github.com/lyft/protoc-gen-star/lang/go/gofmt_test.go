package pgsgo

import (
	"testing"

	pgs "github.com/lyft/protoc-gen-star"

	"github.com/stretchr/testify/assert"
)

func TestGoFmt_Match(t *testing.T) {
	t.Parallel()

	pp := GoFmt()

	tests := []struct {
		n string
		a pgs.Artifact
		m bool
	}{
		{"GenFile", pgs.GeneratorFile{Name: "foo.go"}, true},
		{"GenFileNonGo", pgs.GeneratorFile{Name: "bar.txt"}, false},

		{"GenTplFile", pgs.GeneratorTemplateFile{Name: "foo.go"}, true},
		{"GenTplFileNonGo", pgs.GeneratorTemplateFile{Name: "bar.txt"}, false},

		{"CustomFile", pgs.CustomFile{Name: "foo.go"}, true},
		{"CustomFileNonGo", pgs.CustomFile{Name: "bar.txt"}, false},

		{"CustomTplFile", pgs.CustomTemplateFile{Name: "foo.go"}, true},
		{"CustomTplFileNonGo", pgs.CustomTemplateFile{Name: "bar.txt"}, false},

		{"NonMatch", pgs.GeneratorAppend{FileName: "foo.go"}, false},
	}

	for _, test := range tests {
		tc := test
		t.Run(tc.n, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.m, pp.Match(tc.a))
		})
	}
}

func TestGoFmt_Process(t *testing.T) {
	t.Parallel()

	src := []byte("// test\n package           foo\n\nvar          bar          int = 123\n")
	exp := []byte("// test\npackage foo\n\nvar bar int = 123\n")

	out, err := GoFmt().Process(src)
	assert.NoError(t, err)
	assert.Equal(t, exp, out)
}
