package php_yaml

const msgTpl = `
{{ if not (ignored .) -}}
{{ qualifiedName . }}:
  properties:
    {{- template "msgInner" . }}
{{- end -}}
`

const msgInnerTpl = `
{{- range .NonOneOfFields -}}
    {{ renderConstants (context .) }}
{{- end -}}
{{- range .OneOfs -}}
    {{- template "oneOfConst" . -}}
{{- end -}}

{{ if disabled . }}
    # Validation is disabled for {{ simpleName . }}
{{- else -}}
{{ range .NonOneOfFields }}
	{{ .Name }}:
	{{- render (context .) }}
{{- end -}}
{{- end }}
`
