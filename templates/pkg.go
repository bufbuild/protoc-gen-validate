package templates

import (
	"errors"
	"plugin"

	"github.com/envoyproxy/protoc-gen-validate/templates/cc"
	golang "github.com/envoyproxy/protoc-gen-validate/templates/go"
	"github.com/envoyproxy/protoc-gen-validate/templates/java"
	"github.com/envoyproxy/protoc-gen-validate/templates/shared"
	pgs "github.com/lyft/protoc-gen-star"
)

type makePluginFn func(params pgs.Parameters) *shared.TemplatePlugin

func MakeTemplateForLang(params pgs.Parameters, lang string) *shared.TemplatePlugin {
	switch lang {
	case cc.PluginName:
		return cc.MakePlugin(params)
	case golang.PluginName:
		return golang.MakePlugin(params)
	case java.PluginName:
		return java.MakePlugin(params)
	default:
		return nil
	}
}

func MakeTemplateFromPlugin(path string, params pgs.Parameters) (*shared.TemplatePlugin, error) {
	plug, err := plugin.Open(path)
	if err != nil {
		return nil, err
	}

	symPlugin, err := plug.Lookup("Plugin")
	if err != nil {
		return nil, err
	}

	makePlugin, ok := symPlugin.(makePluginFn)
	if !ok {
		return nil, errors.New("loaded object has an incorrect type, expected: *TemplatePlugin")
	}

	return makePlugin(params), nil
}
