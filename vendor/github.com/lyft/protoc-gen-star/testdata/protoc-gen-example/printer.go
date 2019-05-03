package main

import (
	"fmt"
	"io"
	"strings"

	"bytes"

	pgs "github.com/lyft/protoc-gen-star"
)

type PrinterModule struct {
	*pgs.ModuleBase
}

func ASTPrinter() *PrinterModule { return &PrinterModule{ModuleBase: &pgs.ModuleBase{}} }

func (p *PrinterModule) Name() string { return "printer" }

func (p *PrinterModule) Execute(targets map[string]pgs.File, packages map[string]pgs.Package) []pgs.Artifact {
	buf := &bytes.Buffer{}

	for _, f := range targets {
		p.printFile(f, buf)
	}

	return p.Artifacts()
}

func (p *PrinterModule) printFile(f pgs.File, buf *bytes.Buffer) {
	p.Push(f.Name().String())
	defer p.Pop()

	buf.Reset()
	v := initPrintVisitor(buf, "")
	p.CheckErr(pgs.Walk(v, f), "unable to print AST tree")

	out := buf.String()

	if ok, _ := p.Parameters().Bool("log_tree"); ok {
		p.Logf("Proto Tree:\n%s", out)
	}

	p.AddGeneratorFile(
		f.InputPath().SetExt(".tree.txt").String(),
		out,
	)
}

const (
	startNodePrefix = "┳ "
	subNodePrefix   = "┃"
	leafNodePrefix  = "┣"
	leafNodeSpacer  = "━ "
)

type PrinterVisitor struct {
	pgs.Visitor
	prefix string
	w      io.Writer
}

func initPrintVisitor(w io.Writer, prefix string) pgs.Visitor {
	v := PrinterVisitor{
		prefix: prefix,
		w:      w,
	}
	v.Visitor = pgs.PassThroughVisitor(&v)
	return v
}

func (v PrinterVisitor) leafPrefix() string {
	if strings.HasSuffix(v.prefix, subNodePrefix) {
		return strings.TrimSuffix(v.prefix, subNodePrefix) + leafNodePrefix
	}
	return v.prefix
}

func (v PrinterVisitor) writeSubNode(str string) pgs.Visitor {
	fmt.Fprintf(v.w, "%s%s%s\n", v.leafPrefix(), startNodePrefix, str)
	return initPrintVisitor(v.w, fmt.Sprintf("%s%v", v.prefix, subNodePrefix))
}

func (v PrinterVisitor) writeLeaf(str string) {
	fmt.Fprintf(v.w, "%s%s%s\n", v.leafPrefix(), leafNodeSpacer, str)
}

func (v PrinterVisitor) VisitFile(f pgs.File) (pgs.Visitor, error) {
	return v.writeSubNode("File: " + f.Name().String()), nil
}

func (v PrinterVisitor) VisitMessage(m pgs.Message) (pgs.Visitor, error) {
	return v.writeSubNode("Message: " + m.Name().String()), nil
}

func (v PrinterVisitor) VisitEnum(e pgs.Enum) (pgs.Visitor, error) {
	return v.writeSubNode("Enum: " + e.Name().String()), nil
}

func (v PrinterVisitor) VisitService(s pgs.Service) (pgs.Visitor, error) {
	return v.writeSubNode("Service: " + s.Name().String()), nil
}

func (v PrinterVisitor) VisitEnumValue(ev pgs.EnumValue) (pgs.Visitor, error) {
	v.writeLeaf(ev.Name().String())
	return nil, nil
}

func (v PrinterVisitor) VisitField(f pgs.Field) (pgs.Visitor, error) {
	v.writeLeaf(f.Name().String())
	return nil, nil
}

func (v PrinterVisitor) VisitMethod(m pgs.Method) (pgs.Visitor, error) {
	v.writeLeaf(m.Name().String())
	return nil, nil
}
