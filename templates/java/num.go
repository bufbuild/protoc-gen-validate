package java

const numConstTpl = `{{ $f := .Field }}{{ $r := .Rules -}}
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
			com.lyft.pgv.ConstantValidation.constant("{{ $f.FullyQualifiedName }}", proto.{{ accessor . }}, {{ $r.GetConst }});
{{- end -}}
{{- if $r.Lt }}
			com.lyft.pgv.NumericValidation.lessThan("{{ $f.FullyQualifiedName }}", proto.{{ accessor . }}, {{ $r.GetLt }});
{{- end -}}
{{- if $r.Lte }}
			com.lyft.pgv.NumericValidation.lessThanOrEqual("{{ $f.FullyQualifiedName }}", proto.{{ accessor . }}, {{ $r.GetLte }});
{{- end -}}
{{- if $r.Gt }}
			com.lyft.pgv.NumericValidation.greaterThan("{{ $f.FullyQualifiedName }}", proto.{{ accessor . }}, {{ $r.GetGt }});
{{- end -}}
{{- if $r.Gte }}
			com.lyft.pgv.NumericValidation.greaterThanOrEqual("{{ $f.FullyQualifiedName }}", proto.{{ accessor . }}, {{ $r.GetGte }});
{{- end -}}
{{- if $r.In }}
			com.lyft.pgv.NumericValidation.in("{{ $f.FullyQualifiedName }}", proto.{{ accessor . }}, {{ constantName $f "In" }});
{{- end -}}
{{- if $r.NotIn }}
			com.lyft.pgv.NumericValidation.notIn("{{ $f.FullyQualifiedName }}", proto.{{ accessor . }}, {{ constantName $f "NotIn" }});
{{- end -}}
`
