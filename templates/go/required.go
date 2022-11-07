package golang

const requiredTpl = `
	{{ if .Rules.GetRequired }}{{ $t := .Typ}}
		if {{ accessor . }} == nil {
			err := {{ err . (print $t ".required") "value is required" }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ end }}
`
