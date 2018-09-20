package templates

import (
	"text/template"

	"github.com/lyft/protoc-gen-star"
	"github.com/lyft/protoc-gen-validate/templates/cc"
	"github.com/lyft/protoc-gen-validate/templates/go"
	"github.com/lyft/protoc-gen-validate/templates/gogo"
	"github.com/lyft/protoc-gen-validate/templates/shared"
)

type RegisterFn func(tpl *template.Template, params pgs.Parameters)

func makeTemplate(ext string, fn RegisterFn, params pgs.Parameters) *template.Template {
	tpl := template.New(ext)
	shared.RegisterFunctions(tpl, params)
	fn(tpl, params)
	return tpl
}

func Template(params pgs.Parameters) map[string][]*template.Template {
	return map[string][]*template.Template{
		"cc":   {makeTemplate("h", cc.RegisterHeader, params), makeTemplate("cc", cc.RegisterModule, params)},
		"go":   {makeTemplate("go", golang.Register, params)},
		"gogo": {makeTemplate("go", gogo.Register, params)},
	}
}
