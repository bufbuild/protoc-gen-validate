package pgsgo

import (
	"testing"

	pgs "github.com/lyft/protoc-gen-star"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPGGUpperCamelCase(t *testing.T) {
	t.Parallel()

	tests := []struct {
		in string
		ex string
	}{
		{"foo_bar", "FooBar"},
		{"myJSON", "MyJSON"},
		{"PDFTemplate", "PDFTemplate"},
		{"_my_field_name_2", "XMyFieldName_2"},
		{"my.field", "My.field"},
		{"my_Field", "My_Field"},
	}

	for _, tc := range tests {
		assert.Equal(t, tc.ex, PGGUpperCamelCase(pgs.Name(tc.in)).String())
	}
}

func TestName(t *testing.T) {
	t.Parallel()

	ast := buildGraph(t, "names", "entities")
	ctx := loadContext(t, "names", "entities")

	f := ast.Targets()["entities.proto"]
	assert.Equal(t, pgs.Name("BAD_pack__age_name_"), ctx.Name(f))
	assert.Equal(t, pgs.Name("BAD_pack__age_name_"), ctx.Name(f.Package()))

	assert.Panics(t, func() {
		ctx.Name(nil)
	})

	tests := []struct {
		entity   string
		expected pgs.Name
	}{
		// Top-Level Messages
		{"UpperCamelCaseMessage", "UpperCamelCaseMessage"},
		{"lowerCamelCaseMessage", "LowerCamelCaseMessage"},
		{"SCREAMING_SNAKE_CASE", "SCREAMING_SNAKE_CASE"},
		{"Upper_Snake_Case", "Upper_Snake_Case"},
		{"lower_snake_case", "LowerSnakeCase"},
		{"lowercase", "Lowercase"},
		{"UPPERCASE", "UPPERCASE"},
		{"_underscore", "XUnderscore"},
		{"__DoubleUnderscore", "X_DoubleUnderscore"},
		{"String", "String"},
		{"MsgWith3dInside", "MsgWith3DInside"},
		{"MsgEndsWith3d", "MsgEndsWith3D"},

		// Nested Messages
		{"Nested.Message", "Nested_Message"},
		{"Nested._underscore", "Nested_XUnderscore"},
		{"Nested.String", "Nested_String"},
		{"Nested.Message.Message", "Nested_Message_Message"},
		{"Nested.lowerMsg", "NestedLowerMsg"},

		// Enums
		{"UpperCamelCaseEnum", "UpperCamelCaseEnum"},
		{"lowerCamelCaseEnum", "LowerCamelCaseEnum"},
		{"SCREAMING_SNAKE_ENUM", "SCREAMING_SNAKE_ENUM"},
		{"lower_snake_enum", "LowerSnakeEnum"},
		{"Upper_Snake_Enum", "Upper_Snake_Enum"},

		// EnumValues
		{"UpperCamelCaseEnum.SCREAMING_SNAKE_CASE_ENUM_VALUE", "UpperCamelCaseEnum_SCREAMING_SNAKE_CASE_ENUM_VALUE"},
		{"UpperCamelCaseEnum.lower_snake_case_enum_value", "UpperCamelCaseEnum_lower_snake_case_enum_value"},
		{"UpperCamelCaseEnum.Upper_Snake_Case_Enum_Value", "UpperCamelCaseEnum_Upper_Snake_Case_Enum_Value"},
		{"UpperCamelCaseEnum.UpperCamelCaseEnumValue", "UpperCamelCaseEnum_UpperCamelCaseEnumValue"},
		{"UpperCamelCaseEnum.lowerCamelCaseEnumValue", "UpperCamelCaseEnum_lowerCamelCaseEnumValue"},
		{"lowerCamelCaseEnum.LCC_Value", "LowerCamelCaseEnum_LCC_Value"},
		{"SCREAMING_SNAKE_ENUM.SS_Value", "SCREAMING_SNAKE_ENUM_SS_Value"},
		{"lower_snake_enum.LS_Value", "LowerSnakeEnum_LS_Value"},
		{"Upper_Snake_Enum.US_Value", "Upper_Snake_Enum_US_Value"},

		// Nested Enums
		{"Nested.Enum", "Nested_Enum"},
		{"Nested.Enum.VALUE", "Nested_VALUE"},
		{"Nested.Message.Enum", "Nested_Message_Enum"},
		{"Nested.Message.Enum.NESTED", "Nested_Message_NESTED"},
		{"Nested.lowercase", "NestedLowercase"},
		{"Nested.lowercase.lower", "Nested_lower"},

		// Field Names
		{"Fields.lower_snake_case", "LowerSnakeCase"},
		{"Fields.Upper_Snake_Case", "Upper_Snake_Case"},
		{"Fields.SCREAMING_SNAKE_CASE", "SCREAMING_SNAKE_CASE"},
		{"Fields.lowerCamelCase", "LowerCamelCase"},
		{"Fields.UpperCamelCase", "UpperCamelCase"},
		{"Fields.string", "String_"},

		// OneOfs
		{"Oneofs.lower_snake_case", "LowerSnakeCase"},
		{"Oneofs.Upper_Snake_Case", "Upper_Snake_Case"},
		{"Oneofs.SCREAMING_SNAKE_CASE", "SCREAMING_SNAKE_CASE"},
		{"Oneofs.lowerCamelCase", "LowerCamelCase"},
		{"Oneofs.UpperCamelCase", "UpperCamelCase"},
		{"Oneofs.string", "String_"},
		{"Oneofs.oneof", "Oneof"},

		// Services (always the Server name)
		{"UpperCamelService", "UpperCamelServiceServer"},
		{"lowerCamelService", "LowerCamelServiceServer"},
		{"lower_snake_service", "LowerSnakeServiceServer"},
		{"Upper_Snake_Service", "Upper_Snake_ServiceServer"},
		{"SCREAMING_SNAKE_SERVICE", "SCREAMING_SNAKE_SERVICEServer"},
		{"reset", "ResetServer"},

		// Methods
		{"Service.UpperCamel", "UpperCamel"},
		{"Service.lowerCamel", "LowerCamel"},
		{"Service.lower_snake", "LowerSnake"},
		{"Service.Upper_Snake", "Upper_Snake"},
		{"Service.SCREAMING_SNAKE", "SCREAMING_SNAKE"},
		{"Service.Reset", "Reset"},
	}

	for _, test := range tests {
		tc := test
		t.Run(tc.entity, func(t *testing.T) {
			t.Parallel()

			e, ok := ast.Lookup(".names.entities." + tc.entity)
			require.True(t, ok, "could not locate entity")
			assert.Equal(t, tc.expected, ctx.Name(e))
		})
	}
}

func TestContext_OneofOption(t *testing.T) {
	t.Parallel()

	ast := buildGraph(t, "names", "entities")
	ctx := loadContext(t, "names", "entities")

	tests := []struct {
		field    string
		expected pgs.Name
	}{
		{"LS", "Oneofs_LS"},
		{"US", "Oneofs_US"},
		{"SS", "Oneofs_SS"},
		{"LC", "Oneofs_LC"},
		{"UC", "Oneofs_UC"},
		{"S", "Oneofs_S"},
		{"lower_snake_case_o", "Oneofs_LowerSnakeCaseO"},
		{"Upper_Snake_Case_O", "Oneofs_Upper_Snake_Case_O"},
		{"SCREAMING_SNAKE_CASE_O", "Oneofs_SCREAMING_SNAKE_CASE_O"},
		{"lowerCamelCaseO", "Oneofs_LowerCamelCaseO"},
		{"UpperCamelCaseO", "Oneofs_UpperCamelCaseO"},
		{"reset", "Oneofs_Reset_"},
		{"some_msg", "Oneofs_SomeMsg_"},
		{"some_enum", "Oneofs_SomeEnum_"},
	}

	for _, test := range tests {
		tc := test
		t.Run(tc.field, func(t *testing.T) {
			t.Parallel()

			e, ok := ast.Lookup(".names.entities.Oneofs." + tc.field)
			require.True(t, ok, "could not find field")
			f := e.(pgs.Field)
			assert.Equal(t, tc.expected, ctx.OneofOption(f))
		})
	}

}

func TestContext_ClientName(t *testing.T) {
	t.Parallel()

	ast := buildGraph(t, "names", "entities")
	ctx := loadContext(t, "names", "entities")

	tests := []struct {
		service  string
		expected pgs.Name
	}{
		{"UpperCamelService", "UpperCamelServiceClient"},
		{"lowerCamelService", "LowerCamelServiceClient"},
		{"lower_snake_service", "LowerSnakeServiceClient"},
		{"Upper_Snake_Service", "Upper_Snake_ServiceClient"},
		{"SCREAMING_SNAKE_SERVICE", "SCREAMING_SNAKE_SERVICEClient"},
		{"reset", "ResetClient"},
	}

	for _, test := range tests {
		tc := test
		t.Run(tc.service, func(t *testing.T) {
			t.Parallel()

			e, ok := ast.Lookup(".names.entities." + tc.service)
			require.True(t, ok, "could not find service")
			s := e.(pgs.Service)
			assert.Equal(t, tc.expected, ctx.ClientName(s))
		})
	}
}

func TestContext_ServerStream(t *testing.T) {
	t.Parallel()

	ast := buildGraph(t, "names", "entities")
	ctx := loadContext(t, "names", "entities")

	tests := []struct {
		method   string
		expected pgs.Name
	}{
		{"UpperCamel", "Service_UpperCamelServer"},
		{"lowerCamel", "Service_LowerCamelServer"},
		{"lower_snake", "Service_LowerSnakeServer"},
	}

	for _, test := range tests {
		tc := test
		t.Run(tc.method, func(t *testing.T) {
			t.Parallel()

			e, ok := ast.Lookup(".names.entities.Service." + tc.method)
			require.True(t, ok, "could not find method")
			m := e.(pgs.Method)
			assert.Equal(t, tc.expected, ctx.ServerStream(m))
		})
	}
}
