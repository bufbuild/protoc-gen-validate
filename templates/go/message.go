package golang

const messageTpl = `
	{{ $f := .Field }}{{ $r := .Rules }}
	{{ template "required" . }}

	{{ if .MessageRules.GetSkip }}
		// skipping validation for {{ $f.Name }}
	{{ else }}
		if v, ok := interface{}({{ accessor . }}).(interface{
			ValidateWithMask(*field_mask.FieldMask) error
		}); m.maskHas(mask, "{{ $f.Name }}") && ok {
			mask = m.updateMask(mask, "{{ $f.Name }}")

			if err := v.ValidateWithMask(mask); err != nil {
				return {{ errCause . "err" "embedded message failed validation" }}
			}
		}
	{{ end }}
`
