package goshared

const inTpl = `{{ $f := .Field }}{{ $r := .Rules }}
	{{ if $r.In }}
		if _, ok := {{ lookup $f "InLookup" }}[{{ accessor . }}]; !ok {
			err := {{ err . (t "<prefix>.in" "value must be in list {{$1}}" $r.In) }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ else if $r.NotIn }}
		if _, ok := {{ lookup $f "NotInLookup" }}[{{ accessor . }}]; ok {
			err := {{ err . (t "<prefix>.not_in" "value must not be in list {{$1}}" $r.NotIn) }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ end }}
`
