package golang

import (
	"text/template"

	"github.com/envoyproxy/protoc-gen-validate/templates/shared"
	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
)

const PluginName = "go"

func MakePlugin(params pgs.Parameters) *shared.TemplatePlugin {
	return &shared.TemplatePlugin{
		Templates: []*template.Template{shared.MakeTemplate("go", register, params)},
		PathFunction: func(f pgs.File, ctx pgsgo.Context, tpl *template.Template) *pgs.FilePath {
			out := ctx.OutputPath(f)
			out = out.SetExt(".validate." + tpl.Name())
			return &out
		},
		Name: PluginName,
	}
}
