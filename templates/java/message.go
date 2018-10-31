package java

const messageTpl = `{{ $f := .Field }}{{ $r := .Rules }}
	{{- if $r.GetSkip }}
			// skipping validation for {{ $f.Name }}
	{{- else -}}
		{{- template "required" . }}
		{{- if (isOfMessageType $f) }}
			// Validate {{ $f.Name }}
			com.lyft.pgv.ValidatorIndex.validatorFor(proto.{{ accessor . }}).assertValid(proto.{{ accessor . }});
		{{- end -}}
	{{- end -}}
`
