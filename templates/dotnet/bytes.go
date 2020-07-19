package dotnet

const bytesConstTpl = `
{{- $f := .Field -}}{{- $r := .Rules -}}
{{ if $r.Const }}
	private static readonly Google.Protobuf.ByteString {{ constant $f "Constant" }} = {{ literal $r.GetConst }};
{{ end }}
{{ if $r.In }}
	private static readonly HashSet<Google.Protobuf.ByteString> {{ constant $f "In" }} = new HashSet<Google.Protobuf.ByteString>
	{
		{{ range $r.In }}
			{{ literal . }},
		{{ end }}
	};
{{ end }}
{{ if $r.NotIn }}
	private static readonly HashSet<Google.Protobuf.ByteString> {{ constant $f "NotIn" }} = new HashSet<Google.Protobuf.ByteString>
	{
		{{ range $r.NotIn }}
			{{ literal . }},
		{{ end }}
	};
{{ end }}
{{ if $r.Prefix }}
	private static readonly Google.Protobuf.ByteString {{ constant $f "Prefix" }} = {{ literal $r.GetPrefix }};
{{ end }}
{{ if $r.Suffix }}
	private static readonly Google.Protobuf.ByteString {{ constant $f "Suffix" }} = {{ literal $r.GetSuffix }};
{{ end }}
{{ if $r.Contains }}
	private static readonly Google.Protobuf.ByteString {{ constant $f "Contains" }} = {{ literal $r.GetContains }};
{{ end }}
{{ if $r.Pattern }}
	private static readonly IronRe2.Regex {{ constant $f "Pattern" }} = new IronRe2.Regex({{ literal $r.GetPattern }});
{{ end }}
`

const bytesTpl = `
{{- $f := .Field -}}{{- $r := .Rules -}}
{{ if $r.Const }}
	if ({{ accessor . }} != {{ constant $f "Constant" }})
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
		throw {{ err . "length must be " $r.GetLen " byte(s)" }};
{{ else if and $r.MinLen $r.MaxLen (eq $r.GetMinLen $r.GetMaxLen) }}
	if ({{ accessor . }}.Length != {{ $r.GetMinLen }})
		throw {{ err . "length must be " $r.GetMinLen " byte(s)" }};
{{ else if $r.MinLen }}
	{{ if $r.MaxLen }}
		{
			var length = {{ accessor . }}.Length;
			if (length < {{ $r.GetMinLen }} || length > {{ $r.GetMaxLen }})
				throw {{ err . "length must be between " $r.GetMinLen " and " $r.GetMaxLen " bytes" }};
		}
	{{ else }}
		if ({{ accessor . }}.Length < {{ $r.GetMinLen }})
			throw {{ err . "length must be at least " $r.GetMinLen " byte(s)" }};
	{{ end }}
{{ else if $r.MaxLen }}
	if ({{ accessor . }}.Length > {{ $r.GetMaxLen }})
		throw {{ err . "length must be at most " $r.GetMaxLen " byte(s)" }};
{{ end }}
{{ if $r.Prefix }}
	if (!{{ accessor . }}.Span.StartsWith({{ constant $f "Prefix" }}.Span))
		throw {{ err . "value must start with " (byteStr $r.GetPrefix) }};
{{ end }}
{{ if $r.Suffix }}
	if (!{{ accessor . }}.Span.EndsWith({{ constant $f "Suffix" }}.Span))
		throw {{ err . "value must end with " (byteStr $r.GetSuffix) }};
{{ end }}
{{ if $r.Contains }}
	// ReadOnlySpan<T>.Contains requires netstandard2.1
	if ({{ accessor . }}.Span.IndexOf({{ constant $f "Contains" }}.Span) == -1)
		throw {{ err . "value must contain " (byteStr $r.GetContains) }};
{{ end }}
{{ if $r.GetIp }}
	{
		var length = {{ accessor . }}.Length;
		if (length != 4 && length != 16)
			throw {{ err . "value must be a valid IP address" }};
	}
{{ else if $r.GetIpv4 }}
	if ({{ accessor . }}.Length != 4)
		throw {{ err . "value must be a valid IPv4 address" }};
{{ else if $r.GetIpv6 }}
	if ({{ accessor . }}.Length != 16)
		throw {{ err . "value must be a valid IPv6 address" }};
{{ end }}
{{ if $r.Pattern }}
	if (!{{ constant $f "Pattern" }}.IsMatch({{ accessor . }}.ToStringUtf8()))
		throw {{ err . "value must match pattern " (literal $r.GetPattern) }};
{{ end }}
`
