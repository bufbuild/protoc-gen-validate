package golang

const requiredTpl = `
	{{ if .Rules.GetRequired }}
		if {{ accessor . }} == nil {
			err := {{ err . "value is required" }}
			if stopOnError { return err }
			multiErr = multierror.Append(multiErr, err)
			
		}
	{{ end }}
`
