package java

const oneOfTpl = `
			switch (proto.get{{camelCase .Name }}Case()) {
				{{ range .Fields -}}
				case {{ oneof . }}:
					{{ render (context .) }}
					break;
				{{ end -}}
				{{- if required . }}
				default: 
					com.lyft.pgv.RequiredValidation.required("{{ .FullyQualifiedName }}", null);
				{{- end }}
			}
`
