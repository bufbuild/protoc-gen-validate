package java

const boolTpl = `{{ $f := .Field }}{{ $r := .Rules -}}
{{- if $r.Const }}
			com.lyft.pgv.ConstantValidation.constant("{{ $f.FullyQualifiedName }}", proto.{{ accessor . }}, {{ $r.GetConst }});
{{- end }}`
