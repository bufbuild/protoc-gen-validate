package goshared

const inTpl = `{{ $f := .Field }}{{ $r := .Rules }}
	{{ if $r.In }}
		if _, ok := {{ lookup $f "InLookup" }}[{{ accessor . }}]; !ok {
			err := {{ err . "value must be in list " (inList $f $r.In) }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ else if $r.NotIn }}
		if _, ok := {{ lookup $f "NotInLookup" }}[{{ accessor . }}]; ok {
			err := {{ err . "value must not be in list " (inList $f $r.NotIn) }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ end }}
`
