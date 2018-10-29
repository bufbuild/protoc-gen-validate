package java

const msgTpl = `	
	/**
	 * Validates {@code {{ simpleName . }}} protobuf objects.
	 */
	public static class {{ simpleName . }}Validator extends com.lyft.pgv.Validator<{{ qualifiedName . }}> {
		public void assertValid({{ qualifiedName . }} proto) throws com.lyft.pgv.ValidationException {
			{{ range .NonOneOfFields -}}
				{{- render (context .) -}}
			{{- end }}
		}
	}
`
