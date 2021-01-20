package php

const requiredTpl = `{{ $f := .Field }}
	{{- if .Rules.GetRequired }}
			$metadata->addPropertyConstraint('{{ $f.FullyQualifiedName }}', new Assert\NotBlank());
	{{- end -}}
`
