package pgs

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/golang/protobuf/proto"
	plugin_go "github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/stretchr/testify/assert"
)

func TestStandardWorkflow_Init(t *testing.T) {
	t.Parallel()

	req := &plugin_go.CodeGeneratorRequest{FileToGenerate: []string{"foo"}}
	b, err := proto.Marshal(req)
	assert.NoError(t, err)

	mutated := false

	g := Init(ProtocInput(bytes.NewReader(b)), MutateParams(func(p Parameters) { mutated = true }))
	g.workflow.Init(g)

	assert.True(t, mutated)
}

func TestStandardWorkflow_Run(t *testing.T) {
	t.Parallel()

	g := Init()
	g.workflow = &standardWorkflow{Generator: g}
	g.params = Parameters{}

	m := newMockModule()
	m.name = "foo"

	g.RegisterModule(m)
	g.workflow.Run(&graph{})

	assert.True(t, m.executed)
}

func TestStandardWorkflow_Persist(t *testing.T) {
	t.Parallel()

	g := Init(ProtocOutput(ioutil.Discard))
	g.workflow = &standardWorkflow{Generator: g}
	g.persister = dummyPersister(g.Debugger)

	assert.NotPanics(t, func() { g.workflow.Persist(nil) })
}

func TestOnceWorkflow(t *testing.T) {
	t.Parallel()

	d := &dummyWorkflow{
		AST:       &graph{},
		Artifacts: []Artifact{&CustomFile{}},
	}
	wf := &onceWorkflow{workflow: d}

	ast := wf.Init(nil)
	arts := wf.Run(ast)
	wf.Persist(arts)

	assert.True(t, d.initted)
	assert.True(t, d.run)
	assert.True(t, d.persisted)

	d = &dummyWorkflow{}
	wf.workflow = d

	assert.Equal(t, ast, wf.Init(nil))
	assert.Equal(t, arts, wf.Run(ast))
	wf.Persist(arts)

	assert.False(t, d.initted)
	assert.False(t, d.run)
	assert.False(t, d.persisted)
}

type dummyWorkflow struct {
	AST       AST
	Artifacts []Artifact

	initted, run, persisted bool
}

func (wf *dummyWorkflow) Init(g *Generator) AST   { wf.initted = true; return wf.AST }
func (wf *dummyWorkflow) Run(ast AST) []Artifact  { wf.run = true; return wf.Artifacts }
func (wf *dummyWorkflow) Persist(arts []Artifact) { wf.persisted = true }
