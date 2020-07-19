package dotnet

const stringConstTpl = `
{{- $f := .Field -}}{{- $r := .Rules -}}
{{ if $r.In }}
	private static readonly HashSet<string> {{ constant $f "In" }} = new HashSet<string>
	{
		{{ range $r.In }}
			{{ literal . }},
		{{ end }}
	};
{{ end }}
{{ if $r.NotIn }}
	private static readonly HashSet<string> {{ constant $f "NotIn" }} = new HashSet<string>
	{
		{{ range $r.NotIn }}
			{{ literal . }},
		{{ end }}
	};
{{ end }}
{{ if $r.Pattern }}
	private static readonly IronRe2.Regex {{ constant $f "Pattern" }} = new IronRe2.Regex({{ literal $r.GetPattern }});
{{ end }}
`

const stringTpl = `
{{- $f := .Field -}}{{- $r := .Rules -}}
{{ if $r.Const }}
	if ({{ accessor . }} != {{ literal $r.GetConst }})
		throw {{ err . "value must be equal " $r.GetConst }};
{{ end }}
{{ if $r.In }}
	if (!{{ constant $f "In" }}.Contains({{ accessor . }}))
		throw {{ err . "value must be in list " $r.In }};
{{ end }}
{{ if $r.NotIn }}
	if ({{ constant $f "NotIn" }}.Contains({{ accessor . }}))
		throw {{ err . "value must not be in list " $r.NotIn }};
{{ end }}
{{ if $r.Len }}
	if ({{ accessor . }}.Length != {{ $r.GetLen }})
		throw {{ err . "length must be " $r.GetLen " rune(s)" }};
{{ else if and $r.MinLen $r.MaxLen (eq $r.GetMinLen $r.GetMaxLen) }}
	if ({{ accessor . }}.Length != {{ $r.GetMinLen }})
		throw {{ err . "length must be " $r.GetMinLen " rune(s)" }};
{{ else if $r.MinLen }}
	{{ if $r.MaxLen }}
		{
			var length = {{ accessor . }}.Length;
			if (length < {{ $r.GetMinLen }} || length > {{ $r.GetMaxLen }})
				throw {{ err . "length must be between " $r.GetMinLen " and " $r.GetMaxLen " runes" }};
		}
	{{ else }}
		if ({{ accessor . }}.Length < {{ $r.GetMinLen }})
			throw {{ err . "length must be  at least " $r.GetMinLen " rune(s)" }};
	{{ end }}
{{ else if $r.MaxLen }}
	if ({{ accessor . }}.Length > {{ $r.GetMaxLen }})
		throw {{ err . "length must be  at most " $r.GetMaxLen " rune(s)" }};
{{ end }}
{{ if $r.LenBytes }}
	if (Encoding.UTF8.GetByteCount({{ accessor . }}) != {{ $r.LenBytes }})
		throw {{ err . "length must be " $r.LenBytes " bytes(s)" }};
{{ else if and $r.MinBytes $r.MaxBytes (eq $r.GetMinBytes $r.GetMaxBytes) }}
	if (Encoding.UTF8.GetByteCount({{ accessor . }}) != {{ $r.GetMinBytes }})
		throw {{ err . "length must be " $r.GetMinBytes " bytes(s)" }};
{{ else if $r.MinBytes }}
	{{ if $r.MaxBytes }}
		{
			var length = Encoding.UTF8.GetByteCount({{ accessor . }});
			if (length < {{ $r.GetMinBytes }} || length > {{ $r.GetMaxBytes }})
				throw {{ err . "length must be between " $r.GetMinBytes " and " $r.GetMaxBytes " bytes" }};
		}
	{{ else }}
		if (Encoding.UTF8.GetByteCount({{ accessor . }}) < {{ $r.GetMinBytes }})
			throw {{ err . "length must be at least " $r.GetMinBytes " bytes(s)" }};
	{{ end }}
{{ else if $r.MaxBytes }}
	if (Encoding.UTF8.GetByteCount({{ accessor . }}) > {{ $r.GetMaxBytes }})
		throw {{ err . "length must be at most " $r.GetMaxBytes " bytes(s)" }};
{{ end }}
{{ if $r.Prefix }}
	if (!{{ accessor . }}.StartsWith({{ literal $r.GetPrefix }}))
		throw {{ err . "value must start with " $r.GetPrefix }};
{{ end }}
{{ if $r.Suffix }}
	if (!{{ accessor . }}.EndsWith({{ literal $r.GetSuffix }}))
		throw {{ err . "value must end with " $r.GetSuffix }};
{{ end }}
{{ if $r.Contains }}
	if (!{{ accessor . }}.Contains({{ literal $r.GetContains }}))
		throw {{ err . "value must contain " $r.GetContains }};
{{ end }}
{{ if $r.NotContains }}
	if ({{ accessor . }}.Contains({{ literal $r.GetNotContains }}))
		throw {{ err . "value must not contain " $r.GetNotContains }};
{{ end }}
{{ if $r.GetEmail }}
	if (!Envoyproxy.Validator.StringValidation.ValidateEmail({{ accessor . }}))
		throw {{ err . "value must be a valid email address" }};
{{ else if $r.GetAddress }}
	if (!Envoyproxy.Validator.StringValidation.ValidateAddress({{ accessor . }}))
		throw {{ err . "value must be a valid host name or IP address" }};
{{ else if $r.GetHostname }}
	if (!Envoyproxy.Validator.StringValidation.ValidateHostname({{ accessor . }}))
		throw {{ err . "value must be a valid host name" }};
{{ else if $r.GetIp }}
	if (!Envoyproxy.Validator.StringValidation.ValidateIP({{ accessor . }}))
		throw {{ err . "value must be a valid IP address" }};
{{ else if $r.GetIpv4 }}
	if (!Envoyproxy.Validator.StringValidation.ValidateIPv4({{ accessor . }}))
		throw {{ err . "value must be a valid IPv4 address" }};
{{ else if $r.GetIpv6 }}
	if (!Envoyproxy.Validator.StringValidation.ValidateIPv6({{ accessor . }}))
		throw {{ err . "value must be a valid IPv4 address" }};
{{ else if $r.GetUri }}
	if (!Envoyproxy.Validator.StringValidation.ValidateUri({{ accessor . }}))
		throw {{ err . "value must be an absolute URI" }};
{{ else if $r.GetUriRef }}
	if (!Envoyproxy.Validator.StringValidation.ValidateUriRef({{ accessor . }}))
		throw {{ err . "value must be an URI" }};
{{ else if $r.GetUuid }}
	if (!Envoyproxy.Validator.StringValidation.ValidateUuid({{ accessor . }}))
		throw {{ err . "value must be a UUID" }};
{{ end }}
{{ if $r.Pattern }}
	if (!{{ constant $f "Pattern" }}.IsMatch({{ accessor . }}))
		throw {{ err . "value must match pattern " (literal $r.GetPattern) }};
{{ end }}
`
