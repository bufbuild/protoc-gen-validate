package goshared

const constTpl = `{{ $r := .Rules }}
	{{ if $r.Const }}
		if {{ accessor . }} != {{ lit $r.GetConst }} {
			err := {{ err . "value must equal " $r.GetConst }}
			if stopOnError { return err }
			multiErr = multierror.Append(multiErr, err)
		}
	{{ end }}
`
