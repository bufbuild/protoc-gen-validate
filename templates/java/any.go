package java

const anyConstTpl = `{{ $f := .Field }}{{ $r := .Rules }}
{{- if $r.In }}
		private final String[] {{ constantName $f "In" }} = new String[]{
			{{- range $r.In }}
			"{{ . }}",
			{{- end }}
		};
{{- end -}}
{{- if $r.NotIn }}
		private final String[] {{ constantName $f "NotIn" }} = new String[]{
			{{- range $r.NotIn }}
			"{{ . }}",
			{{- end }}
		};
{{- end -}}`

const anyTpl = `{{ $f := .Field }}{{ $r := .Rules }}
	{{- template "required" . -}}

	{{- if $r.In }}
			com.lyft.pgv.AnyValidation.in("{{ $f.FullyQualifiedName }}", proto.{{ accessor . }}, {{ constantName $f "In" }});
	{{- end -}}
	{{- if $r.NotIn }}
			com.lyft.pgv.AnyValidation.notIn("{{ $f.FullyQualifiedName }}", proto.{{ accessor . }}, {{ constantName $f "NotIn" }});
	{{- end -}}
`
