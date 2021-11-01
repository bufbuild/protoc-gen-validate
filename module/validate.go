package module

import (
	"strings"

	"github.com/envoyproxy/protoc-gen-validate/templates"
	"github.com/envoyproxy/protoc-gen-validate/templates/java"
	"github.com/envoyproxy/protoc-gen-validate/templates/shared"
	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
)

const (
	validatorName   = "validator"
	langParam       = "lang"
	langPluginParam = "lang-plugin"
	moduleParam     = "module"
)

type Module struct {
	*pgs.ModuleBase
	ctx pgsgo.Context
}

func Validator() pgs.Module { return &Module{ModuleBase: &pgs.ModuleBase{}} }

func (m *Module) InitContext(ctx pgs.BuildContext) {
	m.ModuleBase.InitContext(ctx)
	m.ctx = pgsgo.InitContext(ctx.Parameters())
}

func (m *Module) Name() string { return validatorName }

func (m *Module) Execute(targets map[string]pgs.File, pkgs map[string]pgs.Package) []pgs.Artifact {
	module := m.Parameters().Str(moduleParam)
	plugin := m.resolveTemplate()

	for _, f := range targets {
		m.Push(f.Name().String())

		for _, msg := range f.AllMessages() {
			m.CheckRules(msg)
		}

		for _, tpl := range plugin.Templates {
			out := plugin.PathFunction(f, m.ctx, tpl)

			// A nil path means no output should be generated for this file - as controlled by
			// implementation-specific FilePathFor implementations.
			// Ex: Don't generate Java validators for files that don't reference PGV.
			if out != nil {
				outPath := strings.TrimLeft(strings.ReplaceAll(out.String(), module, ""), "/")

				if opts := f.Descriptor().GetOptions(); opts != nil && opts.GetJavaMultipleFiles() && plugin.Name == "java" {
					// TODO: Only Java supports multiple file generation. If more languages add multiple file generation
					// support, the implementation should be made more inderect.
					for _, msg := range f.Messages() {
						m.AddGeneratorTemplateFile(java.JavaMultiFilePath(f, msg).String(), tpl, msg)
					}
				} else {
					m.AddGeneratorTemplateFile(outPath, tpl, f)
				}
			}
		}

		m.Pop()
	}

	return m.Artifacts()
}

func (m *Module) resolveTemplate() *shared.TemplatePlugin {
	lang := m.Parameters().Str(langParam)
	langPlugin := m.Parameters().Str(langPluginParam)
	m.Assert(lang != "" || langPlugin != "", "`lang` parameter or `lang-plugin` must be set")

	if lang != "" {
		plugin := templates.MakeTemplateForLang(m.Parameters(), lang)
		m.Assert(plugin != nil, "could not find templates for `lang`: ", lang)
		return plugin
	}

	plugin, err := templates.MakeTemplateFromPlugin(langPlugin, m.Parameters())
	m.Assert(err == nil, err)
	return plugin
}

var _ pgs.Module = (*Module)(nil)
