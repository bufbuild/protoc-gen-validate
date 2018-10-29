package java

const requiredTpl = `
	{{ if .Rules.GetRequired }}
		if (!{{ hasAccessor . }}) {
			 throw new ValidationException("value is required ");
		}
	{{ end }}
`
