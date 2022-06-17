package java

const stringConstTpl = `{{ $f := .Field }}{{ $r := .Rules -}}
{{- if $r.In }}
		private final {{ javaTypeFor . }}[] {{ constantName . "In" }} = new {{ javaTypeFor . }}[]{
			{{- range $r.In -}}
				"{{- sprintf "%v" . -}}",
			{{- end -}}
		};
{{- end -}}
{{- if $r.NotIn }}
		private final {{ javaTypeFor . }}[] {{ constantName . "NotIn" }} = new {{ javaTypeFor . }}[]{
			{{- range $r.NotIn -}}
				"{{- sprintf "%v" . -}}",
			{{- end -}}
		};
{{- end -}}
{{- if $r.Pattern }}
		com.google.re2j.Pattern {{ constantName . "Pattern" }} = com.google.re2j.Pattern.compile({{ javaStringEscape $r.GetPattern }});
{{- end -}}`

const stringTpl = `{{ $f := .Field }}{{ $r := .Rules -}}
{{- if $r.GetIgnoreEmpty }}
			if ( !{{ accessor . }}.isEmpty() ) {
{{- end -}}
{{- if $r.Const }}
			valctx.getValidatorInterceptor().validate( ({{ safeName . "value"}}) -> {
				io.envoyproxy.pgv.ConstantValidation.constant("{{ $f.FullyQualifiedName }}", {{ accessor . }}, "{{ $r.GetConst }}");
			},proto);
{{- end -}}
{{- if $r.In }}
			valctx.getValidatorInterceptor().validate( ({{ safeName . "value"}}) -> {
				io.envoyproxy.pgv.CollectiveValidation.in("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ constantName . "In" }});
			},proto);
{{- end -}}
{{- if $r.NotIn }}
			valctx.getValidatorInterceptor().validate( ({{ safeName . "value"}}) -> {
				io.envoyproxy.pgv.CollectiveValidation.notIn("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ constantName . "NotIn" }});
			},proto);
{{- end -}}
{{- if $r.Len }}
			valctx.getValidatorInterceptor().validate( ({{ safeName . "value"}}) -> {
				io.envoyproxy.pgv.StringValidation.length("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ $r.GetLen }});
			},proto);
{{- end -}}
{{- if $r.MinLen }}
			valctx.getValidatorInterceptor().validate( ({{ safeName . "value"}}) -> {
				io.envoyproxy.pgv.StringValidation.minLength("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ $r.GetMinLen }});
			},proto);
{{- end -}}
{{- if $r.MaxLen }}
			valctx.getValidatorInterceptor().validate( ({{ safeName . "value"}}) -> {
				io.envoyproxy.pgv.StringValidation.maxLength("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ $r.GetMaxLen }});
			},proto);
{{- end -}}
{{- if $r.LenBytes }}
			valctx.getValidatorInterceptor().validate( ({{ safeName . "value"}}) -> {
				io.envoyproxy.pgv.StringValidation.lenBytes("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ $r.GetLenBytes }});
			},proto);
{{- end -}}
{{- if $r.MinBytes }}
			valctx.getValidatorInterceptor().validate( ({{ safeName . "value"}}) -> {
				io.envoyproxy.pgv.StringValidation.minBytes("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ $r.GetMinBytes }});
			},proto);
{{- end -}}
{{- if $r.MaxBytes }}
			valctx.getValidatorInterceptor().validate( ({{ safeName . "value"}}) -> {
				io.envoyproxy.pgv.StringValidation.maxBytes("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ $r.GetMaxBytes }});
			},proto);
{{- end -}}
{{- if $r.Pattern }}
			valctx.getValidatorInterceptor().validate( ({{ safeName . "value"}}) -> {
				io.envoyproxy.pgv.StringValidation.pattern("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ constantName . "Pattern" }});
			},proto);
{{- end -}}
{{- if $r.Prefix }}
			valctx.getValidatorInterceptor().validate( ({{ safeName . "value"}}) -> {
				io.envoyproxy.pgv.StringValidation.prefix("{{ $f.FullyQualifiedName }}", {{ accessor . }}, "{{ $r.GetPrefix }}");
			},proto);
{{- end -}}
{{- if $r.Contains }}
			valctx.getValidatorInterceptor().validate( ({{ safeName . "value"}}) -> {
				io.envoyproxy.pgv.StringValidation.contains("{{ $f.FullyQualifiedName }}", {{ accessor . }}, "{{ $r.GetContains }}");
			},proto);
{{- end -}}
{{- if $r.NotContains }}
			valctx.getValidatorInterceptor().validate( ({{ safeName . "value"}}) -> {
				io.envoyproxy.pgv.StringValidation.notContains("{{ $f.FullyQualifiedName }}", {{ accessor . }}, "{{ $r.GetNotContains }}");
			},proto);
{{- end -}}
{{- if $r.Suffix }}
			valctx.getValidatorInterceptor().validate( ({{ safeName . "value"}}) -> {
				io.envoyproxy.pgv.StringValidation.suffix("{{ $f.FullyQualifiedName }}", {{ accessor . }}, "{{ $r.GetSuffix }}");
			},proto);
{{- end -}}
{{- if $r.GetEmail }}
			valctx.getValidatorInterceptor().validate( ({{ safeName . "value"}}) -> {
				io.envoyproxy.pgv.StringValidation.email("{{ $f.FullyQualifiedName }}", {{ accessor . }});
			},proto);
{{- end -}}
{{- if $r.GetAddress }}
			valctx.getValidatorInterceptor().validate( ({{ safeName . "value"}}) -> {
				io.envoyproxy.pgv.StringValidation.address("{{ $f.FullyQualifiedName }}", {{ accessor . }});
			},proto);
{{- end -}}
{{- if $r.GetHostname }}
			valctx.getValidatorInterceptor().validate( ({{ safeName . "value"}}) -> {
				io.envoyproxy.pgv.StringValidation.hostName("{{ $f.FullyQualifiedName }}", {{ accessor . }});
			},proto);
{{- end -}}
{{- if $r.GetIp }}
			valctx.getValidatorInterceptor().validate( ({{ safeName . "value"}}) -> {
				io.envoyproxy.pgv.StringValidation.ip("{{ $f.FullyQualifiedName }}", {{ accessor . }});
			},proto);
{{- end -}}
{{- if $r.GetIpv4 }}
			valctx.getValidatorInterceptor().validate( ({{ safeName . "value"}}) -> {
				io.envoyproxy.pgv.StringValidation.ipv4("{{ $f.FullyQualifiedName }}", {{ accessor . }});
			},proto);
{{- end -}}
{{- if $r.GetIpv6 }}
			valctx.getValidatorInterceptor().validate( ({{ safeName . "value"}}) -> {
				io.envoyproxy.pgv.StringValidation.ipv6("{{ $f.FullyQualifiedName }}", {{ accessor . }});
			},proto);
{{- end -}}
{{- if $r.GetUri }}
			valctx.getValidatorInterceptor().validate( ({{ safeName . "value"}}) -> {
				io.envoyproxy.pgv.StringValidation.uri("{{ $f.FullyQualifiedName }}", {{ accessor . }});
			},proto);
{{- end -}}
{{- if $r.GetUriRef }}
			valctx.getValidatorInterceptor().validate( ({{ safeName . "value"}}) -> {
				io.envoyproxy.pgv.StringValidation.uriRef("{{ $f.FullyQualifiedName }}", {{ accessor . }});
			},proto);
{{- end -}}
{{- if $r.GetUuid }}
			valctx.getValidatorInterceptor().validate( ({{ safeName . "value"}}) -> {
				io.envoyproxy.pgv.StringValidation.uuid("{{ $f.FullyQualifiedName }}", {{ accessor . }});
			},proto);
{{- end -}}
{{- if $r.GetIgnoreEmpty }}
			}
{{- end -}}
`
