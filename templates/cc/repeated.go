package tpl

const repTpl = `
	{{ $f := .Field }}{{ $r := .Rules }}{{ $typ := $f.Type | ctype }}

	{{ if $r.GetMinItems }}
		{{ if eq $r.GetMinItems $r.GetMaxItems }}
			if ({{ accessor . }}.size() != {{ $r.GetMinItems }}) {
				{{ err . "value must contain exactly " $r.GetMinItems " item(s)" }}
			}
		{{ else if $r.MaxItems }}
			if ({{ accessor . }}.size() < {{ $r.GetMinItems }} || {{ accessor . }}.size() > {{ $r.GetMaxItems }}) {
			 	{{ err . "value must contain between " $r.GetMinItems " and " $r.GetMaxItems " items, inclusive" }}
			}
		{{ else }}
			if ({{ accessor . }}.size() < {{ $r.GetMinItems }}) {
				{{ err . "value must contain at least " $r.GetMinItems " item(s)" }}
			}
		{{ end }}
	{{ else if $r.MaxItems }}
		if ({{ accessor . }}.size() > {{ $r.GetMaxItems }}) {
			{{ err . "value must contain no more than " $r.GetMaxItems " item(s)" }}
		}
	{{ end }}

	{{ if $r.GetUnique }}
	std::set<std::reference_wrapper<const {{ $typ }}>> {{ lookup $f "Unique" }};
	{{ end }}

	{{ if or $r.GetUnique (ne (.Elem "" "").Typ "none") }}
		for (int i = 0; i < {{ accessor . }}.size(); i++) {
			const {{ $typ }}& item = {{ accessor . }}.Get(i);

			{{ if $r.GetUnique }}
				auto p = {{ lookup $f "Unique" }}.emplace(item);
				if (p.second == false) {
					{{ errIdx . "idx" "repeated value must contain unique items" }}
				}
			{{ end }}

			{{ render (.Elem "item" "i") }}
		}
	{{ end }}
`
