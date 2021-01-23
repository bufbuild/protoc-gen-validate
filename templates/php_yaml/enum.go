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
          # message:  .
{{- end -}}
{{- if $r.In }}
      - TODOEnumIn: {{ constantName . "In" }}
{{- end -}}
{{- if $r.NotIn }}
      - TODOEnumIn: {{ constantName . "NotIn" }}
{{- end -}}
`
