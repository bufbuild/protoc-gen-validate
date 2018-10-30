package java

const messageTpl = `{{ $f := .Field }}{{ $r := .Rules }}
	{{- if $r.GetSkip -}}
		// skipping validation for {{ $f.Name }}
	{{ else -}}
		{{- template "required" . -}}
	{{- end -}}
`
