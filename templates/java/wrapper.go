package java

const wrapperConstTpl = `{{ $f := .Field }}{{ $r := .Rules }}			
			{{- renderConstants (unwrap .) }}`

const wrapperTpl = `{{ $f := .Field }}{{ $r := .Rules }}			
			if ({{ hasAccessor . }}) {
				{{- render (unwrap .) }}
			}`
