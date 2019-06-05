package golang

const requiredTpl = `
	{{ if or (.Rules.GetRequired) (hasrequired .Field) }}
		if {{ accessor . }} == nil {
			return {{ err . "value is required" }}
		}
	{{ end }}
`
