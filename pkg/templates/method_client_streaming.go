package templates

const methodClientStreamingTpl = `// Here goes a message {{ name . }}
func (s *service) {{ name . }}(stream {{ .Service.Descriptor.Name }}_{{ name . }}Server) error {
  {{ render (context .) }}
}
`
