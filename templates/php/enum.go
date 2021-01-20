package php

const enumConstTpl = `{{ $ctx := . }}{{ $f := .Field }}{{ $r := .Rules -}}
{{- if $r.In }}
		private final {{ phpTypeFor . }}[] {{ constantName . "In" }} = new {{ phpTypeFor . }}[]{
			{{- range $r.In }}
			{{ phpTypeFor $ctx }}.forNumber({{- sprintf "%v" . -}}),
			{{- end }}
		};
{{- end -}}
{{- if $r.NotIn }}
		private final {{ phpTypeFor . }}[] {{ constantName . "NotIn" }} = new {{ phpTypeFor . }}[]{
			{{- range $r.NotIn }}
			{{ phpTypeFor $ctx }}.forNumber({{- sprintf "%v" . -}}),
			{{- end }}
		};
{{- end -}}`

const enumTpl = `{{ $f := .Field }}{{ $r := .Rules -}}
{{- if $r.Const }}
			io.envoyproxy.pgv.ConstantValidation.constant("{{ $f.FullyQualifiedName }}", {{ accessor . }}, 
				{{ phpTypeFor . }}.forNumber({{ $r.GetConst }}));
{{- end -}}
{{- if $r.GetDefinedOnly }}
			io.envoyproxy.pgv.EnumValidation.definedOnly("{{ $f.FullyQualifiedName }}", {{ accessor . }});
{{- end -}}
{{- if $r.In }}
			io.envoyproxy.pgv.CollectiveValidation.in("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ constantName . "In" }});
{{- end -}}
{{- if $r.NotIn }}
			io.envoyproxy.pgv.CollectiveValidation.notIn("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ constantName . "NotIn" }});
{{- end -}}
`
