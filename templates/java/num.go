package java

const numConstTpl = `{{ $f := .Field }}{{ $r := .Rules -}}
{{- if $r.Const }}
		private final {{ javaTypeFor .}} {{ constantName . "Const" }} = {{ $r.GetConst }}{{ javaTypeLiteralSuffixFor . }};
{{- end -}}
{{- if $r.Lt }}
		private final {{ javaTypeFor .}} {{ constantName . "Lt" }} = {{ $r.GetLt }}{{ javaTypeLiteralSuffixFor . }};
{{- end -}}
{{- if $r.Lte }}
		private final {{ javaTypeFor .}} {{ constantName . "Lte" }} = {{ $r.GetLte }}{{ javaTypeLiteralSuffixFor . }};
{{- end -}}
{{- if $r.Gt }}
		private final {{ javaTypeFor .}} {{ constantName . "Gt" }} = {{ $r.GetGt }}{{ javaTypeLiteralSuffixFor . }};
{{- end -}}
{{- if $r.Gte }}
		private final {{ javaTypeFor .}} {{ constantName . "Gte" }} = {{ $r.GetGte }}{{ javaTypeLiteralSuffixFor . }};
{{- end -}}
{{- if $r.In }}
		private final {{ javaTypeFor . }}[] {{ constantName . "In" }} = new {{ javaTypeFor . }}[]{
			{{- range $r.In -}}
				{{- sprintf "%v" . -}}{{ javaTypeLiteralSuffixFor $ }},
			{{- end -}}
		};
{{- end -}}
{{- if $r.NotIn }}
		private final {{ javaTypeFor . }}[] {{ constantName . "NotIn" }} = new {{ javaTypeFor . }}[]{
			{{- range $r.NotIn -}}
				{{- sprintf "%v" . -}}{{ javaTypeLiteralSuffixFor $ }},
			{{- end -}}
		};
{{- end -}}`

const numTpl = `{{ $f := .Field }}{{ $r := .Rules -}}
{{- if $r.GetIgnoreEmpty }}
			if ( {{ accessor . }} != 0 ) {
{{- end -}}
{{- if $r.Const }}
			valctx.getValidationCollector().assertValid( ({{ safeName . "value"}}) -> {
				io.envoyproxy.pgv.ConstantValidation.constant("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ constantName . "Const" }});
			},proto);
{{- end -}}
{{- if and (or $r.Lt $r.Lte) (or $r.Gt $r.Gte)}}
			valctx.getValidationCollector().assertValid( ({{ safeName . "value"}}) -> {
				io.envoyproxy.pgv.ComparativeValidation.range("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ if $r.Lt }}{{ constantName . "Lt" }}{{ else }}null{{ end }}, {{ if $r.Lte }}{{ constantName . "Lte" }}{{ else }}null{{ end }}, {{ if $r.Gt }}{{ constantName . "Gt" }}{{ else }}null{{ end }}, {{ if $r.Gte }}{{ constantName . "Gte" }}{{ else }}null{{ end }}, java.util.Comparator.naturalOrder());
			},proto);
{{- else -}}
{{- if $r.Lt }}
			valctx.getValidationCollector().assertValid( ({{ safeName . "value"}}) -> {
				io.envoyproxy.pgv.ComparativeValidation.lessThan("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ constantName . "Lt" }}, java.util.Comparator.naturalOrder());
			},proto);
{{- end -}}
{{- if $r.Lte }}
			valctx.getValidationCollector().assertValid( ({{ safeName . "value"}}) -> {
				io.envoyproxy.pgv.ComparativeValidation.lessThanOrEqual("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ constantName . "Lte" }}, java.util.Comparator.naturalOrder());
			},proto);
{{- end -}}
{{- if $r.Gt }}
			valctx.getValidationCollector().assertValid( ({{ safeName . "value"}}) -> {
				io.envoyproxy.pgv.ComparativeValidation.greaterThan("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ constantName . "Gt" }}, java.util.Comparator.naturalOrder());
			},proto);
{{- end -}}
{{- if $r.Gte }}
			valctx.getValidationCollector().assertValid( ({{ safeName . "value"}}) -> {
				io.envoyproxy.pgv.ComparativeValidation.greaterThanOrEqual("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ constantName . "Gte" }}, java.util.Comparator.naturalOrder());
			},proto);
{{- end -}}
{{- end -}}
{{- if $r.In }}
			valctx.getValidationCollector().assertValid( ({{ safeName . "value"}}) -> {
				io.envoyproxy.pgv.CollectiveValidation.in("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ constantName . "In" }});
			},proto);
{{- end -}}
{{- if $r.NotIn }}
			valctx.getValidationCollector().assertValid( ({{ safeName . "value"}}) -> {
				io.envoyproxy.pgv.CollectiveValidation.notIn("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ constantName . "NotIn" }});
			},proto);
{{- end -}}
{{- if $r.GetIgnoreEmpty }}
			}
{{- end -}}
`
