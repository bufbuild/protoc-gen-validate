package python

import (
	"fmt"
	"github.com/lyft/protoc-gen-validate/templates/shared"
	"reflect"
	"strings"
	"text/template"

	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
)


func Register(tpl *template.Template, params pgs.Parameters) {
	fns := PythonFuncs{pgsgo.InitContext(params)}
	tpl.Funcs(map[string]interface{}{
		"accessor": fns.accessor,
		"err": fns.err,
	})
	template.Must(tpl.New("bool").Parse(constTpl))

}

type PythonFuncs struct{ pgsgo.Context }


func (fns PythonFuncs) accessor(ctx shared.RuleContext, reason ...interface{}) string {
	if ctx.AccessorOverride != "" {
		return ctx.AccessorOverride
	}

	return fmt.Sprintf("m.%s()", fns.Name(ctx.Field))
}

func (fns PythonFuncs) err(ctx shared.RuleContext, reason ...interface{}) string {
	return "return False"
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
	default:
		return fmt.Sprint(x)
	}
}