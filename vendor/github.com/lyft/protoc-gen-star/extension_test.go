package pgs

import (
	"bytes"
	"errors"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
)

func TestExt_FullyQualifiedName(t *testing.T) {
	t.Parallel()

	e := &ext{fqn: "foo"}
	assert.Equal(t, e.fqn, e.FullyQualifiedName())
}

func TestExt_Syntax(t *testing.T) {
	t.Parallel()

	msg := dummyMsg()
	e := &ext{parent: msg}
	assert.Equal(t, msg.Syntax(), e.Syntax())
}

func TestExt_Package(t *testing.T) {
	t.Parallel()

	msg := dummyMsg()
	e := &ext{parent: msg}
	assert.Equal(t, msg.Package(), e.Package())
}

func TestExt_File(t *testing.T) {
	t.Parallel()

	msg := dummyMsg()
	e := &ext{parent: msg}
	assert.Equal(t, msg.File(), e.File())
}

func TestExt_BuildTarget(t *testing.T) {
	t.Parallel()

	msg := dummyMsg()
	e := &ext{parent: msg}
	assert.Equal(t, msg.BuildTarget(), e.BuildTarget())
}

func TestExt_ParentEntity(t *testing.T) {
	t.Parallel()

	msg := dummyMsg()
	e := &ext{parent: msg}
	assert.Equal(t, msg, e.DefinedIn())
}

func TestExt_Extendee(t *testing.T) {
	t.Parallel()

	msg := dummyMsg()
	e := &ext{}
	e.setExtendee(msg)
	assert.Equal(t, msg, e.Extendee())
}

func TestExt_Message(t *testing.T) {
	t.Parallel()

	e := &ext{}
	assert.Nil(t, e.Message())
}

func TestExt_InOneOf(t *testing.T) {
	t.Parallel()

	e := &ext{}
	assert.False(t, e.InOneOf())
}

func TestExt_OneOf(t *testing.T) {
	t.Parallel()

	e := &ext{}
	assert.Nil(t, e.OneOf())
}

func TestExt_Accept(t *testing.T) {
	t.Parallel()

	e := &ext{}

	assert.NoError(t, e.accept(nil))

	v := &mockVisitor{err: errors.New("")}
	assert.Error(t, e.accept(v))
	assert.Equal(t, 1, v.extension)
}

type mockExtractor struct {
	has bool
	get interface{}
	err error
}

func (e *mockExtractor) HasExtension(proto.Message, *proto.ExtensionDesc) bool { return e.has }

func (e *mockExtractor) GetExtension(proto.Message, *proto.ExtensionDesc) (interface{}, error) {
	return e.get, e.err
}

var testExtractor = &mockExtractor{}

func init() { extractor = testExtractor }

func TestExtension(t *testing.T) {
	// cannot be parallel

	defer func() { testExtractor.get = nil }()

	found, err := extension(nil, nil, nil)
	assert.False(t, found)
	assert.NoError(t, err)

	found, err = extension(proto.Message(nil), nil, nil)
	assert.False(t, found)
	assert.NoError(t, err)

	opts := &struct{ proto.Message }{}

	found, err = extension(opts, nil, nil)
	assert.False(t, found)
	assert.Error(t, err)

	desc := &proto.ExtensionDesc{}

	found, err = extension(opts, desc, nil)
	assert.False(t, found)
	assert.Error(t, err)

	type myExt struct{ Name string }

	found, err = extension(opts, desc, &myExt{})
	assert.False(t, found)
	assert.NoError(t, err)

	testExtractor.has = true

	found, err = extension(opts, desc, &myExt{})
	assert.False(t, found)
	assert.NoError(t, err)

	testExtractor.err = errors.New("foo")

	found, err = extension(opts, desc, &myExt{})
	assert.False(t, found)
	assert.Error(t, err)

	testExtractor.err = nil
	testExtractor.get = &myExt{"bar"}

	out := myExt{}

	found, err = extension(opts, desc, out)
	assert.False(t, found)
	assert.Error(t, err)

	found, err = extension(opts, desc, &out)
	assert.True(t, found)
	assert.NoError(t, err)
	assert.Equal(t, "bar", out.Name)

	var ref *myExt
	found, err = extension(opts, desc, &ref)
	assert.True(t, found)
	assert.NoError(t, err)
	assert.Equal(t, "bar", ref.Name)

	found, err = extension(opts, desc, &bytes.Buffer{})
	assert.True(t, found)
	assert.Error(t, err)
}

func TestProtoExtExtractor(t *testing.T) {
	e := protoExtExtractor{}
	assert.NotPanics(t, func() { e.HasExtension(nil, nil) })
	assert.NotPanics(t, func() { e.GetExtension(nil, nil) })
}

// needed to wrapped since there is a Extension method
type mExt interface {
	Extension
}

type mockExtension struct {
	mExt
	err error
}

func (e *mockExtension) accept(v Visitor) error {
	_, err := v.VisitExtension(e)
	if e.err != nil {
		return e.err
	}
	return err
}
