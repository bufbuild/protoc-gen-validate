package java

const requiredTpl = `
	{{ if .Rules.GetRequired }}
		if (!{{ hasAccessor . }}) {
			 throw new com.lyft.pgv.ValidationException("{{ fieldName . }}", "value is required ");
		}
	{{ end }}
`
