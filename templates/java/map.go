package java

const mapConstTpl = `{{ $f := .Field }}{{ $r := .Rules -}}
{{ if or (ne (.Elem "" "").Typ "none") (ne (.Key "" "").Typ "none") }}
		{{ renderConstants (.Key "key" "Key") }}
		{{ renderConstants (.Elem "value" "Value") }}
{{- end -}}
`

const mapTpl = `{{ $f := .Field }}{{ $r := .Rules -}}
{{- if $r.GetIgnoreEmpty }}
	if ( !{{ accessor . }}.isEmpty() ) {
{{- end -}}
{{- if $r.GetMinPairs }}
	valctx.getValidatorInterceptor().validate( ({{ safeName . "value"}}) -> {
		io.envoyproxy.pgv.MapValidation.min("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ $r.GetMinPairs }});
	},proto);
{{- end -}}
{{- if $r.GetMaxPairs }}
	valctx.getValidatorInterceptor().validate( ({{ safeName . "value"}}) -> {
		io.envoyproxy.pgv.MapValidation.max("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ $r.GetMaxPairs }});
	},proto);
{{- end -}}
{{- if $r.GetNoSparse }}
	valctx.getValidatorInterceptor().validate( ({{ safeName . "value"}}) -> {
		io.envoyproxy.pgv.MapValidation.noSparse("{{ $f.FullyQualifiedName }}", {{ accessor . }});
	},proto);
{{- end -}}
{{ if or (ne (.Elem "" "").Typ "none") (ne (.Key "" "").Typ "none") }}
			io.envoyproxy.pgv.MapValidation.validateParts({{ accessor . }}.keySet(), key -> {
				{{ render (.Key "key" "Key") }}
			});
			io.envoyproxy.pgv.MapValidation.validateParts({{ accessor . }}.values(), value -> {
				{{ render (.Elem "value" "Value") }}
			});
{{- end -}}
{{- if $r.GetIgnoreEmpty }}
			}
{{- end -}}
`
