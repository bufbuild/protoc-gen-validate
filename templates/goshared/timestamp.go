package goshared

const timestampcmpTpl = `{{ $f := .Field }}{{ $r := .Rules }}
			{{  if $r.Const }}
				if !ts.Equal({{ tsLit $r.Const }}) {
					{{ if ne $r.GetErrorMsg "" }}
					err := {{ err . $r.GetErrorMsg }}
					{{ else }}
					err := {{ err . "value must equal " (tsStr $r.Const) }}
					{{ end }}
					if !all { return err }
					errors = append(errors, err)
				}
			{{ end }}

			{{ if or $r.LtNow $r.GtNow $r.Within }} now := time.Now(); {{ end }}
			{{- if $r.Lt }}  lt  := {{ tsLit $r.Lt }};  {{ end }}
			{{- if $r.Lte }} lte := {{ tsLit $r.Lte }}; {{ end }}
			{{- if $r.Gt }}  gt  := {{ tsLit $r.Gt }};  {{ end }}
			{{- if $r.Gte }} gte := {{ tsLit $r.Gte }}; {{ end }}
			{{- if $r.Within }} within := {{ durLit $r.Within }}; {{ end }}

			{{ if $r.Lt }}
				{{ if $r.Gt }}
					{{  if tsGt $r.GetLt $r.GetGt }}
						if ts.Sub(gt) <= 0 || ts.Sub(lt) >= 0 {
							{{ if ne $r.GetErrorMsg "" }}
							err := {{ err . $r.GetErrorMsg }}
							{{ else }}
							err := {{ err . "value must be inside range (" (tsStr $r.GetGt) ", " (tsStr $r.GetLt) ")" }}
							{{ end }}
							if !all { return err }
							errors = append(errors, err)
						}
					{{ else }}
						if ts.Sub(lt) >= 0 && ts.Sub(gt) <= 0 {
							{{ if ne $r.GetErrorMsg "" }}
							err := {{ err . $r.GetErrorMsg }}
							{{ else }}
							err := {{ err . "value must be outside range [" (tsStr $r.GetLt) ", " (tsStr $r.GetGt) "]" }}
							{{ end }}
							if !all { return err }
							errors = append(errors, err)
						}
					{{ end }}
				{{ else if $r.Gte }}
					{{  if tsGt $r.GetLt $r.GetGte }}
						if ts.Sub(gte) < 0 || ts.Sub(lt) >= 0 {
							{{ if ne $r.GetErrorMsg "" }}
							err := {{ err . $r.GetErrorMsg }}
							{{ else }}
							err := {{ err . "value must be inside range [" (tsStr $r.GetGte) ", " (tsStr $r.GetLt) ")" }}
							{{ end }}
							if !all { return err }
							errors = append(errors, err)
						}
					{{ else }}
						if ts.Sub(lt) >= 0 && ts.Sub(gte) < 0 {
							{{ if ne $r.GetErrorMsg "" }}
							err := {{ err . $r.GetErrorMsg }}
							{{ else }}
							err := {{ err . "value must be outside range [" (tsStr $r.GetLt) ", " (tsStr $r.GetGte) ")" }}
							{{ end }}
							if !all { return err }
							errors = append(errors, err)
						}
					{{ end }}
				{{ else }}
					if ts.Sub(lt) >= 0 {
						{{ if ne $r.GetErrorMsg "" }}
						err := {{ err . $r.GetErrorMsg }}
						{{ else }}
						err := {{ err . "value must be less than " (tsStr $r.GetLt) }}
						{{ end }}
						if !all { return err }
						errors = append(errors, err)
					}
				{{ end }}
			{{ else if $r.Lte }}
				{{ if $r.Gt }}
					{{  if tsGt $r.GetLte $r.GetGt }}
						if ts.Sub(gt) <= 0 || ts.Sub(lte) > 0 {
							{{ if ne $r.GetErrorMsg "" }}
							err := {{ err . $r.GetErrorMsg }}
							{{ else }}
							err := {{ err . "value must be inside range (" (tsStr $r.GetGt) ", " (tsStr $r.GetLte) "]" }}
							{{ end }}
							if !all { return err }
							errors = append(errors, err)
						}
					{{ else }}
						if ts.Sub(lte) > 0 && ts.Sub(gt) <= 0 {
							{{ if ne $r.GetErrorMsg "" }}
							err := {{ err . $r.GetErrorMsg }}
							{{ else }}
							err := {{ err . "value must be outside range (" (tsStr $r.GetLte) ", " (tsStr $r.GetGt) "]" }}
							{{ end }}
							if !all { return err }
							errors = append(errors, err)
						}
					{{ end }}
				{{ else if $r.Gte }}
					{{ if tsGt $r.GetLte $r.GetGte }}
						if ts.Sub(gte) < 0 || ts.Sub(lte) > 0 {
							{{ if ne $r.GetErrorMsg "" }}
							err := {{ err . $r.GetErrorMsg }}
							{{ else }}
							err := {{ err . "value must be inside range [" (tsStr $r.GetGte) ", " (tsStr $r.GetLte) "]" }}
							{{ end }}
							if !all { return err }
							errors = append(errors, err)
						}
					{{ else }}
						if ts.Sub(lte) > 0 && ts.Sub(gte) < 0 {
							{{ if ne $r.GetErrorMsg "" }}
							err := {{ err . $r.GetErrorMsg }}
							{{ else }}
							err := {{ err . "value must be outside range (" (tsStr $r.GetLte) ", " (tsStr $r.GetGte) ")" }}
							{{ end }}
							if !all { return err }
							errors = append(errors, err)
						}
					{{ end }}
				{{ else }}
					if ts.Sub(lte) > 0 {
						{{ if ne $r.GetErrorMsg "" }}
						err := {{ err . $r.GetErrorMsg }}
						{{ else }}
						err := {{ err . "value must be less than or equal to " (tsStr $r.GetLte) }}
						{{ end }}
						if !all { return err }
						errors = append(errors, err)
					}
				{{ end }}
			{{ else if $r.Gt }}
				if ts.Sub(gt) <= 0 {
					{{ if ne $r.GetErrorMsg "" }}
					err := {{ err . $r.GetErrorMsg }}
					{{ else }}
					err := {{ err . "value must be greater than " (tsStr $r.GetGt) }}
					{{ end }}
					if !all { return err }
					errors = append(errors, err)
				}
			{{ else if $r.Gte }}
				if ts.Sub(gte) < 0 {
					{{ if ne $r.GetErrorMsg "" }}
					err := {{ err . $r.GetErrorMsg }}
					{{ else }}
					err := {{ err . "value must be greater than or equal to " (tsStr $r.GetGte) }}
					{{ end }}
					if !all { return err }
					errors = append(errors, err)
				}
			{{ else if $r.LtNow }}
				{{ if $r.Within }}
					if ts.Sub(now) >= 0 || ts.Sub(now.Add(-within)) < 0 {
						{{ if ne $r.GetErrorMsg "" }}
						err := {{ err . $r.GetErrorMsg }}
						{{ else }}
						err := {{ err . "value must be less than now within " (durStr $r.GetWithin) }}
						{{ end }}
						if !all { return err }
						errors = append(errors, err)
					}
				{{ else }}
					if ts.Sub(now) >= 0 {
						{{ if ne $r.GetErrorMsg "" }}
						err := {{ err . $r.GetErrorMsg }}
						{{ else }}
						err := {{ err . "value must be less than now" }}
						{{ end }}
						if !all { return err }
						errors = append(errors, err)
					}
				{{ end }}
			{{ else if $r.GtNow }}
				{{ if $r.Within }}
					if ts.Sub(now) <= 0 || ts.Sub(now.Add(within)) > 0 {
						{{ if ne $r.GetErrorMsg "" }}
						err := {{ err . $r.GetErrorMsg }}
						{{ else }}
						err := {{ err . "value must be greater than now within " (durStr $r.GetWithin) }}
						{{ end }}
						if !all { return err }
						errors = append(errors, err)
					}
				{{ else }}
					if ts.Sub(now) <= 0 {
						{{ if ne $r.GetErrorMsg "" }}
						err := {{ err . $r.GetErrorMsg }}
						{{ else }}
						err := {{ err . "value must be greater than now" }}
						{{ end }}
						if !all { return err }
						errors = append(errors, err)
					}
				{{ end }}
			{{ else if $r.Within }}
				if ts.Sub(now.Add(within)) >= 0 || ts.Sub(now.Add(-within)) <= 0 {
					{{ if ne $r.GetErrorMsg "" }}
					err := {{ err . $r.GetErrorMsg }}
					{{ else }}
					err := {{ err . "value must be within " (durStr $r.GetWithin) " of now" }}
					{{ end }}
					if !all { return err }
					errors = append(errors, err)
				}
			{{ end }}
`
