package goshared

const inTpl = `{{ $f := .Field }}{{ $r := .Rules }}
	{{ if $r.In }}
		if _, ok := {{ lookup $f "InLookup" }}[{{ accessor . }}]; !ok {
			return {{ err . (t "<prefix>.in" value must be in list {{$1}}" $r.In) }}
		}
	{{ else if $r.NotIn }}
		if _, ok := {{ lookup $f "NotInLookup" }}[{{ accessor . }}]; ok {
			return {{ err . (t "<prefix>.not_in" "value must not be in list {{$1}}" $r.NotIn) }}
		}
	{{ end }}
`
