package java

const durationTpl = `{{ $f := .Field }}{{ $r := .Rules -}}
{{ template "required" . }}

{{- if $r.Const }}
			com.lyft.pgv.DurationValidation.constant("{{ $f.FullyQualifiedName }}", proto.{{ accessor $f }}, {{ durLit $r.GetConst }});
{{- end -}}
{{- if $r.Lt }}
			com.lyft.pgv.DurationValidation.lessThan("{{ $f.FullyQualifiedName }}", proto.{{ accessor $f }}, {{ durLit $r.Lt }});
{{- end -}}
{{- if $r.Lte }}
			com.lyft.pgv.DurationValidation.lessThanOrEqual("{{ $f.FullyQualifiedName }}", proto.{{ accessor $f }}, {{ durLit $r.Lte }});
{{- end -}}
{{- if $r.Gt }}
			com.lyft.pgv.DurationValidation.greaterThan("{{ $f.FullyQualifiedName }}", proto.{{ accessor $f }}, {{ durLit $r.Gt }});
{{- end -}}
{{- if $r.Gte }}
			com.lyft.pgv.DurationValidation.greaterThanOrEqual("{{ $f.FullyQualifiedName }}", proto.{{ accessor $f }}, {{ durLit $r.Gte }});
{{- end -}}
`
