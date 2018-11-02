package java

const enumConstTpl = `{{ $f := .Field }}{{ $r := .Rules -}}
{{- if $r.In }}
		private final {{ javaTypeFor $f }}[] {{ constantName $f "In" }} = new {{ javaTypeFor $f }}[]{
			{{- range $r.In }}
			{{ javaTypeFor $f }}.forNumber({{- sprintf "%v" . -}}),
			{{- end }}
		};
{{- end -}}
{{- if $r.NotIn }}
		private final {{ javaTypeFor $f }}[] {{ constantName $f "NotIn" }} = new {{ javaTypeFor $f }}[]{
			{{- range $r.NotIn }}
			{{ javaTypeFor $f }}.forNumber({{- sprintf "%v" . -}}),
			{{- end }}
		};
{{- end -}}`

const enumTpl = `{{ $f := .Field }}{{ $r := .Rules -}}
{{- if $r.Const }}
			com.lyft.pgv.ConstantValidation.constant("{{ $f.FullyQualifiedName }}", proto.{{ accessor . }}, 
				{{ javaTypeFor $f }}.forNumber({{ $r.GetConst }}));
{{- end -}}
{{- if $r.GetDefinedOnly }}
			com.lyft.pgv.EnumValidation.definedOnly("{{ $f.FullyQualifiedName }}", proto.{{ accessor . }});
{{- end -}}
{{- if $r.In }}
			com.lyft.pgv.CollectiveValidation.in("{{ $f.FullyQualifiedName }}", proto.{{ accessor . }}, {{ constantName $f "In" }});
{{- end -}}
{{- if $r.NotIn }}
			com.lyft.pgv.CollectiveValidation.notIn("{{ $f.FullyQualifiedName }}", proto.{{ accessor . }}, {{ constantName $f "NotIn" }});
{{- end -}}
`
