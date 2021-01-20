package php

const stringConstTpl = `{{ $f := .Field }}{{ $r := .Rules -}}
{{- if $r.In }}
		${{ constantName . "In" }} = new Assert\Choice([
			'choices' => [
				{{- range $r.In -}}
					"{{- sprintf "%v" . -}}",
				{{- end -}}
			],
			
		]);
{{- end -}}
{{- if $r.NotIn }}
		${{ constantName . "NotIn" }} = new Assert\Choice([
			'choices' => [
				{{- range $r.In -}}
					"{{- sprintf "%v" . -}}",
				{{- end -}}
			],
			
		]);
{{- end -}}
{{- if $r.Pattern }}
		${{ constantName . "Pattern" }} = new Assert\Regex([
            'pattern' => '{{ phpStringEscape $r.GetPattern }}'
        ];
{{- end -}}`

const stringTpl = `{{ $f := .Field }}{{ $r := .Rules -}}
{{- if $r.Const }}
			io.envoyproxy.pgv.ConstantValidation.constant("{{ $f.FullyQualifiedName }}", {{ accessor . }}, "{{ $r.GetConst }}");
{{- end -}}
{{- if $r.In }}
			$metadata->addPropertyConstraint("{{ $f.FullyQualifiedName }}", {{ constantName . "In" }});
{{- end -}}
{{- if $r.NotIn }}
			io.envoyproxy.pgv.CollectiveValidation.notIn("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ constantName . "NotIn" }});
{{- end -}}
{{- if $r.Len }}
			io.envoyproxy.pgv.StringValidation.length("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ $r.GetLen }});
{{- end -}}
{{- if $r.MinLen }}
			io.envoyproxy.pgv.StringValidation.minLength("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ $r.GetMinLen }});
{{- end -}}
{{- if $r.MaxLen }}
			io.envoyproxy.pgv.StringValidation.maxLength("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ $r.GetMaxLen }});
{{- end -}}
{{- if $r.LenBytes }}
			io.envoyproxy.pgv.StringValidation.lenBytes("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ $r.GetLenBytes }});
{{- end -}}
{{- if $r.MinBytes }}
			io.envoyproxy.pgv.StringValidation.minBytes("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ $r.GetMinBytes }});
{{- end -}}
{{- if $r.MaxBytes }}
			io.envoyproxy.pgv.StringValidation.maxBytes("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ $r.GetMaxBytes }});
{{- end -}}
{{- if $r.Pattern }}
			$metadata->addPropertyConstraint("{{ $f.FullyQualifiedName }}", ${{ constantName . "Pattern" }});
{{- end -}}
{{- if $r.Prefix }}
			io.envoyproxy.pgv.StringValidation.prefix("{{ $f.FullyQualifiedName }}", {{ accessor . }}, "{{ $r.GetPrefix }}");
{{- end -}}
{{- if $r.Contains }}
			io.envoyproxy.pgv.StringValidation.contains("{{ $f.FullyQualifiedName }}", {{ accessor . }}, "{{ $r.GetContains }}");
{{- end -}}
{{- if $r.NotContains }}
			io.envoyproxy.pgv.StringValidation.notContains("{{ $f.FullyQualifiedName }}", {{ accessor . }}, "{{ $r.GetNotContains }}");
{{- end -}}
{{- if $r.Suffix }}
			io.envoyproxy.pgv.StringValidation.suffix("{{ $f.FullyQualifiedName }}", {{ accessor . }}, "{{ $r.GetSuffix }}");
{{- end -}}
{{- if $r.GetEmail }}
			$metadata->addPropertyConstraint('{{ $f.FullyQualifiedName }}', new Assert\Email([
				'message' => 'The email is not a valid.',
			]));
{{- end -}}
{{- if $r.GetAddress }}
			io.envoyproxy.pgv.StringValidation.address("{{ $f.FullyQualifiedName }}", {{ accessor . }});
{{- end -}}
{{- if $r.GetHostname }}
			io.envoyproxy.pgv.StringValidation.hostName("{{ $f.FullyQualifiedName }}", {{ accessor . }});
{{- end -}}
{{- if $r.GetIp }}
			io.envoyproxy.pgv.StringValidation.ip("{{ $f.FullyQualifiedName }}", {{ accessor . }});
{{- end -}}
{{- if $r.GetIpv4 }}
			io.envoyproxy.pgv.StringValidation.ipv4("{{ $f.FullyQualifiedName }}", {{ accessor . }});
{{- end -}}
{{- if $r.GetIpv6 }}
			io.envoyproxy.pgv.StringValidation.ipv6("{{ $f.FullyQualifiedName }}", {{ accessor . }});
{{- end -}}
{{- if $r.GetUri }}
			io.envoyproxy.pgv.StringValidation.uri("{{ $f.FullyQualifiedName }}", {{ accessor . }});
{{- end -}}
{{- if $r.GetUriRef }}
			io.envoyproxy.pgv.StringValidation.uriRef("{{ $f.FullyQualifiedName }}", {{ accessor . }});
{{- end -}}
{{- if $r.GetUuid }}
			io.envoyproxy.pgv.StringValidation.uuid("{{ $f.FullyQualifiedName }}", {{ accessor . }});
{{- end -}}
`
