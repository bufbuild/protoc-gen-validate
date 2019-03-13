package python

const timestampTpl = `
{{ $f := .Field }}{{ $r := .Rules }}
	{{ template "required" . }}
	{{ if or $r.Lt $r.Lte $r.Gt $r.Gte $r.LtNow $r.GtNow $r.Within $r.Const }}
	if {{ accessor . }} != EMPTY_TIMESTAMP:
		try:
			ts = ({{ accessor . }}.seconds + (10**-9 * {{ accessor . }}.nanos))
			{{ template "timestampcmp" . }}
		except:
			{{ err . "err" "value is not a valid timestamp" }}
	{{ end }}`

const timestampcmp = `
{{ $f := .Field }}{{ $r := .Rules }}
			{{ if $r.Const }}
			if ts != {{ tsLit $r.Const }}:
				{{ err . "value must equal " (tsStr $r.Const) }}
			{{ end }}
			{{ if or $r.LtNow $r.GtNow $r.Within }}
			now = datetime.datetime.now()
			{{ end }}
			{{- if $r.Lt }}
			lt  = {{ tsLit $r.Lt }}
			{{ end }}
			{{- if $r.Lte }}
			lte = {{ tsLit $r.Lte }}
			{{ end }}
			{{- if $r.Gt }}
			gt  = {{ tsLit $r.Gt }}
			{{ end }}
			{{- if $r.Gte }}
			gte = {{ tsLit $r.Gte }}
			{{ end }}
			{{- if $r.Within }}
			within = {{ durLit $r.Within }}
			{{ end }}

			{{ if $r.Lt }}
				{{ if $r.Gt }}
					{{  if tsGt $r.GetLt $r.GetGt }}
			if ts <= gt or ts >= lt:
				{{ err . "value must be inside range (" (tsStr $r.GetGt) ", " (tsStr $r.GetLt) ")" }}
					{{ else }}
			if ts >= lt and ts <= gt:
				{{ err . "value must be outside range [" (tsStr $r.GetLt) ", " (tsStr $r.GetGt) "]" }}
					{{ end }}
				{{ else if $r.Gte }}
					{{  if tsGt $r.GetLt $r.GetGte }}
			if ts < gte or ts >= lt:
				{{ err . "value must be inside range [" (tsStr $r.GetGte) ", " (tsStr $r.GetLt) ")" }}
					{{ else }}
			if ts >= lt and ts < gte:
				{{ err . "value must be outside range [" (tsStr $r.GetLt) ", " (tsStr $r.GetGte) ")" }}
					{{ end }}
				{{ else }}
			if ts >= lt:
				{{ err . "value must be less than " (tsStr $r.GetLt) }}
				{{ end }}
			{{ else if $r.Lte }}
				{{ if $r.Gt }}
					{{  if tsGt $r.GetLte $r.GetGt }}
			if ts <= gt or ts > lte:
				{{ err . "value must be inside range (" (tsStr $r.GetGt) ", " (tsStr $r.GetLte) "]" }}
					{{ else }}
			if ts > lte and ts <= gt:
				{{ err . "value must be outside range (" (tsStr $r.GetLte) ", " (tsStr $r.GetGt) "]" }}
					{{ end }}
				{{ else if $r.Gte }}
					{{ if tsGt $r.GetLte $r.GetGte }}
			if ts < gte or ts > lte:
				{{ err . "value must be inside range [" (tsStr $r.GetGte) ", " (tsStr $r.GetLte) "]" }}
					{{ else }}
			if ts > lte and ts < gte:
				{{ err . "value must be outside range (" (tsStr $r.GetLte) ", " (tsStr $r.GetGte) ")" }}
					{{ end }}
				{{ else }}
			if ts > lte:
				{{ err . "value must be less than or equal to " (tsStr $r.GetLte) }}
				{{ end }}
			{{ else if $r.Gt }}
			if ts <= gt:
				{{ err . "value must be greater than " (tsStr $r.GetGt) }}
			{{ else if $r.Gte }}
			if ts < gte:
				{{ err . "value must be greater than or equal to " (tsStr $r.GetGte) }}
			{{ else if $r.LtNow }}
				{{ if $r.Within }}
			if ts >= now or ts < now - within:
				{{ err . "value must be less than now within " (durStr $r.GetWithin) }}
				{{ else }}
			if ts >= now:
				{{ err . "value must be less than now" }}
				{{ end }}
			{{ else if $r.GtNow }}
				{{ if $r.Within }}
			if ts >= now or ts > now + within:
				{{ err . "value must be greater than now within " (durStr $r.GetWithin) }}
				{{ else }}
			if ts <= now:
				{{ err . "value must be greater than now" }}
				{{ end }}
			{{ else if $r.Within }}
			if ts >= now + within or ts <= now - within:
				{{ err . "value must be within " (durStr $r.GetWithin) " of now" }}
			{{ end }}`
