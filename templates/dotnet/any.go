package dotnet

const anyConstTpl = `
{{- $f := .Field -}}{{- $r := .Rules -}}
{{ if $r.In }}
	private static readonly HashSet<string> {{ constant $f "In" }} = new HashSet<string>
	{
		{{ range $r.In }}
			{{ literal . }},
		{{ end }}
	};
{{ end }}
{{ if $r.NotIn }}
	private static readonly HashSet<string> {{ constant $f "NotIn" }} = new HashSet<string>
	{
		{{ range $r.NotIn }}
			{{ literal . }},
		{{ end }}
	};
{{ end }}
`

const anyTpl = `
{{- $f := .Field -}}{{- $r := .Rules -}}
{{ if $r.GetRequired }}
	if ({{ accessor . }} == null)
		throw {{ err . "value is required" }};
{{ end }}
{{ if or $r.In $r.NotIn }}
	if ({{ accessor . }} != null)
	{
		{{ if $r.In }}
			if (!{{ constant $f "In" }}.Contains({{ accessor . }}.TypeUrl))
				throw {{ err . "type URL must be in list " $r.In }};
		{{ end }}
		{{ if $r.NotIn }}
			if ({{ constant $f "NotIn" }}.Contains({{ accessor . }}.TypeUrl))
				throw {{ err . "type URL must not be in list " $r.NotIn }};
		{{ end }}
	}
{{ end }}
`
