package templates

const methodServerStreamingTpl = `// Here goes a message {{ name . }}
func (s *service) {{ name . }}(req *{{ name . }}_{{ .Input.Name }}, stream {{ .Service.Descriptor.Name }}_{{ name . }}Server) error {
	{{ render (context .) }}
}
`
