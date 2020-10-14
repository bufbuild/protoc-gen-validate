package golang

const requiredTpl = `
	{{ if .Rules.GetRequired }}
		if {{ accessor . }} == nil {
			return {{ err . (t "<prefix>.required" "value is required") }}
		}
	{{ end }}
`
