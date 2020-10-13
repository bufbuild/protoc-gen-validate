package goshared

const bytesTpl = `
	{{ $f := .Field }}{{ $r := .Rules }}

	{{ if or $r.Len (and $r.MinLen $r.MaxLen (eq $r.GetMinLen $r.GetMaxLen)) }}
		{{ if $r.Len }}
			if len({{ accessor . }}) != {{ $r.GetLen }} {
				return {{ err . (t "bytes.len" "value length must be {{$1}} bytes" $r.GetLen) }}
			}
		{{ else }}
			if len({{ accessor . }}) != {{ $r.GetMinLen }} {
				return {{ err . (t "bytes.len" "value length must be {{$1}} bytes" $r.GetMinLen) }}
			}
		{{ end }}
	{{ else if $r.MinLen }}
		{{ if $r.MaxLen }}
			if l := len({{ accessor . }}); l < {{ $r.GetMinLen }} || l > {{ $r.GetMaxLen }} {
				return {{ err . (t "bytes.len_between" "value length must be between {{$1}} and {{$2}} bytes, inclusive" $r.GetMinLen $r.GetMaxLen) }}
			}
		{{ else }}
			if len({{ accessor . }}) < {{ $r.GetMinLen }} {
				return {{ err . (t "bytes.min_len" "value length must be at least {{$1}} bytes" $r.GetMinLen) }}
			}
		{{ end }}
	{{ else if $r.MaxLen }}
		if len({{ accessor . }}) > {{ $r.GetMaxLen }} {
			return {{ err . (t "bytes.max_len" "value length must be at most {{$1}} bytes" $r.GetMaxLen) }}
		}
	{{ end }}

	{{ if $r.Prefix }}
		if !bytes.HasPrefix({{ accessor . }}, {{ lit $r.GetPrefix }}) {
			return {{ err . (t "bytes.prefix" "value does not have prefix {{$1}}" (byteStr $r.GetPrefix)) }}
		}
	{{ end }}

	{{ if $r.Suffix }}
		if !bytes.HasSuffix({{ accessor . }}, {{ lit $r.GetSuffix }}) {
			return {{ err . (t "bytes.suffix" "value does not have suffix {{$1}}" (byteStr $r.GetSuffix)) }}
		}
	{{ end }}

	{{ if $r.Contains }}
		if !bytes.Contains({{ accessor . }}, {{ lit $r.GetContains }}) {
			return {{ err . (t "bytes.contains" "value does not contain {{$1}}" (byteStr $r.GetContains)) }}
		}
	{{ end }}

	{{ if $r.In }}
		if _, ok := {{ lookup $f "InLookup" }}[string({{ accessor . }})]; !ok {
			return {{ err . (t "bytes.in" "value must be in list {{$1}}" $r.In) }}
		}
	{{ else if $r.NotIn }}
		if _, ok := {{ lookup $f "NotInLookup" }}[string({{ accessor . }})]; ok {
			return {{ err . (t "bytes.not_in" "value must not be in list {{$1}}" $r.NotIn) }}
		}
	{{ end }}

	{{ if $r.Const }}
		if !bytes.Equal({{ accessor . }}, {{ lit $r.Const }}) {
			return {{ err . (t "bytes.const" "value must equal {{$1}}" $r.Const) }}
		}
	{{ end }}

	{{ if $r.GetIp }}
		if ip := net.IP({{ accessor . }}); ip.To16() == nil {
			return {{ err . (t "bytes.ip" "value must be a valid IP address") }}
		}
	{{ else if $r.GetIpv4 }}
		if ip := net.IP({{ accessor . }}); ip.To4() == nil {
			return {{ err . (t "bytes.ipv4" "value must be a valid IPv4 address") }}
		}
	{{ else if $r.GetIpv6 }}
		if ip := net.IP({{ accessor . }}); ip.To16() == nil || ip.To4() != nil {
			return {{ err . (t "bytes.ipv6" "value must be a valid IPv6 address") }}
		}
	{{ end }}

	{{ if $r.Pattern }}
	if !{{ lookup $f "Pattern" }}.Match({{ accessor . }}) {
		return {{ err . (t "bytes.pattern" "value does not match regex pattern {{$1}}" (lit $r.GetPattern)) }}
	}
	{{ end }}
`
