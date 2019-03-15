package python

const constTpl = `{{ $r := .Rules }}
	{{ if $r.Const }}
		if {{ accessor . }} != {{ lit $r.GetConst }} :
			{{err . "Cannot validate" }}
	{{ end }}
`
