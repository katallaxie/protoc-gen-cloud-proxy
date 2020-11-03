package templates

const lambdaTpl = `// Encode the Message into JSON representation
b, err := req.MarshalJSON()
if err != nil {
	return nil, err
}

svc := lambda.New(session.New())
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

return nil, nil
`
