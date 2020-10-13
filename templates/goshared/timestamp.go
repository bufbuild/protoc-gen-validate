package goshared

const timestampcmpTpl = `{{ $f := .Field }}{{ $r := .Rules }}
			{{  if $r.Const }}
				if !ts.Equal({{ tsLit $r.Const }}) {
					return {{ err . "value must equal " (tsStr $r.Const) }}
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
							return {{ err . (t "timestamp.between_open" "value must be inside range ({{$1}}, {{$2}})" (tsStr $r.GetGt) (tsStr $r.GetLt)) }}
						}
					{{ else }}
						if ts.Sub(lt) >= 0 && ts.Sub(gt) <= 0 {
							return {{ err . (t "timestamp.outside_closed" "value must be outside range [{{$1}}, {{$2}}]" (tsStr $r.GetLt) (tsStr $r.GetGt)) }}
						}
					{{ end }}
				{{ else if $r.Gte }}
					{{  if tsGt $r.GetLt $r.GetGte }}
						if ts.Sub(gte) < 0 || ts.Sub(lt) >= 0 {
							return {{ err . (t "timestamp.between_include_left" "value must be inside range [{{$1}}, {{$2}})" (tsStr $r.GetGte) (tsStr $r.GetLt)) }}
						}
					{{ else }}
						if ts.Sub(lt) >= 0 && ts.Sub(gte) < 0 {
							return {{ err . (t "timestamp.outside_include_left" "value must be outside range [{{$1}}, {{$2}})" (tsStr $r.GetLt) (tsStr $r.GetGte)) }}
						}
					{{ end }}
				{{ else }}
					if ts.Sub(lt) >= 0 {
						return {{ err . (t "timestamp.lt" "value must be less than {{$1}}" (tsStr $r.GetLt)) }}
					}
				{{ end }}
			{{ else if $r.Lte }}
				{{ if $r.Gt }}
					{{  if tsGt $r.GetLte $r.GetGt }}
						if ts.Sub(gt) <= 0 || ts.Sub(lte) > 0 {
							return {{ err . (t "timestamp.between_include_right" "value must be inside range ({{$1}}, {{$2}}]" (tsStr $r.GetGt) (tsStr $r.GetLte)) }}
						}
					{{ else }}
						if ts.Sub(lte) > 0 && ts.Sub(gt) <= 0 {
							return {{ err . (t "timestamp.outside_include_right" "value must be outside range ({{$1}}, {{$2}}]" (tsStr $r.GetLte) (tsStr $r.GetGt)) }}
						}
					{{ end }}
				{{ else if $r.Gte }}
					{{ if tsGt $r.GetLte $r.GetGte }}
						if ts.Sub(gte) < 0 || ts.Sub(lte) > 0 {
							return {{ err . (t "timestamp.between_closed" "value must be inside range [{{$1}}, {{$2}}]" (tsStr $r.GetGte) (tsStr $r.GetLte)) }}
						}
					{{ else }}
						if ts.Sub(lte) > 0 && ts.Sub(gte) < 0 {
							return {{ err . (t "timestamp.outside_open" "value must be outside range ({{$1}}, {{$2}})" (tsStr $r.GetLte) (tsStr $r.GetGte)) }}
						}
					{{ end }}
				{{ else }}
					if ts.Sub(lte) > 0 {
						return {{ err . (t "timestamp.lte" "value must be less than or equal to {{$1}}" (tsStr $r.GetLte)) }}
					}
				{{ end }}
			{{ else if $r.Gt }}
				if ts.Sub(gt) <= 0 {
					return {{ err . (t "timestamp.gt" "value must be greater than {{$1}}" (tsStr $r.GetGt)) }}
				}
			{{ else if $r.Gte }}
				if ts.Sub(gte) < 0 {
					return {{ err . (t "timestamp.gte" "value must be greater than or equal to {{$1}}" (tsStr $r.GetGte)) }}
				}
			{{ else if $r.LtNow }}
				{{ if $r.Within }}
					if ts.Sub(now) >= 0 || ts.Sub(now.Add(-within)) < 0 {
						return {{ err . (t "timestamp.lt_now_within" "value must be less than now within {{$1}}" (durStr $r.GetWithin)) }}
					}
				{{ else }}
					if ts.Sub(now) >= 0 {
						return {{ err . (t "timestamp.lt_now" "value must be less than now") }}
					}
				{{ end }}
			{{ else if $r.GtNow }}
				{{ if $r.Within }}
					if ts.Sub(now) <= 0 || ts.Sub(now.Add(within)) > 0 {
						return {{ err . (t "timestamp.gt_now_within" "value must be greater than now within {{$1}}" (durStr $r.GetWithin)) }}
					}
				{{ else }}
					if ts.Sub(now) <= 0 {
						return {{ err . (t "timestamp.gt_now" "value must be greater than now") }}
					}
				{{ end }}
			{{ else if $r.Within }}
				if ts.Sub(now.Add(within)) >= 0 || ts.Sub(now.Add(-within)) <= 0 {
					return {{ err . (t "timestamp.within" "value must be within {{$1}} of now" (durStr $r.GetWithin)) }}
				}
			{{ end }}
`
