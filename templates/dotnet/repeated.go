package dotnet

const repeatedConstTpl = `
{{- $f := .Field -}}{{- $r := .Rules -}}
{{ if $r.Items }}
	{{ renderConst (.Elem "item" "") }}
{{ end }}
`

const repeatedTpl = `
{{- $f := .Field -}}{{- $r := .Rules -}}
{{ if $r.MinItems }}
	{{ if $r.MaxItems }}
		{{ if eq $r.GetMinItems $r.GetMaxItems }}
			if ({{ accessor . }}.Count != {{ $r.GetMinItems }})
				throw {{ err . "value must contain exactly " $r.GetMinItems " item(s)" }};
		{{ else }}
			{
				var count = {{ accessor . }}.Count;
				if (count < {{ $r.GetMinItems }} || count > {{ $r.GetMaxItems }})
					throw {{ err . "value must contain between " $r.GetMinItems " and " $r.GetMaxItems " items" }};
			}
		{{ end }}
	{{ else }}
		if ({{ accessor . }}.Count < {{ $r.GetMinItems }})
			throw {{ err . "value must contain at least " $r.GetMinItems " item(s)" }};
	{{ end }}
{{ else if $r.MaxItems }}
	if ({{ accessor . }}.Count > {{ $r.GetMaxItems }})
		throw {{ err . "value must contain at most " $r.GetMaxItems " item(s)" }};
{{ end }}
{{ if or $r.GetUnique (ne (.Elem "" "").Typ "none") }}
	{
		{{ if or $r.GetUnique }}
			var seen = new HashSet<{{ fieldType $f }}>();
		{{ end }}

		foreach (var item in {{ accessor . }})
		{
			{{ if or $r.GetUnique }}
				if (!seen.Add(item))
					throw {{ err . "value must contain unique items" }};
			{{ end }}

			{{ if ne (.Elem "" "").Typ "none" }}
				{{ render (.Elem "item" "") }}
			{{ end }}
		}
	}
{{ end }}
`
