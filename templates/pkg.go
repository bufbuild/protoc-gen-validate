package templates

import (
	"text/template"

	cctpl "github.com/lyft/protoc-gen-validate/templates/cc"
	gotpl "github.com/lyft/protoc-gen-validate/templates/go"
	"github.com/lyft/protoc-gen-validate/templates/shared"
)

func Template() *template.Template {
	tpl := template.New("protoc-gen-validate")
	shared.RegisterFunctions(tpl)

	cctpl.Register(tpl.New("cc"))
	gotpl.Register(tpl.New("go"))

	return tpl
}
