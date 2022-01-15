package goshared

const anyTpl = `{{ $f := .Field }}{{ $r := .Rules }}
	{{ template "required" . }}

	if a := {{ accessor . }}; a != nil {
		{{ if $r.In }}
			if _, ok := {{ lookup $f "InLookup" }}[a.GetTypeUrl()]; !ok {
				{{ if ne $r.GetErrorMsg "" }}
				err := {{ err . $r.GetErrorMsg }}
				{{ else }}
				err := {{ err . "type URL must be in list " $r.In }}
				{{ end }}
				if !all { return err }
				errors = append(errors, err)
			}
		{{ else if $r.NotIn }}
			if _, ok := {{ lookup $f "NotInLookup" }}[a.GetTypeUrl()]; ok {
				{{ if ne $r.GetErrorMsg "" }}
				err := {{ err . $r.GetErrorMsg }}
				{{ else }}
				err := {{ err . "type URL must not be in list " $r.NotIn }}
				{{ end }}
				if !all { return err }
				errors = append(errors, err)
			}
		{{ end }}
	}
`
