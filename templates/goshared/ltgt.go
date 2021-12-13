package goshared

const ltgtTpl = `{{ $f := .Field }}{{ $r := .Rules }}
	{{ if $r.Lt }}
		{{ if $r.Gt }}
			{{  if gt $r.GetLt $r.GetGt }}
				if val := {{ accessor . }};  val <= {{ $r.Gt }} || val >= {{ $r.Lt }} {
					err := {{ err . (t "<prefix>.between_open" "value must be inside range ({{$1}}, {{$2}})" $r.GetGt $r.GetLt) }}
					if !all { return err }
					errors = append(errors, err)
				}
			{{ else }}
				if val := {{ accessor . }}; val >= {{ $r.Lt }} && val <= {{ $r.Gt }} {
					err := {{ err . (t "<prefix>.outside_closed" "value must be outside range [{{$1}}, {{$2}}]" $r.GetLt $r.GetGt) }}
					if !all { return err }
					errors = append(errors, err)
				}
			{{ end }}
		{{ else if $r.Gte }}
			{{  if gt $r.GetLt $r.GetGte }}
				if val := {{ accessor . }};  val < {{ $r.Gte }} || val >= {{ $r.Lt }} {
					err := {{ err . (t "<prefix>.between_include_left" "value must be inside range [{{$1}}, {{$2}})" $r.GetGte $r.GetLt) }}
					if !all { return err }
					errors = append(errors, err)
				}
			{{ else }}
				if val := {{ accessor . }}; val >= {{ $r.Lt }} && val < {{ $r.Gte }} {
					err := {{ err . (t "<prefix>.outside_include_left" "value must be outside range [{{$1}}, {{$2}})" $r.GetLt $r.GetGte) }}
					if !all { return err }
					errors = append(errors, err)
				}
			{{ end }}
		{{ else }}
			if {{ accessor . }} >= {{ $r.Lt }} {
				err := {{ err . (t "<prefix>.lt" "value must be less than {{$1}}" $r.GetLt) }}
				if !all { return err }
				errors = append(errors, err)
			}
		{{ end }}
	{{ else if $r.Lte }}
		{{ if $r.Gt }}
			{{  if gt $r.GetLte $r.GetGt }}
				if val := {{ accessor . }};  val <= {{ $r.Gt }} || val > {{ $r.Lte }} {
					err := {{ err . (t "<prefix>.between_include_right" "value must be inside range ({{$1}}, {{$2}}]" $r.GetGt $r.GetLte) }}
					if !all { return err }
					errors = append(errors, err)
				}
			{{ else }}
				if val := {{ accessor . }}; val > {{ $r.Lte }} && val <= {{ $r.Gt }} {
					err := {{ err . (t "<prefix>.outside_include_right" "value must be outside range ({{$1}}, {{$2}}]" $r.GetLte $r.GetGt) }}
					if !all { return err }
					errors = append(errors, err)
				}
			{{ end }}
		{{ else if $r.Gte }}
			{{ if gt $r.GetLte $r.GetGte }}
				if val := {{ accessor . }};  val < {{ $r.Gte }} || val > {{ $r.Lte }} {
					err := {{ err . (t "<prefix>.between_closed" "value must be inside range [{{$1}}, {{$2}}]" $r.GetGte $r.GetLte) }}
					if !all { return err }
					errors = append(errors, err)
				}
			{{ else }}
				if val := {{ accessor . }}; val > {{ $r.Lte }} && val < {{ $r.Gte }} {
					err := {{ err . (t "<prefix>.outside_open" "value must be outside range ({{$1}}, {{$2}})" $r.GetLte $r.GetGte) }}
					if !all { return err }
					errors = append(errors, err)
				}
			{{ end }}
		{{ else }}
			if {{ accessor . }} > {{ $r.Lte }} {
				err := {{ err . (t "<prefix>.lte" "value must be less than or equal to {{$1}}" $r.GetLte) }}
				if !all { return err }
				errors = append(errors, err)
			}
		{{ end }}
	{{ else if $r.Gt }}
		if {{ accessor . }} <= {{ $r.Gt }} {
			err := {{ err . (t "<prefix>.gt" "value must be greater than {{$1}}" $r.GetGt) }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ else if $r.Gte }}
		if {{ accessor . }} < {{ $r.Gte }} {
			err := {{ err . (t "<prefix>.gte" "value must be greater than or equal to {{$1}}" $r.GetGte) }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ end }}
`
