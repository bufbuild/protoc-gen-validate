package python

const strTpl = `
	{{ $f := .Field }}{{ $r := .Rules }}
	{{ template "const" . }}
	{{ template "in" . }}

	{{ if or $r.Len (and $r.MinLen $r.MaxLen (eq $r.GetMinLen $r.GetMaxLen)) }}
		{{ if $r.Len }}
	if unicode_len({{ accessor . }}) != {{ $r.GetLen }}:
		{{ err . "value length must be " $r.GetLen " chars" }}
		{{ else }}
	if unicode_len({{ accessor . }}) != {{ $r.GetMinLen }}:
		{{ err . "value length must be " $r.GetMinLen " chars" }}
		{{ end }}
	{{ else if $r.MinLen }}
		{{ if $r.MaxLen }}
	if unicode_len({{ accessor . }}) < {{ $r.GetMinLen }} or unicode_len({{ accessor . }}) > {{ $r.GetMaxLen }}:
		{{ err . "value length must be between " $r.GetMinLen " and " $r.GetMaxLen " runes, inclusive" }}
		{{ else }}
	if unicode_len({{ accessor . }}) < {{ $r.GetMinLen }}:
		{{ err . "value length must be at least " $r.GetMinLen " runes" }}
		{{ end }}
	{{ else if $r.MaxLen }}
	if unicode_len({{ accessor . }}) > {{ $r.GetMaxLen }}:
		{{ err . "value length must be at most " $r.GetMaxLen " runes" }}
	{{ end }}

	{{ if or $r.LenBytes (and $r.MinBytes $r.MaxBytes (eq $r.GetMinBytes $r.GetMaxBytes)) }}
		{{ if $r.LenBytes }}
	if byte_length({{ accessor . }}) != {{ $r.GetLenBytes }}:
		{{ err . "value length must be " $r.GetLenBytes " bytes" }}
		{{ else }}
	if byte_length({{ accessor . }}) != {{ $r.GetMinBytes }}:
		{{ err . "value length must be " $r.GetMinBytes " bytes" }}
		{{ end }}
	{{ else if $r.MinBytes }}
		{{ if $r.MaxBytes }}
	if byte_length({{ accessor . }}) < {{ $r.GetMinBytes }} or byte_length({{ accessor . }}) > {{ $r.GetMaxBytes }}:
		{{ err . "value length must be between " $r.GetMinBytes " and " $r.GetMaxBytes " bytes, inclusive" }}
		{{ else }}
	if byte_length({{ accessor . }}) < {{ $r.GetMinBytes }}:
		{{ err . "value length must be at least " $r.GetMinBytes " bytes" }}
		{{ end }}
	{{ else if $r.MaxBytes }}
	if byte_length({{ accessor . }}) > {{ $r.GetMaxBytes }}:
		{{ err . "value length must be at most " $r.GetMaxBytes " bytes" }}
	{{ end }}

	{{ if $r.Prefix }}
	if not {{ accessor . }}.startswith( {{ lit $r.GetPrefix }} ):
		{{ err . "value does not have prefix " (lit $r.GetPrefix) }}
	{{ end }}

	{{ if $r.Suffix }}
	if not {{ accessor . }}.endswith( {{ lit $r.GetSuffix }} ):
		{{ err . "value does not have prefix " (lit $r.GetPrefix) }}
	{{ end }}

	{{ if $r.Contains }}
	if {{ lit $r.GetContains }} not in {{ accessor . }}:
		{{ err . "value does not contain substring " (lit $r.GetContains) }}
	{{ end }}


	{{ if $r.Pattern }}
	if {{ lookup $f "Pattern" }}.match({{ accessor . }}) is None:
		{{ err . "value does not match regex pattern " (lit $r.GetPattern) }}
	{{ end }}
`

const wip = `
	{{ if $r.GetIp }}
		if ip := net.ParseIP({{ accessor . }}); ip == nil {
			return {{ err . "value must be a valid IP address" }}
		}
	{{ else if $r.GetIpv4 }}
		if ip := net.ParseIP({{ accessor . }}); ip == nil || ip.To4() == nil {
			return {{ err . "value must be a valid IPv4 address" }}
		}
	{{ else if $r.GetIpv6 }}
		if ip := net.ParseIP({{ accessor . }}); ip == nil || ip.To4() != nil {
			return {{ err . "value must be a valid IPv6 address" }}
		}
	{{ else if $r.GetEmail }}
		if err := m._validateEmail({{ accessor . }}); err != nil {
			return {{ errCause . "err" "value must be a valid email address" }}
		}
	{{ else if $r.GetHostname }}
		if err := m._validateHostname({{ accessor . }}); err != nil {
			return {{ errCause . "err" "value must be a valid hostname" }}
		}
	{{ else if $r.GetUri }}
		if uri, err := url.Parse({{ accessor . }}); err != nil {
			return {{ errCause . "err" "value must be a valid URI" }}
		} else if !uri.IsAbs() {
			return {{ err . "value must be absolute" }}
		}
	{{ else if $r.GetUriRef }}
		if _, err := url.Parse({{ accessor . }}); err != nil {
			return {{ errCause . "err" "value must be a valid URI" }}
		}
	{{ end }}
`