package shared

import (
	"fmt"
	"text/template"

	api "github.com/katallaxie/protoc-gen-cloud-proxy/api"

	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
)

type sharedFuncs struct {
	pgsgo.Context
}

func Register(tpl *template.Template, params pgs.Parameters) {
	fns := sharedFuncs{pgsgo.InitContext(params)}

	tpl.Funcs(map[string]interface {
	}{
		"pkg":     fns.PackageName,
		"name":    fns.Name,
		"context": proxyContext,
		"render":  Render(tpl),
	})
}

func proxyContext(m pgs.Method) (ProxyContext, error) {
	ctx := ProxyContext{}
	ctx.Method = m

	var services api.Methods
	if _, err := m.Extension(api.E_Methods, &services); err != nil {
		return ctx, err
	}

	ctx.Typ = resolveMethods(&services, m)
	ctx.Methods = &services

	if ctx.Typ == "error" {
		return ctx, fmt.Errorf("unknown template type")
	}

	return ctx, nil
}

func resolveMethods(a *api.Methods, m pgs.Method) string {
	if a.GetLambda() != nil {
		return "lambda"
	}

	if a.GetDynamodb() != nil {
		return "dynamodb"
	}

	if a.GetSqs() != nil && m.ServerStreaming() {
		return "sqs_server_streaming"
	}

	if a.GetSqs() != nil && m.ClientStreaming() {
		return "sqs_client_streaming"
	}

	return "error"
}
