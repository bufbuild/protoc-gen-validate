package tpl

const strTpl = `
	{{ $f := .Field }}{{ $r := .Rules }}
	{{ template "const" . }}
	{{/* template "in" . */}}
{{/* // TODO(akonradi) implement regular expression constraints.
	{{ if $r.Pattern }}
		if !{{ lookup $f "Pattern" }}.MatchString({{ accessor . }}) {
			return {{ err . "value does not match regex pattern " (lit $r.GetPattern) }}
		}
	{{ end }}
*/}}
	{{ if $r.Prefix }}
	{
		const std::string prefix = {{ lit $r.GetPrefix }};
		if ({{ accessor . }}.compare(0, prefix.size(), prefix) != 0) {
			{{ err . "value does not have prefix " (lit $r.GetPrefix) }}
		}
	}
	{{ end }}

	{{ if $r.Suffix }}
	{
		const std::string suffix = {{ lit $r.GetSuffix }};
		const std::string& value = {{ accessor . }};
		if ((value.size() < suffix.size()) ||
			(value.compare(value.size() - suffix.size(), suffix.size(), suffix) != 0)) {
			{{ err . "value does not have suffix " (lit $r.GetSuffix) }}
		}
	}
	{{ end }}

	{{ if $r.Contains }}
	{
		if ({{ accessor . }}.find({{ lit $r.GetContains }}) == std::string::npos) {
			{{ err . "value does not contain substring " (lit $r.GetContains) }}
		}
	}
	{{ end }}

{{/* // TODO(akonradi) implement the below constraints
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

	{{ if $r.MinLen }}
		{{ if $r.MaxLen }}
			{{ if eq $r.GetMinLen $r.GetMaxLen }}
				if utf8.RuneCountInString({{ accessor . }}) != {{ $r.GetMinLen }} {
					return {{ err . "value length must be " $r.GetMinLen " runes" }}
				}
			{{ else }}
				if l := utf8.RuneCountInString({{ accessor . }}); l < {{ $r.GetMinLen }} || l > {{ $r.GetMaxLen }} {
					return {{ err . "value length must be between " $r.GetMinLen " and " $r.GetMaxLen " runes, inclusive" }}
				}
			{{ end }}
		{{ else }}
			if utf8.RuneCountInString({{ accessor . }}) < {{ $r.GetMinLen }} {
				return {{ err . "value length must be at least " $r.GetMinLen " runes" }}
			}
		{{ end }}
	{{ else if $r.MaxLen }}
		if utf8.RuneCountInString({{ accessor . }}) > {{ $r.GetMaxLen }} {
			return {{ err . "value length must be at most " $r.GetMaxLen " runes" }}
		}
	{{ end }}
*/}}

	{{ if $r.MinBytes }}
	{
		const auto length = {{ accessor . }}.size();
		{{ if $r.MaxBytes }}
			{{ if eq $r.GetMinBytes $r.GetMaxBytes }}
				if (length != {{ $r.GetMinBytes }}) {
					{{ err . "value length must be " $r.GetMinBytes " bytes" }}
				}
			{{ else }}
				if (length < {{ $r.GetMinBytes }} || length > {{ $r.GetMaxBytes }}) {
					{{ err . "value length must be between " $r.GetMinBytes " and " $r.GetMaxBytes " bytes, inclusive" }}
				}
			{{ end }}
		{{ else }}
			if (length < {{ $r.GetMinBytes }}) {
				{{ err . "value length must be at least " $r.GetMinBytes " bytes" }}
			}
		{{ end }}
	}
	{{ else if $r.MaxBytes }}
		if ({{ accessor . }}.size() > {{ $r.GetMaxBytes }}) {
			{{ err . "value length must be at most " $r.GetMaxBytes " bytes" }}
		}
	{{ end }}
`
