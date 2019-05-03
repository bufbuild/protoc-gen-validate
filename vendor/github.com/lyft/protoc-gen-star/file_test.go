package pgs

import (
	"errors"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/golang/protobuf/protoc-gen-go/generator"
	"github.com/stretchr/testify/assert"
)

func TestFile_Name(t *testing.T) {
	t.Parallel()

	f := &file{desc: &descriptor.FileDescriptorProto{
		Name: proto.String("foobar"),
	}}

	assert.Equal(t, Name("foobar"), f.Name())
}

func TestFile_FullyQualifiedName(t *testing.T) {
	t.Parallel()

	f := &file{fqn: "foo"}
	assert.Equal(t, f.fqn, f.FullyQualifiedName())
}

func TestFile_Syntax(t *testing.T) {
	t.Parallel()

	f := &file{desc: &descriptor.FileDescriptorProto{}}

	assert.Equal(t, Proto2, f.Syntax())
}

func TestFile_Package(t *testing.T) {
	t.Parallel()

	f := &file{pkg: &pkg{comments: "fizz/buzz"}}
	assert.Equal(t, f.pkg, f.Package())
}

func TestFile_File(t *testing.T) {
	t.Parallel()

	f := &file{buildTarget: true}
	assert.Equal(t, f, f.File())
}

func TestFile_BuildTarget(t *testing.T) {
	t.Parallel()

	f := &file{buildTarget: true}
	assert.True(t, f.BuildTarget())
	f.buildTarget = false
	assert.False(t, f.BuildTarget())
}

func TestFile_Descriptor(t *testing.T) {
	t.Parallel()

	f := &file{desc: &descriptor.FileDescriptorProto{}}
	assert.Equal(t, f.desc, f.Descriptor())
}

func TestFile_InputPath(t *testing.T) {
	t.Parallel()

	f := &file{desc: &descriptor.FileDescriptorProto{Name: proto.String("foo.bar")}}
	assert.Equal(t, "foo.bar", f.InputPath().String())
}

func TestFile_Enums(t *testing.T) {
	t.Parallel()

	f := &file{}

	assert.Empty(t, f.Enums())

	e := &enum{}
	f.addEnum(e)
	assert.Len(t, f.Enums(), 1)
	assert.Equal(t, e, f.Enums()[0])
}

func TestFile_AllEnums(t *testing.T) {
	t.Parallel()

	f := &file{}

	assert.Empty(t, f.AllEnums())

	f.addEnum(&enum{})
	m := &msg{}
	m.addEnum(&enum{})
	f.addMessage(m)

	assert.Len(t, f.Enums(), 1)
	assert.Len(t, f.AllEnums(), 2)
}

func TestFile_Messages(t *testing.T) {
	t.Parallel()

	f := &file{}

	assert.Empty(t, f.Messages())

	m := &msg{}
	f.addMessage(m)
	assert.Len(t, f.Messages(), 1)
	assert.Equal(t, m, f.Messages()[0])
}

func TestFile_MapEntries(t *testing.T) {
	t.Parallel()
	f := &file{}

	assert.Panics(t, func() { f.addMapEntry(&msg{}) })
	assert.Empty(t, f.MapEntries())
}

func TestFile_AllMessages(t *testing.T) {
	t.Parallel()

	f := &file{}

	assert.Empty(t, f.AllMessages())

	m := &msg{}
	m.addMessage(&msg{})
	f.addMessage(m)

	assert.Len(t, f.Messages(), 1)
	assert.Len(t, f.AllMessages(), 2)
}

func TestFile_Services(t *testing.T) {
	t.Parallel()

	f := &file{}

	assert.Empty(t, f.Services())

	s := &service{}
	f.addService(s)

	assert.Len(t, f.Services(), 1)
	assert.Equal(t, s, f.Services()[0])
}

func TestFile_Imports(t *testing.T) {
	t.Parallel()

	m := &msg{}
	m.addMessage(&mockMessage{i: []File{&file{}}, Message: &msg{}})
	svc := &mockService{i: []File{&file{}}, Service: &service{}}

	f := &file{}
	assert.Empty(t, f.Imports())

	f.addMessage(m)
	f.addService(svc)
	assert.Len(t, f.Imports(), 1)

	nf := &file{desc: &descriptor.FileDescriptorProto{
		Name: proto.String("foobar"),
	}}
	f.addMessage(&mockMessage{i: []File{nf}, Message: &msg{}})
	assert.Len(t, f.Imports(), 2)
}

func TestFile_Accept(t *testing.T) {
	t.Parallel()

	f := &file{}

	assert.Nil(t, f.accept(nil))

	v := &mockVisitor{}
	assert.NoError(t, f.accept(v))
	assert.Equal(t, 1, v.file)

	v.Reset()
	v.v = v
	v.err = errors.New("foo")
	assert.Equal(t, v.err, f.accept(v))
	assert.Equal(t, 1, v.file)
	assert.Zero(t, v.enum)
	assert.Zero(t, v.message)
	assert.Zero(t, v.service)
	assert.Zero(t, v.extension)

	v.Reset()
	f.addEnum(&enum{})
	f.addMessage(&msg{})
	f.addService(&service{})
	f.addDefExtension(&ext{})
	assert.NoError(t, f.accept(v))
	assert.Equal(t, 1, v.file)
	assert.Equal(t, 1, v.enum)
	assert.Equal(t, 1, v.message)
	assert.Equal(t, 1, v.service)
	assert.Equal(t, 1, v.extension)

	v.Reset()
	f.addDefExtension(&mockExtension{err: errors.New("fizz")})
	assert.EqualError(t, f.accept(v), "fizz")
	assert.Equal(t, 1, v.file)
	assert.Equal(t, 1, v.enum)
	assert.Equal(t, 1, v.message)
	assert.Equal(t, 1, v.service)
	assert.Equal(t, 2, v.extension)

	v.Reset()
	f.addService(&mockService{err: errors.New("fizz")})
	assert.EqualError(t, f.accept(v), "fizz")
	assert.Equal(t, 1, v.file)
	assert.Equal(t, 1, v.enum)
	assert.Equal(t, 1, v.message)
	assert.Equal(t, 2, v.service)
	assert.Zero(t, v.extension)

	v.Reset()
	f.addMessage(&mockMessage{err: errors.New("bar")})
	assert.EqualError(t, f.accept(v), "bar")
	assert.Equal(t, 1, v.file)
	assert.Equal(t, 1, v.enum)
	assert.Equal(t, 2, v.message)
	assert.Zero(t, v.service)
	assert.Zero(t, v.extension)

	v.Reset()
	f.addEnum(&mockEnum{err: errors.New("baz")})
	assert.EqualError(t, f.accept(v), "baz")
	assert.Equal(t, 1, v.file)
	assert.Equal(t, 2, v.enum)
	assert.Zero(t, v.message)
	assert.Zero(t, v.service)
	assert.Zero(t, v.extension)
}

func TestFile_Extension(t *testing.T) {
	// cannot be parallel

	assert.NotPanics(t, func() {
		(&file{
			desc: &descriptor.FileDescriptorProto{},
		}).Extension(nil, nil)
	})
}

func TestFile_DefinedExtensions(t *testing.T) {
	t.Parallel()

	f := &file{}
	assert.Empty(t, f.DefinedExtensions())

	ext := &ext{}
	f.addDefExtension(ext)
	assert.Len(t, f.DefinedExtensions(), 1)
}

// needed to wrap since there is a File method
type mFile interface {
	File
}

type mockFile struct {
	mFile
	pkg Package
	err error
}

func (f *mockFile) setPackage(p Package) {
	f.pkg = p
}

func (f *mockFile) accept(v Visitor) error {
	_, err := v.VisitFile(f)
	if f.err != nil {
		return f.err
	}
	return err
}

func dummyFile() *file {
	pkg := dummyPkg()
	f := &file{
		pkg: pkg,
		desc: &descriptor.FileDescriptorProto{
			Package: proto.String(pkg.ProtoName().String()),
			Syntax:  proto.String(string(Proto3)),
			Name:    proto.String("file.proto"),
		},
	}
	pkg.addFile(f)

	return f
}

func dummyGenFile() (*file, *generator.FileDescriptor) {
	f := dummyFile()
	return f, &generator.FileDescriptor{FileDescriptorProto: f.desc}
}
