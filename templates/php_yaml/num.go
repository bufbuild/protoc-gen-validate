package php_yaml

const numTpl = `{{ $f := .Field }}{{ $r := .Rules -}}
{{- if $r.Const }}
      - IdenticalTo:
          value: {{ $r.GetConst }}{{ phpTypeLiteralSuffixFor . }}
{{- end -}}
{{- if and (or $r.Lt $r.Lte) (or $r.Gt $r.Gte)}}
      - TODONumLteGteRange: ~
{{- else -}}
{{- if $r.Lt }}
      - LessThan: {{ $r.GetLt }}{{ phpTypeLiteralSuffixFor . }}
{{- end -}}
{{- if $r.Lte }}
      - LessThanOrEqual: {{ $r.GetLte }}{{ phpTypeLiteralSuffixFor . }}
{{- end -}}
{{- if $r.Gt }}
      - GreaterThan: {{ $r.GetGt }}{{ phpTypeLiteralSuffixFor . }}
{{- end -}}
{{- if $r.Gte }}
      - GreaterThanOrEqual: {{ $r.GetGte }}{{ phpTypeLiteralSuffixFor . }}
{{- end -}}
{{- end -}}
{{- template "choice" . -}}
`
