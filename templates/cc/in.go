package cc

const inTpl = `{{ $f := .Field -}}{{ $r := .Rules -}}
	{{- if $r.In }}
	{{- if eq .AccessorOverride "key" }}
	    if ({{ lookup $f "Key_InLookup" }}.find({{ accessor . }}) == {{ lookup $f "Key_InLookup" }}.end()) {
			{{ err . "map key must be in list " $r.In }}
		}
	{{- else if eq .AccessorOverride "val" }}
	    if ({{ lookup $f "Val_InLookup" }}.find({{ accessor . }}) == {{ lookup $f "Val_InLookup" }}.end()) {
			{{ err . "map value must be in list " $r.In }}
		}
	{{- else }}
		if ({{ lookup $f "InLookup" }}.find({{ accessor . }}) == {{ lookup $f "InLookup" }}.end()) {
			{{ err . "value must be in list " $r.In }}
		}
	{{- end}}
	{{- else if $r.NotIn }}
	{{- if eq .AccessorOverride "key" }}
		if ({{ lookup $f "Key_NotInLookup" }}.find({{ accessor . }}) != {{ lookup $f "Key_NotInLookup" }}.end()) {
			{{ err . "map key must not be in list " $r.NotIn }}
		}
	{{- else if eq .AccessorOverride "val" }}
		if ({{ lookup $f "Val_NotInLookup" }}.find({{ accessor . }}) != {{ lookup $f "Val_NotInLookup" }}.end()) {
			{{ err . "map value must not be in list " $r.NotIn }}
		}
	{{- else }}
		if ({{ lookup $f "NotInLookup" }}.find({{ accessor . }}) != {{ lookup $f "NotInLookup" }}.end()) {
			{{ err . "value must not be in list " $r.NotIn }}
		}
	{{- end }}
	{{- end }}
`
