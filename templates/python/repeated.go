package python

const repTpl = `
	{{ $f := .Field }}{{ $r := .Rules }}

	{{ if $r.GetMinItems }}
		{{ if eq $r.GetMinItems $r.GetMaxItems }}
	if len({{ accessor . }}) != {{ $r.GetMinItems }}:
		{{ err . "value must contain exactly " $r.GetMinItems " item(s)" }}
		{{ else if $r.MaxItems }}
	if len({{ accessor . }}) < {{ $r.GetMinItems }} or len({{ accessor . }}) > {{ $r.GetMaxItems }}:
		{{ err . "value must contain between " $r.GetMinItems " and " $r.GetMaxItems " items, inclusive" }}
		{{ else }}
	if len({{ accessor . }}) < {{ $r.GetMinItems }}:
		{{ err . "value must contain at least " $r.GetMinItems " item(s)" }}
		{{ end }}
	{{ else if $r.MaxItems }}
	if len({{ accessor . }}) > {{ $r.GetMaxItems }}:
		{{ err . "value must contain no more than " $r.GetMaxItems " item(s)" }}
	{{ end }}

	{{ if $r.GetUnique }}
	{{ lookup $f "Unique" }} = set()
	{{ end }}

	{{ if or $r.GetUnique (ne (.Elem "" "").Typ "none") }}
	for idx, item in enumerate({{ accessor . }}):
		pass
		{{ if $r.GetUnique }}
		if item in {{ lookup $f "Unique" }}:
			{{ err . "idx" "repeated value must contain unique items" }}
		else:
			{{ lookup $f "Unique" }}.add(item)
		{{ end }}

	{{ end }}`

