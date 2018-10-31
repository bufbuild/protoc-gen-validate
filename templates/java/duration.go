package java

const durationTpl = `{{ $f := .Field }}{{ $r := .Rules -}}
{{- template "required" . -}}

{{- if $r.Const }}
			com.lyft.pgv.DurationValidation.constant("{{ $f.FullyQualifiedName }}", proto.{{ accessor . }}, {{ durLit $r.GetConst }});
{{- end -}}
{{- if $r.Lt }}
			com.lyft.pgv.DurationValidation.lessThan("{{ $f.FullyQualifiedName }}", proto.{{ accessor . }}, {{ durLit $r.Lt }});
{{- end -}}
{{- if $r.Lte }}
			com.lyft.pgv.DurationValidation.lessThanOrEqual("{{ $f.FullyQualifiedName }}", proto.{{ accessor . }}, {{ durLit $r.Lte }});
{{- end -}}
{{- if $r.Gt }}
			com.lyft.pgv.DurationValidation.greaterThan("{{ $f.FullyQualifiedName }}", proto.{{ accessor . }}, {{ durLit $r.Gt }});
{{- end -}}
{{- if $r.Gte }}
			com.lyft.pgv.DurationValidation.greaterThanOrEqual("{{ $f.FullyQualifiedName }}", proto.{{ accessor . }}, {{ durLit $r.Gte }});
{{- end -}}
{{- if $r.In }}
			{
				com.google.protobuf.Duration[] set = new com.google.protobuf.Duration[]{
					{{- range $r.In }}
					{{ durLit . }},
					{{- end }}
				};
				com.lyft.pgv.DurationValidation.in("{{ $f.FullyQualifiedName }}", proto.{{ accessor . }}, set);
			}
{{- end -}}
{{- if $r.NotIn }}
			{
				com.google.protobuf.Duration[] set = new com.google.protobuf.Duration[]{
					{{- range $r.NotIn }}
					{{ durLit . }},
					{{- end }}
				};
				com.lyft.pgv.DurationValidation.notIn("{{ $f.FullyQualifiedName }}", proto.{{ accessor . }}, set);
			}
{{- end -}}
`
