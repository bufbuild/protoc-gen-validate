package php_yaml

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"text/template"
	"unicode"

	"github.com/envoyproxy/protoc-gen-validate/templates/shared"
	"github.com/golang/protobuf/ptypes/duration"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/iancoleman/strcase"
	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
)

func Register(tpl *template.Template, params pgs.Parameters) {
	fns := phpFuncs{pgsgo.InitContext(params)}

	tpl.Funcs(map[string]interface{}{
		"phpNamespace":            phpNamespace,
		"classNameFile":           classNameFile,
		"classNameMessage":        classNameMessage,
		"simpleName":              fns.Name,
		"qualifiedName":           fns.qualifiedName,
		"accessor":                fns.accessor,
		"camelCase":               fns.camelCase,
		"fieldName":               fns.fieldName,
		"phpStringEscape":         fns.phpStringEscape,
		"phpTypeFor":              fns.phpTypeFor,
		"phpTypeLiteralSuffixFor": fns.phpTypeLiteralSuffixFor,
		"hasAccessor":             fns.hasAccessor,
		"oneof":                   fns.oneofTypeName,
		"sprintf":                 fmt.Sprintf,
		"tsLit":                   fns.tsLit,
		"durLit":                  fns.durLit,
		"importsPvg":              importsPvg,
		"isOfFileType":            fns.isOfFileType,
		"isOfMessageType":         fns.isOfMessageType,
		"isOfStringType":          fns.isOfStringType,
		"renderConstants":         fns.renderConstants(tpl),
		"constantName":            fns.constantName,
	})

	template.Must(tpl.Parse(fileTpl))
	template.Must(tpl.New("msg").Parse(msgTpl))
	template.Must(tpl.New("msgInner").Parse(msgInnerTpl))

	template.Must(tpl.New("none").Parse(noneTpl))

	// FIXME In case of declarative validation rules, what do we need constTpl for?

	template.Must(tpl.New("float").Parse(numTpl))
	template.Must(tpl.New("double").Parse(numTpl))
	template.Must(tpl.New("int32").Parse(numTpl))
	template.Must(tpl.New("int64").Parse(numTpl))
	template.Must(tpl.New("uint32").Parse(numTpl))
	template.Must(tpl.New("uint64").Parse(numTpl))
	template.Must(tpl.New("sint32").Parse(numTpl))
	template.Must(tpl.New("sint64").Parse(numTpl))
	template.Must(tpl.New("fixed32").Parse(numTpl))
	template.Must(tpl.New("fixed64").Parse(numTpl))
	template.Must(tpl.New("sfixed32").Parse(numTpl))
	template.Must(tpl.New("sfixed64").Parse(numTpl))

	template.Must(tpl.New("bool").Parse(boolTpl))
	template.Must(tpl.New("string").Parse(stringTpl))

	template.Must(tpl.New("any").Parse(anyTpl))
	template.Must(tpl.New("choice").Parse(choiceTpl))
	template.Must(tpl.New("enum").Parse(enumTpl))
	template.Must(tpl.New("message").Parse(messageTpl))
	template.Must(tpl.New("repeated").Parse(repeatedTpl))
	template.Must(tpl.New("repeatedConst").Parse(repeatedConstTpl))
	template.Must(tpl.New("map").Parse(mapTpl))
	template.Must(tpl.New("mapConst").Parse(mapConstTpl))
	template.Must(tpl.New("oneOf").Parse(oneOfTpl))
	template.Must(tpl.New("oneOfConst").Parse(oneOfConstTpl))

	template.Must(tpl.New("required").Parse(requiredTpl))
	template.Must(tpl.New("timestamp").Parse(timestampTpl))
	template.Must(tpl.New("duration").Parse(durationTpl))
}

type phpFuncs struct{ pgsgo.Context }

func PhpYamlFilePath(f pgs.File, ctx pgsgo.Context, tpl *template.Template) *pgs.FilePath {
	// Don't generate validators for files that don't import PGV
	if !importsPvg(f) {
		return nil
	}

	packagePath := f.Package().ProtoName().String()
	rulesPath := pgs.FilePath(f.Name().String()).BaseName() + ".yaml"
	filePath := pgs.JoinPaths(packagePath + "." + rulesPath)
	return &filePath
}

func importsPvg(f pgs.File) bool {
	for _, dep := range f.Descriptor().Dependency {
		if strings.HasSuffix(dep, "validate.proto") {
			return true
		}
	}
	return false
}

func classNameFile(f pgs.File) string {
	// TODO Add class prefix
	//options := f.Descriptor().GetOptions()
	//if options != nil && options.GetPhpClassPrefix() != "" {
	//}

	protoName := pgs.FilePath(f.Name().String()).BaseName()

	className := sanitizeClassName(protoName)

	return className
}

func classNameMessage(m pgs.Message) string {
	className := m.Name().String()
	// This is really silly, but when the multiple files option is true, protoc puts underscores in file names.
	// When multiple files is false, underscores are stripped. Short of rewriting all the name sanitization
	// logic for php_yaml, using "UnderscoreUnderscoreUnderscore" is an escape sequence seems to work with an extremely
	// small likelihood of name conflict.
	className = strings.Replace(className, "_", "UnderscoreUnderscoreUnderscore", -1)
	className = sanitizeClassName(className)
	className = strings.Replace(className, "UnderscoreUnderscoreUnderscore", "_", -1)
	return className
}

func sanitizeClassName(className string) string {
	className = makeInvalidClassnameCharactersUnderscores(className)
	className = underscoreBetweenConsecutiveUppercase(className)
	className = strcase.ToCamel(strcase.ToSnake(className))
	className = upperCaseAfterNumber(className)
	return className
}

func phpNamespace(file pgs.File) string {
	// Explicit php_yaml package overrides implicit package
	options := file.Descriptor().GetOptions()
	if options != nil && options.PhpNamespace != nil {
		return options.GetPhpNamespace()
	}

	nsEntries := strings.Split(file.Package().ProtoName().String(), ".")
	nsParts := make([]string, len(nsEntries))

	for i, nsPart := range nsEntries {
		nsParts[i] = strcase.ToCamel(nsPart)
	}

	phpPackageName := strings.Join(nsParts, "\\")

	return phpPackageName
}

func (fns phpFuncs) qualifiedName(entity pgs.Entity) string {
	file, isFile := entity.(pgs.File)
	if isFile {
		name := phpNamespace(file) + "." + classNameFile(file)
		return name
	}

	message, isMessage := entity.(pgs.Message)
	if isMessage && message.Parent() != nil {
		parent, isFileParent := message.Parent().(pgs.File)

		// Qualified name doesn't include proto file name
		if isFileParent {
			return phpNamespace(parent) + "\\" + entity.Name().String()
		}

		// recurse
		return fns.qualifiedName(message.Parent()) + "\\" + entity.Name().String()
	}

	enum, isEnum := entity.(pgs.Enum)
	if isEnum && enum.Parent() != nil {
		// recurse
		return fns.qualifiedName(enum.Parent()) + "." + entity.Name().String()
	}

	return entity.Name().String()
}

// Replace invalid identifier characters with an underscore
func makeInvalidClassnameCharactersUnderscores(name string) string {
	var sb string
	for _, c := range name {
		switch {
		case c >= '0' && c <= '9':
			sb += string(c)
		case c >= 'a' && c <= 'z':
			sb += string(c)
		case c >= 'A' && c <= 'Z':
			sb += string(c)
		default:
			sb += "_"
		}
	}
	return sb
}

func upperCaseAfterNumber(name string) string {
	var sb string
	var p rune

	for _, c := range name {
		if unicode.IsDigit(p) {
			sb += string(unicode.ToUpper(c))
		} else {
			sb += string(c)
		}
		p = c
	}
	return sb
}

func underscoreBetweenConsecutiveUppercase(name string) string {
	var sb string
	var p rune

	for _, c := range name {
		if unicode.IsUpper(p) && unicode.IsUpper(c) {
			sb += "_" + string(c)
		} else {
			sb += string(c)
		}
		p = c
	}
	return sb
}

func (fns phpFuncs) accessor(ctx shared.RuleContext) string {
	if ctx.AccessorOverride != "" {
		return ctx.AccessorOverride
	}
	return fns.fieldAccessor(ctx.Field)
}

func (fns phpFuncs) fieldAccessor(f pgs.Field) string {
	fieldName := strcase.ToCamel(f.Name().String())
	if f.Type().IsMap() {
		fieldName += "Map"
	}
	if f.Type().IsRepeated() {
		fieldName += "List"
	}

	fieldName = upperCaseAfterNumber(fieldName)
	return fmt.Sprintf("proto.get%s()", fieldName)
}

func (fns phpFuncs) hasAccessor(ctx shared.RuleContext) string {
	if ctx.AccessorOverride != "" {
		return "true"
	}
	fieldName := strcase.ToCamel(ctx.Field.Name().String())
	fieldName = upperCaseAfterNumber(fieldName)
	return "proto.has" + fieldName + "()"
}

func (fns phpFuncs) fieldName(ctx shared.RuleContext) string {
	return ctx.Field.Name().String()
}

func (fns phpFuncs) phpTypeFor(ctx shared.RuleContext) string {
	t := ctx.Field.Type()

	// Map key and value types
	if t.IsMap() {
		switch ctx.AccessorOverride {
		case "key":
			return fns.phpTypeForProtoType(t.Key().ProtoType())
		case "value":
			return fns.phpTypeForProtoType(t.Element().ProtoType())
		}
	}

	if t.IsEmbed() {
		if embed := t.Embed(); embed.IsWellKnown() {
			switch embed.WellKnownType() {
			case pgs.AnyWKT:
				return "String"
			case pgs.DurationWKT:
				return "com.google.protobuf.Duration"
			case pgs.TimestampWKT:
				return "com.google.protobuf.Timestamp"
			case pgs.Int32ValueWKT, pgs.UInt32ValueWKT:
				return "Integer"
			case pgs.Int64ValueWKT, pgs.UInt64ValueWKT:
				return "Long"
			case pgs.DoubleValueWKT:
				return "Double"
			case pgs.FloatValueWKT:
				return "Float"
			}
		}
	}

	if t.IsRepeated() {
		if t.ProtoType() == pgs.MessageT {
			return fns.qualifiedName(t.Element().Embed())
		} else if t.ProtoType() == pgs.EnumT {
			return fns.qualifiedName(t.Element().Enum())
		}
	}

	if t.IsEnum() {
		return fns.qualifiedName(t.Enum())
	}

	return fns.phpTypeForProtoType(t.ProtoType())
}

func (fns phpFuncs) phpTypeForProtoType(t pgs.ProtoType) string {

	switch t {
	case pgs.Int32T, pgs.UInt32T, pgs.SInt32, pgs.Fixed32T, pgs.SFixed32:
		return "int"
	case pgs.Int64T, pgs.UInt64T, pgs.SInt64, pgs.Fixed64T, pgs.SFixed64:
		return "int"
	case pgs.DoubleT:
		return "float"
	case pgs.FloatT:
		return "float"
	case pgs.BoolT:
		return "bool"
	case pgs.StringT:
		return "string"
	case pgs.BytesT:
		return "com.google.protobuf.ByteString"
	default:
		return "object"
	}
}

func (fns phpFuncs) phpTypeLiteralSuffixFor(ctx shared.RuleContext) string {
	t := ctx.Field.Type()

	if t.IsMap() {
		switch ctx.AccessorOverride {
		case "key":
			return fns.phpTypeLiteralSuffixForPrototype(t.Key().ProtoType())
		case "value":
			return fns.phpTypeLiteralSuffixForPrototype(t.Element().ProtoType())
		}
	}

	if t.IsEmbed() {
		if embed := t.Embed(); embed.IsWellKnown() {
			switch embed.WellKnownType() {
			case pgs.Int64ValueWKT, pgs.UInt64ValueWKT:
				return "L"
			case pgs.FloatValueWKT:
				return "F"
			case pgs.DoubleValueWKT:
				return "D"
			}
		}
	}

	return fns.phpTypeLiteralSuffixForPrototype(t.ProtoType())
}

func (fns phpFuncs) phpTypeLiteralSuffixForPrototype(t pgs.ProtoType) string {
	switch t {
	case pgs.Int64T, pgs.UInt64T, pgs.SInt64, pgs.Fixed64T, pgs.SFixed64:
		return "L"
	case pgs.FloatT:
		return "F"
	case pgs.DoubleT:
		return "D"
	default:
		return ""
	}
}

func (fns phpFuncs) phpStringEscape(s string) string {
	s = fmt.Sprintf("%q", s)
	s = s[1 : len(s)-1]
	s = strings.Replace(s, `\u00`, `\x`, -1)
	s = strings.Replace(s, `\x`, `\\x`, -1)
	// s = strings.Replace(s, `\`, `\\`, -1)
	s = strings.Replace(s, `"`, `\"`, -1)
	return `"` + s + `"`
}

func (fns phpFuncs) camelCase(name pgs.Name) string {
	return strcase.ToCamel(name.String())
}

func (fns phpFuncs) durLit(dur *duration.Duration) string {
	// TODO Figure out how to make use of nanos
	return strconv.FormatInt(dur.GetSeconds(), 10)
}

func (fns phpFuncs) tsLit(ts *timestamp.Timestamp) string {
	// TODO Figure out how to make use of nanos
	return strconv.FormatInt(ts.GetSeconds(), 10)
}

func (fns phpFuncs) oneofTypeName(f pgs.Field) pgsgo.TypeName {
	return pgsgo.TypeName(fmt.Sprintf("%s", strings.ToUpper(f.Name().String())))
}

func (fns phpFuncs) isOfFileType(o interface{}) bool {
	switch o.(type) {
	case pgs.File:
		return true
	default:
		return false
	}
}

func (fns phpFuncs) isOfMessageType(f pgs.Field) bool {
	return f.Type().ProtoType() == pgs.MessageT
}

func (fns phpFuncs) isOfStringType(f pgs.Field) bool {
	return f.Type().ProtoType() == pgs.StringT
}

func (fns phpFuncs) renderConstants(tpl *template.Template) func(ctx shared.RuleContext) (string, error) {
	return func(ctx shared.RuleContext) (string, error) {
		var b bytes.Buffer
		var err error

		hasConstTemplate := false
		for _, t := range tpl.Templates() {
			if t.Name() == ctx.Typ+"Const" {
				hasConstTemplate = true
			}
		}

		if hasConstTemplate {
			err = tpl.ExecuteTemplate(&b, ctx.Typ+"Const", ctx)
		}

		return b.String(), err
	}
}

func (fns phpFuncs) constantName(ctx shared.RuleContext, rule string) string {
	return strcase.ToScreamingSnake(ctx.Field.Name().String() + "_" + ctx.Index + "_" + rule)
}
