package php_yaml

const durationTpl = `{{ $f := .Field }}{{ $r := .Rules -}}
{{- template "required" . -}}

{{- if $r.Const }}
      - TODODurationConst: {{ tsLit $r.GetConst }}
{{- end -}}
{{- if and (or $r.Lt $r.Lte) (or $r.Gt $r.Gte)}}
      - TODODurationLteGte: ~
{{- else -}}
{{- if $r.Lt }}
      - TODODurationLt: {{ durLit $r.GetLt }}
{{- end -}}
{{- if $r.Lte }}
      - TODODurationLte: {{ durLit $r.Lte }}
{{- end -}}
{{- if $r.Gt }}
      - TODODurationLte: {{ durLit $r.GetGt }}
{{- end -}}
{{- if $r.Gte }}
      - TODODurationGte: {{ durLit $r.GetGte }}
{{- end -}}
{{- end -}}
`
