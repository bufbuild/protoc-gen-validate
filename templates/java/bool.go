package java

const boolTpl = `{{ $f := .Field }}{{ $r := .Rules -}}
{{- if $r.Const }}
			com.lyft.pgv.BoolValidation.constant("{{ $f.FullyQualifiedName }}", proto.{{ accessor $f }}, {{ $r.GetConst }});
{{- end }}`
