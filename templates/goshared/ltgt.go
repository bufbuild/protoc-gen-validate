package goshared

const ltgtTpl = `{{ $f := .Field }}{{ $r := .Rules }}{{ $t := .Typ}}
	{{ if $r.Lt }}
		{{ if $r.Gt }}
			{{  if gt $r.GetLt $r.GetGt }}
				if val := {{ accessor . }};  val <= {{ $r.Gt }} || val >= {{ $r.Lt }} {
					err := {{ err . (print $t ".in_range_exclusive") "value must be inside range (" $r.GetGt ", " $r.GetLt ")" }}
					if !all { return err }
					errors = append(errors, err)
				}
			{{ else }}
				if val := {{ accessor . }}; val >= {{ $r.Lt }} && val <= {{ $r.Gt }} {
					err := {{ err . (print $t ".out_of_range") "value must be outside range [" $r.GetLt ", " $r.GetGt "]" }}
					if !all { return err }
					errors = append(errors, err)
				}
			{{ end }}
		{{ else if $r.Gte }}
			{{  if gt $r.GetLt $r.GetGte }}
				if val := {{ accessor . }};  val < {{ $r.Gte }} || val >= {{ $r.Lt }} {
					err := {{ err . (print $t ".in_range_upper_exclusive") "value must be inside range [" $r.GetGte ", " $r.GetLt ")" }}
					if !all { return err }
					errors = append(errors, err)
				}
			{{ else }}
				if val := {{ accessor . }}; val >= {{ $r.Lt }} && val < {{ $r.Gte }} {
					err := {{ err . (print $t ".out_of_range_upper_inclusive") "value must be outside range [" $r.GetLt ", " $r.GetGte ")" }}
					if !all { return err }
					errors = append(errors, err)
				}
			{{ end }}
		{{ else }}
			if {{ accessor . }} >= {{ $r.Lt }} {
				err := {{ err . (print $t ".lt") "value must be less than " $r.GetLt }}
				if !all { return err }
				errors = append(errors, err)
			}
		{{ end }}
	{{ else if $r.Lte }}
		{{ if $r.Gt }}
			{{  if gt $r.GetLte $r.GetGt }}
				if val := {{ accessor . }};  val <= {{ $r.Gt }} || val > {{ $r.Lte }} {
					err := {{ err . (print $t ".in_range_lower_exclusive") "value must be inside range (" $r.GetGt ", " $r.GetLte "]" }}
					if !all { return err }
					errors = append(errors, err)
				}
			{{ else }}
				if val := {{ accessor . }}; val > {{ $r.Lte }} && val <= {{ $r.Gt }} {
					err := {{ err . (print $t ".out_of_range_lower_inclusive") "value must be outside range (" $r.GetLte ", " $r.GetGt "]" }}
					if !all { return err }
					errors = append(errors, err)
				}
			{{ end }}
		{{ else if $r.Gte }}
			{{ if gt $r.GetLte $r.GetGte }}
				if val := {{ accessor . }};  val < {{ $r.Gte }} || val > {{ $r.Lte }} {
					err := {{ err . (print $t ".in_range") "value must be inside range [" $r.GetGte ", " $r.GetLte "]" }}
					if !all { return err }
					errors = append(errors, err)
				}
			{{ else }}
				if val := {{ accessor . }}; val > {{ $r.Lte }} && val < {{ $r.Gte }} {
					err := {{ err . (print $t ".out_of_range_inclusive") "value must be outside range (" $r.GetLte ", " $r.GetGte ")" }}
					if !all { return err }
					errors = append(errors, err)
				}
			{{ end }}
		{{ else }}
			if {{ accessor . }} > {{ $r.Lte }} {
				err := {{ err . (print $t ".lte") "value must be less than or equal to " $r.GetLte }}
				if !all { return err }
				errors = append(errors, err)
			}
		{{ end }}
	{{ else if $r.Gt }}
		if {{ accessor . }} <= {{ $r.Gt }} {
			err := {{ err . (print $t "." "gt") "value must be greater than " $r.GetGt }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ else if $r.Gte }}
		if {{ accessor . }} < {{ $r.Gte }} {
			err := {{ err . (print $t "." "gte") "value must be greater than or equal to " $r.GetGte }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ end }}
`
