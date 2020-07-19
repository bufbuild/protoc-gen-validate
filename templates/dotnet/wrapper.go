package dotnet

const wrapperTpl = `
{{- $f := .Field -}}{{- $r := .Rules -}}
{{ if or (eq .WrapperTyp "string") (eq .WrapperTyp "bytes") }}
	if ({{ accessor . }} != null)
	{
		{{ render (unwrap .) }}
	}
{{ else }}
	if ({{ accessor . }}.HasValue)
	{
		{{ render (unwrap .) }}
	}
{{ end }}
{{ if .MessageRules.GetRequired }}
	else
	{
		throw {{ err . "value is required" }};
	}
{{ end }}
`
