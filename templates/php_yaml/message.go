package php_yaml

const messageTpl = `{{ $f := .Field }}{{ $r := .Rules }}
{{- if .MessageRules.GetSkip }}
    # Skipping validation for {{ $f.Name }}
{{- else }}
      {{- if (isOfMessageType $f) }}
      # Validate {{ $f.Name }}
      {{- end -}}
      {{- template "required" . }}
{{- end -}}
`
