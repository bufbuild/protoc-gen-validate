package goshared

const strTpl = `
	{{ $f := .Field }}{{ $r := .Rules }}
	{{ template "const" . }}
	{{ template "in" . }}

	{{ if or $r.Len (and $r.MinLen $r.MaxLen (eq $r.GetMinLen $r.GetMaxLen)) }}
		{{ if $r.Len }}
			if utf8.RuneCountInString({{ accessor . }}) != {{ $r.GetLen }} {
			return {{ err . (t "string.len" "value length must be {{$1}} runes" $r.GetLen) }}
		{{ else }}
		if utf8.RuneCountInString({{ accessor . }}) != {{ $r.GetMinLen }} {
			return {{ err . (t "string.len" "value length must be {{$1}} runes" $r.GetMinLen) }}
		{{ end }}
	}
	{{ else if $r.MinLen }}
		{{ if $r.MaxLen }}
			if l := utf8.RuneCountInString({{ accessor . }}); l < {{ $r.GetMinLen }} || l > {{ $r.GetMaxLen }} {
				return {{ err . (t "string.len_between" "value length must be between {{$1}} and {{$2}} runes, inclusive" $r.GetMinLen $r.GetMaxLen) }}
			}
		{{ else }}
			if utf8.RuneCountInString({{ accessor . }}) < {{ $r.GetMinLen }} {
				return {{ err . (t "string.min_len" "value length must be at least {{$1}} runes" $r.GetMinLen) }}
			}
		{{ end }}
	{{ else if $r.MaxLen }}
		if utf8.RuneCountInString({{ accessor . }}) > {{ $r.GetMaxLen }} {
			return {{ err . (t "string.max_len" "value length must be at most {{$1}} runes" $r.GetMaxLen) }}
		}
	{{ end }}

	{{ if or $r.LenBytes (and $r.MinBytes $r.MaxBytes (eq $r.GetMinBytes $r.GetMaxBytes)) }}
		{{ if $r.LenBytes }}
			if len({{ accessor . }}) != {{ $r.GetLenBytes }} {
				return {{ err . (t "string.len_bytes" "value length must be {{$1}} bytes" $r.GetLenBytes) }}
			}
		{{ else }}
			if len({{ accessor . }}) != {{ $r.GetMinBytes }} {
				return {{ err . (t "string.len_bytes" "value length must be {{$1}} bytes" $r.GetMinBytes) }}
			}
		{{ end }}
	{{ else if $r.MinBytes }}
		{{ if $r.MaxBytes }}
			if l := len({{ accessor . }}); l < {{ $r.GetMinBytes }} || l > {{ $r.GetMaxBytes }} {
					return {{ err . (t "string.bytes_between" "value length must be between {{$1}} and {{$2}} bytes, inclusive" $r.GetMinBytes $r.GetMaxBytes) }}
			}
		{{ else }}
			if len({{ accessor . }}) < {{ $r.GetMinBytes }} {
				return {{ err . (t "string.min_bytes" "value length must be at least {{$1}} bytes" $r.GetMinBytes) }}
			}
		{{ end }}
	{{ else if $r.MaxBytes }}
		if len({{ accessor . }}) > {{ $r.GetMaxBytes }} {
			return {{ err . (t "string.max_bytes" "value length must be at most {{$1}} bytes" $r.GetMaxBytes) }}
		}
	{{ end }}

	{{ if $r.Prefix }}
		if !strings.HasPrefix({{ accessor . }}, {{ lit $r.GetPrefix }}) {
			return {{ err . (t "string.prefix" "value does not have prefix {{$1}}" (lit $r.GetPrefix)) }}
		}
	{{ end }}

	{{ if $r.Suffix }}
		if !strings.HasSuffix({{ accessor . }}, {{ lit $r.GetSuffix }}) {
			return {{ err . (t "string.suffix" "value does not have suffix {{$1}}" (lit $r.GetSuffix)) }}
		}
	{{ end }}

	{{ if $r.Contains }}
		if !strings.Contains({{ accessor . }}, {{ lit $r.GetContains }}) {
			return {{ err . (t "string.contains" "value does not contain substring {{$1}}" (lit $r.GetContains)) }}
		}
	{{ end }}

	{{ if $r.NotContains }}
		if strings.Contains({{ accessor . }}, {{ lit $r.GetNotContains }}) {
			return {{ err . (t "string.not_contains" "value contains substring {{$1}}" (lit $r.GetNotContains)) }}
		}
	{{ end }}

	{{ if $r.GetIp }}
		if ip := net.ParseIP({{ accessor . }}); ip == nil {
			return {{ err . (t "string.ip" "value must be a valid IP address") }}
		}
	{{ else if $r.GetIpv4 }}
		if ip := net.ParseIP({{ accessor . }}); ip == nil || ip.To4() == nil {
			return {{ err . (t "string.ipv4" "value must be a valid IPv4 address") }}
		}
	{{ else if $r.GetIpv6 }}
		if ip := net.ParseIP({{ accessor . }}); ip == nil || ip.To4() != nil {
			return {{ err . (t "string.ipv6" "value must be a valid IPv6 address") }}
		}
	{{ else if $r.GetEmail }}
		if err := m._validateEmail({{ accessor . }}); err != nil {
			return {{ errCause . "err" (t "string.email" "value must be a valid email address") }}
		}
	{{ else if $r.GetHostname }}
		if err := m._validateHostname({{ accessor . }}); err != nil {
			return {{ errCause . "err" (t "string.hostname" "value must be a valid hostname") }}
		}
	{{ else if $r.GetAddress }}
		if err := m._validateHostname({{ accessor . }}); err != nil {
			if ip := net.ParseIP({{ accessor . }}); ip == nil {
				return {{ err . (t "string.address" "value must be a valid hostname, or ip address") }}
			}
		}
	{{ else if $r.GetUri }}
		if uri, err := url.Parse({{ accessor . }}); err != nil {
			return {{ errCause . "err" (t "string.uri" "value must be a valid URI") }}
		} else if !uri.IsAbs() {
			return {{ err . (t "string.uri_absolute" "value must be absolute") }}
		}
	{{ else if $r.GetUriRef }}
		if _, err := url.Parse({{ accessor . }}); err != nil {
			return {{ errCause . "err" (t "string.uri_ref" "value must be a valid URI") }}
		}
	{{ else if $r.GetUuid }}
		if err := m._validateUuid({{ accessor . }}); err != nil {
			return {{ errCause . "err" (t "string.uuid" "value must be a valid UUID") }}
		}
	{{ end }}

	{{ if $r.Pattern }}
	if !{{ lookup $f "Pattern" }}.MatchString({{ accessor . }}) {
		return {{ err . (t "string.pattern" "value does not match regex pattern {{$1}}" (lit $r.GetPattern)) }}
	}
{{ end }}
`
