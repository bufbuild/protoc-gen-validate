package dotnet

const numConstTpl = `
{{- $f := .Field -}}{{- $r := .Rules -}}
{{ if $r.In }}
	private static readonly HashSet<{{ fieldType $f }}> {{ constant $f "In" }} = new HashSet<{{ fieldType $f }}>
	{
		{{ range $r.In }}
			{{ literal . }},
		{{ end }}
	};
{{ end }}
{{ if $r.NotIn }}
	private static readonly HashSet<{{ fieldType $f }}> {{ constant $f "NotIn" }} = new HashSet<{{ fieldType $f }}>
	{
		{{ range $r.NotIn }}
			{{ literal . }},
		{{ end }}
	};
{{ end }}
`

const numTpl = `
{{- $f := .Field -}}{{- $r := .Rules -}}
{{ if $r.Const }}
	if ({{ accessor . }} != {{ literal $r.GetConst }})
		throw {{ err . "value must be equal " $r.GetConst }};
{{ end }}
{{ if $r.In }}
	if (!{{ constant $f "In" }}.Contains({{ accessor . }}))
		throw {{ err . "value must be in list " $r.In }};
{{ end }}
{{ if $r.NotIn }}
	if ({{ constant $f "NotIn" }}.Contains({{ accessor . }}))
		throw {{ err . "value must not be in list " $r.NotIn }};
{{ end }}
{{ if $r.Lt }}
	{{ if $r.Gt }}
		{{ if gt $r.GetLt $r.GetGt }}
			if ({{ accessor . }} <= {{ literal $r.GetGt }} || {{ accessor . }} >= {{ literal $r.GetLt }})
				throw {{ err . "value must be inside range (" $r.GetGt ", " $r.GetLt ")" }};
		{{ else }}
			if ({{ accessor . }} >= {{ literal $r.GetLt }} && {{ accessor . }} <= {{ literal $r.GetGt }})
				throw {{ err . "value must be outside range [" $r.GetLt ", " $r.GetGt "]" }};
		{{ end }}
	{{ else if $r.Gte }}
		{{ if gt $r.GetLt $r.GetGte }}
			if ({{ accessor . }} < {{ literal $r.GetGte }} || {{ accessor . }} >= {{ literal $r.GetLt }})
				throw {{ err . "value must be inside range [" $r.GetGte ", " $r.GetLt ")" }};
		{{ else }}
			if ({{ accessor . }} >= {{ literal $r.GetLt }} && {{ accessor . }} < {{ literal $r.GetGte }})
				throw {{ err . "value must be outside range [" $r.GetLt ", " $r.GetGte ")" }};
		{{ end }}
	{{ else }}
		if ({{ accessor . }} >= {{ literal $r.GetLt }})
			throw {{ err . "value must be less than " $r.GetLt }};
	{{ end }}
{{ else if $r.Lte }}
	{{ if $r.Gt }}
		{{ if gt $r.GetLte $r.GetGt }}
			if ({{ accessor . }} <= {{ literal $r.GetGt }} || {{ accessor . }} > {{ literal $r.GetLte }})
				throw {{ err . "value must be inside range (" $r.GetGt ", " $r.GetLte "]" }};
		{{ else }}
			if ({{ accessor . }} > {{ literal $r.GetLte }} && {{ accessor . }} <= {{ literal $r.GetGt }})
				throw {{ err . "value must be outside range (" $r.GetLte ", " $r.GetGt "]" }};
		{{ end }}
	{{ else if $r.Gte }}
		{{ if gt $r.GetLte $r.GetGte }}
			if ({{ accessor . }} < {{ literal $r.GetGte }} || {{ accessor . }} > {{ literal $r.GetLte }})
				throw {{ err . "value must be inside range [" $r.GetGte ", " $r.GetLte "]" }};
		{{ else }}
			if ({{ accessor . }} > {{ literal $r.GetLte }} && {{ accessor . }} < {{ literal $r.GetGte }})
				throw {{ err . "value must be outside range (" $r.GetLte ", " $r.Gte ")" }};
		{{ end }}
	{{ else }}
		if ({{ accessor . }} > {{ literal $r.GetLte }})
			throw {{ err . "value must be less than or equal to " $r.GetLte }};
	{{ end }}
{{ else if $r.Gt }}
	if ({{ accessor . }} <= {{ literal $r.GetGt }})
		throw {{ err . "value must be greater than " $r.GetGt }};
{{ else if $r.Gte }}
	if ({{ accessor . }} < {{ literal $r.GetGte }})
		throw {{ err . "value must be greater than or equal to " $r.GetGte }};
{{ end }}
`
