package shared

import (
	"text/template"

	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
)

type RegisterFn func(tpl *template.Template, params pgs.Parameters)
type FilePathFn func(f pgs.File, ctx pgsgo.Context, tpl *template.Template) *pgs.FilePath

type TemplatePlugin struct {
	Templates    []*template.Template
	PathFunction FilePathFn
	Name         string
}

func MakeTemplate(ext string, fn RegisterFn, params pgs.Parameters) *template.Template {
	tpl := template.New(ext)
	RegisterFunctions(tpl, params)
	fn(tpl, params)
	return tpl
}
