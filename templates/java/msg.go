package java

const msgTpl = `
	/**
	 * Validates {@code {{ simpleName . }}} protobuf objects.
	 */
	public static class {{ simpleName . }}Validator extends com.lyft.pgv.Validator<{{ qualifiedName . }}> {
		{{- range .NonOneOfFields }}
			{{ renderConstants (context .) }}
		{{ end }}
		{{ range .OneOfs }}
			{{ template "oneOfConst" . }}
		{{ end }}

		public void assertValid({{ qualifiedName . }} proto, com.lyft.pgv.ValidatorIndex index) throws com.lyft.pgv.ValidationException {
		{{ if disabled . }}
			// Validate is disabled for {{ simpleName . }}
			return;
		{{- else -}}
		{{ range .NonOneOfFields -}}
			{{ render (context .) }}
		{{ end -}}
		{{ range .OneOfs }}
			{{ template "oneOf" . }}
		{{- end -}}
		{{- end }}
		}
	}
`
