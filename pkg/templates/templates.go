package templates

import (
	"text/template"

	"github.com/awslabs/protoc-gen-aws-service-proxy/pkg/templates/shared"
	pgs "github.com/lyft/protoc-gen-star"
)

type RegisterFn func(tpl *template.Template, params pgs.Parameters)

func makeTemplate(ext string, fn RegisterFn, params pgs.Parameters) *template.Template {
	tpl := template.New(ext)
	shared.Register(tpl, params)

	template.Must(tpl.New("lambda").Parse(lambdaTpl))
	template.Must(tpl.New("dynamodb").Parse(dynamoDBTpl))

	fn(tpl, params)
	return tpl
}

func File(params pgs.Parameters) (*template.Template, error) {
	tpl := makeTemplate("file", func(tpl *template.Template, params pgs.Parameters) {}, params)
	return tpl.Parse(fileTpl)
}

func Service(params pgs.Parameters) (*template.Template, error) {
	tpl := makeTemplate("service", func(tpl *template.Template, params pgs.Parameters) {}, params)
	return tpl.Parse(serviceTpl)
}

func Method(params pgs.Parameters) (*template.Template, error) {
	tpl := makeTemplate("method", func(tpl *template.Template, params pgs.Parameters) {}, params)
	return tpl.Parse(methodTpl)
}

func Message(params pgs.Parameters) (*template.Template, error) {
	tpl := makeTemplate("message", func(tpl *template.Template, params pgs.Parameters) {}, params)
	return tpl.Parse(messageTpl)
}
