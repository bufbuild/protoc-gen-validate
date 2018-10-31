package java

const timestampTpl = `{{ $f := .Field }}{{ $r := .Rules -}}
{{- template "required" . -}}

{{- if $r.Const }}
			com.lyft.pgv.TimestampValidation.constant("{{ $f.FullyQualifiedName }}", proto.{{ accessor . }}, {{ tsLit $r.GetConst }});
{{- end -}}
{{- if $r.Lt }}
			com.lyft.pgv.TimestampValidation.lessThan("{{ $f.FullyQualifiedName }}", proto.{{ accessor . }}, {{ tsLit $r.Lt }});
{{- end -}}
{{- if $r.Lte }}
			com.lyft.pgv.TimestampValidation.lessThanOrEqual("{{ $f.FullyQualifiedName }}", proto.{{ accessor . }}, {{ tsLit $r.Lte }});
{{- end -}}
{{- if $r.Gt }}
			com.lyft.pgv.TimestampValidation.greaterThan("{{ $f.FullyQualifiedName }}", proto.{{ accessor . }}, {{ tsLit $r.Gt }});
{{- end -}}
{{- if $r.Gte }}
			com.lyft.pgv.TimestampValidation.greaterThanOrEqual("{{ $f.FullyQualifiedName }}", proto.{{ accessor . }}, {{ tsLit $r.Gte }});
{{- end -}}
{{- if $r.LtNow }}
			com.lyft.pgv.TimestampValidation.lessThan("{{ $f.FullyQualifiedName }}", proto.{{ accessor . }}, com.lyft.pgv.TimestampValidation.currentTimestamp());
{{- end -}}
{{- if $r.GtNow }}
			com.lyft.pgv.TimestampValidation.greaterThan("{{ $f.FullyQualifiedName }}", proto.{{ accessor . }}, com.lyft.pgv.TimestampValidation.currentTimestamp());
{{- end -}}
{{- if $r.Within }}
			com.lyft.pgv.TimestampValidation.within("{{ $f.FullyQualifiedName }}", proto.{{ accessor . }}, {{ durLit $r.Within }}, com.lyft.pgv.TimestampValidation.currentTimestamp());
{{- end -}}
`
