package java

const requiredTpl = `{{ $f := .Field }}
	{{- if .Rules.GetRequired }}
		if ({{ hasAccessor . }}) {
			valctx.getValidatorInterceptor().validate( ({{ safeName . "value"}}) -> {
				io.envoyproxy.pgv.RequiredValidation.required("{{ $f.FullyQualifiedName }}", {{ accessor . }});
			},proto);
		} else {
			valctx.getValidatorInterceptor().validate( ({{ safeName . "value"}}) -> {
				io.envoyproxy.pgv.RequiredValidation.required("{{ $f.FullyQualifiedName }}", null);
			},proto);
		};
	{{- end -}}
`
