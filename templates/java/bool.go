package java

const boolTpl = `{{ $f := .Field }}{{ $r := .Rules -}}
{{- if $r.Const }}
			com.lyft.pgv.BooleanValidation.constant("{{ $f.FullyQualifiedName }}", proto.{{ wrappedAccessor . }}, {{ $r.GetConst }});
{{- end }}`
