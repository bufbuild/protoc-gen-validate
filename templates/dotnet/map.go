package dotnet

const mapConstTpl = `
{{- $f := .Field -}}{{- $r := .Rules -}}
{{ if $r.Keys }}
	{{ renderConst (.Key "key" "") }}
{{ end }}
{{ if $r.Values }}
	{{ renderConst (.Elem "value" "") }}
{{ end }}
`

const mapTpl = `
{{- $f := .Field -}}{{- $r := .Rules -}}
{{ if $r.MinPairs }}
	{{ if $r.MaxPairs }}
		{{ if eq $r.GetMinPairs $r.GetMaxPairs }}
			if ({{ accessor . }}.Count != {{ $r.GetMinPairs }})
				throw {{ err . "value must contain exactly " $r.GetMinPairs " pair(s)" }};
		{{ else }}
			{
				var count = {{ accessor . }}.Count;
				if (count < {{ $r.GetMinPairs }} || count > {{ $r.GetMaxPairs }})
					throw {{ err . "value must contain between " $r.GetMinPairs " and " $r.GetMaxPairs " pairs" }};
			}
		{{ end }}
	{{ else }}
		if ({{ accessor . }}.Count < {{ $r.GetMinPairs }})
			throw {{ err . "value must contain at least " $r.GetMinPairs " pair(s)" }};
	{{ end }}
{{ else if $r.MaxPairs }}
	if ({{ accessor . }}.Count > {{ $r.GetMaxPairs }})
		throw {{ err . "value must contain at most " $r.GetMaxPairs " pair(s)" }};
{{ end }}
{{ if or $r.GetNoSparse (ne (.Key "" "").Typ "none") (ne (.Elem "" "").Typ "none") }}
	foreach (var item in {{ accessor . }})
	{
		{{ if $r.GetNoSparse }}
			if (item.Value == null)
				throw {{ err . "value most not be sparse" }};
		{{ end }}
		{{ render (.Key "item.Key" "item.Key") }}
		{{ render (.Elem "item.Value" "item.Key") }}
	}
{{ end }}
`
