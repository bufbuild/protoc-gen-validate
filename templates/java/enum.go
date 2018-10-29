package java

const enumTpl = `{{ $f := .Field }}{{ $r := .Rules -}}
{{- if $r.Const }}
			com.lyft.pgv.EnumValidation.constant("{{ $f.FullyQualifiedName }}", proto.{{ accessor $f }}, {{ $r.GetConst }});
{{- end -}}
{{- if $r.GetDefinedOnly }}
			com.lyft.pgv.EnumValidation.definedOnly("{{ $f.FullyQualifiedName }}", proto.{{ accessor $f }});
{{- end -}}
{{- if $r.In }}
			{
				{{ javaTypeFor $f }}[] set = new {{ javaTypeFor $f }}[]{
					{{- range $r.In -}}
						{{- sprintf "%v" . -}}{{ javaTypeLiteralSuffixFor $f }},
					{{- end -}}
				};
				com.lyft.pgv.EnumValidation.in("{{ $f.FullyQualifiedName }}", proto.{{ accessor $f }}, set);
			}
{{- end -}}
{{- if $r.NotIn }}
			{
				{{ javaTypeFor $f }}[] set = new {{ javaTypeFor $f }}[]{
					{{- range $r.NotIn -}}
						{{- sprintf "%v" . -}}{{ javaTypeLiteralSuffixFor $f }},
					{{- end -}}
				};
				com.lyft.pgv.EnumValidation.notIn("{{ $f.FullyQualifiedName }}", proto.{{ accessor $f }}, set);
			}
{{- end -}}
`
