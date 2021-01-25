package php_yaml

const choiceTpl = `{{ $f := .Field }}{{ $r := .Rules -}}
{{- if $r.In }}
      - Choice:
          choices:
            {{- range $f.In }}
            - {{ sprintf "%v" . }}{{ phpTypeLiteralSuffixFor $ }}
            {{- end }}
{{- end -}}
{{- if $r.NotIn }}
      - NotInChoice:
          choices:
            {{- range $f.NotIn }}
            - {{ sprintf "%v" . }}{{ phpTypeLiteralSuffixFor $ }}
            {{- end }}
{{- end -}}
`
