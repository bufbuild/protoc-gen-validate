package python

import (
	"fmt"
	"reflect"
	"strings"
	"text/template"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/duration"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/lyft/protoc-gen-star"
	"github.com/lyft/protoc-gen-star/lang/go"
	"github.com/envoyproxy/protoc-gen-validate/templates/shared"
)


func Register(tpl *template.Template, params pgs.Parameters) {
	fns := PythonFuncs{pgsgo.InitContext(params)}

	tpl.Funcs(map[string]interface{}{
		"accessor": fns.accessor,
		"err": fns.err,
		"lit": fns.lit,
		"output": fns.output,
		"lookup": fns.lookup,
		"inKey": fns.inKey,
		"byteStr": fns.byteStr,
		"typ": fns.Type,
		"oneof": fns.oneofTypeName,
		"enumValues": fns.enumValues,
		"name": fns.Name,
		"externalEnums": fns.externalEnums,
		"enumPackages": fns.enumPackages,
		"durLit": fns.durLit,
		"durStr": fns.durStr,
		"durGt": fns.durGt,
		"tsLit": fns.tsLit,
		"tsStr": fns.tsStr,
		"tsGt": fns.tsGt,
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
	template.Must(tpl.New("durationcmpTpl").Parse(durationcmpTpl))


	template.Must(tpl.New("timestampcmp").Parse(timestampcmp))
	template.Must(tpl.New("timestamp").Parse(timestampTpl))

	template.Must(tpl.New("wrapper").Parse(wrapperTpl))
}

type PythonFuncs struct{ pgsgo.Context }


func (fns PythonFuncs) accessor(ctx shared.RuleContext, reason ...interface{}) string {
	if ctx.AccessorOverride != "" {
		return ctx.AccessorOverride
	}
	return fmt.Sprintf("m.%s", ctx.Field.Name())
}

func (fns PythonFuncs) err(ctx shared.RuleContext, reason ...interface{}) string {
	return fmt.Sprintf("return False, '''%s'''", fmt.Sprint(reason))
}

func (fns PythonFuncs) lit(x interface{}) string {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Interface {
		val = val.Elem()
	}

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	switch val.Kind() {
	case reflect.String:
		return fmt.Sprintf("%q", x)
	case reflect.Bool:
		return strings.Title(fmt.Sprint(x))
	case reflect.Slice:
		els := make([]string, val.Len())
		switch reflect.TypeOf(x).Elem().Kind() {
		case reflect.Uint8:
			for i, l := 0, val.Len(); i < l; i++ {
				els[i] = fmt.Sprintf("\\x%x", val.Index(i).Interface())
			}
			return fmt.Sprintf("\"%s\"", strings.Join(els, ""))
		default:
			panic(fmt.Sprintf("don't know how to format literals of type %v", val.Kind()))
		}
	case reflect.Float32:
		return fmt.Sprintf("struct.unpack(\"f\", struct.pack(\"f\", %f))[0]", x) // Python numbers default to double
	default:
		return fmt.Sprint(x)
	}
}

func (fns PythonFuncs) durLit(dur *duration.Duration) string {
	return fmt.Sprintf(
		"(%d + (10**-9 * %d))",
		dur.GetSeconds(), dur.GetNanos())
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

func (fns PythonFuncs) lookup(f pgs.Field, name string) string {
	return fmt.Sprintf(
		"_%s_%s_%s",
		pgsgo.PGGUpperCamelCase(f.Message().Name()),
		pgsgo.PGGUpperCamelCase(f.Name()),
		name,
	)
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

func (fns PythonFuncs) oneofTypeName(f pgs.Field) pgsgo.TypeName {
	return pgsgo.TypeName(fns.OneofOption(f))
}

func (fns PythonFuncs) enumValues(f pgs.Field) string {
	enumVals := f.Type().Enum().Values()
	arr := "["
	for _, enum := range enumVals {
		arr = arr + fmt.Sprint(enum.Value()) + ","
	}
	arr = arr + "]"
	return arr
}

func (fns PythonFuncs) pythonizePath(path pgs.FilePath, extension string) string {
	return strings.Replace(path.SetExt(extension).String(), "/", ".", -1)
}

func (fns PythonFuncs) output(file pgs.File) string {
	return fns.pythonizePath(pgs.FilePath(file.Name().String()), "_pb2")
}

func (fns PythonFuncs) externalEnums(file pgs.File) []pgs.Enum {
	var out []pgs.Enum

	for _, msg := range file.AllMessages() {
		for _, fld := range msg.Fields() {
			if en := fld.Type().Enum(); fld.Type().IsEnum() && en.Package().ProtoName() != fld.Package().ProtoName() && fns.PackageName(en) != fns.PackageName(fld) {
				out = append(out, en)
			}
		}
	}
	return out
}

func (fns PythonFuncs) enumPackages(enums []pgs.Enum) []string {
	out := make([]string, len(enums))
	for idx, en := range enums {
		out[idx] = fns.pythonizePath(pgs.FilePath(en.File().Name().String()), "_pb2")
	}
	return out
}

func PythonFilePath(f pgs.File, ctx pgsgo.Context, tpl *template.Template) *pgs.FilePath {
	out := pgs.FilePath(f.Name().String())
	out = out.SetExt("_pb_validate.py")
	return &out
}