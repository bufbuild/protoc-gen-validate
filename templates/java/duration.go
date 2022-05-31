package java

const durationConstTpl = `{{ $f := .Field }}{{ $r := .Rules -}}
{{- if $r.Const }}
		private final com.google.protobuf.Duration {{ constantName . "Const" }} = {{ durLit $r.GetConst }};
{{- end -}}
{{- if $r.Lt }}
		private final com.google.protobuf.Duration {{ constantName . "Lt" }} = {{ durLit $r.GetLt }};
{{- end -}}
{{- if $r.Lte }}
		private final com.google.protobuf.Duration {{ constantName . "Lte" }} = {{ durLit $r.GetLte }};
{{- end -}}
{{- if $r.Gt }}
		private final com.google.protobuf.Duration {{ constantName . "Gt" }} = {{ durLit $r.GetGt }};
{{- end -}}
{{- if $r.Gte }}
		private final com.google.protobuf.Duration {{ constantName . "Gte" }} = {{ durLit $r.GetGte }};
{{- end -}}
{{- if $r.In }}
		private final com.google.protobuf.Duration[] {{ constantName . "In" }} = new com.google.protobuf.Duration[]{
			{{- range $r.In }}
			{{ durLit . }},
			{{- end }}
		};
{{- end -}}
{{- if $r.NotIn }}
		private final com.google.protobuf.Duration[] {{ constantName . "NotIn" }} = new com.google.protobuf.Duration[]{
			{{- range $r.NotIn }}
			{{ durLit . }},
			{{- end }}
		};
{{- end -}}`

const durationTpl = `{{ $f := .Field }}{{ $r := .Rules -}}
{{- template "required" . -}}

{{- if $r.Const }}
	valctx.getValidationCollector().assertValid( ({{ safeName . "value"}}) -> {
		if ({{ hasAccessor . }}) io.envoyproxy.pgv.ConstantValidation.constant("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ constantName . "Const" }});
	},proto);
{{- end -}}
{{- if and (or $r.Lt $r.Lte) (or $r.Gt $r.Gte)}}
	valctx.getValidationCollector().assertValid( ({{ safeName . "value"}}) -> {
		if ({{ hasAccessor . }}) io.envoyproxy.pgv.ComparativeValidation.range("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ if $r.Lt }}{{ constantName . "Lt" }}{{ else }}null{{ end }}, {{ if $r.Lte }}{{ constantName . "Lte" }}{{ else }}null{{ end }}, {{ if $r.Gt }}{{ constantName . "Gt" }}{{ else }}null{{ end }}, {{ if $r.Gte }}{{ constantName . "Gte" }}{{ else }}null{{ end }}, com.google.protobuf.util.Durations.comparator());
	},proto);
{{- else -}}
{{- if $r.Lt }}
	valctx.getValidationCollector().assertValid( ({{ safeName . "value"}}) -> {
		if ({{ hasAccessor . }}) io.envoyproxy.pgv.ComparativeValidation.lessThan("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ constantName . "Lt" }}, com.google.protobuf.util.Durations.comparator());
	},proto);
{{- end -}}
{{- if $r.Lte }}
	valctx.getValidationCollector().assertValid( ({{ safeName . "value"}}) -> {
		if ({{ hasAccessor . }}) io.envoyproxy.pgv.ComparativeValidation.lessThanOrEqual("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ constantName . "Lte" }}, com.google.protobuf.util.Durations.comparator());
	},proto);
{{- end -}}
{{- if $r.Gt }}
	valctx.getValidationCollector().assertValid( ({{ safeName . "value"}}) -> {
		if ({{ hasAccessor . }}) io.envoyproxy.pgv.ComparativeValidation.greaterThan("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ constantName . "Gt" }}, com.google.protobuf.util.Durations.comparator());
	},proto);
{{- end -}}
{{- if $r.Gte }}
	valctx.getValidationCollector().assertValid( ({{ safeName . "value"}}) -> {
		if ({{ hasAccessor . }}) io.envoyproxy.pgv.ComparativeValidation.greaterThanOrEqual("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ constantName . "Gte" }}, com.google.protobuf.util.Durations.comparator());
	},proto);
{{- end -}}
{{- end -}}
{{- if $r.In }}
	valctx.getValidationCollector().assertValid( ({{ safeName . "value"}}) -> {
		if ({{ hasAccessor . }}) io.envoyproxy.pgv.CollectiveValidation.in("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ constantName . "In" }});
	},proto);
{{- end -}}
{{- if $r.NotIn }}
	valctx.getValidationCollector().assertValid( ({{ safeName . "value"}}) -> {
		if ({{ hasAccessor . }}) io.envoyproxy.pgv.CollectiveValidation.notIn("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ constantName . "NotIn" }});
	},proto);
{{- end -}}
`
