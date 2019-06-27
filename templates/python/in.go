package python

const inTpl = `{{ $f := .Field -}}{{ $r := .Rules -}}
	{{- if $r.In }}
	if {{ accessor . }} not in {{ lookup .Field "InLookup" }}:
		{{ err . "value must be in list " $r.In }}
	{{- else if $r.NotIn }}
	if {{ accessor . }} in {{ lookup .Field "NotInLookup" }}:
		{{ err . "value must be in list " $r.NotIn }}
	{{- end }}
`
