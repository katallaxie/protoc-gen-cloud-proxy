package templates

const lambdaTpl = `b, err := req.MarshalJSON()
if err != nil {
	return nil, err
}

session, err := s.getSession()
if err != nil {
	return nil, err
}

svc := lambda.New(session)
input := &lambda.InvokeInput{
	FunctionName: aws.String("{{ .Methods.GetLambda.FunctionName }}"),
  Payload:      b,
  {{ if .Methods.GetLambda.Qualifier }} Qualifier: aws.String("{{ .Methods.GetLambda.Qualifier }}"), {{ end }}
}

result, err := svc.InvokeWithContext(ctx, input)
if err != nil {
	return nil, err
}

var payload {{ name .Method }}_{{ .Method.Output.Name }}
if err := payload.UnmarshalJSON(result.Payload); err != nil {
	return nil, err
}

return &payload, nil
`
