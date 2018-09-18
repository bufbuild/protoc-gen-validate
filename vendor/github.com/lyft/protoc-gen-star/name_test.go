package pgs

import (
	"fmt"
	"testing"

	"strings"

	"github.com/stretchr/testify/assert"
)

func TestName_Split(t *testing.T) {
	t.Parallel()

	tests := []struct {
		in    string
		parts []string
	}{
		// camel case
		{"fooBar", []string{"foo", "Bar"}},
		{"FooBar", []string{"Foo", "Bar"}},
		{"myJSON", []string{"my", "JSON"}},
		{"JSONStringFooBar", []string{"JSON", "String", "Foo", "Bar"}},
		{"myJSONString", []string{"my", "JSON", "String"}},

		// snake case
		{"FOO_BAR", []string{"FOO", "BAR"}},
		{"foo_bar_baz", []string{"foo", "bar", "baz"}},
		{"Foo_Bar", []string{"Foo", "Bar"}},
		{"JSONString_Foo_Bar", []string{"JSONString", "Foo", "Bar"}},

		// dot notation
		{"foo.bar", []string{"foo", "bar"}},
		{".foo.bar", []string{"", "foo", "bar"}},
		{".JSONString.Foo.Bar", []string{"", "JSONString", "Foo", "Bar"}},

		// leading underscore
		{"_Privatish", []string{"_Privatish"}},
		{"_privatish", []string{"_privatish"}},
		{"_foo_bar", []string{"_foo", "bar"}},
		{"_Foo_Bar", []string{"_Foo", "Bar"}},
		{"_JSON_String", []string{"_JSON", "String"}},
		{"_JString", []string{"_J", "String"}},
		{"__Double", []string{"_", "Double"}},

		// numbers
		{"abc123", []string{"abc", "123"}},
		{"123def", []string{"123", "def"}},
		{"abc1def", []string{"abc", "1", "def"}},
		{"ABC1DEF", []string{"ABC", "1", "DEF"}},

		// empty
		{"", []string{""}},
	}

	for _, test := range tests {
		tc := test
		t.Run(tc.in, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.parts, Name(tc.in).Split())
		})
	}
}

func TestName(t *testing.T) {
	t.Parallel()

	tests := []struct {
		in  []string
		ucc string
		lcc string
		ssc string
		lsc string
		usc string
		ldn string
		udn string
	}{
		{
			in:  []string{"fooBar", "FooBar", "foo_bar", "Foo_Bar", "foo_Bar", "foo.Bar", "Foo.Bar"},
			ucc: "FooBar",
			lcc: "fooBar",
			ssc: "FOO_BAR",
			lsc: "foo_bar",
			usc: "Foo_Bar",
			ldn: "foo.bar",
			udn: "Foo.Bar",
		},
		{
			in:  []string{"JSONString", "JSON_String", "JSON_string", "JSON.string"},
			ucc: "JSONString",
			lcc: "jsonString",
			ssc: "JSON_STRING",
			lsc: "json_string",
			usc: "JSON_String",
			ldn: "json.string",
			udn: "JSON.String",
		},
		{
			in:  []string{"myJSON", "my_JSON", "My_JSON", "my.JSON"},
			ucc: "MyJSON",
			lcc: "myJSON",
			ssc: "MY_JSON",
			lsc: "my_json",
			usc: "My_JSON",
			ldn: "my.json",
			udn: "My.JSON",
		},
	}

	for _, test := range tests {
		tc := test
		for _, in := range tc.in {
			n := Name(in)
			t.Run(string(n), func(t *testing.T) {
				t.Parallel()

				t.Run("UpperCamelCase", func(t *testing.T) {
					t.Parallel()
					assert.Equal(t, tc.ucc, n.UpperCamelCase().String())
				})

				t.Run("lowerCamelCase", func(t *testing.T) {
					t.Parallel()
					assert.Equal(t, tc.lcc, n.LowerCamelCase().String())
				})

				t.Run("SCREAMING_SNAKE_CASE", func(t *testing.T) {
					t.Parallel()
					assert.Equal(t, tc.ssc, n.ScreamingSnakeCase().String())
				})

				t.Run("lower_snake_case", func(t *testing.T) {
					t.Parallel()
					assert.Equal(t, tc.lsc, n.LowerSnakeCase().String())
				})

				t.Run("Upper_Snake_Case", func(t *testing.T) {
					t.Parallel()
					assert.Equal(t, tc.usc, n.UpperSnakeCase().String())
				})

				t.Run("lower.dot.notation", func(t *testing.T) {
					t.Parallel()
					assert.Equal(t, tc.ldn, n.LowerDotNotation().String())
				})

				t.Run("Upper.Dot.Notation", func(t *testing.T) {
					t.Parallel()
					assert.Equal(t, tc.udn, n.UpperDotNotation().String())
				})
			})
		}
	}
}

func TestNameTransformer_Chain(t *testing.T) {
	t.Parallel()

	nt := NameTransformer(strings.ToUpper)
	nt = nt.Chain(func(s string) string { return "_" + s })

	assert.Equal(t, "_FOO", nt("foo"))
}

func TestFilePath(t *testing.T) {
	t.Parallel()

	fp := FilePath("alpha/beta/gamma.proto")
	assert.Equal(t, "alpha/beta/gamma.proto", fp.String())
	assert.Equal(t, "alpha/beta", fp.Dir().String())
	assert.Equal(t, "gamma.proto", fp.Base())
	assert.Equal(t, ".proto", fp.Ext())
	assert.Equal(t, "gamma", fp.BaseName())
	assert.Equal(t, "alpha/beta/gamma.foo", fp.SetExt(".foo").String())
	assert.Equal(t, "alpha/beta/delta.bar", fp.SetBase("delta.bar").String())
	assert.Equal(t, "alpha/beta", fp.Pop().String())
	assert.Equal(t, "alpha/beta/delta", fp.Dir().Push("delta").String())
}

func ExampleName_UpperCamelCase() {
	names := []string{
		"foo_bar",
		"myJSON",
		"PDFTemplate",
	}

	for _, n := range names {
		fmt.Println(Name(n).UpperCamelCase())
	}

	// Output:
	// FooBar
	// MyJSON
	// PDFTemplate
}

func ExampleName_LowerCamelCase() {
	names := []string{
		"foo_bar",
		"myJSON",
		"PDFTemplate",
	}

	for _, n := range names {
		fmt.Println(Name(n).LowerCamelCase())
	}

	// Output:
	// fooBar
	// myJSON
	// pdfTemplate
}

func ExampleName_ScreamingSnakeCase() {
	names := []string{
		"foo_bar",
		"myJSON",
		"PDFTemplate",
	}

	for _, n := range names {
		fmt.Println(Name(n).ScreamingSnakeCase())
	}

	// Output:
	// FOO_BAR
	// MY_JSON
	// PDF_TEMPLATE
}

func ExampleName_LowerSnakeCase() {
	names := []string{
		"foo_bar",
		"myJSON",
		"PDFTemplate",
	}

	for _, n := range names {
		fmt.Println(Name(n).LowerSnakeCase())
	}

	// Output:
	// foo_bar
	// my_json
	// pdf_template
}

func ExampleName_UpperSnakeCase() {
	names := []string{
		"foo_bar",
		"myJSON",
		"PDFTemplate",
	}

	for _, n := range names {
		fmt.Println(Name(n).UpperSnakeCase())
	}

	// Output:
	// Foo_Bar
	// My_JSON
	// PDF_Template
}

func ExampleName_SnakeCase() {
	names := []string{
		"foo_bar",
		"myJSON",
		"PDFTemplate",
	}

	for _, n := range names {
		fmt.Println(Name(n).SnakeCase())
	}

	// Output:
	// foo_bar
	// my_JSON
	// PDF_Template
}

func ExampleName_LowerDotNotation() {
	names := []string{
		"foo_bar",
		"myJSON",
		"PDFTemplate",
	}

	for _, n := range names {
		fmt.Println(Name(n).LowerDotNotation())
	}

	// Output:
	// foo.bar
	// my.json
	// pdf.template
}

func ExampleName_UpperDotNotation() {
	names := []string{
		"foo_bar",
		"myJSON",
		"PDFTemplate",
	}

	for _, n := range names {
		fmt.Println(Name(n).UpperDotNotation())
	}

	// Output:
	// Foo.Bar
	// My.JSON
	// PDF.Template
}
