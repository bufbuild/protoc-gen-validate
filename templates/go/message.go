package golang

const messageTpl = `
	{{ $f := .Field }}{{ $r := .Rules }}
	{{ template "required" . }}

	{{ if .MessageRules.GetSkip }}
		// skipping validation for {{ $f.Name }}
	{{ else }}
		if v, ok := interface{}({{ accessor . }}).(interface{ Validate(bool) error }); ok {
			if err := v.Validate(all); err != nil {
				err = {{ errCause . "err" "embedded message failed validation" }}
				if !all { return err }
				errors = append(errors, err)
			}
		}
	{{ end }}
`
