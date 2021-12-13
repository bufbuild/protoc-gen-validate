package golang

const requiredTpl = `
	{{ if .Rules.GetRequired }}
		if {{ accessor . }} == nil {
			err := {{ err . (t "<prefix>.required" "value is required") }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ end }}
`
