package python

const enumTpl = `
	{{ $f := .Field }}{{ $r := .Rules }}
	{{ template "const" . }}
	{{ template "in" . }}

	{{ if $r.GetDefinedOnly }}
	if {{ accessor . }} not in {{ enumValues $f }}:
		{{ err . "value must be one of the defined enum values" }}
	{{ end }}
`
