package java

const msgTpl = `
	public static class {{ .Name }}Validator {
		public boolean isValid({{ qualifiedName . }} proto) {
			{{ range .NonOneOfFields }}
				{{ render (context .) }}
			{{ end }}

			return true;
		}
	}
`
