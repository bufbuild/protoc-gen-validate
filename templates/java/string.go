package java

const stringTpl = `{{ $f := .Field }}{{ $r := .Rules -}}
{{- if $r.Const }}
			com.lyft.pgv.StringValidation.constant("{{ $f.FullyQualifiedName }}", proto.{{ accessor $f }}, "{{ $r.GetConst }}");
{{- end -}}
{{- if $r.In }}
			{
				{{ javaTypeFor $f }}[] set = new {{ javaTypeFor $f }}[]{
					{{- range $r.In -}}
						"{{- sprintf "%v" . -}}"{{ javaTypeLiteralSuffixFor $f }},
					{{- end -}}
				};
				com.lyft.pgv.StringValidation.in("{{ $f.FullyQualifiedName }}", proto.{{ accessor $f }}, set);
			}
{{- end -}}
{{- if $r.NotIn }}
			com.lyft.pgv.StringValidation.notIn("{{ $f.FullyQualifiedName }}", proto.{{ accessor $f }}, new {{ javaTypeFor $f }}[]{ });
{{- end -}}
{{- if $r.Len }}
			com.lyft.pgv.StringValidation.length("{{ $f.FullyQualifiedName }}", proto.{{ accessor $f }}, {{ $r.GetLen }});
{{- end -}}
{{- if $r.MinLen }}
			com.lyft.pgv.StringValidation.minLength("{{ $f.FullyQualifiedName }}", proto.{{ accessor $f }}, {{ $r.GetMinLen }});
{{- end -}}
{{- if $r.MaxLen }}
			com.lyft.pgv.StringValidation.maxLength("{{ $f.FullyQualifiedName }}", proto.{{ accessor $f }}, {{ $r.GetMaxLen }});
{{- end -}}
{{- if $r.LenBytes }}
			com.lyft.pgv.StringValidation.lenBytes("{{ $f.FullyQualifiedName }}", proto.{{ accessor $f }}, {{ $r.GetLen }});
{{- end -}}
{{- if $r.MinBytes }}
			com.lyft.pgv.StringValidation.minBytes("{{ $f.FullyQualifiedName }}", proto.{{ accessor $f }}, {{ $r.GetMinLen }});
{{- end -}}
{{- if $r.MaxBytes }}
			com.lyft.pgv.StringValidation.maxBytes("{{ $f.FullyQualifiedName }}", proto.{{ accessor $f }}, {{ $r.GetMaxLen }});
{{- end -}}
{{- if $r.Prefix }}
			com.lyft.pgv.StringValidation.prefix("{{ $f.FullyQualifiedName }}", proto.{{ accessor $f }}, "{{ $r.GetPrefix }}");
{{- end -}}
{{- if $r.Contains }}
			com.lyft.pgv.StringValidation.contains("{{ $f.FullyQualifiedName }}", proto.{{ accessor $f }}, "{{ $r.GetContains }}");
{{- end -}}
{{- if $r.Suffix }}
			com.lyft.pgv.StringValidation.suffix("{{ $f.FullyQualifiedName }}", proto.{{ accessor $f }}, "{{ $r.GetSuffix }}");
{{- end -}}
{{- if $r.GetEmail }}
			com.lyft.pgv.StringValidation.email("{{ $f.FullyQualifiedName }}", proto.{{ accessor $f }}, "{{ $r.GetEmail }}");
{{- end -}}
{{- if $r.GetHostname }}
			com.lyft.pgv.StringValidation.hostName("{{ $f.FullyQualifiedName }}", proto.{{ accessor $f }}, "{{ $r.GetHostname }}");
{{- end -}}
{{- if $r.GetIp }}
			com.lyft.pgv.StringValidation.ip("{{ $f.FullyQualifiedName }}", proto.{{ accessor $f }}, "{{ $r.GetIp }}");
{{- end -}}
{{- if $r.GetIpv4 }}
			com.lyft.pgv.StringValidation.ipv4("{{ $f.FullyQualifiedName }}", proto.{{ accessor $f }}, "{{ $r.GetIpv4 }}");
{{- end -}}
{{- if $r.GetIpv6 }}
			com.lyft.pgv.StringValidation.ipv6("{{ $f.FullyQualifiedName }}", proto.{{ accessor $f }}, "{{ $r.GetIpv6 }}");
{{- end -}}
{{- if $r.GetUri }}
			com.lyft.pgv.StringValidation.uri("{{ $f.FullyQualifiedName }}", proto.{{ accessor $f }}, "{{ $r.GetUri }}");
{{- end -}}
{{- if $r.GetUri }}
			com.lyft.pgv.StringValidation.uriRef("{{ $f.FullyQualifiedName }}", proto.{{ accessor $f }}, "{{ $r.GetUriRef }}");
{{- end -}}
`
