package goshared

const strTpl = `
	{{ $f := .Field }}{{ $r := .Rules }}

	{{ if $r.GetIgnoreEmpty }}
		if {{ accessor . }} != "" {
	{{ end }}

	{{ template "const" . }}
	{{ template "in" . }}

	{{ if or $r.Len (and $r.MinLen $r.MaxLen (eq $r.GetMinLen $r.GetMaxLen)) }}
		{{ if $r.Len }}
		if utf8.RuneCountInString({{ accessor . }}) != {{ $r.GetLen }} {
			err := {{ err . (t "string.len" "value length must be {{$1}} runes" $r.GetLen) }}
			if !all { return err }
			errors = append(errors, err)
		{{ else }}
		if utf8.RuneCountInString({{ accessor . }}) != {{ $r.GetMinLen }} {
			err := {{ err . (t "string.len" "value length must be {{$1}} runes" $r.GetMinLen) }}
			if !all { return err }
			errors = append(errors, err)
		{{ end }}
	}
	{{ else if $r.MinLen }}
		{{ if $r.MaxLen }}
			if l := utf8.RuneCountInString({{ accessor . }}); l < {{ $r.GetMinLen }} || l > {{ $r.GetMaxLen }} {
				err := {{ err . (t "string.len_between" "value length must be between {{$1}} and {{$2}} runes, inclusive" $r.GetMinLen $r.GetMaxLen) }}
				if !all { return err }
				errors = append(errors, err)
			}
		{{ else }}
			if utf8.RuneCountInString({{ accessor . }}) < {{ $r.GetMinLen }} {
				err := {{ err . (t "string.min_len" "value length must be at least {{$1}} runes" $r.GetMinLen) }}
				if !all { return err }
				errors = append(errors, err)
			}
		{{ end }}
	{{ else if $r.MaxLen }}
		if utf8.RuneCountInString({{ accessor . }}) > {{ $r.GetMaxLen }} {
			err := {{ err . (t "string.max_len" "value length must be at most {{$1}} runes" $r.GetMaxLen) }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ end }}

	{{ if or $r.LenBytes (and $r.MinBytes $r.MaxBytes (eq $r.GetMinBytes $r.GetMaxBytes)) }}
		{{ if $r.LenBytes }}
			if len({{ accessor . }}) != {{ $r.GetLenBytes }} {
				err := {{ err . (t "string.len_bytes" "value length must be {{$1}} bytes" $r.GetLenBytes) }}
				if !all { return err }
				errors = append(errors, err)
			}
		{{ else }}
			if len({{ accessor . }}) != {{ $r.GetMinBytes }} {
				err := {{ err . (t "string.len_bytes" "value length must be {{$1}} bytes" $r.GetMinBytes) }}
				if !all { return err }
				errors = append(errors, err)
			}
		{{ end }}
	{{ else if $r.MinBytes }}
		{{ if $r.MaxBytes }}
			if l := len({{ accessor . }}); l < {{ $r.GetMinBytes }} || l > {{ $r.GetMaxBytes }} {
					err := {{ err . (t "string.bytes_between" "value length must be between {{$1}} and {{$2}} bytes, inclusive" $r.GetMinBytes $r.GetMaxBytes) }}
					if !all { return err }
					errors = append(errors, err)
			}
		{{ else }}
			if len({{ accessor . }}) < {{ $r.GetMinBytes }} {
				err := {{ err . (t "string.min_bytes" "value length must be at least {{$1}} bytes" $r.GetMinBytes) }}
				if !all { return err }
				errors = append(errors, err)
			}
		{{ end }}
	{{ else if $r.MaxBytes }}
		if len({{ accessor . }}) > {{ $r.GetMaxBytes }} {
			err := {{ err . (t "string.max_bytes" "value length must be at most {{$1}} bytes" $r.GetMaxBytes) }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ end }}

	{{ if $r.Prefix }}
		if !strings.HasPrefix({{ accessor . }}, {{ lit $r.GetPrefix }}) {
			err := {{ err . (t "string.prefix" "value does not have prefix {{$1}}" (lit $r.GetPrefix)) }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ end }}

	{{ if $r.Suffix }}
		if !strings.HasSuffix({{ accessor . }}, {{ lit $r.GetSuffix }}) {
			err := {{ err . (t "string.suffix" "value does not have suffix {{$1}}" (lit $r.GetSuffix)) }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ end }}

	{{ if $r.Contains }}
		if !strings.Contains({{ accessor . }}, {{ lit $r.GetContains }}) {
			err := {{ err . (t "string.contains" "value does not contain substring {{$1}}" (lit $r.GetContains)) }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ end }}

	{{ if $r.NotContains }}
		if strings.Contains({{ accessor . }}, {{ lit $r.GetNotContains }}) {
			err := {{ err . (t "string.not_contains" "value contains substring {{$1}}" (lit $r.GetNotContains)) }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ end }}

	{{ if $r.GetIp }}
		if ip := net.ParseIP({{ accessor . }}); ip == nil {
			err := {{ err . (t "string.ip" "value must be a valid IP address") }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ else if $r.GetIpv4 }}
		if ip := net.ParseIP({{ accessor . }}); ip == nil || ip.To4() == nil {
			err := {{ err . (t "string.ipv4" "value must be a valid IPv4 address") }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ else if $r.GetIpv6 }}
		if ip := net.ParseIP({{ accessor . }}); ip == nil || ip.To4() != nil {
			err := {{ err . (t "string.ipv6" "value must be a valid IPv6 address")}}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ else if $r.GetEmail }}
		if err := m._validateEmail({{ accessor . }}); err != nil {
			err = {{ errCause . "err" (t "string.email" "value must be a valid email address") }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ else if $r.GetHostname }}
		if err := m._validateHostname({{ accessor . }}); err != nil {
			err = {{ errCause . "err" (t "string.hostname" "value must be a valid hostname")}}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ else if $r.GetAddress }}
		if err := m._validateHostname({{ accessor . }}); err != nil {
			if ip := net.ParseIP({{ accessor . }}); ip == nil {
				err := {{ err . (t "string.address" "value must be a valid hostname, or ip address") }}
				if !all { return err }
				errors = append(errors, err)
			}
		}
	{{ else if $r.GetUri }}
		if uri, err := url.Parse({{ accessor . }}); err != nil {
			err = {{ errCause . "err" (t "string.uri" "value must be a valid URI") }}
			if !all { return err }
			errors = append(errors, err)
		} else if !uri.IsAbs() {
			err := {{ err . (t "string.uri_absolute" "value must be absolute") }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ else if $r.GetUriRef }}
		if _, err := url.Parse({{ accessor . }}); err != nil {
			err = {{ errCause . "err" (t "string.uri_ref" "value must be a valid URI") }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ else if $r.GetUuid }}
		if err := m._validateUuid({{ accessor . }}); err != nil {
			err = {{ errCause . "err" (t "string.uuid" "value must be a valid UUID") }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ end }}

	{{ if $r.Pattern }}
		if !{{ lookup $f "Pattern" }}.MatchString({{ accessor . }}) {
			err := {{ err . (t "string.pattern" "value does not match regex pattern {{$1}}" (lit $r.GetPattern)) }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ end }}

	{{ if $r.GetIgnoreEmpty }}
		}
	{{ end }}

`
