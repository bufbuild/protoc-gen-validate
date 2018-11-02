package java

const bytesConstTpl = `{{ $f := .Field }}{{ $r := .Rules -}}
{{- if $r.Const }}
		private final com.google.protobuf.ByteString {{ constantName $f "Const" }} = com.google.protobuf.ByteString.copyFrom({{ byteArrayLit $r.GetConst }});
{{- end -}}
{{- if $r.In }}
		private final com.google.protobuf.ByteString[] {{ constantName $f "In" }} = new com.google.protobuf.ByteString[]{
			{{- range $r.In }}
			com.google.protobuf.ByteString.copyFrom({{ byteArrayLit . }}),
			{{- end }}
		};
{{- end -}}
{{- if $r.NotIn }}
		private final com.google.protobuf.ByteString[] {{ constantName $f "NotIn" }} = new com.google.protobuf.ByteString[]{
			{{- range $r.NotIn }}
			com.google.protobuf.ByteString.copyFrom({{ byteArrayLit . }}),
			{{- end }}
		};
{{- end -}}
{{- if $r.Pattern }}
		private final com.google.re2j.Pattern {{ constantName $f "Pattern" }} = com.google.re2j.Pattern.compile({{ javaStringEscape $r.GetPattern }});
{{- end -}}
{{- if $r.Prefix }}
		private final byte[] {{ constantName $f "Prefix" }} = {{ byteArrayLit $r.GetPrefix }};
{{- end -}}
{{- if $r.Contains }}
		private final byte[] {{ constantName $f "Contains" }} = {{ byteArrayLit $r.GetContains }};
{{- end -}}
{{- if $r.Suffix }}
		private final byte[] {{ constantName $f "Suffix" }} = {{ byteArrayLit $r.GetSuffix }};
{{- end -}}`

const bytesTpl = `{{ $f := .Field }}{{ $r := .Rules -}}
{{- if $r.Const }}
			com.lyft.pgv.ConstantValidation.constant("{{ $f.FullyQualifiedName }}", proto.{{ accessor . }}, {{ constantName $f "Const" }});
{{- end -}}
{{- if $r.Len }}
			com.lyft.pgv.BytesValidation.length("{{ $f.FullyQualifiedName }}", proto.{{ accessor . }}, {{ $r.GetLen }});
{{- end -}}
{{- if $r.MinLen }}
			com.lyft.pgv.BytesValidation.minLength("{{ $f.FullyQualifiedName }}", proto.{{ accessor . }}, {{ $r.GetMinLen }});
{{- end -}}
{{- if $r.MaxLen }}
			com.lyft.pgv.BytesValidation.maxLength("{{ $f.FullyQualifiedName }}", proto.{{ accessor . }}, {{ $r.GetMaxLen }});
{{- end -}}
{{- if $r.Pattern }}
			com.lyft.pgv.BytesValidation.pattern("{{ $f.FullyQualifiedName }}", proto.{{ accessor . }}, {{ constantName $f "Pattern" }});
{{- end -}}
{{- if $r.Prefix }}
			com.lyft.pgv.BytesValidation.prefix("{{ $f.FullyQualifiedName }}", proto.{{ accessor . }}, {{ constantName $f "Prefix" }});
{{- end -}}
{{- if $r.Contains }}
			com.lyft.pgv.BytesValidation.contains("{{ $f.FullyQualifiedName }}", proto.{{ accessor . }}, {{ constantName $f "Contains" }});
{{- end -}}
{{- if $r.Suffix }}
			com.lyft.pgv.BytesValidation.suffix("{{ $f.FullyQualifiedName }}", proto.{{ accessor . }}, {{ constantName $f "Suffix" }});
{{- end -}}
{{- if $r.GetIp }}
			com.lyft.pgv.BytesValidation.ip("{{ $f.FullyQualifiedName }}", proto.{{ accessor . }});
{{- end -}}
{{- if $r.GetIpv4 }}
			com.lyft.pgv.BytesValidation.ipv4("{{ $f.FullyQualifiedName }}", proto.{{ accessor . }});
{{- end -}}
{{- if $r.GetIpv6 }}
			com.lyft.pgv.BytesValidation.ipv6("{{ $f.FullyQualifiedName }}", proto.{{ accessor . }});
{{- end -}}
{{- if $r.In }}
			com.lyft.pgv.CollectiveValidation.in("{{ $f.FullyQualifiedName }}", proto.{{ accessor . }}, {{ constantName $f "In" }});
{{- end -}}
{{- if $r.NotIn }}
			com.lyft.pgv.CollectiveValidation.notIn("{{ $f.FullyQualifiedName }}", proto.{{ accessor . }}, {{ constantName $f "NotIn" }});
{{- end -}}
`
