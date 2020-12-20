package golang

const messageTpl = `
	{{ $f := .Field }}{{ $r := .Rules }}
	{{ template "required" . }}

	{{ if .MessageRules.GetSkip }}
		// skipping validation for {{ $f.Name }}
	{{ else }}
		if stopOnError {
			if v, ok := interface{}({{ accessor . }}).(interface{ Validate() error }); ok {
				if err := v.Validate(); err != nil {
					return {{ errCause . "err" "embedded message failed validation" }}
				}
			}
		} else if v, ok := interface{}({{ accessor . }}).(interface{ AllErrors() error }); ok {
			if err := v.AllErrors(); err != nil {
				err = {{ errCause . "err" "embedded message failed validation" }}
				multiErr = multierror.Append(multiErr, err)
			}	
		}
	{{ end }}
`
