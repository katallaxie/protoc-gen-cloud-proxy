package templates

const lambdaTpl = `
b, err := req.MarshalJSON()
if err != nil {
	return nil, err
}

svc := lambda.New(session.New())
input := &lambda.InvokeInput{
	FunctionName: aws.String("dartTest"),
	Payload:      b,
	Qualifier:    aws.String("$LATEST"),
}

result, err := svc.Invoke(input)
if err != nil {
	return nil, err
}

var payload {{ name .Method }}_{{ .Method.Output.Name }}
if err := payload.UnmarshalJSON(result.Payload); err != nil {
	return nil, err
}

return nil, nil
`
