package php_yaml

const requiredTpl = `{{ $f := .Field }}
{{- if .Rules.GetRequired }}
      - NotBlank: ~
{{- end -}}
`
