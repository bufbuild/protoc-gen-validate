package java

import (
	"text/template"

	"github.com/envoyproxy/protoc-gen-validate/templates/shared"
	pgs "github.com/lyft/protoc-gen-star"
)

const PluginName = "java"

func MakePlugin(params pgs.Parameters) *shared.TemplatePlugin {
	return &shared.TemplatePlugin{
		Templates:    []*template.Template{shared.MakeTemplate("java", register, params)},
		PathFunction: javaFilePath,
		Name:         PluginName,
	}
}
