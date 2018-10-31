package java

const msgTpl = `
	/**
	 * Validates {@code {{ simpleName . }}} protobuf objects.
	 */
	public static class {{ simpleName . }}Validator extends com.lyft.pgv.Validator<{{ qualifiedName . }}> {
		public void assertValid({{ qualifiedName . }} proto) throws com.lyft.pgv.ValidationException {
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
