package goshared

const inTpl = `{{ $f := .Field }}{{ $r := .Rules }}
	{{ if $r.In }}
		if _, ok := {{ lookup $f "InLookup" }}[{{ accessor . }}]; m.maskHas(mask, "{{ $f.Name }}") && !ok {
			return {{ err . "value must be in list " $r.In }}
		}
	{{ else if $r.NotIn }}
		if _, ok := {{ lookup $f "NotInLookup" }}[{{ accessor . }}]; m.maskHas(mask, "{{ $f.Name }}") && ok {
			return {{ err . "value must not be in list " $r.NotIn }}
		}
	{{ end }}
`
