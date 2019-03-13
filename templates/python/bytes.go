package python

const bytesTpl = `
	{{ $f := .Field }}{{ $r := .Rules }}
	{{ template "const" . }}
	{{ template "in" . }}

	{{ if or $r.Len (and $r.MinLen $r.MaxLen (eq $r.GetMinLen $r.GetMaxLen)) }}
		{{ if $r.Len }}
	if len({{ accessor . }}) != {{ $r.GetLen }}:
		{{ err . "value length must be " $r.GetLen " bytes" }}
		{{ else }}
	if len({{ accessor . }}) != {{ $r.GetMinLen }}:
		{{ err . "value length must be " $r.GetMinLen " bytes" }}
		{{ end }}
	{{ else if $r.MinLen }}
		{{ if $r.MaxLen }}
	if len({{ accessor . }}) < {{ $r.GetMinLen }} or len({{ accessor . }}) > {{ $r.GetMaxLen }}:
		{{ err . "value length must be between " $r.GetMinLen " and " $r.GetMaxLen " bytes, inclusive" }}
		{{ else }}
	if len({{ accessor . }}) < {{ $r.GetMinLen }}:
		{{ err . "value length must be at least " $r.GetMinLen " bytes" }}
		{{ end }}
	{{ else if $r.MaxLen }}
	if len({{ accessor . }}) > {{ $r.GetMaxLen }}:
		{{ err . "value length must be at most " $r.GetMaxLen " bytes" }}
	{{ end }}

	{{ if $r.Prefix }}
	if not {{ accessor . }}.startswith({{ lit $r.GetPrefix }}):
		{{ err . "value does not have prefix " (byteStr $r.GetPrefix) }}
	{{ end }}

	{{ if $r.Suffix }}
	if not {{ accessor . }}.endswith({{ lit $r.GetSuffix }}):
		{{ err . "value does not have suffix " (byteStr $r.GetSuffix) }}
	{{ end }}

	{{ if $r.Contains }}
	if {{ lit $r.GetContains }} not in {{ accessor . }}:
		{{ err . "value does not contain " (byteStr $r.GetContains) }}
	{{ end }}

	{{ if $r.Pattern }}
	if {{ lookup $f "Pattern" }}.match({{ accessor . }}) is None:
		{{ err . "value does not match regex pattern " (lit $r.GetPattern) }}
	{{ end }}
`

const wipp = `
	{{ if $r.GetIp }}
		if ip := net.IP({{ accessor . }}); ip.To16() == nil {
			return {{ err . "value must be a valid IP address" }}
		}
	{{ else if $r.GetIpv4 }}
		if ip := net.IP({{ accessor . }}); ip.To4() == nil {
			return {{ err . "value must be a valid IPv4 address" }}
		}
	{{ else if $r.GetIpv6 }}
		if ip := net.IP({{ accessor . }}); ip.To16() == nil || ip.To4() != nil {
			return {{ err . "value must be a valid IPv6 address" }}
		}
	{{ end }}
`
