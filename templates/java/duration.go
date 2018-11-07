package java

const durationConstTpl = `{{ $f := .Field }}{{ $r := .Rules -}}
{{- if $r.Const }}
		private final com.google.protobuf.Duration {{ constantName $f "Const" }} = {{ durLit $r.GetConst }};
{{- end -}}
{{- if $r.Lt }}
		private final com.google.protobuf.Duration {{ constantName $f "Lt" }} = {{ durLit $r.GetLt }};
{{- end -}}
{{- if $r.Lte }}
		private final com.google.protobuf.Duration {{ constantName $f "Lte" }} = {{ durLit $r.GetLte }};
{{- end -}}
{{- if $r.Gt }}
		private final com.google.protobuf.Duration {{ constantName $f "Gt" }} = {{ durLit $r.GetGt }};
{{- end -}}
{{- if $r.Gte }}
		private final com.google.protobuf.Duration {{ constantName $f "Gte" }} = {{ durLit $r.GetGte }};
{{- end -}}
{{- if $r.In }}
		private final com.google.protobuf.Duration[] {{ constantName $f "In" }} = new com.google.protobuf.Duration[]{
			{{- range $r.In }}
			{{ durLit . }},
			{{- end }}
		};
{{- end -}}
{{- if $r.NotIn }}
		private final com.google.protobuf.Duration[] {{ constantName $f "NotIn" }} = new com.google.protobuf.Duration[]{
			{{- range $r.NotIn }}
			{{ durLit . }},
			{{- end }}
		};
{{- end -}}`

const durationTpl = `{{ $f := .Field }}{{ $r := .Rules -}}
{{- template "required" . -}}

{{- if $r.Const }}
			if ({{ hasAccessor . }}) com.lyft.pgv.ConstantValidation.constant("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ constantName $f "Const" }});
{{- end -}}
{{- if and (or $r.Lt $r.Lte) (or $r.Gt $r.Gte)}}
			if ({{ hasAccessor . }}) com.lyft.pgv.ComparativeValidation.range("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ if $r.Lt }}{{ constantName $f "Lt" }}{{ else }}null{{ end }}, {{ if $r.Lte }}{{ constantName $f "Lte" }}{{ else }}null{{ end }}, {{ if $r.Gt }}{{ constantName $f "Gt" }}{{ else }}null{{ end }}, {{ if $r.Gte }}{{ constantName $f "Gte" }}{{ else }}null{{ end }}, com.google.protobuf.util.Durations.comparator());
{{- else -}}
{{- if $r.Lt }}
			if ({{ hasAccessor . }}) com.lyft.pgv.ComparativeValidation.lessThan("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ constantName $f "Lt" }}, com.google.protobuf.util.Durations.comparator());
{{- end -}}
{{- if $r.Lte }}
			if ({{ hasAccessor . }}) com.lyft.pgv.ComparativeValidation.lessThanOrEqual("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ constantName $f "Lte" }}, com.google.protobuf.util.Durations.comparator());
{{- end -}}
{{- if $r.Gt }}
			if ({{ hasAccessor . }}) com.lyft.pgv.ComparativeValidation.greaterThan("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ constantName $f "Gt" }}, com.google.protobuf.util.Durations.comparator());
{{- end -}}
{{- if $r.Gte }}
			if ({{ hasAccessor . }}) com.lyft.pgv.ComparativeValidation.greaterThanOrEqual("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ constantName $f "Gte" }}, com.google.protobuf.util.Durations.comparator());
{{- end -}}
{{- end -}}
{{- if $r.In }}
			if ({{ hasAccessor . }}) com.lyft.pgv.CollectiveValidation.in("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ constantName $f "In" }});
{{- end -}}
{{- if $r.NotIn }}
			if ({{ hasAccessor . }}) com.lyft.pgv.CollectiveValidation.notIn("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ constantName $f "NotIn" }});
{{- end -}}
`
