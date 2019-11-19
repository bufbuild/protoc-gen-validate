package goshared

const inTpl = `{{ $f := .Field }}{{ $r := .Rules }}
	{{ if $r.In }}
	    {{ if eq .AccessorOverride "key" }}
		if _, ok := {{ lookup $f "Key_InLookup" }}[{{ accessor . }}]; !ok {
			return {{ err . "map key must be in list " $r.In }}
		}
		{{ else if eq .AccessorOverride "val" }}
		if _, ok := {{ lookup $f "Val_InLookup" }}[{{ accessor . }}]; !ok {
			return {{ err . "map value must be in list " $r.In }}
		}
		{{ else }}
		if _, ok := {{ lookup $f "InLookup" }}[{{ accessor . }}]; !ok {
			return {{ err . "value must be in list " $r.In }}
		}
		{{ end }}
	{{ else if $r.NotIn }}
	    {{ if eq .AccessorOverride "key" }}
		if _, ok := {{ lookup $f "Key_NotInLookup" }}[{{ accessor . }}]; ok {
			return {{ err . "map key must not be in list " $r.NotIn }}
		}
		{{ else if eq .AccessorOverride "val" }}
		if _, ok := {{ lookup $f "Val_NotInLookup" }}[{{ accessor . }}]; ok {
			return {{ err . "map value must not be in list " $r.NotIn }}
		}
		{{ else }}
		if _, ok := {{ lookup $f "NotInLookup" }}[{{ accessor . }}]; ok {
			return {{ err . "value must not be in list " $r.NotIn }}
		}
		{{ end }}
	{{ end }}
`
