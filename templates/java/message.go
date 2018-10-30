package java

const messageTpl = `{{ $f := .Field }}{{ $r := .Rules }}
	{{- if $r.GetSkip -}}
		// skipping validation for {{ $f.Name }}
	{{ else -}}
		{{- if $r.Required -}}
			com.lyft.pgv.MessageValidation.required("{{ $f.FullyQualifiedName }}", proto.{{ accessor $f }});
		{{- end -}}
	{{- end -}}
`
