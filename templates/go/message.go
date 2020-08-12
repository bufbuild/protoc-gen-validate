package golang

const messageTpl = `
	{{ $f := .Field }}{{ $r := .Rules }}
	{{ template "required" . }}

	{{ if .MessageRules.GetSkip }}
		// skipping validation for {{ $f.Name }}
	{{ else }}
		if _, ok := interface{}({{ accessor . }}).(interface{ Validate() error }); ok {
			if err := {{ accessor . }}.Validate(); err != nil {
				return {{ errCause . "err" "embedded message failed validation" }}
			}
		}
	{{ end }}
`
