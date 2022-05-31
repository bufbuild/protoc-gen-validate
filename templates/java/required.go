package java

const requiredTpl = `{{ $f := .Field }}
	{{- if .Rules.GetRequired }}
		if ({{ hasAccessor . }}) {
			valctx.getValidationCollector().assertValid( ({{ safeName . "value"}}) -> {
				io.envoyproxy.pgv.RequiredValidation.required("{{ $f.FullyQualifiedName }}", {{ accessor . }});
			},proto);
		} else {
			valctx.getValidationCollector().assertValid( ({{ safeName . "value"}}) -> {
				io.envoyproxy.pgv.RequiredValidation.required("{{ $f.FullyQualifiedName }}", null);
			},proto);
		};
	{{- end -}}
`
