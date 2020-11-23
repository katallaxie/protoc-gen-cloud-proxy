package templates

const methodServerStreamingTpl = `// Here goes a message {{ name . }}
func (s *service) {{ name . }}(req *Empty, stream {{ .Service.Descriptor.Name }}_{{ name . }}Server) error {
	{{ render (context .) }}
}
`
