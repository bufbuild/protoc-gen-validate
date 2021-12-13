package goshared

const constTpl = `{{ $r := .Rules }}
	{{ if $r.Const }}
		if {{ accessor . }} != {{ lit $r.GetConst }} {
			err := {{ err . (t "<prefix>.const" "value must equal {{$1}}" $r.GetConst) }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ end }}
`
