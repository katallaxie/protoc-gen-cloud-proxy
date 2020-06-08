package templates

const messageTpl = `
// {{ name . }}JSONMarshaler describes the default jsonpb.Marshaler used by all 
// instances of {{ name . }}JSONMarshaler This struct is safe to replace or modify but 
// should not be done so concurrently.
var {{ name . }}JSONMarshaler = new(jsonpb.Marshaler)
// MarshalJSON satisfies the encoding/json Marshaler interface. This method 
// uses the more correct jsonpb package to correctly marshal the message.
func (m *{{ name . }}) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}
	buf := &bytes.Buffer{}
	if err := {{ name . }}JSONMarshaler.Marshal(buf, m); err != nil {
	  return nil, err
	}
	return buf.Bytes(), nil
}

var _ json.Marshaler = (*{{ name . }})(nil)
// {{ name . }}JSONUnmarshaler describes the default jsonpb.Unmarshaler used by all 
// instances of {{ name . }}. This struct is safe to replace or modify but 
// should not be done so concurrently.
var {{ name . }}JSONUnmarshaler = new(jsonpb.Unmarshaler)
// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method 
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *{{ name . }}) UnmarshalJSON(b []byte) error {
	return {{ name . }}JSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}
var _ json.Unmarshaler = (*{{ name . }})(nil)
`
