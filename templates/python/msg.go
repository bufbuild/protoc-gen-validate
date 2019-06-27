package python

const msgTpl = `
{{ range .Fields }}{{ with (context .) }}{{ $f := .Field }}
	{{ if has .Rules "In" }}{{ if .Rules.In }}
{{ lookup .Field "InLookup" }} = {
	{{- range .Rules.In }}
		{{ inKey $f . }},
	{{- end }}
}
	{{ end }}{{ end }}
 	{{ if has .Rules "NotIn" }}{{ if .Rules.NotIn }}
{{ lookup .Field "NotInLookup" }} = {
	{{- range .Rules.NotIn }}
		{{ inKey $f . }},
	{{- end }}
}
	{{ end }}{{ end }}
	{{ if has .Rules "Pattern"}}{{ if .Rules.Pattern }}
{{ lookup .Field "Pattern" }} = re.compile({{ lit .Rules.GetPattern }})
	{{ end }}{{ end }}
 {{ end }}{{ end }}

{{ if needs . "hostname" }}{{ template "hostname" . }}{{ end }}

{{ if needs . "email" }}{{ template "email" . }}{{ end }}

def validate_{{ .Name }}(m):
	{{ if disabled . -}}
	return True, ""
	{{ else -}}
	{{ range .NonOneOfFields }}
		{{ render (context .) }}
	{{ end }}
	{{ range .OneOfs }}
	{{ unimplemented }}
	{{ end }}
	return True, ""
	{{ end -}}
`
