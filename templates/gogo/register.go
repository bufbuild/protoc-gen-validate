package gogo

import (
	"text/template"

	"github.com/lyft/protoc-gen-star"

	shared "github.com/lyft/protoc-gen-validate/templates/goshared"
)

func Register(tpl *template.Template, params pgs.Parameters) {
	shared.Register(tpl, params)
	template.Must(tpl.Parse(fileTpl))
	template.Must(tpl.New("required").Parse(requiredTpl))
	template.Must(tpl.New("timestamp").Parse(timestampTpl))
	template.Must(tpl.New("duration").Parse(durationTpl))
}
