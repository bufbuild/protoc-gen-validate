package dotnet

const timestampConstTpl = `
{{- $f := .Field -}}{{- $r := .Rules -}}
{{ if $r.Const }}
	private static readonly Google.Protobuf.WellKnownTypes.Timestamp {{ constant $f "Const" }} = {{ literal $r.GetConst }};
{{ end }}
{{ if $r.Lt }}
	private static readonly Google.Protobuf.WellKnownTypes.Timestamp {{ constant $f "Lt" }} = {{ literal $r.GetLt }};
{{ end }}
{{ if $r.Lte }}
	private static readonly Google.Protobuf.WellKnownTypes.Timestamp {{ constant $f "Lte" }} = {{ literal $r.GetLte }};
{{ end }}
{{ if $r.Gt }}
	private static readonly Google.Protobuf.WellKnownTypes.Timestamp {{ constant $f "Gt" }} = {{ literal $r.GetGt }};
{{ end }}
{{ if $r.Gte }}
	private static readonly Google.Protobuf.WellKnownTypes.Timestamp {{ constant $f "Gte" }} = {{ literal $r.GetGte }};
{{ end }}
{{ if $r.Within }}
	private static readonly Google.Protobuf.WellKnownTypes.Duration {{ constant $f "Within" }} = {{ literal $r.GetWithin }};
{{ end }}
`

const timestampTpl = `
{{- $f := .Field -}}{{- $r := .Rules -}}
{{ if $r.GetRequired }}
	if ({{ accessor . }} == null)
		throw {{ err . "value is required" }};
{{ end }}
{{ if or $r.Const $r.Lt $r.Lte $r.Gt $r.Gte $r.LtNow $r.GtNow $r.Within }}
	if ({{ accessor . }} != null)
	{
		{{ if $r.Const }}
			if ({{ accessor . }} != {{ constant $f "Const" }})
				throw {{ err . "value must be equal " (tsStr $r.GetConst) }};
		{{ end }}
		{{ if $r.Lt }}
			{{ if $r.Gt }}
				{{ if tsGt $r.GetLt $r.GetGt }}
					if ({{ accessor . }} <= {{ constant $f "Gt" }} || {{ accessor . }} >= {{ constant $f "Lt" }})
						throw {{ err . "value must be inside range (" (tsStr $r.GetGt) ", " (tsStr $r.GetLt) ")" }};
				{{ else }}
					if ({{ accessor . }} >= {{ constant $f "Lt" }} && {{ accessor . }} <= {{ constant $f "Gt" }})
						throw {{ err . "value must be outside range [" (tsStr $r.GetLt) ", " (tsStr $r.GetGt) "]" }};
				{{ end }}
			{{ else if $r.Gte }}
				{{ if tsGt $r.GetLt $r.GetGte }}
					if ({{ accessor . }} < {{ constant $f "Gte" }} || {{ accessor . }} >= {{ constant $f "Lt" }})
						throw {{ err . "value must be inside range [" (tsStr $r.GetGte) ", " (tsStr $r.GetLt) ")" }};
				{{ else }}
					if ({{ accessor . }} >= {{ constant $f "Lt" }} && {{ accessor . }} < {{ constant $f "Gte" }})
						throw {{ err . "value must be outside range [" (tsStr $r.GetLt) ", " (tsStr $r.GetGte) ")" }};
				{{ end }}
			{{ else }}
				if ({{ accessor . }} >= {{ constant $f "Lt" }})
					throw {{ err . "value must be less than " (tsStr $r.GetLt) }};
			{{ end }}
		{{ else if $r.Lte }}
			{{ if $r.Gt }}
				{{ if tsGt $r.GetLte $r.GetGt }}
					if ({{ accessor . }} <= {{ constant $f "Gt" }} || {{ accessor . }} > {{ constant $f "Lte" }})
						throw {{ err . "value must be inside range (" (tsStr $r.GetGt) ", " (tsStr $r.GetLte) "]" }};
				{{ else }}
					if ({{ accessor . }} > {{ constant $f "Lte" }} && {{ accessor . }} <= {{ constant $f "Gt" }})
						throw {{ err . "value must be outside range (" (tsStr $r.GetLte) ", " (tsStr $r.GetGt) "]" }};
				{{ end }}
			{{ else if $r.Gte }}
				{{ if tsGt $r.GetLte $r.GetGte }}
					if ({{ accessor . }} < {{ constant $f "Gte" }} || {{ accessor . }} > {{ constant $f "Lte" }})
						throw {{ err . "value must be inside range [" (tsStr $r.GetGte) ", " (tsStr $r.GetLte) "]" }};
				{{ else }}
					if ({{ accessor . }} > {{ constant $f "Lte" }} && {{ accessor . }} < {{ constant $f "Gte" }})
						throw {{ err . "value must be outside range (" (tsStr $r.GetLte) ", " (tsStr $r.GetGte) ")" }};
				{{ end }}
			{{ else }}
				if ({{ accessor . }} > {{ constant $f "Lte" }})
					throw {{ err . "value must be less than or equal to " (tsStr $r.GetLte) }};
			{{ end }}
		{{ else if $r.Gt }}
			if ({{ accessor . }} <= {{ constant $f "Gt" }})
				throw {{ err . "value must be greater than " (tsStr $r.GetGt) }};
		{{ else if $r.Gte }}
			if ({{ accessor . }} < {{ constant $f "Gte" }})
				throw {{ err . "value must be greater than or equal to " (tsStr $r.GetGte) }};
		{{ else if $r.LtNow }}
			{{ if $r.Within }}
				{
					var now = Envoyproxy.Validator.Operator.Now();
					if ({{ accessor . }} >= now || {{ accessor . }} < now - {{ constant $f "Within" }})
						throw {{ err . "value must be less than now within " (durStr $r.GetWithin) }};
				}
			{{ else }}
				{
					var now = Envoyproxy.Validator.Operator.Now();
					if ({{ accessor . }} >= now)
						throw {{ err . "value must be less than now" }};
				}
			{{ end }}
		{{ else if $r.GtNow }}
			{{ if $r.Within }}
				{
					var now = Envoyproxy.Validator.Operator.Now();
					if ({{ accessor . }} <= now || {{ accessor . }} > now + {{ constant $f "Within" }})
						throw {{ err . "value must be greater than now within " (durStr $r.GetWithin) }};
				}
			{{ else }}
				{
					var now = Envoyproxy.Validator.Operator.Now();
					if ({{ accessor . }} <= now)
						throw {{ err . "value must be greater than now" }};
				}
			{{ end }}
		{{ else if $r.Within }}
			{
				var now = Envoyproxy.Validator.Operator.Now();
				if ({{ accessor . }} < now - {{ constant $f "Within" }} || {{ accessor . }} > now + {{ constant $f "Within" }})
					throw {{ err . "value must be within " (durStr $r.GetWithin) " of now" }};
			}
		{{ end }}
	}
{{ end }}
`
