package python

const messageTpl = `
	{{ $f := .Field }}{{ $r := .Rules }}
	{{ template "required" . }}

	{{ if $r.GetSkip }}
		# skipping validation for {{ $f.Name }}
	{{ else }}
	try:
		if {{ accessor . }}.validate():
			{{ err . "err" "embedded message failed validation" }}
	except AttributeError:
		{{ err . }}

	{{ end }}
`

const requiredTpl = `
	{{ if .Rules.GetRequired }}
	if {{ accessor . }} is None:
		{{ err . "value is required" }}
	{{ end }}
`
