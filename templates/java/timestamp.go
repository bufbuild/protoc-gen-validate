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
			if ({{ hasAccessor . }}) com.lyft.pgv.ConstantValidation.constant("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ constantName $f "Const" }});
{{- end -}}
{{- if and (or $r.Lt $r.Lte) (or $r.Gt $r.Gte)}}
			if ({{ hasAccessor . }}) com.lyft.pgv.ComparativeValidation.range("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ if $r.Lt }}{{ constantName $f "Lt" }}{{ else }}null{{ end }}, {{ if $r.Lte }}{{ constantName $f "Lte" }}{{ else }}null{{ end }}, {{ if $r.Gt }}{{ constantName $f "Gt" }}{{ else }}null{{ end }}, {{ if $r.Gte }}{{ constantName $f "Gte" }}{{ else }}null{{ end }}, com.google.protobuf.util.Timestamps.comparator());
{{- else -}}
{{- if $r.Lt }}
			if ({{ hasAccessor . }}) com.lyft.pgv.ComparativeValidation.lessThan("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ constantName $f "Lt" }}, com.google.protobuf.util.Timestamps.comparator());
{{- end -}}
{{- if $r.Lte }}
			if ({{ hasAccessor . }}) com.lyft.pgv.ComparativeValidation.lessThanOrEqual("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ constantName $f "Lte" }}, com.google.protobuf.util.Timestamps.comparator());
{{- end -}}
{{- if $r.Gt }}
			if ({{ hasAccessor . }}) com.lyft.pgv.ComparativeValidation.greaterThan("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ constantName $f "Gt" }}, com.google.protobuf.util.Timestamps.comparator());
{{- end -}}
{{- if $r.Gte }}
			if ({{ hasAccessor . }}) com.lyft.pgv.ComparativeValidation.greaterThanOrEqual("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ constantName $f "Gte" }}, com.google.protobuf.util.Timestamps.comparator());
{{- end -}}
{{- end -}}
{{- if $r.LtNow }}
			if ({{ hasAccessor . }}) com.lyft.pgv.ComparativeValidation.lessThan("{{ $f.FullyQualifiedName }}", {{ accessor . }}, com.lyft.pgv.TimestampValidation.currentTimestamp(), com.google.protobuf.util.Timestamps.comparator());
{{- end -}}
{{- if $r.GtNow }}
			if ({{ hasAccessor . }}) com.lyft.pgv.ComparativeValidation.greaterThan("{{ $f.FullyQualifiedName }}", {{ accessor . }}, com.lyft.pgv.TimestampValidation.currentTimestamp(), com.google.protobuf.util.Timestamps.comparator());
{{- end -}}
{{- if $r.Within }}
			if ({{ hasAccessor . }}) com.lyft.pgv.TimestampValidation.within("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ constantName $f "Within" }}, com.lyft.pgv.TimestampValidation.currentTimestamp());
{{- end -}}
`
