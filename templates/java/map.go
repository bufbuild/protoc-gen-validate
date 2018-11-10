package java

const mapConstTpl = `{{ $f := .Field }}{{ $r := .Rules -}}
{{ if or (ne (.Elem "" "").Typ "none") (ne (.Key "" "").Typ "none") }}
		{{ renderConstants (.Key "key" "Key") }}
		{{ renderConstants (.Elem "value" "Value") }}
{{- end -}}
`

const mapTpl = `{{ $f := .Field }}{{ $r := .Rules -}}
{{- if $r.GetMinPairs }}
			com.lyft.pgv.MapValidation.min("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ $r.GetMinPairs }});
{{- end -}}
{{- if $r.GetMaxPairs }}
			com.lyft.pgv.MapValidation.max("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ $r.GetMaxPairs }});
{{- end -}}
{{- if $r.GetNoSparse }}
			com.lyft.pgv.MapValidation.noSparse("{{ $f.FullyQualifiedName }}", {{ accessor . }});
{{- end -}}
{{ if or (ne (.Elem "" "").Typ "none") (ne (.Key "" "").Typ "none") }}
			com.lyft.pgv.MapValidation.validateParts({{ accessor . }}.keySet(), key -> {
				{{ render (.Key "key" "Key") }}
			});
			com.lyft.pgv.MapValidation.validateParts({{ accessor . }}.values(), value -> {
				{{ render (.Elem "value" "Value") }}
			});
{{- end -}}
`
