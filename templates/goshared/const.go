package goshared

const constTpl = `{{ $r := .Rules }}
	{{ if $r.Const }}
		if m.maskHas(mask, "{{ .Field.Name }}") && {{ accessor . }} != {{ lit $r.GetConst }} {
			return {{ err . "value must equal " $r.GetConst }}
		}
	{{ end }}
`
