package java

const requiredTpl = `{{ $f := .Field }}
	{{- if .Rules.GetRequired }}
			com.lyft.pgv.RequiredValidation.required("{{ $f.FullyQualifiedName }}", {{ accessor . }});
	{{- end -}}
`
