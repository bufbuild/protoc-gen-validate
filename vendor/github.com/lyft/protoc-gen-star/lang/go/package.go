package pgsgo

import (
	"go/token"
	"regexp"
	"strings"

	"github.com/lyft/protoc-gen-star"
)

var nonAlphaNumPattern = regexp.MustCompile("[^a-zA-Z0-9]")

func (c context) PackageName(node pgs.Node) pgs.Name {
	e, ok := node.(pgs.Entity)
	if !ok {
		e = node.(pgs.Package).Files()[0]
	}

	_, pkg := c.optionPackage(e)

	if ip := c.p.Str("import_path"); ip != "" {
		pkg = ip
	}

	// if the package name is a Go keyword, prefix with '_'
	if token.Lookup(pkg).IsKeyword() {
		pkg = "_" + pkg
	}

	// package name is kosher
	return pgs.Name(pkg)
}

func (c context) ImportPath(e pgs.Entity) pgs.FilePath {
	path, _ := c.optionPackage(e)
	path = c.p.Str("import_prefix") + path
	return pgs.FilePath(path)
}

func (c context) OutputPath(e pgs.Entity) pgs.FilePath {
	out := e.File().InputPath().SetExt(".pb.go")

	// source relative doesn't try to be fancy
	if Paths(c.p) == SourceRelative {
		return out
	}

	path, _ := c.optionPackage(e)

	// Import relative ignores the existing file structure
	return pgs.FilePath(path).Push(out.Base())
}

func (c context) optionPackage(e pgs.Entity) (path, pkg string) {
	// M mapping param overrides everything IFF the entity is not a build target
	if override, ok := c.p["M"+e.File().InputPath().String()]; ok && !e.BuildTarget() {
		path = override
		pkg = override
		if idx := strings.LastIndex(pkg, "/"); idx > -1 {
			pkg = pkg[idx+1:]
		}
		return
	}

	// check if there's a go_package option specified
	pkg = e.File().Descriptor().GetOptions().GetGoPackage()
	path = e.File().InputPath().Dir().String()

	if pkg == "" {
		// have a proto package name, so use that
		if n := e.Package().ProtoName(); n != "" {
			pkg = n.SnakeCase().String()
		} else { // no other info, then replace all non-alphanumerics from the input file name
			pkg = nonAlphaNumPattern.ReplaceAllString(e.File().InputPath().BaseName(), "_")
		}
		return
	}

	// go_package="example.com/foo/bar;baz" should have a package name of `baz`
	if idx := strings.LastIndex(pkg, ";"); idx > -1 {
		path = pkg[:idx]
		pkg = pkg[idx+1:]
		return
	}

	// go_package="example.com/foo/bar" should have a package name of `bar`
	if idx := strings.LastIndex(pkg, "/"); idx > -1 {
		path = pkg
		pkg = pkg[idx+1:]
		return
	}

	pkg = nonAlphaNumPattern.ReplaceAllString(pkg, "_")

	return
}
