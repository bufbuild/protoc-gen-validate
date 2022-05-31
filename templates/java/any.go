package java

const anyConstTpl = `{{ $f := .Field }}{{ $r := .Rules }}
{{- if $r.In }}
		private final String[] {{ constantName . "In" }} = new String[]{
			{{- range $r.In }}
			"{{ . }}",
			{{- end }}
		};
{{- end -}}
{{- if $r.NotIn }}
		private final String[] {{ constantName . "NotIn" }} = new String[]{
			{{- range $r.NotIn }}
			"{{ . }}",
			{{- end }}
		};
{{- end -}}`

const anyTpl = `{{ $f := .Field }}{{ $r := .Rules }}
	{{- template "required" . -}}

	{{- if $r.In }}
			valctx.getValidationCollector().assertValid( ({{ safeName . "value"}}) -> {
				if ({{ hasAccessor . }}) io.envoyproxy.pgv.CollectiveValidation.in("{{ $f.FullyQualifiedName }}", {{ accessor . }}.getTypeUrl(), {{ constantName . "In" }});
			},proto);
	{{- end -}}
	{{- if $r.NotIn }}
			valctx.getValidationCollector().assertValid( ({{ safeName . "value"}}) -> {
				if ({{ hasAccessor . }}) io.envoyproxy.pgv.CollectiveValidation.notIn("{{ $f.FullyQualifiedName }}", {{ accessor . }}.getTypeUrl(), {{ constantName . "NotIn" }});
			},proto);
	{{- end -}}
`
