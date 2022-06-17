package java

const oneOfConstTpl = `
{{ range .Fields }}{{ renderConstants (context .) }}{{ end }}
`

const oneOfTpl = `
			switch (proto.get{{camelCase .Name }}Case()) {
				{{ range .Fields -}}
				case {{ oneof . }}:
					{{ render (context .) }}
					break;
				{{ end -}}
				{{- if required . }}
				default: 
					valctx.getValidatorInterceptor().validate( (value) -> {
						io.envoyproxy.pgv.RequiredValidation.required("{{ .FullyQualifiedName }}", null);
					},proto);
				{{- end }}
			}
`
