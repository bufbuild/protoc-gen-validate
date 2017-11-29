package tpl

const fileTpl = `// Code generated by protoc-gen-validate
// source: {{ .InputPath }}
// DO NOT EDIT!!!
#include <string>
#include <sstream>

#include "{{ .File.OutputPath.SetExt ".h" }}"

{{ range .Package.ProtoName.Split }}
namespace {{ . }} {
{{- end }}

using std::string;

{{ range .AllMessages }}
	{{- template "decl" . }}
{{ end }}

{{ range .AllMessages }}
	{{- template "msg" . }}
{{ end }}

{{ range .Package.ProtoName.Split -}}
} // namespace
{{ end }}

#define X_{{ .Package.ProtoName.ScreamingSnakeCase }}_{{ .File.InputPath.BaseName | upper }}(X) \
{{ range .AllMessages -}}
	X({{class . }}) \
{{ end }}
`
