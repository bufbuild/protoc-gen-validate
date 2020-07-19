package dotnet

const msgTpl = `
{{ range parentClassNames . }}
	public {{ if eq . "Types" }}static{{ else }}sealed{{ end }} partial class {{ . }}
	{
{{ end }}
public sealed partial class {{ className . }}: Envoyproxy.Validator.IValidateable
{
	{{ range .Fields }}
		{{ renderConst (context .) }}
	{{ end }}

	void Envoyproxy.Validator.IValidateable.Validate()
	{
		{{ if disabled . }}
			// Validate is disabled for {{ className . }}
			return;
		{{ else }}
			{{ range .NonOneOfFields }}
				{{ render (context .) }}
			{{ end }}
			{{ range .OneOfs }}
				switch ({{ oneofAccessor . }}) {
					{{ range .Fields -}}
						case {{ oneofCase . }}:
							{{ render (context .) }}
							break;
					{{ end -}}
					{{ if required . }}
						default:
							throw {{ unboundErr . "value is required" }};
					{{ end }}
				}
			{{ end }}
		{{ end }}
	}
}
{{ range parentClassNames . }}
	}
{{ end }}
`
