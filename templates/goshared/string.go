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
			err := {{ err . "string.len" "value length must be " $r.GetLen " runes" }}
			if !all { return err }
			errors = append(errors, err)
		{{ else }}
		if utf8.RuneCountInString({{ accessor . }}) != {{ $r.GetMinLen }} {
			err := {{ err . "string.len" "value length must be " $r.GetMinLen " runes" }}
			if !all { return err }
			errors = append(errors, err)
		{{ end }}
	}
	{{ else if $r.MinLen }}
		{{ if $r.MaxLen }}
			if l := utf8.RuneCountInString({{ accessor . }}); l < {{ $r.GetMinLen }} || l > {{ $r.GetMaxLen }} {
				err := {{ err . "string.between" "value length must be between " $r.GetMinLen " and " $r.GetMaxLen " runes, inclusive" }}
				if !all { return err }
				errors = append(errors, err)
			}
		{{ else }}
			if utf8.RuneCountInString({{ accessor . }}) < {{ $r.GetMinLen }} {
				err := {{ err . "string.min_len" "value length must be at least " $r.GetMinLen " runes" }}
				if !all { return err }
				errors = append(errors, err)
			}
		{{ end }}
	{{ else if $r.MaxLen }}
		if utf8.RuneCountInString({{ accessor . }}) > {{ $r.GetMaxLen }} {
			err := {{ err . "string.max_len" "value length must be at most " $r.GetMaxLen " runes" }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ end }}

	{{ if or $r.LenBytes (and $r.MinBytes $r.MaxBytes (eq $r.GetMinBytes $r.GetMaxBytes)) }}
		{{ if $r.LenBytes }}
			if len({{ accessor . }}) != {{ $r.GetLenBytes }} {
				err := {{ err . "string.len_bytes" "value length must be " $r.GetLenBytes " bytes" }}
				if !all { return err }
				errors = append(errors, err)
			}
		{{ else }}
			if len({{ accessor . }}) != {{ $r.GetMinBytes }} {
				err := {{ err . "string.len_bytes" "value length must be " $r.GetMinBytes " bytes" }}
				if !all { return err }
				errors = append(errors, err)
			}
		{{ end }}
	{{ else if $r.MinBytes }}
		{{ if $r.MaxBytes }}
			if l := len({{ accessor . }}); l < {{ $r.GetMinBytes }} || l > {{ $r.GetMaxBytes }} {
					err := {{ err . "string.between_bytes" "value length must be between " $r.GetMinBytes " and " $r.GetMaxBytes " bytes, inclusive" }}
					if !all { return err }
					errors = append(errors, err)
			}
		{{ else }}
			if len({{ accessor . }}) < {{ $r.GetMinBytes }} {
				err := {{ err . "string.min_bytes" "value length must be at least " $r.GetMinBytes " bytes" }}
				if !all { return err }
				errors = append(errors, err)
			}
		{{ end }}
	{{ else if $r.MaxBytes }}
		if len({{ accessor . }}) > {{ $r.GetMaxBytes }} {
			err := {{ err . "string.max_bytes" "value length must be at most " $r.GetMaxBytes " bytes" }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ end }}

	{{ if $r.Prefix }}
		if !strings.HasPrefix({{ accessor . }}, {{ lit $r.GetPrefix }}) {
			err := {{ err . "string.prefix" "value does not have prefix " (lit $r.GetPrefix) }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ end }}

	{{ if $r.Suffix }}
		if !strings.HasSuffix({{ accessor . }}, {{ lit $r.GetSuffix }}) {
			err := {{ err . "string.suffix" "value does not have suffix " (lit $r.GetSuffix) }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ end }}

	{{ if $r.Contains }}
		if !strings.Contains({{ accessor . }}, {{ lit $r.GetContains }}) {
			err := {{ err . "string.contains" "value does not contain substring " (lit $r.GetContains) }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ end }}

	{{ if $r.NotContains }}
		if strings.Contains({{ accessor . }}, {{ lit $r.GetNotContains }}) {
			err := {{ err . "string.not_contains" "value contains substring " (lit $r.GetNotContains) }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ end }}

	{{ if $r.GetIp }}
		if ip := net.ParseIP({{ accessor . }}); ip == nil {
			err := {{ err . "string.ip" "value must be a valid IP address" }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ else if $r.GetIpv4 }}
		if ip := net.ParseIP({{ accessor . }}); ip == nil || ip.To4() == nil {
			err := {{ err . "string.ipv4" "value must be a valid IPv4 address" }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ else if $r.GetIpv6 }}
		if ip := net.ParseIP({{ accessor . }}); ip == nil || ip.To4() != nil {
			err := {{ err . "string.ipv6" "value must be a valid IPv6 address" }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ else if $r.GetEmail }}
		if err := m._validateEmail({{ accessor . }}); err != nil {
			err = {{ errCause . "string.email" "err" "value must be a valid email address" }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ else if $r.GetHostname }}
		if err := m._validateHostname({{ accessor . }}); err != nil {
			err = {{ errCause . "string.hostname" "err" "value must be a valid hostname" }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ else if $r.GetAddress }}
		if err := m._validateHostname({{ accessor . }}); err != nil {
			if ip := net.ParseIP({{ accessor . }}); ip == nil {
				err := {{ err . "string.address" "value must be a valid hostname, or ip address" }}
				if !all { return err }
				errors = append(errors, err)
			}
		}
	{{ else if $r.GetUri }}
		if uri, err := url.Parse({{ accessor . }}); err != nil {
			err = {{ errCause . "string.uri" "err" "value must be a valid URI" }}
			if !all { return err }
			errors = append(errors, err)
		} else if !uri.IsAbs() {
			err := {{ err . "value must be absolute" }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ else if $r.GetUriRef }}
		if _, err := url.Parse({{ accessor . }}); err != nil {
			err = {{ errCause . "string.uri_ref" "err" "value must be a valid URI" }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ else if $r.GetUuid }}
		if err := m._validateUuid({{ accessor . }}); err != nil {
			err = {{ errCause . "string.uuid" "err" "value must be a valid UUID" }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ end }}

	{{ if $r.Pattern }}
		if !{{ lookup $f "Pattern" }}.MatchString({{ accessor . }}) {
			err := {{ err . "string.pattern" "value does not match regex pattern " (lit $r.GetPattern) }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ end }}

	{{ if $r.GetIgnoreEmpty }}
		}
	{{ end }}

`
