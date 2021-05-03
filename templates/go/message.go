package golang

// Embedded message validation.
const messageTpl = `
	{{ $f := .Field }}{{ $r := .Rules }}
	{{ template "required" . }}

	{{ if .MessageRules.GetSkip }}
		// skipping validation for {{ $f.Name }}
	{{ else }}
		if v, ok := interface{}({{ accessor . }}).(interface{ ValidateWith(bool) error }); ok {
			if err := v.ValidateWith(all); err != nil {
				err = {{ errCause . "err" "embedded message failed validation" }}
				if !all { return err }
				errors = append(errors, err)
			}
		} else if v, ok := interface{}({{ accessor . }}).(interface{ Validate() error }); ok {
			{{- /* Support legacy validation for repos that were generated prior to extended validation using ValidateWith() */ -}}
			if err := v.Validate(); err != nil {
				err = {{ errCause . "err" "embedded message failed validation" }}
				if !all { return err }
				errors = append(errors, err)
			}
		}
	{{ end }}
`
