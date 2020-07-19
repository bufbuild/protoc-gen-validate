package dotnet

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"
	"unicode"

	"github.com/golang/protobuf/ptypes/timestamp"

	"github.com/envoyproxy/protoc-gen-validate/templates/shared"
	"github.com/golang/protobuf/ptypes/duration"
	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
)

func Register(tpl *template.Template, params pgs.Parameters) {
	tpl.Funcs(map[string]interface{}{
		"renderConst": renderWithSuffixIfExists(tpl, "Const"),
		"err":         err,
		"unboundErr":  unboundErr,

		"accessor":  accessor,
		"constant":  constant,
		"fieldType": fieldType,
		"literal":   literal,

		"unwrap": unwrap,

		"parentClassNames": parentClassNames,
		"className":        className,
		"namespace":        namespace,

		"oneofAccessor": oneofAccessor,
		"oneofCase":     oneofCase,

		"enumType": enumType,
	})

	template.Must(tpl.Parse(fileTpl))
	template.Must(tpl.New("msg").Parse(msgTpl))

	template.Must(tpl.New("none").Parse(noneTpl))

	template.Must(tpl.New("float").Parse(numTpl))
	template.Must(tpl.New("floatConst").Parse(numConstTpl))
	template.Must(tpl.New("double").Parse(numTpl))
	template.Must(tpl.New("doubleConst").Parse(numConstTpl))
	template.Must(tpl.New("int32").Parse(numTpl))
	template.Must(tpl.New("int32Const").Parse(numConstTpl))
	template.Must(tpl.New("int64").Parse(numTpl))
	template.Must(tpl.New("int64Const").Parse(numConstTpl))
	template.Must(tpl.New("uint32").Parse(numTpl))
	template.Must(tpl.New("uint32Const").Parse(numConstTpl))
	template.Must(tpl.New("uint64").Parse(numTpl))
	template.Must(tpl.New("uint64Const").Parse(numConstTpl))
	template.Must(tpl.New("sint32").Parse(numTpl))
	template.Must(tpl.New("sint32Const").Parse(numConstTpl))
	template.Must(tpl.New("sint64").Parse(numTpl))
	template.Must(tpl.New("sint64Const").Parse(numConstTpl))
	template.Must(tpl.New("fixed32").Parse(numTpl))
	template.Must(tpl.New("fixed32Const").Parse(numConstTpl))
	template.Must(tpl.New("fixed64").Parse(numTpl))
	template.Must(tpl.New("fixed64Const").Parse(numConstTpl))
	template.Must(tpl.New("sfixed32").Parse(numTpl))
	template.Must(tpl.New("sfixed32Const").Parse(numConstTpl))
	template.Must(tpl.New("sfixed64").Parse(numTpl))
	template.Must(tpl.New("sfixed64Const").Parse(numConstTpl))

	template.Must(tpl.New("bool").Parse(boolTpl))
	template.Must(tpl.New("string").Parse(stringTpl))
	template.Must(tpl.New("stringConst").Parse(stringConstTpl))
	template.Must(tpl.New("bytes").Parse(bytesTpl))
	template.Must(tpl.New("bytesConst").Parse(bytesConstTpl))

	template.Must(tpl.New("enum").Parse(enumTpl))
	template.Must(tpl.New("enumConst").Parse(enumConstTpl))
	template.Must(tpl.New("repeated").Parse(repeatedTpl))
	template.Must(tpl.New("repeatedConst").Parse(repeatedConstTpl))
	template.Must(tpl.New("map").Parse(mapTpl))
	template.Must(tpl.New("mapConst").Parse(mapConstTpl))

	template.Must(tpl.New("any").Parse(anyTpl))
	template.Must(tpl.New("anyConst").Parse(anyConstTpl))
	template.Must(tpl.New("message").Parse(messageTpl))

	template.Must(tpl.New("timestamp").Parse(timestampTpl))
	template.Must(tpl.New("timestampConst").Parse(timestampConstTpl))
	template.Must(tpl.New("duration").Parse(durationTpl))
	template.Must(tpl.New("durationConst").Parse(durationConstTpl))

	template.Must(tpl.New("wrapper").Parse(wrapperTpl))
}

func DotnetFilePath(f pgs.File, ctx pgsgo.Context, tpl *template.Template) *pgs.FilePath {
	var fileName strings.Builder
	cap := true

	for _, c := range f.InputPath().BaseName() {
		upperC := unicode.ToUpper(c)
		isASCIILetter := 'A' <= upperC && upperC <= 'Z'

		if isASCIILetter || ('0' <= c && c <= '9') {
			if cap {
				fileName.WriteRune(upperC)
			} else {
				fileName.WriteRune(c)
			}
		}

		cap = !isASCIILetter
	}

	fileName.WriteString("Validator.cs")

	filePath := pgs.FilePath(fileName.String())
	return &filePath
}

type name struct {
	namespaces []string
	classNames []string
}

func dotnetName(e pgs.Entity) name {
	fqn := e.FullyQualifiedName()
	rawNamespaces := e.Package().ProtoName().String()
	rawClassNames := fqn[len(rawNamespaces)+2:] // +2 <=> strip leading and trailing dot of namespace

	namespaces := strings.Split(rawNamespaces, ".")
	classNames := strings.Split(rawClassNames, ".")

	for idx, component := range namespaces {
		namespaces[idx] = pgs.Name(component).UpperCamelCase().String()
	}
	for idx, component := range classNames {
		classNames[idx] = pgs.Name(component).UpperCamelCase().String()
	}

	if len(classNames) != 1 {
		intermediary := classNames
		classNames = make([]string, len(classNames)*2)

		for idx, cls := range intermediary {
			classNames[2*idx] = cls
			classNames[2*idx+1] = "Types"
		}

		classNames = classNames[:len(classNames)-1]
	}

	return name{
		namespaces,
		classNames,
	}
}

func (n name) className() string {
	return n.classNames[len(n.classNames)-1]
}

func (n name) fullyQualifiedName() string {
	return fmt.Sprintf("%s.%s", strings.Join(n.namespaces, "."), strings.Join(n.classNames, "."))
}

func (n name) parentClassNames() []string {
	return n.classNames[:len(n.classNames)-1]
}

func renderWithSuffixIfExists(tpl *template.Template, suffix string) func(ctx shared.RuleContext) (string, error) {
	return func(ctx shared.RuleContext) (string, error) {
		var b bytes.Buffer
		var err error

		if tpl.Lookup(ctx.Typ+suffix) != nil {
			err = tpl.ExecuteTemplate(&b, ctx.Typ+suffix, ctx)
		}

		return b.String(), err
	}
}

func err(ctx shared.RuleContext, reason ...interface{}) string {
	return fmt.Sprintf(
		"Envoyproxy.Validator.ValidationException.New(%q, %s, %q)",
		ctx.Field.FullyQualifiedName(),
		accessor(ctx),
		fmt.Sprint(reason...))
}

func unboundErr(e pgs.Entity, reason ...interface{}) string {
	return fmt.Sprintf(
		"Envoyproxy.Validator.ValidationException.New(%q, null, %q)",
		e.FullyQualifiedName(),
		fmt.Sprint(reason...))
}

func accessor(ctx shared.RuleContext) string {
	if ctx.AccessorOverride != "" {
		return ctx.AccessorOverride
	}

	return ctx.Field.Name().UpperCamelCase().String()
}

func constant(f pgs.Field, suffix string) string {
	return fmt.Sprintf("__Validation_%s%s", f.Name().UpperCamelCase().String(), suffix)
}

func fieldType(f pgs.Field) string {
	switch f.Type().ProtoType() {
	case pgs.Int32T, pgs.SInt32, pgs.SFixed32:
		return "int"

	case pgs.UInt32T, pgs.Fixed32T:
		return "uint"

	case pgs.Int64T, pgs.SInt64, pgs.SFixed64:
		return "long"

	case pgs.UInt64T, pgs.Fixed64T:
		return "ulong"

	case pgs.FloatT:
		return "float"

	case pgs.DoubleT:
		return "double"

	case pgs.BoolT:
		return "bool"

	case pgs.StringT:
		return "string"

	case pgs.BytesT:
		return "Google.Protobuf.ByteString"

	case pgs.EnumT:
		if f.Type().IsRepeated() {
			return f.Type().Element().Enum().Name().UpperCamelCase().String()
		}
		return f.Type().Enum().Name().UpperCamelCase().String()

	case pgs.MessageT:
		switch f.Type().Embed().Name() {
		case "Any":
			return "string"

		case "Duration":
			return "Google.Protobuf.WellKnownTypes.Duration"

		case "Timestamp":
			return "Google.Protobuf.WellKnownTypes.Timestamp"

		case "Int32Value":
			return "int"

		case "UInt32Value":
			return "uint"

		case "Int64Value":
			return "long"

		case "UInt64Value":
			return "ulong"

		case "FloatValue":
			return "float"

		case "DoubleValue":
			return "double"

		case "BoolValue":
			return "bool"

		case "StringValue":
			return "string"
		}
	}

	panic(fmt.Sprintf("type %v not defined", f.Type().Embed().Name()))
}

func literal(v interface{}) string {
	switch v.(type) {
	case int32:
		return fmt.Sprintf("%d", v)

	case uint32:
		return fmt.Sprintf("%dU", v)

	case int64:
		return fmt.Sprintf("%dL", v)

	case uint64:
		return fmt.Sprintf("%dUL", v)

	case float32:
		return fmt.Sprintf("%gF", v)

	case float64:
		return fmt.Sprintf("%gD", v)

	case bool:
		return fmt.Sprintf("%t", v)

	case string:
		return fmt.Sprintf("%q", v)

	case *string:
		v := v.(*string)
		return fmt.Sprintf("%q", *v)

	case []byte:
		v := v.([]byte)
		l := len(v)
		e := make([]string, l)

		for i := 0; i < l; i++ {
			e[i] = fmt.Sprintf("0x%X", v[i])
		}

		return fmt.Sprintf("Google.Protobuf.ByteString.CopyFrom(%s)", strings.Join(e, ", "))

	case *duration.Duration:
		v := v.(*duration.Duration)
		return fmt.Sprintf(
			"new Google.Protobuf.WellKnownTypes.Duration { Seconds = %dL, Nanos = %d }",
			v.Seconds,
			v.Nanos)

	case *timestamp.Timestamp:
		v := v.(*timestamp.Timestamp)
		return fmt.Sprintf(
			"new Google.Protobuf.WellKnownTypes.Timestamp { Seconds = %dL, Nanos = %d }",
			v.Seconds,
			v.Nanos)
	}

	panic(fmt.Sprintf("type %T not defined", v))
}

func unwrap(ctx shared.RuleContext) (shared.RuleContext, error) {
	override := ctx.WrapperTyp != "string" && ctx.WrapperTyp != "bytes"

	ctx, err := ctx.Unwrap("wrapper")
	if err != nil {
		return ctx, err
	}

	if override {
		ctx.AccessorOverride = fmt.Sprintf("%s.Value", ctx.Field.Name().UpperCamelCase().String())
	} else {
		ctx.AccessorOverride = ""
	}
	return ctx, nil
}

func parentClassNames(m pgs.Message) []string {
	return dotnetName(m).parentClassNames()
}

func className(m pgs.Message) string {
	return dotnetName(m).className()
}

func namespace(file pgs.File) string {
	// Explicit .NET namespace overrides implicit package
	options := file.Descriptor().GetOptions()
	if options != nil && options.CsharpNamespace != nil {
		return options.GetCsharpNamespace()
	}

	parts := strings.Split(file.Package().ProtoName().String(), ".")
	l := len(parts)
	for i := 0; i < l; i++ {
		parts[i] = pgs.Name(parts[i]).UpperCamelCase().String()
	}

	return strings.Join(parts, ".")
}

func oneofAccessor(o pgs.OneOf) string {
	return fmt.Sprintf("%sCase", o.Name().UpperCamelCase().String())
}

func oneofCase(f pgs.Field) string {
	return fmt.Sprintf("%sOneofCase.%s", f.OneOf().Name().UpperCamelCase().String(), f.Name().UpperCamelCase().String())
}

func enumType(ctx shared.RuleContext) string {
	t := ctx.Field.Type()

	ctx.Field.Package().ProtoName()

	var e pgs.Enum

	switch {
	case t.IsMap():
		if ctx.OnKey {
			e = t.Key().Enum()
		} else {
			e = t.Element().Enum()
		}
		break

	case t.IsRepeated():
		e = t.Element().Enum()
		break

	case t.IsEnum():
		e = t.Enum()

	default:
		panic(ctx)
	}

	return dotnetName(e).fullyQualifiedName()
}
