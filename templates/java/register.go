package java

import (
	"fmt"
	"strings"
	"text/template"
	"unicode"

	"github.com/golang/protobuf/ptypes/duration"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/iancoleman/strcase"
	"github.com/lyft/protoc-gen-star"
	"github.com/lyft/protoc-gen-star/lang/go"
	"github.com/lyft/protoc-gen-validate/templates/shared"
)

func RegisterIndex(tpl *template.Template, params pgs.Parameters) {
	fns := javaFuncs{pgsgo.InitContext(params)}

	tpl.Funcs(map[string]interface{}{
		"classNameFile": classNameFile,
		"javaPackage":   fns.javaPackage,
		"simpleName":    fns.Name,
		"qualifiedName": fns.qualifiedName,
	})

	template.Must(tpl.Parse(fileIndexTpl))
}

func Register(tpl *template.Template, params pgs.Parameters) {
	fns := javaFuncs{pgsgo.InitContext(params)}

	tpl.Funcs(map[string]interface{}{
		"accessor":                 fns.accessor,
		"byteArrayLit":             fns.byteArrayLit,
		"camelCase":                fns.camelCase,
		"classNameFile":            classNameFile,
		"durLit":                   fns.durLit,
		"rawPrint":                 fns.rawPrint,
		"fieldName":                fns.fieldName,
		"javaPackage":              fns.javaPackage,
		"javaStringEscape":         fns.javaStringEscape,
		"javaTypeFor":              fns.javaTypeFor,
		"javaTypeLiteralSuffixFor": fns.javaTypeLiteralSuffixFor,
		"hasAccessor":              fns.hasAccessor,
		"oneof":                    fns.oneofTypeName,
		"sprintf":                  fmt.Sprintf,
		"simpleName":               fns.Name,
		"tsLit":                    fns.tsLit,
		"qualifiedName":            fns.qualifiedName,
		"isOfMessageType":          fns.isOfMessageType,
	})

	template.Must(tpl.Parse(fileTpl))
	template.Must(tpl.New("msg").Parse(msgTpl))

	template.Must(tpl.New("none").Parse(noneTpl))

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
	template.Must(tpl.New("bytes").Parse(bytesTpl))

	template.Must(tpl.New("any").Parse(anyTpl))
	template.Must(tpl.New("enum").Parse(enumTpl))
	template.Must(tpl.New("message").Parse(messageTpl))
	template.Must(tpl.New("repeated").Parse(notImplementedTpl))
	template.Must(tpl.New("map").Parse(mapTpl))

	template.Must(tpl.New("required").Parse(requiredTpl))
	template.Must(tpl.New("timestamp").Parse(timestampTpl))
	template.Must(tpl.New("duration").Parse(durationTpl))
	template.Must(tpl.New("wrapper").Parse(notImplementedTpl))
}

type javaFuncs struct{ pgsgo.Context }

func JavaFilePath(f pgs.File, ctx pgsgo.Context, tpl *template.Template) pgs.FilePath {
	fullPath := ctx.OutputPath(f)

	fileName := classNameFile(f).String()
	fileName += "Validator.java"

	return fullPath.SetBase(fileName)
}

func classNameFile(file pgs.File) pgs.Name {
	protoName := pgs.FilePath(file.Name().String()).BaseName()

	className := makeInvalidClassnameCharactersUnderscores(protoName)
	className = strcase.ToCamel(strcase.ToSnake(className))
	className = upperCaseAfterNumber(className)
	className = appendOuterClassName(className, file)

	return pgs.Name(className)
}

func (fns javaFuncs) javaPackage(file pgs.File) pgs.Name {
	return file.Package().ProtoName()
}

func (fns javaFuncs) qualifiedName(entity pgs.Entity) pgs.Name {
	file, isFile := entity.(pgs.File)
	if isFile {
		return fns.javaPackage(file) + "." + classNameFile(file)
	}

	message, isMessage := entity.(pgs.Message)
	if isMessage && message.Parent() != nil {
		// recurse
		return pgs.Name(fns.qualifiedName(message.Parent()) + "." + pgs.Name(entity.Name()))
	}

	enum, isEnum := entity.(pgs.Enum)
	if isEnum && enum.Parent() != nil {
		// recurse
		return pgs.Name(fns.qualifiedName(enum.Parent()) + "." + pgs.Name(entity.Name()))
	}

	return pgs.Name(entity.Name())
}

// Replace invalid identifier characters with an underscore
func makeInvalidClassnameCharactersUnderscores(name string) string {
	var sb strings.Builder
	for _, c := range name {
		switch {
		case c >= '0' && c <= '9':
			sb.WriteRune(c)
		case c >= 'a' && c <= 'z':
			sb.WriteRune(c)
		case c >= 'A' && c <= 'Z':
			sb.WriteRune(c)
		default:
			sb.WriteRune('_')
		}
	}
	return sb.String()
}

func upperCaseAfterNumber(name string) string {
	var sb strings.Builder
	var p rune

	for _, c := range name {
		if unicode.IsDigit(p) {
			sb.WriteRune(unicode.ToUpper(c))
		} else {
			sb.WriteRune(c)
		}
		p = c
	}
	return sb.String()
}

func appendOuterClassName(outerClassName string, file pgs.File) string {
	conflict := false

	for _, enum := range file.Enums() {
		if enum.Name().String() == outerClassName {
			conflict = true
		}
	}

	for _, message := range file.Messages() {
		if message.Name().String() == outerClassName {
			conflict = true
		}
	}

	for _, service := range file.Services() {
		if service.Name().String() == outerClassName {
			conflict = true
		}
	}

	if conflict {
		return outerClassName + "OuterClass"
	} else {
		return outerClassName
	}
}

func (fns javaFuncs) accessor(field pgs.Field) string {
	fieldName := strcase.ToCamel(field.Name().String())
	if field.Type().IsMap() {
		fieldName += "Map"
	}
	return "get" + fieldName + "()"
}

func (fns javaFuncs) rawPrint(instr string) string {
	return fmt.Sprintf("%#v", instr)
}

func (fns javaFuncs) hasAccessor(ctx shared.RuleContext) string {
	if ctx.AccessorOverride != "" {
		return "true"
	}
	fiedlName := strcase.ToCamel(ctx.Field.Name().String())
	return "proto.has" + fiedlName + "()"
}

func (fns javaFuncs) fieldName(ctx shared.RuleContext) string {
	return ctx.Field.Name().String()
}

func (fns javaFuncs) javaTypeFor(f pgs.Field) string {
	t := f.Type()
	switch t.ProtoType() {
	case pgs.Int32T, pgs.UInt32T, pgs.SInt32, pgs.Fixed32T, pgs.SFixed32:
		return "int"
	case pgs.Int64T, pgs.UInt64T, pgs.SInt64, pgs.Fixed64T, pgs.SFixed64:
		return "long"
	case pgs.DoubleT:
		return "double"
	case pgs.FloatT:
		return "float"
	case pgs.BoolT:
		return "boolean"
	case pgs.StringT:
		return "String"
	case pgs.BytesT:
		return "com.google.protobuf.ByteString"
	case pgs.EnumT:
		return fns.qualifiedName(f.Type().Enum()).String()
	case pgs.MessageT:
		if t.IsEmbed() {
			return fns.qualifiedName(t.Embed()).String()
		}
		if t.IsRepeated() {
			if t.ProtoType() == pgs.MessageT {
				return fns.qualifiedName(t.Element().Embed()).String()
			}
		}
		return "Object"
	default:
		return "Object"
	}
}

func (fns javaFuncs) javaTypeLiteralSuffixFor(f pgs.Field) string {
	switch f.Type().ProtoType() {
	case pgs.Int64T, pgs.UInt64T, pgs.SInt64, pgs.Fixed64T, pgs.SFixed64:
		return "L"
	case pgs.FloatT:
		return "F"
	default:
		return ""
	}
}

func (fns javaFuncs) javaStringEscape(s string) string {
	s = fns.rawPrint(s)
	s = s[1 : len(s)-1]
	s = strings.Replace(s, "\\", "\\\\", -1)
	s = strings.Replace(s, "\"", "\\\"", -1)
	return "\"" + s + "\""
}

func (fns javaFuncs) camelCase(name pgs.Name) string {
	return strcase.ToCamel(name.String())
}

func (fns javaFuncs) byteArrayLit(bytes []uint8) string {
	var sb strings.Builder
	sb.WriteString("new byte[]{")
	for _, b := range bytes {
		sb.WriteString(fmt.Sprintf("(byte)%#x,", b))
	}
	sb.WriteString("}")

	return sb.String()
}

func (fns javaFuncs) durLit(dur *duration.Duration) string {
	return fmt.Sprintf(
		"com.lyft.pgv.DurationValidation.toDuration(%d,%d)",
		dur.GetSeconds(), dur.GetNanos())
}

func (fns javaFuncs) tsLit(ts *timestamp.Timestamp) string {
	return fmt.Sprintf(
		"com.lyft.pgv.TimestampValidation.toTimestamp(%d,%d)",
		ts.GetSeconds(), ts.GetNanos())
}

func (fns javaFuncs) oneofTypeName(f pgs.Field) pgsgo.TypeName {
	return pgsgo.TypeName(fmt.Sprintf("%s",
		pgsgo.PGGUpperCamelCase(f.Name()),
	))
}

func (fns javaFuncs) isOfMessageType(f pgs.Field) bool {
	return f.Type().ProtoType() == pgs.MessageT
}
