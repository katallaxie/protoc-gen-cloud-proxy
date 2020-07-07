package module

import (
	"bytes"
	"io"

	"github.com/awslabs/protoc-gen-aws-service-proxy/pkg/templates"

	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
)

type Module struct {
	*pgs.ModuleBase
	ctx pgsgo.Context
}

func Proxy() pgs.Module { return &Module{ModuleBase: &pgs.ModuleBase{}} }

func (m *Module) InitContext(ctx pgs.BuildContext) {
	m.ModuleBase.InitContext(ctx)
	m.ctx = pgsgo.InitContext(ctx.Parameters())
}

func (m *Module) Name() string { return "proxy" }

func (m *Module) Execute(targets map[string]pgs.File, pkgs map[string]pgs.Package) []pgs.Artifact {
	buf := &bytes.Buffer{}

	for _, f := range targets {
		m.genProxy(f, buf, m.Parameters())
	}

	return m.Artifacts()
}

func (m *Module) genProxy(f pgs.File, buf *bytes.Buffer, params pgs.Parameters) {
	m.Push(f.Name().String())
	defer m.Pop()

	buf.Reset()
	v := initProxyVisitor(m, buf, "", params)
	m.CheckErr(pgs.Walk(v, f), "unable to generate AST tree")

	out := buf.String()

	m.AddGeneratorFile(
		m.ctx.OutputPath(f).SetExt(".proxy.go").String(),
		out,
	)
}

type ProxyVisitor struct {
	pgs.Visitor
	pgs.DebuggerCommon
	pgs.Parameters
	prefix string
	w      io.Writer
}

func initProxyVisitor(d pgs.DebuggerCommon, w io.Writer, prefix string, params pgs.Parameters) pgs.Visitor {
	p := ProxyVisitor{
		prefix:         prefix,
		w:              w,
		Parameters:     params,
		DebuggerCommon: d,
	}

	p.Visitor = pgs.PassThroughVisitor(&p)

	return p
}

// VisitFile ...
func (p ProxyVisitor) VisitFile(f pgs.File) (pgs.Visitor, error) {
	tpl, err := templates.File(p.Parameters)
	if err != nil {
		return nil, err
	}

	err = tpl.Execute(p.w, f)
	if err != nil {
		return nil, err
	}

	return p, err
}

// VisitService ...
func (p ProxyVisitor) VisitService(s pgs.Service) (pgs.Visitor, error) {
	tpl, err := templates.Service(p.Parameters)
	if err != nil {
		return nil, err
	}

	err = tpl.Execute(p.w, s)
	if err != nil {
		return nil, err
	}

	return p, err
}

// VisitMethod ...
func (p ProxyVisitor) VisitMethod(m pgs.Method) (pgs.Visitor, error) {
	tpl, err := templates.Method(p.Parameters)
	if err != nil {
		return nil, err
	}

	err = tpl.Execute(p.w, m)
	if err != nil {
		return nil, err
	}

	return p, err
}

// VisitMessage ...
func (p ProxyVisitor) VisitMessage(m pgs.Message) (pgs.Visitor, error) {
	tpl, err := templates.Message(p.Parameters)
	if err != nil {
		return nil, err
	}

	err = tpl.Execute(p.w, m)
	if err != nil {
		return nil, err
	}

	return p, err
}

var _ pgs.Module = (*Module)(nil)
