package java

const repeatedConstTpl = `{{ renderConstants (.Elem "" "") }}`

const repeatedTpl = `{{ $f := .Field }}{{ $r := .Rules -}}
{{- if $r.GetIgnoreEmpty }}
			if ( !{{ accessor . }}.isEmpty() ) {
{{- end -}}
{{- if $r.GetMinItems }}
	valctx.getValidationCollector().assertValid( ({{ safeName . "value"}}) -> {
		io.envoyproxy.pgv.RepeatedValidation.minItems("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ $r.GetMinItems }});
	},proto);
{{- end -}}
{{- if $r.GetMaxItems }}
	valctx.getValidationCollector().assertValid( ({{ safeName . "value"}}) -> {
		io.envoyproxy.pgv.RepeatedValidation.maxItems("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ $r.GetMaxItems }});
	},proto);
{{- end -}}
{{- if $r.GetUnique }}
	valctx.getValidationCollector().assertValid( ({{ safeName . "value"}}) -> {
		io.envoyproxy.pgv.RepeatedValidation.unique("{{ $f.FullyQualifiedName }}", {{ accessor . }});
	},proto);
{{- end }}
			io.envoyproxy.pgv.RepeatedValidation.forEach({{ accessor . }}, item -> {
				{{ render (.Elem "item" "") }}
			});
{{- if $r.GetIgnoreEmpty }}
			}
{{- end -}}
`
