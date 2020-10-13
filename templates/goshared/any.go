package goshared

const anyTpl = `{{ $f := .Field }}{{ $r := .Rules }}
	{{ template "required" . }}

	if a := {{ accessor . }}; a != nil {
		{{ if $r.In }}
			if _, ok := {{ lookup $f "InLookup" }}[a.GetTypeUrl()]; !ok {
				return {{ err . (t "any.in" "type URL must be in list {{$1}}" $r.In) }}
			}
		{{ else if $r.NotIn }}
			if _, ok := {{ lookup $f "NotInLookup" }}[a.GetTypeUrl()]; ok {
				return {{ err . (t "any.not_in" "type URL must not be in list {{$1}}" $r.NotIn) }}
			}
		{{ end }}
	}
`
