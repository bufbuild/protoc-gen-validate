package tpl

const messageTpl = `
	{{ $f := .Field }}{{ $r := .Rules }}
	{{ template "required" . }}

	{{ if $r.GetSkip }}
		// skipping validation for {{ $f.Name }}
	{{ else }}
	{
		string inner_err;
		if ({{ hasAccessor .}} && !pgv::validate::MessageValidator<{{ ctype $f.Type }}>::Check({{ accessor . }}, &inner_err)) {
			{{ errCause . "inner_err" "embedded message failed validation" }}
		}
	}
	{{ end }}
`

const requiredTpl = `
	{{ if .Rules.GetRequired }}
		if (!{{ hasAccessor . }}) {
			{{ err . "value is required" }}
		}
	{{ end }}
`
