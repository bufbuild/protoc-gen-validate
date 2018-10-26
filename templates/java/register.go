package java

import (
	"strings"
	"text/template"
	"unicode"

	"github.com/iancoleman/strcase"
	"github.com/lyft/protoc-gen-star"
	"github.com/lyft/protoc-gen-star/lang/go"
)

func Register(tpl *template.Template, params pgs.Parameters) {
	fns := javaFuncs{pgsgo.InitContext(params)}

	tpl.Funcs(map[string]interface{}{
		"className":   fns.className,
		"javaPackage": fns.javaPackage,
	})

	template.Must(tpl.Parse(fileTpl))
}

func JavaFilePath(f pgs.File, ctx pgsgo.Context, tpl *template.Template) pgs.FilePath {
	fullPath := ctx.OutputPath(f)

	fileName := strings.TrimSuffix(ctx.OutputPath(f).Base(), ".pb.go")
	fileName = toJavaClassName(fileName)
	fileName += "Validator.java"

	return fullPath.SetBase(fileName)
}

type javaFuncs struct{ pgsgo.Context }

func (fns javaFuncs) className(file pgs.File) pgs.Name {
	return pgs.Name(toJavaClassName(pgs.FilePath(file.Name().String()).BaseName()))
}

func (fns javaFuncs) javaPackage(file pgs.File) pgs.Name {
	return file.Package().ProtoName()
}

func toJavaClassName(protoName string) string {
	className := makeInvalidClassnameCharactersUnderscores(protoName)
	className = strcase.ToCamel(strcase.ToSnake(className))
	className = upperCaseAfterNumber(className)

	return className
}

// Replace invalid identifier characters with an underscore
func makeInvalidClassnameCharactersUnderscores(name string) string {
	var sb strings.Builder
	for _, c := range name {
		switch {
		case c >= '0' && c <= '9':
			sb.WriteRune(c)
		case c >= 'a' && c <= 'z':
			sb.WriteRune(c)
		case c >= 'A' && c <= 'Z':
			sb.WriteRune(c)
		default:
			sb.WriteRune('_')
		}
	}
	return sb.String()
}

func upperCaseAfterNumber(name string) string {
	var sb strings.Builder
	var p rune

	for _, c := range name {
		if unicode.IsDigit(p) {
			sb.WriteRune(unicode.ToUpper(c))
		} else {
			sb.WriteRune(c)
		}
		p = c
	}
	return sb.String()
}
