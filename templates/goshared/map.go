package goshared

const mapTpl = `
	{{ $f := .Field }}{{ $r := .Rules }}

	{{ if $r.GetIgnoreEmpty }}
		if len({{ accessor . }}) > 0 {
	{{ end }}

	{{ if $r.GetMinPairs }}
		{{ if eq $r.GetMinPairs $r.GetMaxPairs }}
			if len({{ accessor . }}) != {{ $r.GetMinPairs }} {
				err := {{ err . (t "map.pairs" "value must contain exactly {{$1}} pair(s)" $r.GetMinPairs) }}
				if !all { return err }
				errors = append(errors, err)
			}
		{{ else if $r.MaxPairs }}
			if l := len({{ accessor . }}); l < {{ $r.GetMinPairs }} || l > {{ $r.GetMaxPairs }} {
				err := {{ err . (t "map.pairs_between" "value must contain between {{$1}} and {{$2}} pairs, inclusive" $r.GetMinPairs $r.GetMaxPairs) }}
				if !all { return err }
				errors = append(errors, err)
			}
		{{ else }}
			if len({{ accessor . }}) < {{ $r.GetMinPairs }} {
				err := {{ err . (t "map.min_pairs" "value must contain at least {{$1}} pair(s)" $r.GetMinPairs) }}
				if !all { return err }
				errors = append(errors, err)
			}
		{{ end }}
	{{ else if $r.MaxPairs }}
		if len({{ accessor . }}) > {{ $r.GetMaxPairs }} {
			err := {{ err . (t "map.max_pairs" "value must contain no more than {{$1}} pair(s)" $r.GetMaxPairs) }}
			if !all { return err }
			errors = append(errors, err)
		}
	{{ end }}

	{{ if or $r.GetNoSparse (ne (.Elem "" "").Typ "none") (ne (.Key "" "").Typ "none") }}
		{{- /* Sort the keys to make the iteration order (and therefore failure output) deterministic. */ -}}
		{
			sorted_keys := make([]{{ (typ .Field).Key }}, len({{ accessor . }}))
			i := 0
			for key := range {{ accessor . }} {
				sorted_keys[i] = key
				i++
			}
			sort.Slice(sorted_keys, func (i, j int) bool { return sorted_keys[i] < sorted_keys[j] })
			for _, key := range sorted_keys {
				val := {{ accessor .}}[key]
				_ = val

				{{ if $r.GetNoSparse }}
					if val == nil {
						err := {{ errIdx . (t "map.no_sparse" "value cannot be sparse, all pairs must be non-nil") }}
						if !all { return err }
						errors = append(errors, err)
					}
				{{ end }}

				{{ render (.Key "key" "key") }}

				{{ render (.Elem "val" "key") }}
			}
		}
	{{ end }}

	{{ if $r.GetIgnoreEmpty }}
		}
	{{ end }}

`
