package goshared

const wrapperTpl = `
	{{ $f := .Field }}{{ $r := .Rules }}

	if wrapper := {{ accessor . }}; wrapper != nil {
		{{ render (unwrap . "wrapper") }}
	} {{ if hasrequired $f }} else {
		return {{ err . "value is required and must not be nil." }}
	} {{ end }}
`
