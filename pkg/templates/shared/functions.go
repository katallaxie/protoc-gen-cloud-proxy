package shared

import (
	"fmt"
	"text/template"

	api "github.com/awslabs/protoc-gen-aws-service-proxy/api"

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

	ctx.Typ = resolveMethods(&services)
	ctx.Methods = &services

	if ctx.Typ == "error" {
		return ctx, fmt.Errorf("unknown template type")
	}

	return ctx, nil
}

func resolveMethods(m *api.Methods) string {
	if m.GetLambda() != nil {
		return "lambda"
	}

	if m.GetDynamodb() != nil {
		return "dynamodb"
	}

	return "error"
}
