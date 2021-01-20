package php

const msgTpl = `
{{ if not (ignored .) -}}
	/**
	 * Validates {@code {{ simpleName . }}} protobuf objects.
	 */ 	
	public static function load{{ camelCase (simpleName .) }}ValidatorMetadata(ClassMetadata $metadata)
	{
		{{- template "msgInner" . -}}
	}
{{- end -}}
`

const msgInnerTpl = `
	{{- range .NonOneOfFields }}
		{{ renderConstants (context .) }}
	{{ end }}
	{{ range .OneOfs }}
		{{ template "oneOfConst" . }}
	{{ end }}

	{{ if disabled . }}
		// Validate is disabled for {{ simpleName . }}
		return;
	{{- else -}}
	{{ range .NonOneOfFields -}}
		{{ render (context .) }}
	{{ end -}}
	{{- end }}
	`
	// {{ range .OneOfs }}
	// 	{{ template "oneOf" . }}
	// {{- end -}}
