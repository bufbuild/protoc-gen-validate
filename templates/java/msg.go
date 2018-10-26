package java

const msgTpl = `	
	/**
	 * Validates {@code {{ simpleName . }}} protobuf objects.
	 */
	public static class {{ simpleName . }}Validator {
		/**
		 * Asserts validation rules on a protobuf object.
		 * 
		 * @param proto the protobuf object to validate.
		 * @throws com.lyft.pgv.ValidationException with the first validation error encountered.
		 */
		public void assertValid({{ qualifiedName . }} proto) throws com.lyft.pgv.ValidationException {
			{{ range .NonOneOfFields -}}
				{{- render (context .) -}}
			{{- end }}
		}

		/**
		 * Checks validation rules on a protobuf object.
		 *
		 * @param proto the protobuf object to validate.
		 * @return {@code true} if all rules are valid, {@code false} if not.
		 */
		public boolean isValid({{ qualifiedName . }} proto) {
			try {
				assertValid(proto);
				return true;
			} catch (com.lyft.pgv.ValidationException ex) {
				return false;
			}
		}
	}
`
