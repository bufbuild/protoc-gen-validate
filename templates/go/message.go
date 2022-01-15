package golang

// Embedded message validation.
const messageTpl = `
	{{ $f := .Field }}{{ $r := .Rules }}
	{{ template "required" . }}

	{{ if .MessageRules.GetSkip }}
		// skipping validation for {{ $f.Name }}
	{{ else }}
		if all {
			switch v := interface{}({{ accessor . }}).(type) {
				case interface{ ValidateAll() error }:
					if err := v.ValidateAll(); err != nil {
						{{ if ne $r.GetErrorMsg "" }}
						errors = append(errors, {{ errCause . $r.GetErrorMsg }})
						{{ else }}
						errors = append(errors, {{ errCause . "err" "embedded message failed validation" }})
						{{ end }}
					}
				case interface{ Validate() error }:
					{{- /* Support legacy validation for messages that were generated with a plugin version prior to existence of ValidateAll() */ -}}
					if err := v.Validate(); err != nil {
						{{ if ne $r.GetErrorMsg "" }}
						errors = append(errors, {{ errCause . $r.GetErrorMsg }})
						{{ else }}
						errors = append(errors, {{ errCause . "err" "embedded message failed validation" }})
						{{ end }}
					}
			}
		} else if v, ok := interface{}({{ accessor . }}).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				{{ if ne $r.GetErrorMsg "" }}
				errors = append(errors, {{ errCause . $r.GetErrorMsg }})
				{{ else }}
				errors = append(errors, {{ errCause . "err" "embedded message failed validation" }})
				{{ end }}
			}
		}
	{{ end }}
`
