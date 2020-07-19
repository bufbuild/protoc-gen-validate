package dotnet

const messageTpl = `
{{- $f := .Field -}}{{- $r := .Rules -}}
{{ if $r.GetRequired }}
	if ({{ accessor . }} == null)
		throw {{ err . "value is required" }};
{{ end }}
{{ if .MessageRules.GetSkip }}
	// skipping validation for {{ $f.Name }}
{{ else }}
	{
		if ((object){{ accessor . }} is Envoyproxy.Validator.IValidateable validateable)
			validateable.Validate();
	}
{{ end }}
`
