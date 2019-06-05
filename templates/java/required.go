package java

const requiredTpl = `{{ $f := .Field }}
	{{- if or (.Rules.GetRequired) (hasRequired .Field) }}
		if ({{ hasAccessor . }}) {
			io.envoyproxy.pgv.RequiredValidation.required("{{ $f.FullyQualifiedName }}", {{ accessor . }});
		} else {
			io.envoyproxy.pgv.RequiredValidation.required("{{ $f.FullyQualifiedName }}", null);
		};
	{{- end -}}
`
