package pgs

import (
	"html/template"
	"testing"

	"errors"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestPersister_Persist_Unrecognized(t *testing.T) {
	t.Parallel()

	d := InitMockDebugger()
	p := dummyPersister(d)

	p.Persist(nil)

	assert.True(t, d.Failed())
}

func TestPersister_Persist_GeneratorFile(t *testing.T) {
	t.Parallel()

	d := InitMockDebugger()
	p := dummyPersister(d)
	fs := afero.NewMemMapFs()
	p.SetFS(fs)

	resp := p.Persist(
		GeneratorFile{
			Name:     "foo",
			Contents: "bar",
		},
		GeneratorFile{
			Name:     "quux",
			Contents: "baz",
		},
		GeneratorFile{
			Name:      "foo",
			Contents:  "fizz",
			Overwrite: true,
		})

	assert.Len(t, resp.File, 2)
	assert.Equal(t, "foo", resp.File[0].GetName())
	assert.Equal(t, "fizz", resp.File[0].GetContent())
}

var genTpl = template.Must(template.New("good").Parse("{{ . }}"))

func TestPersister_Persist_GeneratorTemplateFile(t *testing.T) {
	t.Parallel()

	d := InitMockDebugger()
	p := dummyPersister(d)
	fs := afero.NewMemMapFs()
	p.SetFS(fs)

	resp := p.Persist(
		GeneratorTemplateFile{
			Name: "foo",
			TemplateArtifact: TemplateArtifact{
				Template: genTpl,
				Data:     "bar",
			},
		},
		GeneratorTemplateFile{
			Name: "quux",
			TemplateArtifact: TemplateArtifact{
				Template: genTpl,
				Data:     "baz",
			},
		},
		GeneratorTemplateFile{
			Name: "foo",
			TemplateArtifact: TemplateArtifact{
				Template: genTpl,
				Data:     "fizz",
			},
			Overwrite: true,
		},
	)

	assert.Len(t, resp.File, 2)
	assert.Equal(t, "foo", resp.File[0].GetName())
	assert.Equal(t, "fizz", resp.File[0].GetContent())
}

func TestPersister_Persist_GeneratorAppend(t *testing.T) {
	t.Parallel()

	d := InitMockDebugger()
	p := dummyPersister(d)
	fs := afero.NewMemMapFs()
	p.SetFS(fs)

	resp := p.Persist(
		GeneratorFile{Name: "foo"},
		GeneratorFile{Name: "bar"},
		GeneratorAppend{
			FileName: "foo",
			Contents: "baz",
		},
		GeneratorAppend{
			FileName: "bar",
			Contents: "quux",
		},
	)

	assert.Len(t, resp.File, 4)
	assert.Equal(t, "", resp.File[1].GetName())
	assert.Equal(t, "baz", resp.File[1].GetContent())
	assert.Equal(t, "", resp.File[3].GetName())
	assert.Equal(t, "quux", resp.File[3].GetContent())

	p.Persist(GeneratorAppend{FileName: "doesNotExist"})

	assert.True(t, d.Failed())
}

func TestPersister_Persist_GeneratorTemplateAppend(t *testing.T) {
	t.Parallel()

	d := InitMockDebugger()
	p := dummyPersister(d)
	fs := afero.NewMemMapFs()
	p.SetFS(fs)

	resp := p.Persist(
		GeneratorFile{Name: "foo"},
		GeneratorFile{Name: "bar"},
		GeneratorTemplateAppend{
			FileName: "foo",
			TemplateArtifact: TemplateArtifact{
				Template: genTpl,
				Data:     "baz",
			},
		}, GeneratorTemplateAppend{
			FileName: "bar",
			TemplateArtifact: TemplateArtifact{
				Template: genTpl,
				Data:     "quux",
			},
		},
	)

	assert.Len(t, resp.File, 4)
	assert.Equal(t, "", resp.File[1].GetName())
	assert.Equal(t, "baz", resp.File[1].GetContent())
	assert.Equal(t, "", resp.File[3].GetName())
	assert.Equal(t, "quux", resp.File[3].GetContent())

	resp = p.Persist(GeneratorTemplateAppend{
		FileName: "doesNotExist",
		TemplateArtifact: TemplateArtifact{
			Template: genTpl,
			Data:     "baz",
		},
	})

	assert.True(t, d.Failed())
}

func TestPersister_Persist_GeneratorInjection(t *testing.T) {
	t.Parallel()

	d := InitMockDebugger()
	p := dummyPersister(d)
	fs := afero.NewMemMapFs()
	p.SetFS(fs)

	resp := p.Persist(GeneratorInjection{
		FileName:       "foo",
		InsertionPoint: "bar",
		Contents:       "baz",
	})

	assert.Len(t, resp.File, 1)
	assert.Equal(t, "foo", resp.File[0].GetName())
	assert.Equal(t, "bar", resp.File[0].GetInsertionPoint())
	assert.Equal(t, "baz", resp.File[0].GetContent())
}

func TestPersister_Persist_GeneratorTemplateInjection(t *testing.T) {
	t.Parallel()

	d := InitMockDebugger()
	p := dummyPersister(d)
	fs := afero.NewMemMapFs()
	p.SetFS(fs)

	resp := p.Persist(GeneratorTemplateInjection{
		FileName:       "foo",
		InsertionPoint: "bar",
		TemplateArtifact: TemplateArtifact{
			Template: genTpl,
			Data:     "baz",
		},
	})

	assert.Len(t, resp.File, 1)
	assert.Equal(t, "foo", resp.File[0].GetName())
	assert.Equal(t, "bar", resp.File[0].GetInsertionPoint())
	assert.Equal(t, "baz", resp.File[0].GetContent())
}

func TestPersister_Persist_CustomFile(t *testing.T) {
	t.Parallel()

	d := InitMockDebugger()
	p := dummyPersister(d)
	fs := afero.NewMemMapFs()
	p.SetFS(fs)

	p.Persist(CustomFile{
		Name:     "foo/bar/baz.txt",
		Perms:    0655,
		Contents: "fizz",
	})

	b, err := afero.ReadFile(fs, "foo/bar/baz.txt")
	assert.NoError(t, err)
	assert.Equal(t, "fizz", string(b))

	p.Persist(CustomFile{
		Name:     "foo/bar/baz.txt",
		Perms:    0655,
		Contents: "buzz",
	})

	b, err = afero.ReadFile(fs, "foo/bar/baz.txt")
	assert.NoError(t, err)
	assert.Equal(t, "fizz", string(b))

	p.Persist(CustomFile{
		Name:      "foo/bar/baz.txt",
		Perms:     0655,
		Contents:  "buzz",
		Overwrite: true,
	})

	b, err = afero.ReadFile(fs, "foo/bar/baz.txt")
	assert.NoError(t, err)
	assert.Equal(t, "buzz", string(b))
}

func TestPersister_Persist_CustomTemplateFile(t *testing.T) {
	t.Parallel()

	d := InitMockDebugger()
	p := dummyPersister(d)
	fs := afero.NewMemMapFs()
	p.SetFS(fs)

	p.Persist(CustomTemplateFile{
		Name:  "foo/bar/baz.txt",
		Perms: 0655,
		TemplateArtifact: TemplateArtifact{
			Template: genTpl,
			Data:     "fizz",
		},
	})

	b, err := afero.ReadFile(fs, "foo/bar/baz.txt")
	assert.NoError(t, err)
	assert.Equal(t, "fizz", string(b))

	p.Persist(CustomTemplateFile{
		Name:  "foo/bar/baz.txt",
		Perms: 0655,
		TemplateArtifact: TemplateArtifact{
			Template: genTpl,
			Data:     "buzz",
		},
	})

	b, err = afero.ReadFile(fs, "foo/bar/baz.txt")
	assert.NoError(t, err)
	assert.Equal(t, "fizz", string(b))

	p.Persist(CustomTemplateFile{
		Name:  "foo/bar/baz.txt",
		Perms: 0655,
		TemplateArtifact: TemplateArtifact{
			Template: genTpl,
			Data:     "buzz",
		},
		Overwrite: true,
	})

	b, err = afero.ReadFile(fs, "foo/bar/baz.txt")
	assert.NoError(t, err)
	assert.Equal(t, "buzz", string(b))
}

func TestPersister_AddPostProcessor(t *testing.T) {
	t.Parallel()

	p := dummyPersister(InitMockDebugger())

	good := &mockPP{match: true, out: []byte("good")}
	bad := &mockPP{err: errors.New("should not be called")}

	p.AddPostProcessor(good, bad)
	out := p.postProcess(GeneratorFile{}, "")
	assert.Equal(t, "good", out)
}

func dummyPersister(d Debugger) *stdPersister {
	return &stdPersister{
		Debugger: d,
		fs:       afero.NewMemMapFs(),
	}
}

func TestPersister_Persist_GeneratorError(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		input  []Artifact
		output string
	}{
		"no errors": {
			[]Artifact{},
			"",
		},
		"one error": {
			[]Artifact{GeneratorError{Message: "something went wrong"}},
			"something went wrong",
		},
		"two errors": {
			[]Artifact{
				GeneratorError{Message: "something went wrong"},
				GeneratorError{Message: "something else went wrong, too"},
			},
			"something went wrong; something else went wrong, too",
		},
	}
	for desc, tc := range cases {
		t.Run(desc, func(t *testing.T) {
			d := InitMockDebugger()
			p := dummyPersister(d)
			fs := afero.NewMemMapFs()
			p.SetFS(fs)

			resp := p.Persist(tc.input...)

			assert.Len(t, resp.File, 0)
			assert.Equal(t, tc.output, resp.GetError())
		})
	}
}
