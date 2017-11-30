package tpl

const durationTpl = `{{ $f := .Field }}{{ $r := .Rules }}
	{{ template "required" . }}

	{{ if or $r.In $r.NotIn $r.Lt $r.Lte $r.Gt $r.Gte $r.Const }}
	        if (!{{ hasAccessor . }}) {
			return true;
	        }

	        const pgv::protobuf::Duration& dur = {{ accessor . }};

		if (dur.nanos() > 999999999 || dur.nanos() < -999999999 ||
		    dur.seconds() > pgv::protobuf::util::TimeUtil::kDurationMaxSeconds ||
		    dur.seconds() < pgv::protobuf::util::TimeUtil::kDurationMinSeconds)
	                {{ errCause . "err" "value is not a valid duration" }}

		const int64_t nanos = pgv::protobuf::util::TimeUtil::DurationToNanoseconds(dur);

		{{  if $r.Const }}
			if (nanos != {{ durLit $r.Const }})
				{{ err . "value must equal " (durStr $r.Const) }}
		{{ end }}

		{{  if $r.Lt }}  const int64_t lt  = {{ durLit $r.Lt }};  {{ end }}
		{{- if $r.Lte }} const int64_t lte = {{ durLit $r.Lte }}; {{ end }}
		{{- if $r.Gt }}  const int64_t gt  = {{ durLit $r.Gt }};  {{ end }}
		{{- if $r.Gte }} const int64_t gte = {{ durLit $r.Gte }}; {{ end }}

		{{ if $r.Lt }}
			{{ if $r.Gt }}
				{{  if durGt $r.GetLt $r.GetGt }}
					if (nanos <= gt || nanos >= lt)
						{{ err . "value must be inside range (" (durStr $r.GetGt) ", " (durStr $r.GetLt) ")" }}
				{{ else }}
					if (nanos >= lt && nanos <= gt)
						{{ err . "value must be outside range [" (durStr $r.GetLt) ", " (durStr $r.GetGt) "]" }}
				{{ end }}
			{{ else if $r.Gte }}
				{{  if durGt $r.GetLt $r.GetGte }}
					if (nanos < gte || nanos >= lt)
						{{ err . "value must be inside range [" (durStr $r.GetGte) ", " (durStr $r.GetLt) ")" }}
				{{ else }}
					if (nanos >= lt && nanos < gte)
						{{ err . "value must be outside range [" (durStr $r.GetLt) ", " (durStr $r.GetGte) ")" }}
				{{ end }}
			{{ else }}
				if (nanos >= lt)
					{{ err . "value must be less than " (durStr $r.GetLt) }}
			{{ end }}
		{{ else if $r.Lte }}
			{{ if $r.Gt }}
				{{  if durGt $r.GetLte $r.GetGt }}
					if (nanos <= gt || nanos > lte)
						{{ err . "value must be inside range (" (durStr $r.GetGt) ", " (durStr $r.GetLte) "]" }}
				{{ else }}
					if (nanos > lte && nanos <= gt)
						{{ err . "value must be outside range (" (durStr $r.GetLte) ", " (durStr $r.GetGt) "]" }}
				{{ end }}
			{{ else if $r.Gte }}
				{{ if durGt $r.GetLte $r.GetGte }}
					if (nanos < gte || nanos > lte)
						{{ err . "value must be inside range [" (durStr $r.GetGte) ", " (durStr $r.GetLte) "]" }}
				{{ else }}
					if (nanos > lte && nanos < gte)
						{{ err . "value must be outside range (" (durStr $r.GetLte) ", " (durStr $r.GetGte) ")" }}
				{{ end }}
			{{ else }}
				if (nanos > lte)
					{{ err . "value must be less than or equal to " (durStr $r.GetLte) }}
			{{ end }}
		{{ else if $r.Gt }}
			if (nanos <= gt)
				{{ err . "value must be greater than " (durStr $r.GetGt) }}
		{{ else if $r.Gte }}
			if (nanos < gte)
				{{ err . "value must be greater than or equal to " (durStr $r.GetGte) }}
		{{ end }}


		{{ if $r.In }}
			if ({{ lookup $f "InLookup" }}.find(nanos) == {{ lookup $f "InLookup" }}.end())
				{{ err . "value must be in list " $r.In }}
		{{ else if $r.NotIn }}
			if ({{ lookup $f "NotInLookup" }}.find(nanos) != {{ lookup $f "NotInLookup" }}.end())
				{{ err . "value must not be in list " $r.NotIn }}
		{{ end }}
	{{ end }}
`
