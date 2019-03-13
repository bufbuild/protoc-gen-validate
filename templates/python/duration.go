package python

const durationTpl = `{{ $f := .Field }}{{ $r := .Rules }}
	{{ template "required" . }}

	{{ if or $r.In $r.NotIn $r.Lt $r.Lte $r.Gt $r.Gte $r.Const }}
	if {{ accessor . }} != EMPTY_DURATION:
		try:
			dur = ({{ accessor . }}.seconds + (10**-9 * {{ accessor . }}.nanos))
			{{ template "durationcmpTpl" .}}
		except Exception as e:
			{{ err . "err" "value is not a valid duration" }}
	{{ end }}
`

const durationcmpTpl = `
{{ $f := .Field }}{{ $r := .Rules }}

			{{  if $r.Const }}
			if dur != {{ durLit $r.Const }}:
				{{ err . "value must equal " (durStr $r.Const) }}
			{{ end }}

			{{- if $r.Lt }}
			lt  = {{ durLit $r.Lt }}
			{{ end }}
			{{- if $r.Lte }}
			lte = {{ durLit $r.Lte }}
			{{ end }}
			{{- if $r.Gt }}
			gt  = {{ durLit $r.Gt }}
			{{ end }}
			{{- if $r.Gte }}
			gte = {{ durLit $r.Gte }}
			{{ end }}

			{{ if $r.Lt }}
				{{ if $r.Gt }}
					{{  if durGt $r.GetLt $r.GetGt }}
			if dur <= gt or dur >= lt:
				{{ err . "value must be inside range (" (durStr $r.GetGt) ", " (durStr $r.GetLt) ")" }}
					{{ else }}
			if dur >= lt and dur <= gt:
				{{ err . "value must be outside range [" (durStr $r.GetLt) ", " (durStr $r.GetGt) "]" }}
					{{ end }}
				{{ else if $r.Gte }}
					{{  if durGt $r.GetLt $r.GetGte }}
			if dur < gte or dur >= lt:
				{{ err . "value must be inside range [" (durStr $r.GetGte) ", " (durStr $r.GetLt) ")" }}
					{{ else }}
			if dur >= lt and dur < gte:
				{{ err . "value must be outside range [" (durStr $r.GetLt) ", " (durStr $r.GetGte) ")" }}
					{{ end }}
				{{ else }}
			if dur >= lt:
				{{ err . "value must be less than " (durStr $r.GetLt) }}
				{{ end }}
			{{ else if $r.Lte }}
				{{ if $r.Gt }}
					{{  if durGt $r.GetLte $r.GetGt }}
			if dur <= gt or dur > lte:
				{{ err . "value must be inside range (" (durStr $r.GetGt) ", " (durStr $r.GetLte) "]" }}
					{{ else }}
			if dur > lte and dur <= gt:
				{{ err . "value must be outside range (" (durStr $r.GetLte) ", " (durStr $r.GetGt) "]" }}
					{{ end }}
				{{ else if $r.Gte }}
					{{ if durGt $r.GetLte $r.GetGte }}
			if dur < gte or dur > lte:
				{{ err . "value must be inside range [" (durStr $r.GetGte) ", " (durStr $r.GetLte) "]" }}
					{{ else }}
			if dur > lte and dur < gte:
				{{ err . "value must be outside range (" (durStr $r.GetLte) ", " (durStr $r.GetGte) ")" }}
					{{ end }}
				{{ else }}
			if dur > lte:
				{{ err . "value must be less than or equal to " (durStr $r.GetLte) }}
				{{ end }}
			{{ else if $r.Gt }}
			if dur <= gt:
				{{ err . "value must be greater than " (durStr $r.GetGt) }}
			{{ else if $r.Gte }}
			if dur < gte:
				{{ err . "value must be greater than or equal to " (durStr $r.GetGte) }}
			{{ end }}

			{{- if $r.In }}
			if {{ accessor . }} not in {{ lookup .Field "InLookup" }}:
				{{ err . "value must be in list " $r.In }}
			{{- else if $r.NotIn }}
			if {{ accessor . }} in {{ lookup .Field "NotInLookup" }}:
				{{ err . "value must be in list " $r.NotIn }}
			{{- end }}
`

