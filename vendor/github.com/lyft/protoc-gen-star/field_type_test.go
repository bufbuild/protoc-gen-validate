package pgs

import (
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/stretchr/testify/assert"
)

func TestScalarT_Field(t *testing.T) {
	t.Parallel()

	f := dummyField()
	s := &scalarT{}
	f.addType(s)

	assert.Equal(t, f, s.Field())
}

func TestScalarT_IsRepeated(t *testing.T) {
	t.Parallel()
	s := &scalarT{}
	assert.False(t, s.IsRepeated())
}

func TestScalarT_IsMap(t *testing.T) {
	t.Parallel()
	s := &scalarT{}
	assert.False(t, s.IsMap())
}

func TestScalarT_IsEnum(t *testing.T) {
	t.Parallel()
	s := &scalarT{}
	assert.False(t, s.IsEnum())
}

func TestScalarT_IsEmbed(t *testing.T) {
	t.Parallel()
	s := &scalarT{}
	assert.False(t, s.IsEmbed())
}

func TestScalarT_ProtoType(t *testing.T) {
	t.Parallel()
	f := dummyField()
	s := &scalarT{}
	f.addType(s)

	assert.Equal(t, f.desc.GetType(), s.ProtoType().Proto())
}

func TestScalarT_ProtoLabel(t *testing.T) {
	t.Parallel()
	f := dummyField()
	s := &scalarT{}
	f.addType(s)

	assert.Equal(t, f.desc.GetLabel(), s.ProtoLabel().Proto())
}

func TestScalarT_Imports(t *testing.T) {
	t.Parallel()
	assert.Nil(t, (&scalarT{}).Imports())
}

func TestScalarT_Enum(t *testing.T) {
	t.Parallel()
	assert.Nil(t, (&scalarT{}).Enum())
}

func TestScalarT_Embed(t *testing.T) {
	t.Parallel()
	assert.Nil(t, (&scalarT{}).Embed())
}

func TestScalarT_Element(t *testing.T) {
	t.Parallel()
	assert.Nil(t, (&scalarT{}).Element())
}

func TestScalarT_Key(t *testing.T) {
	t.Parallel()
	assert.Nil(t, (&scalarT{}).Key())
}

func TestScalarT_IsOptional(t *testing.T) {
	t.Parallel()

	s := &scalarT{}
	f := dummyField()
	f.addType(s)

	assert.True(t, s.IsOptional())

	fl := dummyFile()
	fl.desc.Syntax = nil
	f.Message().setParent(fl)

	assert.True(t, s.IsOptional())

	req := descriptor.FieldDescriptorProto_LABEL_REQUIRED
	f.desc.Label = &req

	assert.False(t, s.IsOptional())
}

func TestScalarT_IsRequired(t *testing.T) {
	t.Parallel()

	s := &scalarT{}
	f := dummyField()
	f.addType(s)

	assert.False(t, s.IsRequired())

	fl := dummyFile()
	fl.desc.Syntax = nil
	f.Message().setParent(fl)

	assert.False(t, s.IsRequired())

	req := descriptor.FieldDescriptorProto_LABEL_REQUIRED
	f.desc.Label = &req

	assert.True(t, s.IsRequired())
}

func TestScalarT_ToElem(t *testing.T) {
	t.Parallel()

	s := &scalarT{}
	f := dummyField()
	f.addType(s)

	el := s.toElem()
	assert.Equal(t, s, el.ParentType())
	assert.Equal(t, s.ProtoType(), el.ProtoType())
}

func TestEnumT_Enum(t *testing.T) {
	t.Parallel()
	e := &enumT{enum: &enum{}}
	assert.Equal(t, e.enum, e.Enum())
}

func TestEnumT_IsEnum(t *testing.T) {
	t.Parallel()
	e := &enumT{}
	assert.True(t, e.IsEnum())
}

func TestEnumT_Imports(t *testing.T) {
	t.Parallel()

	f := dummyFile()
	en := dummyEnum()
	en.parent = f
	e := &enumT{scalarT: &scalarT{}, enum: en}
	fld := dummyField()
	fld.addType(e)

	assert.Empty(t, e.Imports())

	f.desc.Name = proto.String("some/other/file.proto")
	assert.Len(t, e.Imports(), 1)
	assert.Equal(t, e.enum.File(), e.Imports()[0])
}

func TestEnumT_ToElem(t *testing.T) {
	t.Parallel()

	e := &enumT{
		scalarT: &scalarT{},
		enum:    dummyEnum(),
	}
	f := dummyField()
	f.addType(e)

	el := e.toElem()
	assert.True(t, el.IsEnum())
	assert.Equal(t, e.enum, el.Enum())
	assert.Equal(t, e.ProtoType(), el.ProtoType())
}

func TestEmbedT_IsEmbed(t *testing.T) {
	t.Parallel()
	e := &embedT{}
	assert.True(t, e.IsEmbed())
}

func TestEmbedT_Embed(t *testing.T) {
	t.Parallel()
	e := &embedT{msg: dummyMsg()}
	assert.Equal(t, e.msg, e.Embed())
}

func TestEmbedT_Imports(t *testing.T) {
	t.Parallel()

	msg := dummyMsg()
	f := dummyFile()
	msg.parent = f
	e := &embedT{scalarT: &scalarT{}, msg: msg}
	dummyField().addType(e)

	assert.Empty(t, e.Imports())

	f.desc.Name = proto.String("some/other/file.proto")
	assert.Len(t, e.Imports(), 1)
	assert.Equal(t, e.msg.File(), e.Imports()[0])
}

func TestEmbedT_ToElem(t *testing.T) {
	t.Parallel()

	e := &embedT{
		scalarT: &scalarT{},
		msg:     dummyMsg(),
	}
	f := dummyField()
	f.addType(e)

	el := e.toElem()
	assert.True(t, el.IsEmbed())
	assert.Equal(t, e.msg, el.Embed())
	assert.Equal(t, e.ProtoType(), el.ProtoType())
}

func TestRepT_IsRepeated(t *testing.T) {
	t.Parallel()
	r := &repT{}
	assert.True(t, r.IsRepeated())
}

func TestRepT_Element(t *testing.T) {
	t.Parallel()
	r := &repT{el: &scalarE{}}
	assert.Equal(t, r.el, r.Element())
}

func TestRepT_Imports(t *testing.T) {
	t.Parallel()

	msg := dummyMsg()
	f := dummyFile()
	msg.parent = f
	e := &embedT{scalarT: &scalarT{}, msg: msg}
	dummyField().addType(e)

	fld := dummyField()
	r := &repT{scalarT: &scalarT{}, el: e.toElem()}
	fld.addType(r)

	assert.Empty(t, r.Imports())

	f.desc.Name = proto.String("some/other/file.proto")
	assert.Len(t, r.Imports(), 1)
	assert.Equal(t, r.el.Embed().File(), r.Imports()[0])
}

func TestRepT_ToElem(t *testing.T) {
	t.Parallel()
	assert.Panics(t, func() { (&repT{}).toElem() })
}

func TestMapT_IsRepeated(t *testing.T) {
	t.Parallel()
	assert.False(t, (&mapT{}).IsRepeated())
}

func TestMapT_IsMap(t *testing.T) {
	t.Parallel()
	assert.True(t, (&mapT{}).IsMap())
}

func TestMapT_Key(t *testing.T) {
	t.Parallel()
	m := &mapT{key: &scalarE{}}
	assert.Equal(t, m.key, m.Key())
}

type mockT struct {
	FieldType
	i   []File
	f   Field
	err error
}

func (t *mockT) Imports() []File { return t.i }

func (t *mockT) setField(f Field) { t.f = f }
