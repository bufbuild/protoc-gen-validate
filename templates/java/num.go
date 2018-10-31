package java

const numTpl = `{{ $f := .Field }}{{ $r := .Rules -}}
{{- if $r.Const }}
			com.lyft.pgv.NumericValidation.constant("{{ $f.FullyQualifiedName }}", proto.{{ wrappedAccessor . }}, {{ $r.GetConst }});
{{- end -}}
{{- if $r.Lt }}
			com.lyft.pgv.NumericValidation.lessThan("{{ $f.FullyQualifiedName }}", proto.{{ wrappedAccessor . }}, {{ $r.GetLt }});
{{- end -}}
{{- if $r.Lte }}
			com.lyft.pgv.NumericValidation.lessThanOrEqual("{{ $f.FullyQualifiedName }}", proto.{{ wrappedAccessor . }}, {{ $r.GetLte }});
{{- end -}}
{{- if $r.Gt }}
			com.lyft.pgv.NumericValidation.greaterThan("{{ $f.FullyQualifiedName }}", proto.{{ wrappedAccessor . }}, {{ $r.GetGt }});
{{- end -}}
{{- if $r.Gte }}
			com.lyft.pgv.NumericValidation.greaterThanOrEqual("{{ $f.FullyQualifiedName }}", proto.{{ wrappedAccessor . }}, {{ $r.GetGte }});
{{- end -}}
{{- if $r.In }}
			{
				{{ javaTypeFor $f }}[] set = new {{ javaTypeFor $f }}[]{
					{{- range $r.In -}}
						{{- sprintf "%v" . -}}{{ javaTypeLiteralSuffixFor $f }},
					{{- end -}}
				};
				com.lyft.pgv.NumericValidation.in("{{ $f.FullyQualifiedName }}", proto.{{ wrappedAccessor . }}, set);
			}
{{- end -}}
{{- if $r.NotIn }}
			{
				{{ javaTypeFor $f }}[] set = new {{ javaTypeFor $f }}[]{
					{{- range $r.NotIn -}}
						{{- sprintf "%v" . -}}{{ javaTypeLiteralSuffixFor $f }},
					{{- end -}}
				};
				com.lyft.pgv.NumericValidation.notIn("{{ $f.FullyQualifiedName }}", proto.{{ wrappedAccessor . }}, set);
			}
{{- end -}}
`
