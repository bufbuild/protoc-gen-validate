package golang

const requiredTpl = `
	{{ if or (.Rules.GetRequired) (.MessageRules.GetRequired) }}
		if {{ accessor . }} == nil {
			return {{ err . "value is required" }}
		}
	{{ end }}
`
