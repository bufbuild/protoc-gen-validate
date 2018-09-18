package pgs

import (
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"

	"errors"

	"github.com/stretchr/testify/assert"
)

func TestPkg_ProtoName(t *testing.T) {
	t.Parallel()

	p := dummyPkg()
	assert.Equal(t, p.fd.GetPackage(), p.ProtoName().String())
}

func TestPkg_Files(t *testing.T) {
	t.Parallel()

	p := &pkg{}
	assert.Empty(t, p.Files())

	p.addFile(&file{})
	p.addFile(&file{})
	p.addFile(&file{})

	assert.Len(t, p.Files(), 3)
}

func TestPkg_AddFile(t *testing.T) {
	t.Parallel()

	p := &pkg{}
	f := &file{}
	p.addFile(f)
	assert.Len(t, p.files, 1)
	assert.EqualValues(t, f, p.files[0])
}

func TestPkg_Accept(t *testing.T) {
	t.Parallel()

	p := &pkg{
		files: []File{&mockFile{}},
	}
	assert.Nil(t, p.accept(nil))

	v := &mockVisitor{}
	assert.NoError(t, p.accept(v))
	assert.Equal(t, 1, v.pkg)
	assert.Zero(t, v.file)

	v.Reset()
	v.err = errors.New("foobar")
	assert.EqualError(t, p.accept(v), "foobar")
	assert.Equal(t, 1, v.pkg)
	assert.Zero(t, v.file)

	v.Reset()
	v.v = v
	assert.NoError(t, p.accept(v))
	assert.Equal(t, 1, v.pkg)
	assert.Equal(t, 1, v.file)

	v.Reset()
	p.addFile(&mockFile{err: errors.New("fizzbuzz")})
	assert.EqualError(t, p.accept(v), "fizzbuzz")
	assert.Equal(t, 1, v.pkg)
	assert.Equal(t, 2, v.file)
}

func TestPackage_Comments(t *testing.T) {
	t.Parallel()

	pkg := dummyPkg()
	pkg.setComments("foobar")
	assert.Equal(t, "foobar", pkg.Comments())
}

func dummyPkg() *pkg {
	return &pkg{
		fd: &descriptor.FileDescriptorProto{Package: proto.String("pkg_name")},
	}
}
