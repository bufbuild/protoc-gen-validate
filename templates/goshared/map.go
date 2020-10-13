package goshared

const mapTpl = `
	{{ $f := .Field }}{{ $r := .Rules }}

	{{ if $r.GetMinPairs }}
		{{ if eq $r.GetMinPairs $r.GetMaxPairs }}
			if len({{ accessor . }}) != {{ $r.GetMinPairs }} {
				return {{ err . (t "map.pairs" "value must contain exactly {{$1}} pair(s)" $r.GetMinPairs) }}
			}
		{{ else if $r.MaxPairs }}
			if l := len({{ accessor . }}); l < {{ $r.GetMinPairs }} || l > {{ $r.GetMaxPairs }} {
			 	return {{ err . (t "map.pairs_between" "value must contain between {{$1}} and {{$2}} pairs, inclusive" $r.GetMinPairs $r.GetMaxPairs) }}
			}
		{{ else }}
			if len({{ accessor . }}) < {{ $r.GetMinPairs }} {
				return {{ err . (t "map.min_pairs" "value must contain at least {{$1}} pair(s)" $r.GetMinPairs) }}
			}
		{{ end }}
	{{ else if $r.MaxPairs }}
		if len({{ accessor . }}) > {{ $r.GetMaxPairs }} {
			return {{ err . (t "map.max_pairs" "value must contain no more than {{$1}} pair(s)" $r.GetMaxPairs) }}
		}
	{{ end }}

	{{ if or $r.GetNoSparse (ne (.Elem "" "").Typ "none") (ne (.Key "" "").Typ "none") }}
		for key, val := range {{ accessor . }} {
			_ = val

			{{ if $r.GetNoSparse }}
				if val == nil {
					return {{ errIdx . "key" (t "map.no_sparse" "value cannot be sparse, all pairs must be non-nil") }}
				}
			{{ end }}

			{{ render (.Key "key" "key") }}

			{{ render (.Elem "val" "key") }}
		}
	{{ end }}
`
