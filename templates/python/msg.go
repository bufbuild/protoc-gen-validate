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
		# Validate is disabled. This method will always return nil.
	return True, ""
	{{ else -}}
		# Validate checks the field values with the rules defined in the proto definition for this message. If any rules are violated, an error is returned.
	{{ end }}
	{{ range .NonOneOfFields }}
		{{ render (context .) }}
	{{ end }}
	{{ range .OneOfs }}
	{{ unimplemented }}
	{{ end }}
	return True, ""
`
