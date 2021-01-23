package php_yaml

const anyTpl = `{{ $f := .Field }}{{ $r := .Rules }}
{{- template "required" . -}}
{{- if $r.In }}
      - Choice:
          choices:
            {{- range $r.In -}}- {{- sprintf "%v" . -}}{{- end -}}
		  # message:  .
{{- end -}}
`
