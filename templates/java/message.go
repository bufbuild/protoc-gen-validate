package java

const msgTpl = `
	public static class {{ .Name }}Validator {
		public boolean isValid({{ qualifiedName . }} proto) {
			return true;
		}
	}
`
