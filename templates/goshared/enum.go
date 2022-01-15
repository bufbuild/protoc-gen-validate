package goshared

const enumTpl = `
		{{ $f := .Field }}{{ $r := .Rules }}
		{{ template "const" . }}
		{{ template "in" . }}

		{{ if $r.GetDefinedOnly }}
			if _, ok := {{ (typ $f).Element }}_name[int32({{ accessor . }})]; !ok {
				{{ if ne $r.GetErrorMsg "" }}
				err := {{ err . $r.GetErrorMsg }}
				{{ else }}
				err := {{ err . "value must be one of the defined enum values" }}
				{{ end }}
				if !all { return err }
				errors = append(errors, err)
			}
		{{ end }}
`
