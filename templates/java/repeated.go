package java

const repeatedConstTpl = `{{ renderConstants (.Elem "" "") }}`

const repeatedTpl = `{{ $f := .Field }}{{ $r := .Rules -}}
{{- if $r.GetMinItems }}
			com.lyft.pgv.RepeatedValidation.minItems("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ $r.GetMinItems }});
{{- end -}}
{{- if $r.GetMaxItems }}
			com.lyft.pgv.RepeatedValidation.maxItems("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ $r.GetMaxItems }});
{{- end -}}
{{- if $r.GetUnique }}
			com.lyft.pgv.RepeatedValidation.unique("{{ $f.FullyQualifiedName }}", {{ accessor . }});
{{- end -}}
{{- if $r.GetItems }}
			com.lyft.pgv.RepeatedValidation.forEach({{ accessor . }}, item -> {
				{{- render (.Elem "item" "") }}
			});
{{- end -}}`
