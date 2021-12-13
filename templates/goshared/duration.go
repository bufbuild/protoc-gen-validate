package goshared

const durationcmpTpl = `{{ $f := .Field }}{{ $r := .Rules }}
			{{  if $r.Const }}
				if dur != {{ durLit $r.Const }} {
					err := {{ err . (t "duration.const" "value must equal {{$1}}" (durStr $r.Const)) }}
					if !all { return err }
					errors = append(errors, err)
				}
			{{ end }}


			{{  if $r.Lt }}  lt  := {{ durLit $r.Lt }};  {{ end }}
			{{- if $r.Lte }} lte := {{ durLit $r.Lte }}; {{ end }}
			{{- if $r.Gt }}  gt  := {{ durLit $r.Gt }};  {{ end }}
			{{- if $r.Gte }} gte := {{ durLit $r.Gte }}; {{ end }}

			{{ if $r.Lt }}
				{{ if $r.Gt }}
					{{  if durGt $r.GetLt $r.GetGt }}
						if dur <= gt || dur >= lt {
							err := {{ err . (t "duration.between_open" "value must be inside range ({{$1}}, {{$2}})" (durStr $r.GetGt) (durStr $r.GetLt)) }}
							if !all { return err }
							errors = append(errors, err)
						}
					{{ else }}
						if dur >= lt && dur <= gt {
							err := {{ err . (t "duration.outside_closed" "value must be outside range [{{$1}}, {{$2}}]" (durStr $r.GetLt) (durStr $r.GetGt)) }}
							if !all { return err }
							errors = append(errors, err)
						}
					{{ end }}
				{{ else if $r.Gte }}
					{{  if durGt $r.GetLt $r.GetGte }}
						if dur < gte || dur >= lt {
							err := {{ err . (t "duration.between_include_left" "value must be inside range [{{$1}}, {{$2}})" (durStr $r.GetGte) (durStr $r.GetLt)) }}
							if !all { return err }
							errors = append(errors, err)
						}
					{{ else }}
						if dur >= lt && dur < gte {
							err := {{ err . (t "duration.outside_include_left" "value must be outside range [{{$1}}, {{$2}})" (durStr $r.GetLt) (durStr $r.GetGte)) }}
							if !all { return err }
							errors = append(errors, err)
						}
					{{ end }}
				{{ else }}
					if dur >= lt {
						err := {{ err . (t "duration.lt" "value must be less than {{$1}}" (durStr $r.GetLt)) }}
						if !all { return err }
						errors = append(errors, err)
					}
				{{ end }}
			{{ else if $r.Lte }}
				{{ if $r.Gt }}
					{{  if durGt $r.GetLte $r.GetGt }}
						if dur <= gt || dur > lte {
							err := {{ err . (t "duration.between_include_right" "value must be inside range ({{$1}}, {{$2}}]" (durStr $r.GetGt) (durStr $r.GetLte)) }}
							if !all { return err }
							errors = append(errors, err)
						}
					{{ else }}
						if dur > lte && dur <= gt {
							err := {{ err . (t "duration.outside_include_right" "value must be outside range ({{$1}}, {{$2}}]" (durStr $r.GetLte) (durStr $r.GetGt)) }}
							if !all { return err }
							errors = append(errors, err)
						}
					{{ end }}
				{{ else if $r.Gte }}
					{{ if durGt $r.GetLte $r.GetGte }}
						if dur < gte || dur > lte {
							err := {{ err . (t "duration.between_closed" "value must be inside range [{{$1}}, {{$2}}]" (durStr $r.GetGte) (durStr $r.GetLte)) }}
							if !all { return err }
							errors = append(errors, err)
						}
					{{ else }}
						if dur > lte && dur < gte {
							err := {{ err . (t "duration.outside_open" "value must be outside range ({{$1}}, {{$2}})" (durStr $r.GetLte) (durStr $r.GetGte)) }}
							if !all { return err }
							errors = append(errors, err)
						}
					{{ end }}
				{{ else }}
					if dur > lte {
						err := {{ err . (t "duration.lte" "value must be less than or equal to {{$1}}" (durStr $r.GetLte)) }}
						if !all { return err }
						errors = append(errors, err)
					}
				{{ end }}
			{{ else if $r.Gt }}
				if dur <= gt {
					err := {{ err . (t "duration.gt" "value must be greater than {{$1}}" (durStr $r.GetGt)) }}
					if !all { return err }
					errors = append(errors, err)
				}
			{{ else if $r.Gte }}
				if dur < gte {
					err := {{ err . (t "duration.gte" "value must be greater than or equal to {{$1}}" (durStr $r.GetGte)) }}
					if !all { return err }
					errors = append(errors, err)
				}
			{{ end }}


			{{ if $r.In }}
				if _, ok := {{ lookup $f "InLookup" }}[dur]; !ok {
					err := {{ err . (t "duration.in" "value must be in list {{$1}}" $r.In) }}
					if !all { return err }
					errors = append(errors, err)
				}
			{{ else if $r.NotIn }}
				if _, ok := {{ lookup $f "NotInLookup" }}[dur]; ok {
					err := {{ err . (t "duration.not_in" "value must not be in list {{$1}}" $r.NotIn) }}
					if !all { return err }
					errors = append(errors, err)
				}
			{{ end }}
`
