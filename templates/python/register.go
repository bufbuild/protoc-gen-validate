package python

import (
	"fmt"
	"github.com/envoyproxy/protoc-gen-validate/templates/shared"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/duration"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/lyft/protoc-gen-star"
	"github.com/lyft/protoc-gen-star/lang/go"
	"reflect"
	"strings"
	"text/template"
)

func Register(tpl *template.Template, params pgs.Parameters) {
	fns := PythonFuncs{pgsgo.InitContext(params)}

	tpl.Funcs(map[string]interface{}{
		"accessor":       fns.accessor,
		"ctype":          fns.cType,
		"durGt":          fns.durGt,
		"durLit":         fns.durLit,
		"durStr":         fns.durStr,
		"err":            fns.err,
		"hasAccessor":    fns.hasAccessor,
		"shouldImport":   fns.shouldImport,
		"inKey":          fns.inKey,
		"lit":            fns.lit,
		"lookup":         fns.lookup,
		"msgTyp":         fns.msgTyp,
		"name":           fns.Name,
		"output":         fns.output,
		"shouldValidate": fns.shouldValidate,
		"tsGt":           fns.tsGt,
		"tsLit":          fns.tsLit,
		"tsStr":          fns.tsStr,
		"typ":            fns.Type,
		"unimplemented":  fns.failUnimplemented,
		"unwrap":         fns.unwrap,
	})

	template.Must(tpl.Parse(fileTpl))
	template.Must(tpl.New("msg").Parse(msgTpl))
	template.Must(tpl.New("const").Parse(constTpl))
	template.Must(tpl.New("ltgt").Parse(ltgtTpl))
	template.Must(tpl.New("in").Parse(inTpl))
	template.Must(tpl.New("required").Parse(requiredTpl))

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

	template.Must(tpl.New("bool").Parse(constTpl))
	template.Must(tpl.New("string").Parse(strTpl))
	template.Must(tpl.New("bytes").Parse(bytesTpl))

	template.Must(tpl.New("email").Parse(emailTpl))
	template.Must(tpl.New("hostname").Parse(hostTpl))

	template.Must(tpl.New("enum").Parse(enumTpl))
	template.Must(tpl.New("message").Parse(messageTpl))
	template.Must(tpl.New("repeated").Parse(repTpl))
	template.Must(tpl.New("map").Parse(mapTpl))

	template.Must(tpl.New("any").Parse(anyTpl))
	template.Must(tpl.New("duration").Parse(durationTpl))
	template.Must(tpl.New("timestamp").Parse(timestampTpl))

	template.Must(tpl.New("wrapper").Parse(wrapperTpl))
}

type PythonFuncs struct{ pgsgo.Context }

func PythonFilePath(f pgs.File, ctx pgsgo.Context, tpl *template.Template) *pgs.FilePath {
	out := pgs.FilePath(f.Name().String())
	out = out.SetExt("_pb_validate.py")
	return &out
}

func (fns PythonFuncs) output(file pgs.File) string {
	return fns.pythonizePath(pgs.FilePath(file.Name().String()), "_pb2")
}

func (fns PythonFuncs) pythonizePath(path pgs.FilePath, extension string) string {
	return strings.Replace(path.SetExt(extension).String(), "/", ".", -1)
}

func (fns PythonFuncs) accessor(ctx shared.RuleContext, reason ...interface{}) string {
	if ctx.AccessorOverride != "" {
		return ctx.AccessorOverride
	}
	return fmt.Sprintf("m.%s", ctx.Field.Name())
}

func (fns PythonFuncs) err(ctx shared.RuleContext, reason ...interface{}) string {
	return fmt.Sprintf("return False, '''%s'''", fmt.Sprint(reason))
}

func (fns PythonFuncs) failUnimplemented() string {
	return "raise UnimplementedException()"
}

func (fns PythonFuncs) lookup(f pgs.Field, name string) string {
	return fmt.Sprintf(
		"_%s_%s_%s",
		pgsgo.PGGUpperCamelCase(f.Message().Name()),
		pgsgo.PGGUpperCamelCase(f.Name()),
		name,
	)
}

func (fns PythonFuncs) lit(x interface{}) string {
	val := reflect.ValueOf(x)

	switch val.Kind() {
	case reflect.String:
		return fmt.Sprintf("%q", x)
	case reflect.Bool:
		return strings.Title(fmt.Sprint(x))
	case reflect.Float32:
		// Must convert double to float
		return fmt.Sprintf("struct.unpack(\"f\", struct.pack(\"f\", %f))[0]", x)
	default:
		return fmt.Sprint(x)
	}
}

func (fns PythonFuncs) inKey(f pgs.Field, x interface{}) string {
	switch f.Type().ProtoType() {
	case pgs.BytesT:
		return fns.byteStr(x.([]byte))
	case pgs.MessageT:
		switch x := x.(type) {
		case *duration.Duration:
			return fns.durLit(x)
		default:
			return fns.lit(x)
		}
	default:
		return fns.lit(x)
	}
}

func (fns PythonFuncs) byteStr(x []byte) string {
	elms := make([]string, len(x))
	for i, b := range x {
		elms[i] = fmt.Sprintf(`\x%X`, b)
	}

	return fmt.Sprintf(`"%s"`, strings.Join(elms, ""))
}

func (fns PythonFuncs) durLit(dur *duration.Duration) string {
	return fmt.Sprintf(
		"(%d + (10**-9 * %d))",
		dur.GetSeconds(), dur.GetNanos())
}

func (fns PythonFuncs) unwrap(ctx shared.RuleContext, name string) (shared.RuleContext, error) {
	ctx, err := ctx.Unwrap("wrapper")
	if err != nil {
		return ctx, err
	}

	ctx.AccessorOverride = fmt.Sprintf("%s.Get%s()", name,
		pgsgo.PGGUpperCamelCase(ctx.Field.Type().Embed().Fields()[0].Name()))

	return ctx, nil
}

func (fns PythonFuncs) cType(t pgs.FieldType, usage bool) string {
	if t.IsEmbed() {
		if usage {
			return t.Embed().Name().String()
		} else {
			return fns.importer(t)
		}
	}
	return ""
}

func (fns PythonFuncs) importer(t pgs.FieldType) string {
	return "from " + fns.packageName(t.Embed()) + "." + strings.ToLower(t.Embed().Name().String()) + "_pb_validate import validate_" + t.Embed().Name().String()
}

func (fns PythonFuncs) hasAccessor(ctx shared.RuleContext) string {
	return fmt.Sprintf(
		"m.HasField(\"%s\")",
		ctx.Field.Name())
}

func (fns PythonFuncs) packageName(msg pgs.Entity) string {
	return strings.Join(msg.Package().ProtoName().Split(), ".")
}

func (fns PythonFuncs) shouldImport(f pgs.Field) bool {
	if f.Type().Embed().Package().ProtoName() != f.Message().Package().ProtoName() {
		if f.Type().Embed().Package().ProtoName() != "google.protobuf" {
			return true
		}
	}
	return false
}

func (fns PythonFuncs) shouldValidate(f pgs.Field) bool {
	return f.Type().Embed().Package().ProtoName() != "google.protobuf"
}

func (fns PythonFuncs) msgTyp(message pgs.Message) pgsgo.TypeName {
	return pgsgo.TypeName(fns.Name(message))
}

func (fns PythonFuncs) durStr(dur *duration.Duration) string {
	d, _ := ptypes.Duration(dur)
	return d.String()
}

func (fns PythonFuncs) durGt(a, b *duration.Duration) bool {
	ad, _ := ptypes.Duration(a)
	bd, _ := ptypes.Duration(b)

	return ad > bd
}

func (fns PythonFuncs) tsLit(ts *timestamp.Timestamp) string {
	return fmt.Sprintf(
		"(%d + (10**-9 * %d))",
		ts.GetSeconds(), ts.GetNanos(),
	)
}

func (fns PythonFuncs) tsStr(ts *timestamp.Timestamp) string {
	t, _ := ptypes.Timestamp(ts)
	return t.String()
}

func (fns PythonFuncs) tsGt(a, b *timestamp.Timestamp) bool {
	at, _ := ptypes.Timestamp(a)
	bt, _ := ptypes.Timestamp(b)

	return bt.Before(at)
}
