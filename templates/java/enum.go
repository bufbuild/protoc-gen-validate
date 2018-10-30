package java

const enumTpl = `{{ $f := .Field }}{{ $r := .Rules -}}
{{- if $r.Const }}
			com.lyft.pgv.EnumValidation.constant("{{ $f.FullyQualifiedName }}", proto.{{ accessor $f }},
				{{ javaTypeFor $f }}.forNumber({{ $r.GetConst }}));
{{- end -}}
{{- if $r.GetDefinedOnly }}
			com.lyft.pgv.EnumValidation.definedOnly("{{ $f.FullyQualifiedName }}", proto.{{ accessor $f }});
{{- end -}}
{{- if $r.In }}
			{{ javaTypeFor $f }}[] set = new {{ javaTypeFor $f }}[]{
				{{- range $r.In }}
				{{ javaTypeFor $f }}.forNumber({{- sprintf "%v" . -}}),
				{{- end }}
			};
			com.lyft.pgv.EnumValidation.in("{{ $f.FullyQualifiedName }}", proto.{{ accessor $f }}, set);
{{- end -}}
{{- if $r.NotIn }}
			{{ javaTypeFor $f }}[] set = new {{ javaTypeFor $f }}[]{
				{{- range $r.NotIn }}
				{{ javaTypeFor $f }}.forNumber({{- sprintf "%v" . -}}),
				{{- end }}
			};
			com.lyft.pgv.EnumValidation.notIn("{{ $f.FullyQualifiedName }}", proto.{{ accessor $f }}, set);
{{- end -}}
`
