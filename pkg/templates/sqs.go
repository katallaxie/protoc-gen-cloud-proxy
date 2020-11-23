package templates

const sqsTpl = `
{{ if and (eq .Method.Input.Name "Empty") (.Method.ServerStreaming) }}
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
  if err := proto.Unmarshal([]byte(*msg.Body), &payload); err != nil {
    return err
  }

  stream.Send(&payload)
}

return nil
{{ end }}
`
