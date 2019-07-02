package python

const messageTpl = `
	{{ $f := .Field }}{{ $r := .Rules }}
	{{ template "required" . }}
	{{ if $r.GetSkip }}
	# Skipping validation for {{ $f.Name }}
	{{ else }}
	try:
		if not validate_{{ ctype $f.Type }}({{ accessor . }}):
			{{ err . "embedded message failed validation" }}
	except NameError:
		{{ unimplemented }}
	{{ end }}
`

const requiredTpl = `
	{{ if .Rules.GetRequired }}
	if not m.HasField("val"):
	if not {{ hasAcessor . }}:
		{{ err . "value is required" }}
	{{ end }}
`
