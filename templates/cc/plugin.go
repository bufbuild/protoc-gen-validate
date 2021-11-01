package cc

import (
	"text/template"

	"github.com/envoyproxy/protoc-gen-validate/templates/shared"
	pgs "github.com/lyft/protoc-gen-star"
)

const PluginName = "cc"

func MakePlugin(params pgs.Parameters) *shared.TemplatePlugin {
	return &shared.TemplatePlugin{
		Templates: []*template.Template{
			shared.MakeTemplate("h", registerHeader, params),
			shared.MakeTemplate("cc", registerModule, params),
		},
		PathFunction: ccFilePath,
		Name:         PluginName,
	}
}
