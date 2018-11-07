package java

const numConstTpl = `{{ $f := .Field }}{{ $r := .Rules -}}
{{- if $r.Const }}
		private final {{ javaTypeFor $f}} {{ constantName $f "Const" }} = {{ $r.GetConst }}{{ javaTypeLiteralSuffixFor $f }};
{{- end -}}
{{- if $r.Lt }}
		private final {{ javaTypeFor $f}} {{ constantName $f "Lt" }} = {{ $r.GetLt }}{{ javaTypeLiteralSuffixFor $f }};
{{- end -}}
{{- if $r.Lte }}
		private final {{ javaTypeFor $f}} {{ constantName $f "Lte" }} = {{ $r.GetLte }}{{ javaTypeLiteralSuffixFor $f }};
{{- end -}}
{{- if $r.Gt }}
		private final {{ javaTypeFor $f}} {{ constantName $f "Gt" }} = {{ $r.GetGt }}{{ javaTypeLiteralSuffixFor $f }};
{{- end -}}
{{- if $r.Gte }}
		private final {{ javaTypeFor $f}} {{ constantName $f "Gte" }} = {{ $r.GetGte }}{{ javaTypeLiteralSuffixFor $f }};
{{- end -}}
{{- if $r.In }}
		private final {{ javaTypeFor $f }}[] {{ constantName $f "In" }} = new {{ javaTypeFor $f }}[]{
			{{- range $r.In -}}
				{{- sprintf "%v" . -}}{{ javaTypeLiteralSuffixFor $f }},
			{{- end -}}
		};
{{- end -}}
{{- if $r.NotIn }}
		private final {{ javaTypeFor $f }}[] {{ constantName $f "NotIn" }} = new {{ javaTypeFor $f }}[]{
			{{- range $r.NotIn -}}
				{{- sprintf "%v" . -}}{{ javaTypeLiteralSuffixFor $f }},
			{{- end -}}
		};
{{- end -}}`

const numTpl = `{{ $f := .Field }}{{ $r := .Rules -}}
{{- if $r.Const }}
			com.lyft.pgv.ConstantValidation.constant("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ constantName $f "Const" }});
{{- end -}}
{{- if and (or $r.Lt $r.Lte) (or $r.Gt $r.Gte)}}
			com.lyft.pgv.ComparativeValidation.range("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ if $r.Lt }}{{ constantName $f "Lt" }}{{ else }}null{{ end }}, {{ if $r.Lte }}{{ constantName $f "Lte" }}{{ else }}null{{ end }}, {{ if $r.Gt }}{{ constantName $f "Gt" }}{{ else }}null{{ end }}, {{ if $r.Gte }}{{ constantName $f "Gte" }}{{ else }}null{{ end }}, java.util.Comparator.naturalOrder());
{{- else -}}
{{- if $r.Lt }}
			com.lyft.pgv.ComparativeValidation.lessThan("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ constantName $f "Lt" }}, java.util.Comparator.naturalOrder());
{{- end -}}
{{- if $r.Lte }}
			com.lyft.pgv.ComparativeValidation.lessThanOrEqual("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ constantName $f "Lte" }}, java.util.Comparator.naturalOrder());
{{- end -}}
{{- if $r.Gt }}
			com.lyft.pgv.ComparativeValidation.greaterThan("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ constantName $f "Gt" }}, java.util.Comparator.naturalOrder());
{{- end -}}
{{- if $r.Gte }}
			com.lyft.pgv.ComparativeValidation.greaterThanOrEqual("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ constantName $f "Gte" }}, java.util.Comparator.naturalOrder());
{{- end -}}
{{- end -}}
{{- if $r.In }}
			com.lyft.pgv.CollectiveValidation.in("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ constantName $f "In" }});
{{- end -}}
{{- if $r.NotIn }}
			com.lyft.pgv.CollectiveValidation.notIn("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ constantName $f "NotIn" }});
{{- end -}}
`
