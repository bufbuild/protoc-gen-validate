package tpl

const anyTpl = `{{ $f := .Field }}{{ $r := .Rules }}
	{{ template "required" . }}

	{{ if $r.In }}
		if _, ok := {{ lookup $f "InLookup" }}[{{ accessor . }}.GetTypeUrl()]; !ok {
			return {{ err . "type URL must be in list " $r.In }}
		}
	{{ else if $r.NotIn }}
		if _, ok := {{ lookup $f "NotInLookup" }}[{{ accessor . }}.GetTypeUrl()]; ok {
			return {{ err . "type URL must not be in list " $r.NotIn }}
		}
	{{ end }}
`
