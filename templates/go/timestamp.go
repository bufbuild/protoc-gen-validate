package golang

const timestampTpl = `{{ $f := .Field }}{{ $r := .Rules }}
	{{ template "required" . }}

	{{ if or $r.Lt $r.Lte $r.Gt $r.Gte $r.LtNow $r.GtNow $r.Within $r.Const }}
		if t := {{ accessor . }}; t != nil {
			ts, err := t.AsTime(), t.CheckValid()
			if err != nil { return {{ errCause . "err" "value is not a valid timestamp" }} }

			{{ template "timestampcmp" . }}
		}
	{{ end }}
`
