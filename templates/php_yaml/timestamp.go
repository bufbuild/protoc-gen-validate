package php_yaml

const timestampTpl = `{{ $f := .Field }}{{ $r := .Rules -}}
{{- template "required" . -}}

{{- if $r.Const }}
      - TODOTimestampConst: {{ tsLit $r.GetConst }}
{{- end -}}
{{- if and (or $r.Lt $r.Lte) (or $r.Gt $r.Gte)}}
      - TODOTimestampLteGte: ~
{{- else -}}
{{- if $r.Lt }}
      - TODOTimestampLt: {{ tsLit $r.GetLt }}
{{- end -}}
{{- if $r.Lte }}
      - TODOTimestampLte: {{ tsLit $r.Lte }}
{{- end -}}
{{- if $r.Gt }}
      - TODOTimestampGte: {{ tsLit $r.GetGt }}
{{- end -}}
{{- if $r.Gte }}
      - TODOTimestampGte: {{ tsLit $r.GetGte }}
{{- end -}}
{{- end -}}
{{- if $r.LtNow }}
      - TODOTimestampLtNow: ~
{{- end -}}
{{- if $r.GtNow }}
      - TODOTimestampGtNow: ~
{{- end -}}
{{- if $r.Within }}
      - TODOTimestampWithin: {{ durLit $r.GetWithin }}
{{- end -}}
`
