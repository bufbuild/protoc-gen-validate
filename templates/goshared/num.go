package goshared

const numTpl = `
	{{if .Rules.GetOmitempty}}
		if {{accessor .}} != 0 {
	{{end}}

	{{ template "const" . }}
	{{ template "ltgt" . }}
	{{ template "in" . }}

	{{if .Rules.GetOmitempty}}
		}
	{{end}}

`
