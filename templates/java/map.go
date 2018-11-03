package java

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
`
