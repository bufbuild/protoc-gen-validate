package goshared

const repTpl = `
	{{ $f := .Field }}{{ $r := .Rules }}

	{{ if $r.GetMinItems }}
		{{ if eq $r.GetMinItems $r.GetMaxItems }}
			if len({{ accessor . }}) != {{ $r.GetMinItems }} {
				return {{ err . (t "repeated.items" "value must contain exactly {{$1}} item(s)" $r.GetMinItems) }}
			}
		{{ else if $r.MaxItems }}
			if l := len({{ accessor . }}); l < {{ $r.GetMinItems }} || l > {{ $r.GetMaxItems }} {
			 	return {{ err . (t "repeated.items_between" "value must contain between {{$1}} and {{$2}} items, inclusive" $r.GetMinItems $r.GetMaxItems) }}
			}
		{{ else }}
			if len({{ accessor . }}) < {{ $r.GetMinItems }} {
				return {{ err . (t "repeated.min_items" "value must contain at least {{$1}} item(s)" $r.GetMinItems) }}
			}
		{{ end }}
	{{ else if $r.MaxItems }}
		if len({{ accessor . }}) > {{ $r.GetMaxItems }} {
			return {{ err . (t "repeated.max_items" "value must contain no more than {{$1}} item(s)" $r.GetMaxItems) }}
		}
	{{ end }}

	{{ if $r.GetUnique }}
		{{ lookup $f "Unique" }} := {{ if isBytes $f.Type.Element -}}
			make(map[string]struct{}, len({{ accessor . }}))
		{{ else -}}
			make(map[{{ (typ $f).Element }}]struct{}, len({{ accessor . }}))
		{{ end -}}
	{{ end }}

	{{ if or $r.GetUnique (ne (.Elem "" "").Typ "none") }}
		for idx, item := range {{ accessor . }} {
			_, _ = idx, item
			{{ if $r.GetUnique }}
				if _, exists := {{ lookup $f "Unique" }}[{{ if isBytes $f.Type.Element }}string(item){{ else }}item{{ end }}]; exists {
					return {{ errIdx . "idx" (t "repeated.unique" "repeated value must contain unique items") }}
				} else {
					{{ lookup $f "Unique" }}[{{ if isBytes $f.Type.Element }}string(item){{ else }}item{{ end }}] = struct{}{}
				}
			{{ end }}

			{{ render (.Elem "item" "idx") }}
		}
	{{ end }}
`
