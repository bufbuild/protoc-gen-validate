package dotnet

const durationConstTpl = `
{{- $f := .Field -}}{{- $r := .Rules -}}
{{ if $r.Const }}
	private static readonly Google.Protobuf.WellKnownTypes.Duration {{ constant $f "Const" }} = {{ literal $r.GetConst }};
{{ end }}
{{ if $r.Lt }}
	private static readonly Google.Protobuf.WellKnownTypes.Duration {{ constant $f "Lt" }} = {{ literal $r.GetLt }};
{{ end }}
{{ if $r.Lte }}
	private static readonly Google.Protobuf.WellKnownTypes.Duration {{ constant $f "Lte" }} = {{ literal $r.GetLte }};
{{ end }}
{{ if $r.Gt }}
	private static readonly Google.Protobuf.WellKnownTypes.Duration {{ constant $f "Gt" }} = {{ literal $r.GetGt }};
{{ end }}
{{ if $r.Gte }}
	private static readonly Google.Protobuf.WellKnownTypes.Duration {{ constant $f "Gte" }} = {{ literal $r.GetGte }};
{{ end }}
{{ if $r.In }}
	private static readonly HashSet<Google.Protobuf.WellKnownTypes.Duration> {{ constant $f "In" }} = new HashSet<Google.Protobuf.WellKnownTypes.Duration>
	{
		{{ range $r.In }}
			{{ literal . }},
		{{ end }}
	};
{{ end }}
{{ if $r.NotIn }}
	private static readonly HashSet<Google.Protobuf.WellKnownTypes.Duration> {{ constant $f "NotIn" }} = new HashSet<Google.Protobuf.WellKnownTypes.Duration>
	{
		{{ range $r.NotIn }}
			{{ literal . }},
		{{ end }}
	};
{{ end }}
`

const durationTpl = `
{{- $f := .Field -}}{{- $r := .Rules -}}
{{ if $r.GetRequired }}
	if ({{ accessor . }} == null)
		throw {{ err . "value is required" }};
{{ end }}
{{ if or $r.Const $r.Lt $r.Lte $r.Gt $r.Gte $r.In $r.NotIn }}
	if ({{ accessor . }} != null)
	{
		{{ if $r.Const }}
			if (!Envoyproxy.Validator.Operator.Equal({{ accessor . }}, {{ constant $f "Const" }}))
				throw {{ err . "value must be equal " (durStr $r.GetConst) }};
		{{ end }}
		{{ if $r.Lt }}
			{{ if $r.Gt }}
				{{ if durGt $r.GetLt $r.GetGt }}
					if (Envoyproxy.Validator.Operator.Lte({{ accessor . }}, {{ constant $f "Gt" }}) || Envoyproxy.Validator.Operator.Gte({{ accessor . }}, {{ constant $f "Lt" }}))
						throw {{ err . "value must be inside range (" (durStr $r.GetGt) ", " (durStr $r.GetLt) ")" }};
				{{ else }}
					if (Envoyproxy.Validator.Operator.Gte({{ accessor . }}, {{ constant $f "Lt" }}) && Envoyproxy.Validator.Operator.Lte({{ accessor . }}, {{ constant $f "Gt" }}))
						throw {{ err . "value must be outside range [" (durStr $r.GetLt) ", " (durStr $r.GetGt) "]" }};
				{{ end }}
			{{ else if $r.Gte }}
				{{ if durGt $r.GetLt $r.GetGte }}
					if (Envoyproxy.Validator.Operator.Lt({{ accessor . }}, {{ constant $f "Gte" }}) || Envoyproxy.Validator.Operator.Gte({{ accessor . }}, {{ constant $f "Lt" }}))
						throw {{ err . "value must be inside range [" (durStr $r.GetGte) ", " (durStr $r.GetLt) ")" }};
				{{ else }}
					if (Envoyproxy.Validator.Operator.Gte({{ accessor . }}, {{ constant $f "Lt" }}) && Envoyproxy.Validator.Operator.Lt({{ accessor . }}, {{ constant $f "Gte" }}))
						throw {{ err . "value must be outside range [" (durStr $r.GetLt) ", " (durStr $r.GetGte) ")" }};
				{{ end }}
			{{ else }}
				if (Envoyproxy.Validator.Operator.Gte({{ accessor . }}, {{ constant $f "Lt" }}))
					throw {{ err . "value must be less than " (durStr $r.GetLt) }};
			{{ end }}
		{{ else if $r.Lte }}
			{{ if $r.Gt }}
				{{ if durGt $r.GetLte $r.GetGt }}
					if (Envoyproxy.Validator.Operator.Lte({{ accessor . }}, {{ constant $f "Gt" }}) || Envoyproxy.Validator.Operator.Gt({{ accessor . }}, {{ constant $f "Lte" }}))
						throw {{ err . "value must be inside range (" (durStr $r.GetGt) ", " (durStr $r.GetLte) "]" }};
				{{ else }}
					if (Envoyproxy.Validator.Operator.Gt({{ accessor . }}, {{ constant $f "Lte" }}) && Envoyproxy.Validator.Operator.Lte({{ accessor . }}, {{ constant $f "Gt" }}))
						throw {{ err . "value must be outside range (" (durStr $r.GetLte) ", " (durStr $r.GetGt) "]" }};
				{{ end }}
			{{ else if $r.Gte }}
				{{ if durGt $r.GetLte $r.GetGte }}
					if (Envoyproxy.Validator.Operator.Lt({{ accessor . }}, {{ constant $f "Gte" }}) || Envoyproxy.Validator.Operator.Gt({{ accessor . }}, {{ constant $f "Lte" }}))
						throw {{ err . "value must be inside range [" (durStr $r.GetGte) ", " (durStr $r.GetLte) "]" }};
				{{ else }}
					if (Envoyproxy.Validator.Operator.Gt({{ accessor . }}, {{ constant $f "Lte" }}) && Envoyproxy.Validator.Operator.Lt({{ accessor . }}, {{ constant $f "Gte" }}))
						throw {{ err . "value must be outside range (" (durStr $r.GetLte) ", " (durStr $r.GetGte) ")" }};
				{{ end }}
			{{ else }}
				if (Envoyproxy.Validator.Operator.Gt({{ accessor . }}, {{ constant $f "Lte" }}))
					throw {{ err . "value must be less than or equal to " (durStr $r.GetLte) }};
			{{ end }}
		{{ else if $r.Gt }}
			if (Envoyproxy.Validator.Operator.Lte({{ accessor . }}, {{ constant $f "Gt" }}))
				throw {{ err . "value must be greater than " (durStr $r.GetGt) }};
		{{ else if $r.Gte }}
			if (Envoyproxy.Validator.Operator.Lt({{ accessor . }}, {{ constant $f "Gte" }}))
				throw {{ err . "value must be greater than or equal to " (durStr $r.GetGte) }};
		{{ end }}
		{{ if $r.In }}
			if (!{{ constant $f "In" }}.Contains({{ accessor . }}))
				throw {{ err . "value must be in list " $r.In }};
		{{ end }}
		{{ if $r.NotIn }}
			if ({{ constant $f "NotIn" }}.Contains({{ accessor . }}))
				throw {{ err . "value must not be in list " $r.NotIn }};
		{{ end }}
	}
{{ end }}
`
