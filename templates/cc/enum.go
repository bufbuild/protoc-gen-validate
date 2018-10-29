package cc

const enumTpl = `
		{{ $f := .Field }}{{ $r := .Rules }}
		{{ template "const" . }}
		{{ template "in" . }}

		{{ if $r.GetDefinedOnly }}
			if (!{{ package $f.Type.Enum }}::{{ (typ $f).Element }}_IsValid({{ accessor . }})) {
				{{ err . "value must be one of the defined enum values" }}
			}
		{{ end }}
`
