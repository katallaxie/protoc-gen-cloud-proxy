package templates

const methodTpl = `// Here goes a message {{ name . }}
func (s *service) {{ name . }}(ctx context.Context, req *{{ .Input.Parent.Name }}_{{ .Input.Name }}) (*{{ name . }}_{{ .Output.Name }}, error) {
  {{ render (context .) }}
}
`
