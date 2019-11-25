package goshared

const msgTpl = `
{{ if disabled . -}}
	{{ cmt "Validate is disabled for " (msgTyp .) ". This method will always return nil." }}
{{- else -}}
	{{ cmt "Validate checks the field values on " (msgTyp .) " with the rules defined in the proto definition for this message. If any rules are violated, an error is returned." }}
{{- end -}}

{{ $msg := . }}
func (m {{ (msgTyp .).Pointer }}) Validate() error {
	{{ if disabled . -}}
		return nil
	{{ else -}}
		if m == nil { return nil }

		var validationFunctions = []func() error {
		{{ range .NonOneOfFields }}m.Validate{{ (msgTyp $msg).String }}{{ (name .) }},
		{{ end }}{{ range .OneOfs }}m.Validate{{ (msgTyp $msg).String }}{{ (name .) }},
		{{ end }}
		}

		for _, validationFunction := range validationFunctions {
			if err := validationFunction(); err != nil {
				return err
			}
		}

		return nil
	{{ end -}}
}

func (m {{ (msgTyp .).Pointer }}) ValidateAll() error {
	{{ if disabled . -}}
		return nil
	{{ else -}}
		if m == nil { return nil }

		var validationFunctions = []func() error {
		{{ range .NonOneOfFields }}m.Validate{{ (msgTyp $msg).String }}{{ (name .) }}{{ if or ((.Type).IsRepeated) ((.Type).IsMap) ((.Type).IsEmbed) -}}All{{ end }},
		{{ end }}{{ range .OneOfs }}m.Validate{{ (msgTyp $msg).String }}{{ (name .) }}All,
		{{ end }}
		}

		var wg sync.WaitGroup
		wg.Add(len(validationFunctions))
		var errorsChan = make(chan error, len(validationFunctions))
		for _, validateFunction := range validationFunctions {
			go func(f func() error) {
				defer wg.Done()
				if err := f(); err != nil {
					errorsChan <- err
				}
			}(validateFunction)
		}
		wg.Wait()

		var result {{ errname . }}s
	loop:
		for {
			var err error
			select {
			case err = <- errorsChan:
				err, ok := err.({{ errname . }})
				if !ok {
					return err
				} else {
					result = append(result, err)
				}
			default:
				break loop
			}
		}
	
		if len(result) > 0 {
			return result
		} else {
			return nil
		}

	{{ end -}}
}

{{ range .NonOneOfFields }}
	func (m {{ (msgTyp $msg).Pointer }}) Validate{{ (msgTyp $msg).String }}{{ (name .) }}() error {
		{{ render (context .) }}
		return nil
	}
	{{ if or ((.Type).IsRepeated) ((.Type).IsMap) ((.Type).IsEmbed) -}}
	func (m {{ (msgTyp $msg).Pointer }}) Validate{{ (msgTyp $msg).String }}{{ (name .) }}All() error {
		{{ render ((context .).WithAllErrors) }}
		return nil
	}
	{{ end -}}
{{ end }}

{{ range .OneOfs }}

	{{ $oneOfField := . }}

	func (m {{ (msgTyp $msg).Pointer }}) Validate{{ (msgTyp $msg).String }}{{ (name .) }}() error {
		switch m.{{ name . }}.(type) {
			{{ range .Fields }}
				case {{ oneof . }}:
					if err := m.Validate{{ (msgTyp $msg).String }}{{ (name $oneOfField) }}{{ (name .) }}(); err != nil {
						return err
					}
			{{ end }}
			{{ if required . }}
				default:
					return {{ errname .Message }}{
						field: "{{ name . }}",
						reason: "value is required",
					}
			{{ end }}
		}
		return nil
	}

	func (m {{ (msgTyp $msg).Pointer }}) Validate{{ (msgTyp $msg).String }}{{ (name .) }}All() error {
		switch m.{{ name . }}.(type) {
			{{ range .Fields }}
				case {{ oneof . }}:
					if err := m.Validate{{ (msgTyp $msg).String }}{{ (name $oneOfField) }}{{ (name .) }}{{ if or ((.Type).IsRepeated) ((.Type).IsMap) ((.Type).IsEmbed) -}}All{{ end }}(); err != nil {
						return err
					}
			{{ end }}
			{{ if required . }}
				default:
					return {{ errname .Message }}{
						field: "{{ name . }}",
						reason: "value is required",
					}
			{{ end }}
		}
		return nil
	}

	{{ range .Fields }}
		func (m {{ (msgTyp $msg).Pointer }}) Validate{{ (msgTyp $msg).String }}{{ (name $oneOfField) }}{{ (name .) }}() error {
			{{ render (context .) }}
			return nil
		}
		{{ if or ((.Type).IsRepeated) ((.Type).IsMap) ((.Type).IsEmbed) -}}
			func (m {{ (msgTyp $msg).Pointer }}) Validate{{ (msgTyp $msg).String }}{{ (name $oneOfField) }}{{ (name .) }}All() error {
				{{ render ((context .).WithAllErrors) }}
				return nil
			}
		{{ end -}}
	{{ end }}
{{ end }}

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

type {{ errname . }}s []{{ errname . }}

func (e {{ errname . }}s) Error() string {
	var fields []string
	for _, err := range e {
		fields = append(fields, strings.ToLower(err.Field()))
	}

	return fmt.Sprintf("invalid fields: %s", strings.Join(fields, ", "))
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
	{{ end }}{{ end }}

	{{ if has .Rules "Items"}}{{ if .Rules.Items }}
	{{ if has .Rules.Items.GetString_ "NotIn" }} {{ if .Rules.Items.GetString_.NotIn }}
		var {{ lookup .Field "NotInLookup" }} = map[string]struct{}{
			{{- range .Rules.Items.GetString_.NotIn }}
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
