package goshared

const msgTpl = `
{{ if disabled . -}}
	{{ cmt "Validate is disabled for " (msgTyp .) ". This method will always return nil." }}
{{- else -}}
	{{ cmt "Validate checks the field values on " (msgTyp .) " with the rules defined in the proto definition for this message. If any rules are violated, an error is returned." }}
{{- end -}}
func (m {{ (msgTyp .).Pointer }}) Validate() error {
	return m.ValidateWithMask(nil)
}

func (m {{ (msgTyp .).Pointer }}) ValidateWithMask(mask *field_mask.FieldMask) error {
	{{ if disabled . -}}
		return nil
	{{ else -}}
		if m == nil { return nil }

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
						if m.maskHas(mask, "{{ .Name }}") {
							return {{ errname .Message }}{
								field: "{{ name . }}",
								reason: "value is required",
							}
						}
				{{ end }}
			}
		{{ end }}

		return nil
	{{ end -}}
}

func (m {{ (msgTyp .).Pointer }}) maskHas(mask *field_mask.FieldMask, name string) bool {
	// if we don't have a mask, allow everything
	if mask == nil {
		return true
	}
	for _, path := range mask.GetPaths() {
		if name == path || strings.HasPrefix(path, name+".") {
			return true
		}
	}
	return false
}

func (m {{ (msgTyp .).Pointer }}) updateMask(mask *field_mask.FieldMask, prefix string) *field_mask.FieldMask {
	// update the mask to remove the outer level
	if mask != nil {
		paths := []string{}
		prefix += "."
		for _, path := range mask.GetPaths() {
			if strings.HasPrefix(path, prefix) {
				paths = append(paths, strings.TrimPrefix(path, prefix))
			}
		}
		if len(paths) > 0 {
			// if fields were explicitly given within the sub-message, we only
			// validate those specific fields. We remove the prefix and pass the
			// remaining fields down as a new FieldMask for sub-message validation.
			mask = &field_mask.FieldMask{Paths: paths}
		} else {
			// if a sub-message is specified in the last position of the field mask,
			// then we validate the entire sub-message. This matches the expectation
			// of FieldMask on Update operations to overwrite the entire sub-message.
			// https://developers.google.com/protocol-buffers/docs/reference/csharp/class/google/protobuf/well-known-types/field-mask
			mask = nil
		}
	}
	return mask
}

{{ if needs . "hostname" }}{{ template "hostname" . }}{{ end }}

{{ if needs . "email" }}{{ template "email" . }}{{ end }}

{{ if needs . "uuid" }}{{ template "uuid" . }}{{ end }}

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
`
