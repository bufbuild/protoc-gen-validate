package dotnet

const enumConstTpl = `
{{- $f := .Field -}}{{- $r := .Rules -}}
{{ if $r.In }}
	private static readonly HashSet<int> {{ constant $f "In" }} = new HashSet<int>
	{
		{{ range $r.In }}
			{{ literal . }},
		{{ end }}
	};
{{ end }}
{{ if $r.NotIn }}
	private static readonly HashSet<int> {{ constant $f "NotIn" }} = new HashSet<int>
	{
		{{ range $r.NotIn }}
			{{ literal . }},
		{{ end }}
	};
{{ end }}
`

const enumTpl = `
{{- $f := .Field -}}{{- $r := .Rules -}}
{{ if $r.Const }}
	if ((int){{ accessor . }} != {{ literal $r.GetConst }})
		throw {{ err . "value must be equal " $r.GetConst }};
{{ end }}
{{ if $r.In }}
	if (!{{ constant $f "In" }}.Contains((int){{ accessor . }}))
		throw {{ err . "value must be in list " $r.In }};
{{ end }}
{{ if $r.NotIn }}
	if ({{ constant $f "NotIn" }}.Contains((int){{ accessor . }}))
		throw {{ err . "value must not be in list " $r.NotIn }};
{{ end }}
{{ if $r.GetDefinedOnly }}
	if (!Enum.IsDefined(typeof({{ enumType . }}), {{ accessor . }}))
		throw {{ err . "value must be oen of the defined enum values" }};
{{ end }}
`
