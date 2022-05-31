package java

const bytesConstTpl = `{{ $f := .Field }}{{ $r := .Rules -}}
{{- if $r.Const }}
		private final com.google.protobuf.ByteString {{ constantName . "Const" }} = com.google.protobuf.ByteString.copyFrom({{ byteArrayLit $r.GetConst }});
{{- end -}}
{{- if $r.In }}
		private final com.google.protobuf.ByteString[] {{ constantName . "In" }} = new com.google.protobuf.ByteString[]{
			{{- range $r.In }}
			com.google.protobuf.ByteString.copyFrom({{ byteArrayLit . }}),
			{{- end }}
		};
{{- end -}}
{{- if $r.NotIn }}
		private final com.google.protobuf.ByteString[] {{ constantName . "NotIn" }} = new com.google.protobuf.ByteString[]{
			{{- range $r.NotIn }}
			com.google.protobuf.ByteString.copyFrom({{ byteArrayLit . }}),
			{{- end }}
		};
{{- end -}}
{{- if $r.Pattern }}
		private final com.google.re2j.Pattern {{ constantName . "Pattern" }} = com.google.re2j.Pattern.compile({{ javaStringEscape $r.GetPattern }});
{{- end -}}
{{- if $r.Prefix }}
		private final byte[] {{ constantName . "Prefix" }} = {{ byteArrayLit $r.GetPrefix }};
{{- end -}}
{{- if $r.Contains }}
		private final byte[] {{ constantName . "Contains" }} = {{ byteArrayLit $r.GetContains }};
{{- end -}}
{{- if $r.Suffix }}
		private final byte[] {{ constantName . "Suffix" }} = {{ byteArrayLit $r.GetSuffix }};
{{- end -}}`

const bytesTpl = `{{ $f := .Field }}{{ $r := .Rules -}}
{{- if $r.GetIgnoreEmpty }}
			if ( !{{ accessor . }}.isEmpty() ) {
{{- end -}}
{{- if $r.Const }}
			valctx.getValidationCollector().assertValid( ({{ safeName . "value"}}) -> {
				io.envoyproxy.pgv.ConstantValidation.constant("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ constantName . "Const" }});
			},proto);
{{- end -}}
{{- if $r.Len }}
	valctx.getValidationCollector().assertValid( ({{ safeName . "value"}}) -> {
		io.envoyproxy.pgv.BytesValidation.length("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ $r.GetLen }});
	},proto);
{{- end -}}
{{- if $r.MinLen }}
	valctx.getValidationCollector().assertValid( ({{ safeName . "value"}}) -> {
		io.envoyproxy.pgv.BytesValidation.minLength("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ $r.GetMinLen }});
	},proto);
{{- end -}}
{{- if $r.MaxLen }}
	valctx.getValidationCollector().assertValid( ({{ safeName . "value"}}) -> {
		io.envoyproxy.pgv.BytesValidation.maxLength("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ $r.GetMaxLen }});
	},proto);
{{- end -}}
{{- if $r.Pattern }}
	valctx.getValidationCollector().assertValid( ({{ safeName . "value"}}) -> {
		io.envoyproxy.pgv.BytesValidation.pattern("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ constantName . "Pattern" }});
	},proto);
{{- end -}}
{{- if $r.Prefix }}
	valctx.getValidationCollector().assertValid( ({{ safeName . "value"}}) -> {
		io.envoyproxy.pgv.BytesValidation.prefix("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ constantName . "Prefix" }});
	},proto);
{{- end -}}
{{- if $r.Contains }}
	valctx.getValidationCollector().assertValid( ({{ safeName . "value"}}) -> {
		io.envoyproxy.pgv.BytesValidation.contains("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ constantName . "Contains" }});
	},proto);
{{- end -}}
{{- if $r.Suffix }}
	valctx.getValidationCollector().assertValid( ({{ safeName . "value"}}) -> {
		io.envoyproxy.pgv.BytesValidation.suffix("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ constantName . "Suffix" }});
	},proto);
{{- end -}}
{{- if $r.GetIp }}
	valctx.getValidationCollector().assertValid( ({{ safeName . "value"}}) -> {
		io.envoyproxy.pgv.BytesValidation.ip("{{ $f.FullyQualifiedName }}", {{ accessor . }});
	},proto);
{{- end -}}
{{- if $r.GetIpv4 }}
	valctx.getValidationCollector().assertValid( ({{ safeName . "value"}}) -> {
		io.envoyproxy.pgv.BytesValidation.ipv4("{{ $f.FullyQualifiedName }}", {{ accessor . }});
	},proto);
{{- end -}}
{{- if $r.GetIpv6 }}
	valctx.getValidationCollector().assertValid( ({{ safeName . "value"}}) -> {
		io.envoyproxy.pgv.BytesValidation.ipv6("{{ $f.FullyQualifiedName }}", {{ accessor . }});
	},proto);
{{- end -}}
{{- if $r.In }}
	valctx.getValidationCollector().assertValid( ({{ safeName . "value"}}) -> {
		io.envoyproxy.pgv.CollectiveValidation.in("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ constantName . "In" }});
	},proto);
{{- end -}}
{{- if $r.NotIn }}
	valctx.getValidationCollector().assertValid( ({{ safeName . "value"}}) -> {
		io.envoyproxy.pgv.CollectiveValidation.notIn("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ constantName . "NotIn" }});
	},proto);
{{- end -}}
{{- if $r.GetIgnoreEmpty }}
			}
{{- end -}}
`
