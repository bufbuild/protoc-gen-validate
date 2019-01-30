package cc

const enumTpl = `
		{{ $f := .Field }}{{ $r := .Rules }}
		{{ template "const" . }}
		{{ template "in" . }}

		{{ if $r.GetDefinedOnly }}
			{{ if $f.Type.IsRepeated }}
				if (!{{ class $f.Type.Element.Enum }}_IsValid({{ accessor . }})) {
			{{ else }}
				if (!{{ class $f.Type.Enum }}_IsValid({{ accessor . }})) {
			{{ end }}
				{{ err . "value must be one of the defined enum values" }}
			}
		{{ end }}
`
