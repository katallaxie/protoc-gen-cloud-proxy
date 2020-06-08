package templates

const methodTpl = `// Here goes a message {{ name . }}

func (s *service) {{ name . }}(ctx context.Context, req *{{ name . }}_{{ .Input.Name }}) (*{{ name . }}_{{ .Output.Name }}, error) {
	{{ render (context .) }}
}
`
