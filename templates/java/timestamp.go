package java

const timestampConstTpl = `{{ $f := .Field }}{{ $r := .Rules -}}
{{- if $r.Const }}
		private final com.google.protobuf.Timestamp {{ constantName $f "Const" }} = {{ tsLit $r.GetConst }};
{{- end -}}
{{- if $r.Lt }}
		private final com.google.protobuf.Timestamp {{ constantName $f "Lt" }} = {{ tsLit $r.GetLt }};
{{- end -}}
{{- if $r.Lte }}
		private final com.google.protobuf.Timestamp {{ constantName $f "Lte" }} = {{ tsLit $r.Lte }};
{{- end -}}
{{- if $r.Gt }}
		private final com.google.protobuf.Timestamp {{ constantName $f "Gt" }} = {{ tsLit $r.GetGt }};
{{- end -}}
{{- if $r.Gte }}
		private final com.google.protobuf.Timestamp {{ constantName $f "Gte" }} = {{ tsLit $r.GetGte }};
{{- end -}}
{{- if $r.Within }}
		private final com.google.protobuf.Duration {{ constantName $f "Within" }} = {{ durLit $r.GetWithin }};
{{- end -}}`

const timestampTpl = `{{ $f := .Field }}{{ $r := .Rules -}}
{{- template "required" . -}}

{{- if $r.Const }}
			com.lyft.pgv.ConstantValidation.constant("{{ $f.FullyQualifiedName }}", proto.{{ accessor . }}, {{ constantName $f "Const" }});
{{- end -}}
{{- if $r.Lt }}
			com.lyft.pgv.ComparativeValidation.lessThan("{{ $f.FullyQualifiedName }}", proto.{{ accessor . }}, {{ constantName $f "Lt" }}, com.google.protobuf.util.Timestamps.comparator());
{{- end -}}
{{- if $r.Lte }}
			com.lyft.pgv.ComparativeValidation.lessThanOrEqual("{{ $f.FullyQualifiedName }}", proto.{{ accessor . }}, {{ constantName $f "Lte" }}, com.google.protobuf.util.Timestamps.comparator());
{{- end -}}
{{- if $r.Gt }}
			com.lyft.pgv.ComparativeValidation.greaterThan("{{ $f.FullyQualifiedName }}", proto.{{ accessor . }}, {{ constantName $f "Gt" }}, com.google.protobuf.util.Timestamps.comparator());
{{- end -}}
{{- if $r.Gte }}
			com.lyft.pgv.ComparativeValidation.greaterThanOrEqual("{{ $f.FullyQualifiedName }}", proto.{{ accessor . }}, {{ constantName $f "Gte" }}, com.google.protobuf.util.Timestamps.comparator());
{{- end -}}
{{- if $r.LtNow }}
			com.lyft.pgv.ComparativeValidation.lessThan("{{ $f.FullyQualifiedName }}", proto.{{ accessor . }}, com.lyft.pgv.TimestampValidation.currentTimestamp(), com.google.protobuf.util.Timestamps.comparator());
{{- end -}}
{{- if $r.GtNow }}
			com.lyft.pgv.ComparativeValidation.greaterThan("{{ $f.FullyQualifiedName }}", proto.{{ accessor . }}, com.lyft.pgv.TimestampValidation.currentTimestamp(), com.google.protobuf.util.Timestamps.comparator());
{{- end -}}
{{- if $r.Within }}
			com.lyft.pgv.TimestampValidation.within("{{ $f.FullyQualifiedName }}", proto.{{ accessor . }}, {{ constantName $f "Within" }}, com.lyft.pgv.TimestampValidation.currentTimestamp());
{{- end -}}
`
