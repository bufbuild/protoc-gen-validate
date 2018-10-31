package java

const wrapperTpl = `
	{{ $f := .Field }}{{ $r := .Rules }}
    	if ({{ hasAccessor . }}) {
    		{{- render (unwrap .) }}
    	}
`
