package templates

const sqsTpl = `
{{ if .Method.ServerStreaming }}
svc := sqs.New(s.session)
input := &sqs.ReceiveMessageInput{
  QueueUrl: aws.String(""),
}

output, err := svc.ReceiveMessageWithContext(stream.Context(), input)
if err != nil {
  return err
}

for _, msg := range output.Messages {
  var payload {{ .Method.Output.Name }}
  if err := payload.UnmarshalJSON([]byte(*msg.Body)); err != nil {
    return err
  }

  stream.Send(&payload)
}

return nil
{{ end }}
`
