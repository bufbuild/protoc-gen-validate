package php

const numConstTpl = `{{ $f := .Field }}{{ $r := .Rules -}}
{{- if $r.Const }}
		private final {{ phpTypeFor .}} {{ constantName . "Const" }} = {{ $r.GetConst }}{{ phpTypeLiteralSuffixFor . }};
{{- end -}}
{{- if $r.Lt }}
		private final {{ phpTypeFor .}} {{ constantName . "Lt" }} = {{ $r.GetLt }}{{ phpTypeLiteralSuffixFor . }};
{{- end -}}
{{- if $r.Lte }}
		private final {{ phpTypeFor .}} {{ constantName . "Lte" }} = {{ $r.GetLte }}{{ phpTypeLiteralSuffixFor . }};
{{- end -}}
{{- if $r.Gt }}
		private final {{ phpTypeFor .}} {{ constantName . "Gt" }} = {{ $r.GetGt }}{{ phpTypeLiteralSuffixFor . }};
{{- end -}}
{{- if $r.Gte }}
		private final {{ phpTypeFor .}} {{ constantName . "Gte" }} = {{ $r.GetGte }}{{ phpTypeLiteralSuffixFor . }};
{{- end -}}
{{- if $r.In }}
		private final {{ phpTypeFor . }}[] {{ constantName . "In" }} = new {{ phpTypeFor . }}[]{
			{{- range $r.In -}}
				{{- sprintf "%v" . -}}{{ phpTypeLiteralSuffixFor $ }},
			{{- end -}}
		};
{{- end -}}
{{- if $r.NotIn }}
		private final {{ phpTypeFor . }}[] {{ constantName . "NotIn" }} = new {{ phpTypeFor . }}[]{
			{{- range $r.NotIn -}}
				{{- sprintf "%v" . -}}{{ phpTypeLiteralSuffixFor $ }},
			{{- end -}}
		};
{{- end -}}`

const numTpl = `{{ $f := .Field }}{{ $r := .Rules -}}
{{- if $r.Const }}
			io.envoyproxy.pgv.ConstantValidation.constant("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ constantName . "Const" }});
{{- end -}}
{{- if and (or $r.Lt $r.Lte) (or $r.Gt $r.Gte)}}
			io.envoyproxy.pgv.ComparativeValidation.range("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ if $r.Lt }}{{ constantName . "Lt" }}{{ else }}null{{ end }}, {{ if $r.Lte }}{{ constantName . "Lte" }}{{ else }}null{{ end }}, {{ if $r.Gt }}{{ constantName . "Gt" }}{{ else }}null{{ end }}, {{ if $r.Gte }}{{ constantName . "Gte" }}{{ else }}null{{ end }}, php.util.Comparator.naturalOrder());
{{- else -}}
{{- if $r.Lt }}
			io.envoyproxy.pgv.ComparativeValidation.lessThan("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ constantName . "Lt" }}, php.util.Comparator.naturalOrder());
{{- end -}}
{{- if $r.Lte }}
			io.envoyproxy.pgv.ComparativeValidation.lessThanOrEqual("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ constantName . "Lte" }}, php.util.Comparator.naturalOrder());
{{- end -}}
{{- if $r.Gt }}
			io.envoyproxy.pgv.ComparativeValidation.greaterThan("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ constantName . "Gt" }}, php.util.Comparator.naturalOrder());
{{- end -}}
{{- if $r.Gte }}
			io.envoyproxy.pgv.ComparativeValidation.greaterThanOrEqual("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ constantName . "Gte" }}, php.util.Comparator.naturalOrder());
{{- end -}}
{{- end -}}
{{- if $r.In }}
			io.envoyproxy.pgv.CollectiveValidation.in("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ constantName . "In" }});
{{- end -}}
{{- if $r.NotIn }}
			io.envoyproxy.pgv.CollectiveValidation.notIn("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ constantName . "NotIn" }});
{{- end -}}
`
