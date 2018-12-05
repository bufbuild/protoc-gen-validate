package java

const messageTpl = `{{ $f := .Field }}{{ $r := .Rules }}
	{{- if $r.GetSkip }}
			// skipping validation for {{ $f.Name }}
	{{- else -}}
		{{- template "required" . }}
		{{- if (isOfMessageType $f) }}
			// Validate {{ $f.Name }}
			if ({{ hasAccessor . }}) com.lyft.pgv.ReflectiveValidatorIndex.validatorFor({{ accessor . }}).assertValid({{ accessor . }});
		{{- end -}}
	{{- end -}}
`
