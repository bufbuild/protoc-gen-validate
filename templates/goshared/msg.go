package goshared

const msgTpl = `
{{ if not (ignored .) -}}
{{ if disabled . -}}
	{{ cmt "Validate is disabled for " (msgTyp .) ". This method will always return nil." }}
{{- else -}}
	{{ cmt "Validate checks the field values on " (msgTyp .) " with the rules defined in the proto definition for this message. If any rules are violated, an error is returned. When asked to return all errors, validation continues after first violation, and the result is a list of violation errors wrapped in " (multierrname .) ", or nil if none found. Otherwise, only the first error is returned, if any." }}
{{- end -}}
func (m {{ (msgTyp .).Pointer }}) Validate() error {
	return m.ValidateWith(false)
}

func (m {{ (msgTyp .).Pointer }}) ValidateWith(all bool) error {
	{{ if disabled . -}}
		return nil
	{{ else -}}
		if m == nil { return nil }

		var errors []error

		{{ range .NonOneOfFields }}
			{{ render (context .) }}
		{{ end }}

		{{ range .OneOfs }}
			switch m.{{ name . }}.(type) {
				{{ range .Fields }}
					case {{ oneof . }}:
						{{ render (context .) }}
				{{ end }}
				{{ if required . }}
					default:
						err := {{ errname .Message }}{
							field: "{{ name . }}",
							reason: "value is required",
						}
						if !all { return err }
						errors = append(errors, err)
				{{ end }}
			}
		{{ end }}

		if len(errors) > 0 {
			return {{ multierrname . }}(errors)
		}
		return nil
	{{ end -}}
}

{{ if needs . "hostname" }}{{ template "hostname" . }}{{ end }}

{{ if needs . "email" }}{{ template "email" . }}{{ end }}

{{ if needs . "uuid" }}{{ template "uuid" . }}{{ end }}

{{ cmt (multierrname .) " is an error wrapping multiple validation errors returned by " (msgTyp .) ".Validate(true) if the designated constraints aren't met." -}}
type {{ multierrname . }} []error

// Error returns a concatenation of all the error messages it wraps.
func (m {{ multierrname . }}) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m {{ multierrname . }}) AllErrors() []error { return m }

{{ cmt (errname .) " is the validation error returned by " (msgTyp .) ".Validate if the designated constraints aren't met." -}}
type {{ errname . }} struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e {{ errname . }}) Field() string { return e.field }

// Reason function returns reason value.
func (e {{ errname . }}) Reason() string { return e.reason }

// Cause function returns cause value.
func (e {{ errname . }}) Cause() error { return e.cause }

// Key function returns key value.
func (e {{ errname . }}) Key() bool { return e.key }

// ErrorName returns error name.
func (e {{ errname . }}) ErrorName() string { return "{{ errname . }}" }

// Error satisfies the builtin error interface
func (e {{ errname . }}) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %s{{ (msgTyp .) }}.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = {{ errname . }}{}

var _ interface{
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = {{ errname . }}{}

{{ range .Fields }}{{ with (context .) }}{{ $f := .Field }}
	{{ if has .Rules "In" }}{{ if .Rules.In }}
		var {{ lookup .Field "InLookup" }} = map[{{ inType .Field .Rules.In }}]struct{}{
			{{- range .Rules.In }}
				{{ inKey $f . }}: {},
			{{- end }}
		}
	{{ end }}{{ end }}

	{{ if has .Rules "NotIn" }}{{ if .Rules.NotIn }}
		var {{ lookup .Field "NotInLookup" }} = map[{{ inType .Field .Rules.In }}]struct{}{
			{{- range .Rules.NotIn }}
				{{ inKey $f . }}: {},
			{{- end }}
		}
	{{ end }}{{ end }}

	{{ if has .Rules "Pattern"}}{{ if .Rules.Pattern }}
		var {{ lookup .Field "Pattern" }} = regexp.MustCompile({{ lit .Rules.GetPattern }})
	{{ end }}{{ end }}

	{{ if has .Rules "Items"}}{{ if .Rules.Items }}
	{{ if has .Rules.Items.GetString_ "Pattern" }} {{ if .Rules.Items.GetString_.Pattern }}
		var {{ lookup .Field "Pattern" }} = regexp.MustCompile({{ lit .Rules.Items.GetString_.GetPattern }})
	{{ end }}{{ end }}
	{{ end }}{{ end }}

	{{ if has .Rules "Items"}}{{ if .Rules.Items }}
	{{ if has .Rules.Items.GetString_ "In" }} {{ if .Rules.Items.GetString_.In }}
		var {{ lookup .Field "InLookup" }} = map[string]struct{}{
			{{- range .Rules.Items.GetString_.In }}
				{{ inKey $f . }}: {},
			{{- end }}
		}
	{{ end }}{{ end }}
	{{ if has .Rules.Items.GetEnum "In" }} {{ if .Rules.Items.GetEnum.In }}
		var {{ lookup .Field "InLookup" }} = map[{{ inType .Field .Rules.Items.GetEnum.In }}]struct{}{
			{{- range .Rules.Items.GetEnum.In }}
				{{ inKey $f . }}: {},
			{{- end }}
		}
	{{ end }}{{ end }}
	{{ end }}{{ end }}

	{{ if has .Rules "Items"}}{{ if .Rules.Items }}
	{{ if has .Rules.Items.GetString_ "NotIn" }} {{ if .Rules.Items.GetString_.NotIn }}
		var {{ lookup .Field "NotInLookup" }} = map[string]struct{}{
			{{- range .Rules.Items.GetString_.NotIn }}
				{{ inKey $f . }}: {},
			{{- end }}
		}
	{{ end }}{{ end }}
	{{ if has .Rules.Items.GetEnum "NotIn" }} {{ if .Rules.Items.GetEnum.NotIn }}
		var {{ lookup .Field "NotInLookup" }} = map[{{ inType .Field .Rules.Items.GetEnum.NotIn }}]struct{}{
			{{- range .Rules.Items.GetEnum.NotIn }}
				{{ inKey $f . }}: {},
			{{- end }}
		}
	{{ end }}{{ end }}
	{{ end }}{{ end }}

	{{ if has .Rules "Keys"}}{{ if .Rules.Keys }}
	{{ if has .Rules.Keys.GetString_ "Pattern" }} {{ if .Rules.Keys.GetString_.Pattern }}
		var {{ lookup .Field "Pattern" }} = regexp.MustCompile({{ lit .Rules.Keys.GetString_.GetPattern }})
	{{ end }}{{ end }}
	{{ end }}{{ end }}

	{{ if has .Rules "Values"}}{{ if .Rules.Values }}
	{{ if has .Rules.Values.GetString_ "Pattern" }} {{ if .Rules.Values.GetString_.Pattern }}
		var {{ lookup .Field "Pattern" }} = regexp.MustCompile({{ lit .Rules.Values.GetString_.GetPattern }})
	{{ end }}{{ end }}
	{{ end }}{{ end }}

{{ end }}{{ end }}
{{- end -}}
`
