package cc

const wrapperTpl = `
	{{ $f := .Field }}{{ $r := .Rules }}

	if ({{ hasAccessor . }}) {
		const auto wrapped = {{ accessor . }};
		{{ render (unwrap . "wrapped") }}
	}
	{{ if hasRequired $f }}
		{{ unimplemented }}
	{{ end }}
`
