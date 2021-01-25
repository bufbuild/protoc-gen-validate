package php_yaml

const repeatedConstTpl = `{{ renderConstants (.Elem "" "") }}`

const repeatedTpl = `{{ $f := .Field }}{{ $r := .Rules -}}
{{- if or $r.MinItems $r.MaxItems }}
      - Count:
          {{- if $r.MinItems }}
          min: {{ $r.GetMinItems }}
          {{- end -}}
          {{- if $r.MaxItems }}
          max: {{ $r.GetMaxItems }}
          {{- end -}}
{{- end -}}
{{- if $r.GetUnique }}
      - TODOArrayUnique: ~
{{- end -}}
`
