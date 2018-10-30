package java

const mapTpl = `{{ $f := .Field }}{{ $r := .Rules -}}
{{- if $r.GetMinPairs }}
			com.lyft.pgv.MapValidation.min("{{ $f.FullyQualifiedName }}", proto.{{ accessor $f }}, {{ $r.GetMinPairs }});
{{- end -}}
{{- if $r.GetMaxPairs }}
			com.lyft.pgv.MapValidation.max("{{ $f.FullyQualifiedName }}", proto.{{ accessor $f }}, {{ $r.GetMaxPairs }});
{{- end -}}
{{- if $r.GetNoSparse }}
			com.lyft.pgv.MapValidation.noSparse("{{ $f.FullyQualifiedName }}", proto.{{ accessor $f }});
{{- end -}}
`
