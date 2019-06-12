package gogo

const requiredTpl = `
	{{ if or (.Rules.GetRequired) (.MessageRules.GetRequired) }}
		{{ if .Gogo.Nullable }}
			if {{ accessor . }} == nil {
				return {{ err . "value is required" }}
			}
		{{ end }}
	{{ end }}
`
