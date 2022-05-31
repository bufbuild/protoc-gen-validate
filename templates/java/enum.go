package java

const enumConstTpl = `{{ $ctx := . }}{{ $f := .Field }}{{ $r := .Rules -}}
{{- if $r.In }}
		private final {{ javaTypeFor . }}[] {{ constantName . "In" }} = new {{ javaTypeFor . }}[]{
			{{- range $r.In }}
			{{ javaTypeFor $ctx }}.forNumber({{- sprintf "%v" . -}}),
			{{- end }}
		};
{{- end -}}
{{- if $r.NotIn }}
		private final {{ javaTypeFor . }}[] {{ constantName . "NotIn" }} = new {{ javaTypeFor . }}[]{
			{{- range $r.NotIn }}
			{{ javaTypeFor $ctx }}.forNumber({{- sprintf "%v" . -}}),
			{{- end }}
		};
{{- end -}}`

const enumTpl = `{{ $f := .Field }}{{ $r := .Rules -}}
{{- if $r.Const }}
	valctx.getValidationCollector().assertValid( ({{ safeName . "value"}}) -> {
		io.envoyproxy.pgv.ConstantValidation.constant("{{ $f.FullyQualifiedName }}", {{ accessor . }}, 
					{{ javaTypeFor . }}.forNumber({{ $r.GetConst }}));
	},proto);
{{- end -}}
{{- if $r.GetDefinedOnly }}
	valctx.getValidationCollector().assertValid( ({{ safeName . "value"}}) -> {
		io.envoyproxy.pgv.EnumValidation.definedOnly("{{ $f.FullyQualifiedName }}", {{ accessor . }});
	},proto);
{{- end -}}
{{- if $r.In }}
	valctx.getValidationCollector().assertValid( ({{ safeName . "value"}}) -> {
		io.envoyproxy.pgv.CollectiveValidation.in("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ constantName . "In" }});
	},proto);
{{- end -}}
{{- if $r.NotIn }}
	valctx.getValidationCollector().assertValid( ({{ safeName . "value"}}) -> {
		io.envoyproxy.pgv.CollectiveValidation.notIn("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ constantName . "NotIn" }});
	},proto);
{{- end -}}
`
