package php_yaml

const enumTpl = `{{ $f := .Field }}{{ $r := .Rules -}}
{{- if $r.Const }}
      - IdenticalTo:
          value: {{ $r.Const }}
{{- end -}}
{{- if $r.GetDefinedOnly }}
      - Choice: # Enum.DefinedOnly
          choices:
            {{- range $f.Type.Enum.Values }}
            - {{ sprintf "%v" .Value }} # {{ .Name }}
            {{- end }}
{{- end -}}
{{- if $r.In }}
      - Choice: # Enum.In
          choices:
            {{- range $r.In }}
            - {{ sprintf "%v" . }}
            {{- end }}
{{- end -}}
{{- if $r.NotIn }}
      - NotInChoice: # Enum.NotIn
          choices:
            {{- range $r.NotIn }}
            - {{ sprintf "%v" . }}
            {{- end }}
{{- end -}}
`
