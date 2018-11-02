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
			com.lyft.pgv.ConstantValidation.constant("{{ $f.FullyQualifiedName }}", proto.{{ accessor . }}, {{ constantName $f "Const" }});
{{- end -}}
{{- if $r.Lt }}
			com.lyft.pgv.ComparativeValidation.lessThan("{{ $f.FullyQualifiedName }}", proto.{{ accessor . }}, {{ constantName $f "Lt" }}, com.google.protobuf.util.Durations.comparator());
{{- end -}}
{{- if $r.Lte }}
			com.lyft.pgv.ComparativeValidation.lessThanOrEqual("{{ $f.FullyQualifiedName }}", proto.{{ accessor . }}, {{ constantName $f "Lte" }}, com.google.protobuf.util.Durations.comparator());
{{- end -}}
{{- if $r.Gt }}
			com.lyft.pgv.ComparativeValidation.greaterThan("{{ $f.FullyQualifiedName }}", proto.{{ accessor . }}, {{ constantName $f "Gt" }}, com.google.protobuf.util.Durations.comparator());
{{- end -}}
{{- if $r.Gte }}
			com.lyft.pgv.ComparativeValidation.greaterThanOrEqual("{{ $f.FullyQualifiedName }}", proto.{{ accessor . }}, {{ constantName $f "Gte" }}, com.google.protobuf.util.Durations.comparator());
{{- end -}}
{{- if $r.In }}
			com.lyft.pgv.DurationValidation.in("{{ $f.FullyQualifiedName }}", proto.{{ accessor . }}, {{ constantName $f "In" }});
{{- end -}}
{{- if $r.NotIn }}
			com.lyft.pgv.DurationValidation.notIn("{{ $f.FullyQualifiedName }}", proto.{{ accessor . }}, {{ constantName $f "NotIn" }});
{{- end -}}
`
