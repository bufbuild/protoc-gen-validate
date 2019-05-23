package pgsgo

import (
	"fmt"
	"testing"

	pgs "github.com/lyft/protoc-gen-star"
	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"
)

func TestType(t *testing.T) {
	t.Parallel()

	ast := buildGraph(t, "names", "types")
	ctx := loadContext(t, "names", "types")

	tests := []struct {
		field    string
		expected TypeName
	}{
		// proto2 syntax, optional
		{"Proto2.double", "*float64"},
		{"Proto2.float", "*float32"},
		{"Proto2.int64", "*int64"},
		{"Proto2.sfixed64", "*int64"},
		{"Proto2.sint64", "*int64"},
		{"Proto2.uint64", "*uint64"},
		{"Proto2.fixed64", "*uint64"},
		{"Proto2.int32", "*int32"},
		{"Proto2.sfixed32", "*int32"},
		{"Proto2.sint32", "*int32"},
		{"Proto2.uint32", "*uint32"},
		{"Proto2.fixed32", "*uint32"},
		{"Proto2.bool", "*bool"},
		{"Proto2.string", "*string"},
		{"Proto2.bytes", "[]byte"},
		{"Proto2.enum", "*Proto2_Enum"},
		{"Proto2.ext_enum", "*ptype.Syntax"},
		{"Proto2.msg", "*Proto2_Required"},
		{"Proto2.ext_msg", "*duration.Duration"},
		{"Proto2.repeated_scalar", "[]float64"},
		{"Proto2.repeated_enum", "[]Proto2_Enum"},
		{"Proto2.repeated_ext_enum", "[]ptype.Syntax"},
		{"Proto2.repeated_msg", "[]*Proto2_Required"},
		{"Proto2.repeated_ext_msg", "[]*duration.Duration"},
		{"Proto2.map_scalar", "map[string]float32"},
		{"Proto2.map_enum", "map[int32]Proto2_Enum"},
		{"Proto2.map_ext_enum", "map[uint64]ptype.Syntax"},
		{"Proto2.map_msg", "map[uint32]*Proto2_Required"},
		{"Proto2.map_ext_msg", "map[int64]*duration.Duration"},

		// proto2 syntax, required
		{"Proto2.Required.double", "*float64"},
		{"Proto2.Required.float", "*float32"},
		{"Proto2.Required.int64", "*int64"},
		{"Proto2.Required.sfixed64", "*int64"},
		{"Proto2.Required.sint64", "*int64"},
		{"Proto2.Required.uint64", "*uint64"},
		{"Proto2.Required.fixed64", "*uint64"},
		{"Proto2.Required.int32", "*int32"},
		{"Proto2.Required.sfixed32", "*int32"},
		{"Proto2.Required.sint32", "*int32"},
		{"Proto2.Required.uint32", "*uint32"},
		{"Proto2.Required.fixed32", "*uint32"},
		{"Proto2.Required.bool", "*bool"},
		{"Proto2.Required.string", "*string"},
		{"Proto2.Required.bytes", "[]byte"},
		{"Proto2.Required.enum", "*Proto2_Enum"},
		{"Proto2.Required.ext_enum", "*ptype.Syntax"},
		{"Proto2.Required.msg", "*Proto2_Required"},
		{"Proto2.Required.ext_msg", "*duration.Duration"},

		{"Proto3.double", "float64"},
		{"Proto3.float", "float32"},
		{"Proto3.int64", "int64"},
		{"Proto3.sfixed64", "int64"},
		{"Proto3.sint64", "int64"},
		{"Proto3.uint64", "uint64"},
		{"Proto3.fixed64", "uint64"},
		{"Proto3.int32", "int32"},
		{"Proto3.sfixed32", "int32"},
		{"Proto3.sint32", "int32"},
		{"Proto3.uint32", "uint32"},
		{"Proto3.fixed32", "uint32"},
		{"Proto3.bool", "bool"},
		{"Proto3.string", "string"},
		{"Proto3.bytes", "[]byte"},
		{"Proto3.enum", "Proto3_Enum"},
		{"Proto3.ext_enum", "ptype.Syntax"},
		{"Proto3.msg", "*Proto3_Message"},
		{"Proto3.ext_msg", "*duration.Duration"},
		{"Proto3.repeated_scalar", "[]float64"},
		{"Proto3.repeated_enum", "[]Proto3_Enum"},
		{"Proto3.repeated_ext_enum", "[]ptype.Syntax"},
		{"Proto3.repeated_msg", "[]*Proto3_Message"},
		{"Proto3.repeated_ext_msg", "[]*duration.Duration"},
		{"Proto3.map_scalar", "map[string]float32"},
		{"Proto3.map_enum", "map[int32]Proto3_Enum"},
		{"Proto3.map_ext_enum", "map[uint64]ptype.Syntax"},
		{"Proto3.map_msg", "map[uint32]*Proto3_Message"},
		{"Proto3.map_ext_msg", "map[int64]*duration.Duration"},
	}

	for _, test := range tests {
		tc := test
		t.Run(tc.field, func(t *testing.T) {
			t.Parallel()

			e, ok := ast.Lookup(".names.types." + tc.field)
			require.True(t, ok, "could not find field")

			fld, ok := e.(pgs.Field)
			require.True(t, ok, "entity is not a field")

			assert.Equal(t, tc.expected, ctx.Type(fld))
		})
	}
}

func TestTypeName(t *testing.T) {
	t.Parallel()

	tests := []struct {
		in  string
		el  string
		key string
		ptr string
		val string
	}{
		{
			in:  "int",
			el:  "int",
			ptr: "*int",
			val: "int",
		},
		{
			in:  "*int",
			el:  "*int",
			ptr: "*int",
			val: "int",
		},
		{
			in:  "foo.bar",
			el:  "foo.bar",
			ptr: "*foo.bar",
			val: "foo.bar",
		},
		{
			in:  "*foo.bar",
			el:  "*foo.bar",
			ptr: "*foo.bar",
			val: "foo.bar",
		},
		{
			in:  "[]string",
			el:  "string",
			key: "int",
			ptr: "[]string",
			val: "[]string",
		},
		{
			in:  "[]*string",
			el:  "*string",
			key: "int",
			ptr: "[]*string",
			val: "[]*string",
		},
		{
			in:  "[]foo.bar",
			el:  "foo.bar",
			key: "int",
			ptr: "[]foo.bar",
			val: "[]foo.bar",
		},
		{
			in:  "[]*foo.bar",
			el:  "*foo.bar",
			key: "int",
			ptr: "[]*foo.bar",
			val: "[]*foo.bar",
		},
		{
			in:  "map[string]float64",
			el:  "float64",
			key: "string",
			ptr: "map[string]float64",
			val: "map[string]float64",
		},
		{
			in:  "map[string]*float64",
			el:  "*float64",
			key: "string",
			ptr: "map[string]*float64",
			val: "map[string]*float64",
		},
		{
			in:  "map[string]foo.bar",
			el:  "foo.bar",
			key: "string",
			ptr: "map[string]foo.bar",
			val: "map[string]foo.bar",
		},
		{
			in:  "map[string]*foo.bar",
			el:  "*foo.bar",
			key: "string",
			ptr: "map[string]*foo.bar",
			val: "map[string]*foo.bar",
		},
		{
			in:  "[][]byte",
			el:  "[]byte",
			key: "int",
			ptr: "[][]byte",
			val: "[][]byte",
		},
		{
			in:  "map[int64][]byte",
			el:  "[]byte",
			key: "int64",
			ptr: "map[int64][]byte",
			val: "map[int64][]byte",
		},
	}

	for _, test := range tests {
		tc := test
		t.Run(tc.in, func(t *testing.T) {
			tn := TypeName(tc.in)
			t.Parallel()

			t.Run("Element", func(t *testing.T) {
				t.Parallel()
				assert.Equal(t, tc.el, tn.Element().String())
			})

			t.Run("Key", func(t *testing.T) {
				t.Parallel()
				assert.Equal(t, tc.key, tn.Key().String())
			})

			t.Run("Pointer", func(t *testing.T) {
				t.Parallel()
				assert.Equal(t, tc.ptr, tn.Pointer().String())
			})

			t.Run("Value", func(t *testing.T) {
				t.Parallel()
				assert.Equal(t, tc.val, tn.Value().String())
			})
		})
	}
}

func TestTypeName_Key_Malformed(t *testing.T) {
	t.Parallel()
	tn := TypeName("]malformed")
	assert.Empty(t, tn.Key().String())
}

func TestScalarType_Invalid(t *testing.T) {
	t.Parallel()
	assert.Panics(t, func() {
		scalarType(pgs.ProtoType(0))
	})
}

func ExampleTypeName_Element() {
	types := []string{
		"int",
		"*my.Type",
		"[]string",
		"map[string]*io.Reader",
	}

	for _, t := range types {
		fmt.Println(TypeName(t).Element())
	}

	// Output:
	// int
	// *my.Type
	// string
	// *io.Reader
}

func ExampleTypeName_Key() {
	types := []string{
		"int",
		"*my.Type",
		"[]string",
		"map[string]*io.Reader",
	}

	for _, t := range types {
		fmt.Println(TypeName(t).Key())
	}

	// Output:
	//
	//
	// int
	// string
}

func ExampleTypeName_Pointer() {
	types := []string{
		"int",
		"*my.Type",
		"[]string",
		"map[string]*io.Reader",
	}

	for _, t := range types {
		fmt.Println(TypeName(t).Pointer())
	}

	// Output:
	// *int
	// *my.Type
	// []string
	// map[string]*io.Reader
}

func ExampleTypeName_Value() {
	types := []string{
		"int",
		"*my.Type",
		"[]string",
		"map[string]*io.Reader",
	}

	for _, t := range types {
		fmt.Println(TypeName(t).Value())
	}

	// Output:
	// int
	// my.Type
	// []string
	// map[string]*io.Reader
}
