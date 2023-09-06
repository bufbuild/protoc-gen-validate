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
						errors = append(errors, {{ errCause . "err" "embedded message failed validation" }})
					}
					if err := v.ValidateAll(); err != nil {
						if allErr, ok :=  err.(interface { AllErrors() []error }); ok {
							for _, err := range allErr.AllErrors() {
								if verr, ok := err.(interface {
									Field() string
									Reason() string
									Cause() error
								}); ok {
									errors = append(errors, {{ errIdxCauseReason . (printf "%q+verr.Field()" (printf "%s%s" $f.Name ".") ) "verr.Cause()" "verr.Reason()" }})
								} else {
									errors = append(errors, {{ errCause . "err" "fmt.Sprintf(\"embedded message failed validation: %s\", err.Error())" }})
								}
							}
						} else {
							errors = append(errors, {{ errCause . "err" "fmt.Sprintf(\"embedded message failed validation: %s\", err.Error())" }})
						}
					}
				case interface{ Validate() error }:
					{{- /* Support legacy validation for messages that were generated with a plugin version prior to existence of ValidateAll() */ -}}
					if err := v.Validate(); err != nil {
						errors = append(errors, {{ errCause . "err" "embedded message failed validation" }})
					}
			}
		} else if v, ok := interface{}({{ accessor . }}).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return {{ errCause . "err" "embedded message failed validation" }}
			}
		}
	{{ end }}
`
