package pgs

import (
	"testing"

	"errors"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/stretchr/testify/assert"
)

func TestMethod_Name(t *testing.T) {
	t.Parallel()
	m := &method{desc: &descriptor.MethodDescriptorProto{Name: proto.String("foo")}}
	assert.Equal(t, "foo", m.Name().String())
}

func TestMethod_FullyQualifiedName(t *testing.T) {
	t.Parallel()

	m := &method{fqn: "fizz"}
	assert.Equal(t, m.fqn, m.FullyQualifiedName())
}

func TestMethod_Syntax(t *testing.T) {
	t.Parallel()
	m := &method{}
	s := dummyService()
	s.addMethod(m)
	assert.Equal(t, s.Syntax(), m.Syntax())
}

func TestMethod_Package(t *testing.T) {
	t.Parallel()
	m := &method{}
	s := dummyService()
	s.addMethod(m)

	assert.NotNil(t, m.Package())
	assert.Equal(t, s.Package(), m.Package())
}

func TestMethod_File(t *testing.T) {
	t.Parallel()
	m := &method{}
	s := dummyService()
	s.addMethod(m)

	assert.NotNil(t, m.File())
	assert.Equal(t, s.File(), m.File())
}

func TestMethod_BuildTarget(t *testing.T) {
	t.Parallel()
	m := &method{}
	s := dummyService()
	s.addMethod(m)

	assert.False(t, m.BuildTarget())
	s.setFile(&file{buildTarget: true})
	assert.True(t, m.BuildTarget())
}

func TestMethod_Descriptor(t *testing.T) {
	t.Parallel()
	m := &method{desc: &descriptor.MethodDescriptorProto{}}
	assert.Equal(t, m.desc, m.Descriptor())
}

func TestMethod_Service(t *testing.T) {
	t.Parallel()
	m := &method{}
	s := dummyService()
	s.addMethod(m)

	assert.Equal(t, s, m.Service())
}

func TestMethod_Input(t *testing.T) {
	t.Parallel()
	m := &method{in: &msg{}}
	assert.Equal(t, m.in, m.Input())
}

func TestMethod_Output(t *testing.T) {
	t.Parallel()
	m := &method{out: &msg{}}
	assert.Equal(t, m.out, m.Output())
}

func TestMethod_ClientStreaming(t *testing.T) {
	t.Parallel()

	m := &method{desc: &descriptor.MethodDescriptorProto{}}
	assert.False(t, m.ClientStreaming())
	m.desc.ClientStreaming = proto.Bool(true)
	assert.True(t, m.ClientStreaming())
}

func TestMethod_ServerStreaming(t *testing.T) {
	t.Parallel()

	m := &method{desc: &descriptor.MethodDescriptorProto{}}
	assert.False(t, m.ServerStreaming())
	m.desc.ServerStreaming = proto.Bool(true)
	assert.True(t, m.ServerStreaming())
}

func TestMethod_BiDirStreaming(t *testing.T) {
	t.Parallel()

	m := &method{desc: &descriptor.MethodDescriptorProto{}}
	assert.False(t, m.BiDirStreaming())
	m.desc.ServerStreaming = proto.Bool(true)
	assert.False(t, m.BiDirStreaming())
	m.desc.ServerStreaming = proto.Bool(false)
	m.desc.ClientStreaming = proto.Bool(true)
	assert.False(t, m.BiDirStreaming())
	m.desc.ServerStreaming = proto.Bool(true)
	assert.True(t, m.BiDirStreaming())
}

func TestMethod_Imports(t *testing.T) {
	t.Parallel()

	s := dummyService()
	m := &method{
		in:  dummyMsg(),
		out: dummyMsg(),
	}
	s.addMethod(m)

	f := &file{desc: &descriptor.FileDescriptorProto{
		Name: proto.String("foobar"),
	}}
	assert.Empty(t, m.Imports())
	m.in = &msg{parent: f}
	assert.Len(t, m.Imports(), 1)

	m.out = &msg{parent: &file{}}
	assert.Len(t, m.Imports(), 2)

	m.out = &msg{parent: f}
	assert.Len(t, m.Imports(), 1)
}

func TestMethod_Extension(t *testing.T) {
	// cannot be parallel

	m := &method{desc: &descriptor.MethodDescriptorProto{}}
	assert.NotPanics(t, func() { m.Extension(nil, nil) })
}

func TestMethod_Accept(t *testing.T) {
	t.Parallel()

	m := &method{}

	assert.Nil(t, m.accept(nil))

	v := &mockVisitor{err: errors.New("foo")}
	assert.Error(t, m.accept(v))
	assert.Equal(t, 1, v.method)
}

func TestMethod_ChildAtPath(t *testing.T) {
	t.Parallel()

	m := &method{}

	assert.Equal(t, m, m.childAtPath(nil))
	assert.Nil(t, m.childAtPath([]int32{1}))
}

type mockMethod struct {
	Method
	i   []File
	s   Service
	err error
}

func (m *mockMethod) Imports() []File { return m.i }

func (m *mockMethod) setService(s Service) { m.s = s }

func (m *mockMethod) accept(v Visitor) error {
	_, err := v.VisitMethod(m)
	if m.err != nil {
		return m.err
	}
	return err
}
