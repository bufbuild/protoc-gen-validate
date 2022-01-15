package goshared

const bytesTpl = `
	{{ $f := .Field }}{{ $r := .Rules }}

	{{ if $r.GetIgnoreEmpty }}
		if len({{ accessor . }}) > 0 {
	{{ end }}

	{{ if or $r.Len (and $r.MinLen $r.MaxLen (eq $r.GetMinLen $r.GetMaxLen)) }}
		{{ if $r.Len }}
			if len({{ accessor . }}) != {{ $r.GetLen }} {
				{{ if ne $r.GetErrorMsg "" }}
				err := {{ err . $r.GetErrorMsg }}
				{{ else }}
				err := {{ err . "value length must be " $r.GetLen " bytes" }}
				{{ end }}
				if !all { return err }
				errors = append(errors, err)
			}
		{{ else }}
			if len({{ accessor . }}) != {{ $r.GetMinLen }} {
				{{ if ne $r.GetErrorMsg "" }}
				err := {{ err . $r.GetErrorMsg }}
				{{ else }}
				err := {{ err . "value length must be " $r.GetMinLen " bytes" }}
				{{ end }}
				if !all { return err }
				errors = append(errors, err)
			}
		{{ end }}
	{{ else if $r.MinLen }}
		{{ if $r.MaxLen }}
			if l := len({{ accessor . }}); l < {{ $r.GetMinLen }} || l > {{ $r.GetMaxLen }} {
				{{ if ne $r.GetErrorMsg "" }}
				err := {{ err . $r.GetErrorMsg }}
				{{ else }}
				err := {{ err . "value length must be between " $r.GetMinLen " and " $r.GetMaxLen " bytes, inclusive" }}
				{{ end }}
				if !all { return err }
				errors = append(errors, err)
			}
		{{ else }}
			if len({{ accessor . }}) < {{ $r.GetMinLen }} {
				{{ if ne $r.GetErrorMsg "" }}
				err := {{ err . $r.GetErrorMsg }}
				{{ else }}
				err := {{ err . "value length must be at least " $r.GetMinLen " bytes" }}
				{{ end }}
				if !all { return err }
				errors = append(errors, err)
			}
		{{ end }}
	{{ else if $r.MaxLen }}
		if len({{ accessor . }}) > {{ $r.GetMaxLen }} {
			{{ if ne $r.GetErrorMsg "" }}
			err := {{ err . $r.GetErrorMsg }}
			{{ else }}
			err := {{ err . "value length must be at most " $r.GetMaxLen " bytes" }}
			{{ end }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ end }}

	{{ if $r.Prefix }}
		if !bytes.HasPrefix({{ accessor . }}, {{ lit $r.GetPrefix }}) {
			{{ if ne $r.GetErrorMsg "" }}
			err := {{ err . $r.GetErrorMsg }}
			{{ else }}
			err := {{ err . "value does not have prefix " (byteStr $r.GetPrefix) }}
			{{ end }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ end }}

	{{ if $r.Suffix }}
		if !bytes.HasSuffix({{ accessor . }}, {{ lit $r.GetSuffix }}) {
			{{ if ne $r.GetErrorMsg "" }}
			err := {{ err . $r.GetErrorMsg }}
			{{ else }}
			err := {{ err . "value does not have suffix " (byteStr $r.GetSuffix) }}
			{{ end }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ end }}

	{{ if $r.Contains }}
		if !bytes.Contains({{ accessor . }}, {{ lit $r.GetContains }}) {
			{{ if ne $r.GetErrorMsg "" }}
			err := {{ err . $r.GetErrorMsg }}
			{{ else }}
			err := {{ err . "value does not contain " (byteStr $r.GetContains) }}
			{{ end }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ end }}

	{{ if $r.In }}
		if _, ok := {{ lookup $f "InLookup" }}[string({{ accessor . }})]; !ok {
			{{ if ne $r.GetErrorMsg "" }}
			err := {{ err . $r.GetErrorMsg }}
			{{ else }}
			err := {{ err . "value must be in list " $r.In }}
			{{ end }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ else if $r.NotIn }}
		if _, ok := {{ lookup $f "NotInLookup" }}[string({{ accessor . }})]; ok {
			{{ if ne $r.GetErrorMsg "" }}
			err := {{ err . $r.GetErrorMsg }}
			{{ else }}
			err := {{ err . "value must not be in list " $r.NotIn }}
			{{ end }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ end }}

	{{ if $r.Const }}
		if !bytes.Equal({{ accessor . }}, {{ lit $r.Const }}) {
			{{ if ne $r.GetErrorMsg "" }}
			err := {{ err . $r.GetErrorMsg }}
			{{ else }}
			err := {{ err . "value must equal " $r.Const }}
			{{ end }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ end }}

	{{ if $r.GetIp }}
		if ip := net.IP({{ accessor . }}); ip.To16() == nil {
			{{ if ne $r.GetErrorMsg "" }}
			err := {{ err . $r.GetErrorMsg }}
			{{ else }}
			err := {{ err . "value must be a valid IP address" }}
			{{ end }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ else if $r.GetIpv4 }}
		if ip := net.IP({{ accessor . }}); ip.To4() == nil {
			{{ if ne $r.GetErrorMsg "" }}
			err := {{ err . $r.GetErrorMsg }}
			{{ else }}
			err := {{ err . "value must be a valid IPv4 address" }}
			{{ end }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ else if $r.GetIpv6 }}
		if ip := net.IP({{ accessor . }}); ip.To16() == nil || ip.To4() != nil {
			{{ if ne $r.GetErrorMsg "" }}
			err := {{ err . $r.GetErrorMsg }}
			{{ else }}
			err := {{ err . "value must be a valid IPv6 address" }}
			{{ end }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ end }}

	{{ if $r.Pattern }}
	if !{{ lookup $f "Pattern" }}.Match({{ accessor . }}) {
		{{ if ne $r.GetErrorMsg "" }}
		err := {{ err . $r.GetErrorMsg }}
		{{ else }}
		err := {{ err . "value does not match regex pattern " (lit $r.GetPattern) }}
		{{ end }}
		if !all { return err }
		errors = append(errors, err)
	}
	{{ end }}

	{{ if $r.GetIgnoreEmpty }}
		}
	{{ end }}
`
