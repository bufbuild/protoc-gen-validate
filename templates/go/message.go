package golang

const messageTpl = `
	{{ $f := .Field }}{{ $r := .Rules }}
	{{ template "required" . }}

	{{ if .MessageRules.GetSkip }}
		// skipping validation for {{ $f.Name }}
	{{ else }}
		if v, ok := interface{}({{ accessor . }}).(interface{ Validate{{ if .AllErrors }}All{{ end }}() error }); ok {
			if err := v.Validate{{ if .AllErrors }}All{{ end }}(); err != nil {
				return {{ errCause . "err" "embedded message failed validation" }}
			}
		}
	{{ end }}
`
