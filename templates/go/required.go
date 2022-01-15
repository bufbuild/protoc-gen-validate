package golang

const requiredTpl = `
	{{ if .Rules.GetRequired }}
		if {{ accessor . }} == nil {
			{{ if ne .Rules.GetErrorMsg "" }}
			err := {{ err . .Rules.GetErrorMsg }}
			{{ else }}
			err := {{ err . "value is required" }}
			{{ end }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ end }}
`
