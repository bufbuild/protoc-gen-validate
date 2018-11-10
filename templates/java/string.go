package java

const stringConstTpl = `{{ $f := .Field }}{{ $r := .Rules -}}
{{- if $r.In }}
		private final {{ javaTypeFor $f }}[] {{ constantName . "In" }} = new {{ javaTypeFor $f }}[]{
			{{- range $r.In -}}
				"{{- sprintf "%v" . -}}",
			{{- end -}}
		};
{{- end -}}
{{- if $r.NotIn }}
		private final {{ javaTypeFor $f }}[] {{ constantName . "NotIn" }} = new {{ javaTypeFor $f }}[]{
			{{- range $r.NotIn -}}
				"{{- sprintf "%v" . -}}",
			{{- end -}}
		};
{{- end -}}
{{- if $r.Pattern }}
		com.google.re2j.Pattern {{ constantName . "Pattern" }} = com.google.re2j.Pattern.compile({{ javaStringEscape $r.GetPattern }});
{{- end -}}`

const stringTpl = `{{ $f := .Field }}{{ $r := .Rules -}}
{{- if $r.Const }}
			com.lyft.pgv.ConstantValidation.constant("{{ $f.FullyQualifiedName }}", {{ accessor . }}, "{{ $r.GetConst }}");
{{- end -}}
{{- if $r.In }}
			com.lyft.pgv.CollectiveValidation.in("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ constantName . "In" }});
{{- end -}}
{{- if $r.NotIn }}
			com.lyft.pgv.CollectiveValidation.notIn("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ constantName . "NotIn" }});
{{- end -}}
{{- if $r.Len }}
			com.lyft.pgv.StringValidation.length("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ $r.GetLen }});
{{- end -}}
{{- if $r.MinLen }}
			com.lyft.pgv.StringValidation.minLength("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ $r.GetMinLen }});
{{- end -}}
{{- if $r.MaxLen }}
			com.lyft.pgv.StringValidation.maxLength("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ $r.GetMaxLen }});
{{- end -}}
{{- if $r.LenBytes }}
			com.lyft.pgv.StringValidation.lenBytes("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ $r.GetLenBytes }});
{{- end -}}
{{- if $r.MinBytes }}
			com.lyft.pgv.StringValidation.minBytes("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ $r.GetMinBytes }});
{{- end -}}
{{- if $r.MaxBytes }}
			com.lyft.pgv.StringValidation.maxBytes("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ $r.GetMaxBytes }});
{{- end -}}
{{- if $r.Pattern }}
			com.lyft.pgv.StringValidation.pattern("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ constantName . "Pattern" }});
{{- end -}}
{{- if $r.Prefix }}
			com.lyft.pgv.StringValidation.prefix("{{ $f.FullyQualifiedName }}", {{ accessor . }}, "{{ $r.GetPrefix }}");
{{- end -}}
{{- if $r.Contains }}
			com.lyft.pgv.StringValidation.contains("{{ $f.FullyQualifiedName }}", {{ accessor . }}, "{{ $r.GetContains }}");
{{- end -}}
{{- if $r.Suffix }}
			com.lyft.pgv.StringValidation.suffix("{{ $f.FullyQualifiedName }}", {{ accessor . }}, "{{ $r.GetSuffix }}");
{{- end -}}
{{- if $r.GetEmail }}
			com.lyft.pgv.StringValidation.email("{{ $f.FullyQualifiedName }}", {{ accessor . }});
{{- end -}}
{{- if $r.GetHostname }}
			com.lyft.pgv.StringValidation.hostName("{{ $f.FullyQualifiedName }}", {{ accessor . }});
{{- end -}}
{{- if $r.GetIp }}
			com.lyft.pgv.StringValidation.ip("{{ $f.FullyQualifiedName }}", {{ accessor . }});
{{- end -}}
{{- if $r.GetIpv4 }}
			com.lyft.pgv.StringValidation.ipv4("{{ $f.FullyQualifiedName }}", {{ accessor . }});
{{- end -}}
{{- if $r.GetIpv6 }}
			com.lyft.pgv.StringValidation.ipv6("{{ $f.FullyQualifiedName }}", {{ accessor . }});
{{- end -}}
{{- if $r.GetUri }}
			com.lyft.pgv.StringValidation.uri("{{ $f.FullyQualifiedName }}", {{ accessor . }});
{{- end -}}
{{- if $r.GetUriRef }}
			com.lyft.pgv.StringValidation.uriRef("{{ $f.FullyQualifiedName }}", {{ accessor . }});
{{- end -}}
`
