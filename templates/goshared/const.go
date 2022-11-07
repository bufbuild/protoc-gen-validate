package goshared

const constTpl = `{{ $f := .Field }}{{ $r := .Rules }}{{ $t := .Typ}}
	{{ if $r.Const }}
		if {{ accessor . }} != {{ lit $r.GetConst }} {
			{{- if isEnum $f }}
			err := {{ err . (print $t ".const") "value must equal " (enumVal $f $r.GetConst) }}
			{{- else }}
			err := {{ err . (print $t ".const") "value must equal " $r.GetConst }}
			{{- end }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ end }}
`
