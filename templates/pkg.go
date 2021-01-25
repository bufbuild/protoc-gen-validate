package templates

import (
	"text/template"

	"github.com/envoyproxy/protoc-gen-validate/templates/cc"
	"github.com/envoyproxy/protoc-gen-validate/templates/go"
	"github.com/envoyproxy/protoc-gen-validate/templates/java"
	"github.com/envoyproxy/protoc-gen-validate/templates/php_yaml"
	"github.com/envoyproxy/protoc-gen-validate/templates/shared"
	"github.com/lyft/protoc-gen-star"
	"github.com/lyft/protoc-gen-star/lang/go"
)

type RegisterFn func(tpl *template.Template, params pgs.Parameters)
type FilePathFn func(f pgs.File, ctx pgsgo.Context, tpl *template.Template) *pgs.FilePath

func makeTemplate(ext string, fn RegisterFn, params pgs.Parameters) *template.Template {
	tpl := template.New(ext)
	shared.RegisterFunctions(tpl, params)
	fn(tpl, params)
	return tpl
}

func Template(params pgs.Parameters) map[string][]*template.Template {
	return map[string][]*template.Template{
		"cc":       {makeTemplate("h", cc.RegisterHeader, params), makeTemplate("cc", cc.RegisterModule, params)},
		"go":       {makeTemplate("go", golang.Register, params)},
		"java":     {makeTemplate("java", java.Register, params)},
		"php_yaml": {makeTemplate("php_yaml", php_yaml.Register, params)},
	}
}

func FilePathFor(tpl *template.Template) FilePathFn {
	switch tpl.Name() {
	case "h":
		return cc.CcFilePath
	case "cc":
		return cc.CcFilePath
	case "java":
		return java.JavaFilePath
	case "php_yaml":
		return php_yaml.PhpYamlFilePath
	default:
		return func(f pgs.File, ctx pgsgo.Context, tpl *template.Template) *pgs.FilePath {
			out := ctx.OutputPath(f)
			out = out.SetExt(".validate." + tpl.Name())
			return &out
		}
	}
}
