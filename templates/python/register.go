package python

import (
	"fmt"
	"github.com/envoyproxy/protoc-gen-validate/templates/shared"
	"github.com/golang/protobuf/ptypes/duration"
	"github.com/lyft/protoc-gen-star"
	"github.com/lyft/protoc-gen-star/lang/go"
	"reflect"
	"strings"
	"text/template"
)

func Register(tpl *template.Template, params pgs.Parameters) {
	fns := PythonFuncs{pgsgo.InitContext(params)}

	tpl.Funcs(map[string]interface{}{
		"accessor": fns.accessor,
		"err": fns.err,
		"inKey": fns.inKey,
		"lit": fns.lit,
		"lookup": fns.lookup,
		"name": fns.Name,
		"output": fns.output,
		"typ": fns.Type,
		"unimplemented": fns.failUnimplemented,
	})

	template.Must(tpl.Parse(fileTpl))
	template.Must(tpl.New("msg").Parse(msgTpl))
	template.Must(tpl.New("const").Parse(constTpl))
	template.Must(tpl.New("in").Parse(inTpl))

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
