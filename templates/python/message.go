package python

const messageTpl = `
	{{ $f := .Field }}{{ $r := .Rules }}
	{{ template "required" . }}
	
	{{ if $r.GetSkip }}
	# Skipping validation for {{ $f.Name }}
	{{ else }}

	{{ if shouldImport $f }}
	{{ ctype $f.Type false }}
	{{ end }}
	
	{{ if shouldValidate $f }}
	if {{ hasAccessor . }} and not validate_{{ ctype $f.Type true }}({{ accessor . }})[0]:
		{{ err . "embedded message failed validation" }}
	{{ end }}
	{{ end }}
`

const requiredTpl = `
	{{ if .Rules.GetRequired }}
	if not {{ hasAccessor . }}:
		{{ err . "value is required" }}
	{{ end }}
`

