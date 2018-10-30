package java

const anyTpl = `{{ $f := .Field }}{{ $r := .Rules }}
	{{ template "required" . }}

  {{- if $r.In }}
  			{
  				String[] set = new String[]{
  					{{- range $r.In }}
  					"{{ . }}",
  					{{- end }}
  				};
  				com.lyft.pgv.AnyValidation.in("{{ $f.FullyQualifiedName }}", proto.{{ accessor $f }}, set);
  			}
  {{- end -}}
  {{- if $r.NotIn }}
  			{
  				String[] set = new String[]{
  					{{- range $r.NotIn }}
  					"{{ . }}",
  					{{- end }}
  				};
  				com.lyft.pgv.AnyValidation.notIn("{{ $f.FullyQualifiedName }}", proto.{{ accessor $f }}, set);
  			}
  {{- end -}}
`
