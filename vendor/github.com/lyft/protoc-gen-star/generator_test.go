package pgs

import (
	"bytes"
	"os"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	plugin_go "github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	t.Parallel()

	b := &bytes.Buffer{}
	bb := &bytes.Buffer{}
	g := Init(ProtocInput(b), ProtocOutput(bb), func(g *Generator) { /* noop */ })

	assert.NotNil(t, g)
	assert.Equal(t, g.in, b)
	assert.Equal(t, g.out, bb)

	g = Init()
	assert.Equal(t, os.Stdin, g.in)
	assert.Equal(t, os.Stdout, g.out)

	_, ok := g.workflow.(*onceWorkflow)
	assert.True(t, ok)
}

func TestGenerator_RegisterModule(t *testing.T) {
	t.Parallel()

	d := InitMockDebugger()
	g := &Generator{Debugger: d}

	assert.Empty(t, g.mods)
	g.RegisterModule(&mockModule{name: "foo"})

	assert.False(t, d.Failed())
	assert.Len(t, g.mods, 1)

	assert.Panics(t, func() { g.RegisterModule(nil) })
	assert.True(t, d.Failed())
}

func TestGenerator_RegisterPostProcessor(t *testing.T) {
	t.Parallel()

	d := InitMockDebugger()
	p := newPersister()
	g := &Generator{Debugger: d, persister: p}

	pp := &mockPP{}

	assert.Empty(t, p.procs)
	g.RegisterPostProcessor(pp)

	assert.False(t, d.Failed())
	assert.Len(t, p.procs, 1)

	g.RegisterPostProcessor(nil)
	assert.True(t, d.Failed())
}

func TestGenerator_AST(t *testing.T) {
	t.Parallel()

	g := Init()

	wf := &dummyWorkflow{AST: new(graph)}
	g.workflow = wf

	assert.Equal(t, wf.AST, g.AST())
	assert.True(t, wf.initted)
}

func TestGenerator_Render(t *testing.T) {
	// cannot be parallel

	req := &plugin_go.CodeGeneratorRequest{
		FileToGenerate: []string{"foo"},
		ProtoFile: []*descriptor.FileDescriptorProto{
			{
				Name:    proto.String("foo"),
				Syntax:  proto.String("proto2"),
				Package: proto.String("bar"),
			},
		},
	}
	b, err := proto.Marshal(req)
	assert.NoError(t, err)

	buf := &bytes.Buffer{}
	g := Init(ProtocInput(bytes.NewReader(b)), ProtocOutput(buf))
	assert.NotPanics(t, g.Render)

	var res plugin_go.CodeGeneratorResponse
	assert.NoError(t, proto.Unmarshal(buf.Bytes(), &res))
}

func TestGenerator_PushPop(t *testing.T) {
	t.Parallel()

	g := Init()
	g.push("foo")

	pd, ok := g.Debugger.(prefixedDebugger)
	assert.True(t, ok)
	assert.Equal(t, "[foo]", pd.prefix)

	g.pop()

	_, ok = g.Debugger.(rootDebugger)
	assert.True(t, ok)
}
