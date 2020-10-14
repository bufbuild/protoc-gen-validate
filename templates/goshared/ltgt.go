package goshared

const ltgtTpl = `{{ $f := .Field }}{{ $r := .Rules }}
	{{ if $r.Lt }}
		{{ if $r.Gt }}
			{{  if gt $r.GetLt $r.GetGt }}
				if val := {{ accessor . }};  val <= {{ $r.Gt }} || val >= {{ $r.Lt }} {
					return {{ err . (t "<prefix>.between_open" "value must be inside range ({{$1}}, {{$2}})" $r.GetGt $r.GetLt) }}
				}
			{{ else }}
				if val := {{ accessor . }}; val >= {{ $r.Lt }} && val <= {{ $r.Gt }} {
					return {{ err . (t "<prefix>.outside_closed" "value must be outside range [{{$1}}, {{$2}}]" $r.GetLt $r.GetGt) }}
				}
			{{ end }}
		{{ else if $r.Gte }}
			{{  if gt $r.GetLt $r.GetGte }}
				if val := {{ accessor . }};  val < {{ $r.Gte }} || val >= {{ $r.Lt }} {
					return {{ err . (t "<prefix>.between_include_left" "value must be inside range [{{$1}}, {{$2}})" $r.GetGte $r.GetLt) }}
				}
			{{ else }}
				if val := {{ accessor . }}; val >= {{ $r.Lt }} && val < {{ $r.Gte }} {
					return {{ err . (t "<prefix>.outside_include_left" "value must be outside range [{{$1}}, {{$2}})" $r.GetLt $r.GetGte) }}
				}
			{{ end }}
		{{ else }}
			if {{ accessor . }} >= {{ $r.Lt }} {
				return {{ err . (t "<prefix>.lt" "value must be less than " $r.GetLt) }}
			}
		{{ end }}
	{{ else if $r.Lte }}
		{{ if $r.Gt }}
			{{  if gt $r.GetLte $r.GetGt }}
				if val := {{ accessor . }};  val <= {{ $r.Gt }} || val > {{ $r.Lte }} {
					return {{ err . (t "<prefix>.between_include_right" "value must be inside range ({{$1}}, {{$2}}]" $r.GetGt $r.GetLte) }}
				}
			{{ else }}
				if val := {{ accessor . }}; val > {{ $r.Lte }} && val <= {{ $r.Gt }} {
					return {{ err . (t "<prefix>.outside_include_right" "value must be outside range ({{$1}}, {{$2}}]" $r.GetLte $r.GetGt) }}
				}
			{{ end }}
		{{ else if $r.Gte }}
			{{ if gt $r.GetLte $r.GetGte }}
				if val := {{ accessor . }};  val < {{ $r.Gte }} || val > {{ $r.Lte }} {
					return {{ err . (t "<prefix>.between_closed" "value must be inside range [{{$1}}, {{$2}}]" $r.GetGte $r.GetLte) }}
				}
			{{ else }}
				if val := {{ accessor . }}; val > {{ $r.Lte }} && val < {{ $r.Gte }} {
					return {{ err . (t "<prefix>.outside_open" "value must be outside range ({{$1}}, {{$2}})" $r.GetLte $r.GetGte) }}
				}
			{{ end }}
		{{ else }}
			if {{ accessor . }} > {{ $r.Lte }} {
				return {{ err . (t "<prefix>.lte" "value must be less than or equal to " $r.GetLte) }}
			}
		{{ end }}
	{{ else if $r.Gt }}
		if {{ accessor . }} <= {{ $r.Gt }} {
			return {{ err . (t "<prefix>.gt" "value must be greater than " $r.GetGt) }}
		}
	{{ else if $r.Gte }}
		if {{ accessor . }} < {{ $r.Gte }} {
			return {{ err . (t "<prefix>.gte" "value must be greater than or equal to " $r.GetGte) }}
		}
	{{ end }}
`
