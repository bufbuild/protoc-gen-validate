package goshared

const ltgtTpl = `{{ $f := .Field }}{{ $r := .Rules }}
	{{ if $r.Lt }}
		{{ if $r.Gt }}
			{{  if gt $r.GetLt $r.GetGt }}
				if val := {{ accessor . }};  val <= {{ $r.Gt }} || val >= {{ $r.Lt }} {
					{{ if ne $r.GetErrorMsg "" }}
					err := {{ err . $r.GetErrorMsg }}
					{{ else }}
					err := {{ err . "value must be inside range (" $r.GetGt ", " $r.GetLt ")" }}
					{{ end }}
					if !all { return err }
					errors = append(errors, err)
				}
			{{ else }}
				if val := {{ accessor . }}; val >= {{ $r.Lt }} && val <= {{ $r.Gt }} {
					{{ if ne $r.GetErrorMsg "" }}
					err := {{ err . $r.GetErrorMsg }}
					{{ else }}
					err := {{ err . "value must be outside range [" $r.GetLt ", " $r.GetGt "]" }}
					{{ end }}
					if !all { return err }
					errors = append(errors, err)
				}
			{{ end }}
		{{ else if $r.Gte }}
			{{  if gt $r.GetLt $r.GetGte }}
				if val := {{ accessor . }};  val < {{ $r.Gte }} || val >= {{ $r.Lt }} {
					{{ if ne $r.GetErrorMsg "" }}
					err := {{ err . $r.GetErrorMsg }}
					{{ else }}
					err := {{ err . "value must be inside range [" $r.GetGte ", " $r.GetLt ")" }}
					{{ end }}
					if !all { return err }
					errors = append(errors, err)
				}
			{{ else }}
				if val := {{ accessor . }}; val >= {{ $r.Lt }} && val < {{ $r.Gte }} {
					{{ if ne $r.GetErrorMsg "" }}
					err := {{ err . $r.GetErrorMsg }}
					{{ else }}
					err := {{ err . "value must be outside range [" $r.GetLt ", " $r.GetGte ")" }}
					{{ end }}
					if !all { return err }
					errors = append(errors, err)
				}
			{{ end }}
		{{ else }}
			if {{ accessor . }} >= {{ $r.Lt }} {
				{{ if ne $r.GetErrorMsg "" }}
				err := {{ err . $r.GetErrorMsg }}
				{{ else }}
				err := {{ err . "value must be less than " $r.GetLt }}
				{{ end }}
				if !all { return err }
				errors = append(errors, err)
			}
		{{ end }}
	{{ else if $r.Lte }}
		{{ if $r.Gt }}
			{{  if gt $r.GetLte $r.GetGt }}
				if val := {{ accessor . }};  val <= {{ $r.Gt }} || val > {{ $r.Lte }} {
					{{ if ne $r.GetErrorMsg "" }}
					err := {{ err . $r.GetErrorMsg }}
					{{ else }}
					err := {{ err . "value must be inside range (" $r.GetGt ", " $r.GetLte "]" }}
					{{ end }}
					if !all { return err }
					errors = append(errors, err)
				}
			{{ else }}
				if val := {{ accessor . }}; val > {{ $r.Lte }} && val <= {{ $r.Gt }} {
					{{ if ne $r.GetErrorMsg "" }}
					err := {{ err . $r.GetErrorMsg }}
					{{ else }}
					err := {{ err . "value must be outside range (" $r.GetLte ", " $r.GetGt "]" }}
					{{ end }}
					if !all { return err }
					errors = append(errors, err)
				}
			{{ end }}
		{{ else if $r.Gte }}
			{{ if gt $r.GetLte $r.GetGte }}
				if val := {{ accessor . }};  val < {{ $r.Gte }} || val > {{ $r.Lte }} {
					{{ if ne $r.GetErrorMsg "" }}
					err := {{ err . $r.GetErrorMsg }}
					{{ else }}
					err := {{ err . "value must be inside range [" $r.GetGte ", " $r.GetLte "]" }}
					{{ end }}
					if !all { return err }
					errors = append(errors, err)
				}
			{{ else }}
				if val := {{ accessor . }}; val > {{ $r.Lte }} && val < {{ $r.Gte }} {
					{{ if ne $r.GetErrorMsg "" }}
					err := {{ err . $r.GetErrorMsg }}
					{{ else }}
					err := {{ err . "value must be outside range (" $r.GetLte ", " $r.GetGte ")" }}
					{{ end }}
					if !all { return err }
					errors = append(errors, err)
				}
			{{ end }}
		{{ else }}
			if {{ accessor . }} > {{ $r.Lte }} {
				{{ if ne $r.GetErrorMsg "" }}
				err := {{ err . $r.GetErrorMsg }}
				{{ else }}
				err := {{ err . "value must be less than or equal to " $r.GetLte }}
				{{ end }}
				if !all { return err }
				errors = append(errors, err)
			}
		{{ end }}
	{{ else if $r.Gt }}
		if {{ accessor . }} <= {{ $r.Gt }} {
			{{ if ne $r.GetErrorMsg "" }}
			err := {{ err . $r.GetErrorMsg }}
			{{ else }}
			err := {{ err . "value must be greater than " $r.GetGt }}
			{{ end }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ else if $r.Gte }}
		if {{ accessor . }} < {{ $r.Gte }} {
			{{ if ne $r.GetErrorMsg "" }}
			err := {{ err . $r.GetErrorMsg }}
			{{ else }}
			err := {{ err . "value must be greater than or equal to " $r.GetGte }}
			{{ end }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ end }}
`
