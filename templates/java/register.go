package java

import (
	"text/template"

	"github.com/lyft/protoc-gen-star"
)

func Register(tpl *template.Template, params pgs.Parameters) {
	template.Must(tpl.Parse(fileTpl))
}
