package goshared

const durationcmpTpl = `{{ $f := .Field }}{{ $r := .Rules }}
			{{  if $r.Const }}
				if dur != {{ durLit $r.Const }} {
					{{ if ne $r.GetErrorMsg "" }}
					err := {{ err . $r.GetErrorMsg }}
					{{ else }}
					err := {{ err . "value must equal " (durStr $r.Const) }}
					{{ end }}
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
							{{ if ne $r.GetErrorMsg "" }}
							err := {{ err . $r.GetErrorMsg }}
							{{ else }}
							err := {{ err . "value must be inside range (" (durStr $r.GetGt) ", " (durStr $r.GetLt) ")" }}
							{{ end }}
							if !all { return err }
							errors = append(errors, err)
						}
					{{ else }}
						if dur >= lt && dur <= gt {
							{{ if ne $r.GetErrorMsg "" }}
							err := {{ err . $r.GetErrorMsg }}
							{{ else }}
							err := {{ err . "value must be outside range [" (durStr $r.GetLt) ", " (durStr $r.GetGt) "]" }}
							{{ end }}
							if !all { return err }
							errors = append(errors, err)
						}
					{{ end }}
				{{ else if $r.Gte }}
					{{  if durGt $r.GetLt $r.GetGte }}
						if dur < gte || dur >= lt {
							{{ if ne $r.GetErrorMsg "" }}
							err := {{ err . $r.GetErrorMsg }}
							{{ else }}
							err := {{ err . "value must be inside range [" (durStr $r.GetGte) ", " (durStr $r.GetLt) ")" }}
							{{ end }}
							if !all { return err }
							errors = append(errors, err)
						}
					{{ else }}
						if dur >= lt && dur < gte {
							{{ if ne $r.GetErrorMsg "" }}
							err := {{ err . $r.GetErrorMsg }}
							{{ else }}
							err := {{ err . "value must be outside range [" (durStr $r.GetLt) ", " (durStr $r.GetGte) ")" }}
							{{ end }}
							if !all { return err }
							errors = append(errors, err)
						}
					{{ end }}
				{{ else }}
					if dur >= lt {
						{{ if ne $r.GetErrorMsg "" }}
						err := {{ err . $r.GetErrorMsg }}
						{{ else }}
						err := {{ err . "value must be less than " (durStr $r.GetLt) }}
						{{ end }}
						if !all { return err }
						errors = append(errors, err)
					}
				{{ end }}
			{{ else if $r.Lte }}
				{{ if $r.Gt }}
					{{  if durGt $r.GetLte $r.GetGt }}
						if dur <= gt || dur > lte {
							{{ if ne $r.GetErrorMsg "" }}
							err := {{ err . $r.GetErrorMsg }}
							{{ else }}
							err := {{ err . "value must be inside range (" (durStr $r.GetGt) ", " (durStr $r.GetLte) "]" }}
							{{ end }}
							if !all { return err }
							errors = append(errors, err)
						}
					{{ else }}
						if dur > lte && dur <= gt {
							{{ if ne $r.GetErrorMsg "" }}
							err := {{ err . $r.GetErrorMsg }}
							{{ else }}
							err := {{ err . "value must be outside range (" (durStr $r.GetLte) ", " (durStr $r.GetGt) "]" }}
							{{ end }}
							if !all { return err }
							errors = append(errors, err)
						}
					{{ end }}
				{{ else if $r.Gte }}
					{{ if durGt $r.GetLte $r.GetGte }}
						if dur < gte || dur > lte {
							{{ if ne $r.GetErrorMsg "" }}
							err := {{ err . $r.GetErrorMsg }}
							{{ else }}
							err := {{ err . "value must be inside range [" (durStr $r.GetGte) ", " (durStr $r.GetLte) "]" }}
							{{ end }}
							if !all { return err }
							errors = append(errors, err)
						}
					{{ else }}
						if dur > lte && dur < gte {
							{{ if ne $r.GetErrorMsg "" }}
							err := {{ err . $r.GetErrorMsg }}
							{{ else }}
							err := {{ err . "value must be outside range (" (durStr $r.GetLte) ", " (durStr $r.GetGte) ")" }}
							{{ end }}
							if !all { return err }
							errors = append(errors, err)
						}
					{{ end }}
				{{ else }}
					if dur > lte {
						{{ if ne $r.GetErrorMsg "" }}
						err := {{ err . $r.GetErrorMsg }}
						{{ else }}
						err := {{ err . "value must be less than or equal to " (durStr $r.GetLte) }}
						{{ end }}
						if !all { return err }
						errors = append(errors, err)
					}
				{{ end }}
			{{ else if $r.Gt }}
				if dur <= gt {
					{{ if ne $r.GetErrorMsg "" }}
					err := {{ err . $r.GetErrorMsg }}
					{{ else }}
					err := {{ err . "value must be greater than " (durStr $r.GetGt) }}
					{{ end }}
					if !all { return err }
					errors = append(errors, err)
				}
			{{ else if $r.Gte }}
				if dur < gte {
					{{ if ne $r.GetErrorMsg "" }}
					err := {{ err . $r.GetErrorMsg }}
					{{ else }}
					err := {{ err . "value must be greater than or equal to " (durStr $r.GetGte) }}
					{{ end }}
					if !all { return err }
					errors = append(errors, err)
				}
			{{ end }}


			{{ if $r.In }}
				if _, ok := {{ lookup $f "InLookup" }}[dur]; !ok {
					{{ if ne $r.GetErrorMsg "" }}
					err := {{ err . $r.GetErrorMsg }}
					{{ else }}
					err := {{ err . "value must be in list " $r.In }}
					{{ end }}
					if !all { return err }
					errors = append(errors, err)
				}
			{{ else if $r.NotIn }}
				if _, ok := {{ lookup $f "NotInLookup" }}[dur]; ok {
					{{ if ne $r.GetErrorMsg "" }}
					err := {{ err . $r.GetErrorMsg }}
					{{ else }}
					err := {{ err . "value must not be in list " $r.NotIn }}
					{{ end }}
					if !all { return err }
					errors = append(errors, err)
				}
			{{ end }}
`
