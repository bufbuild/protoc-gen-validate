package dotnet

const boolTpl = `
{{- $r := .Rules -}}
{{ if $r.Const }}
	if ({{ accessor . }} != {{ literal $r.GetConst }})
		throw {{ err . "value must be equal " $r.GetConst }};
{{ end }}
`
