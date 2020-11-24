package templates

const sqsServerStreamingTpl = `
svc := sqs.New(s.session)
input := &sqs.ReceiveMessageInput{
  QueueUrl:              aws.String(""),
  MessageAttributeNames: aws.StringSlice(req.MessageAttributeNames),
}

output, err := svc.ReceiveMessageWithContext(stream.Context(), input)
if err != nil {
  return err
}

for _, msg := range output.Messages {
  var payload Song
  if err := proto.Unmarshal([]byte(*msg.Body), &payload); err != nil {
    return err
  }

  stream.Send(&payload)
}

return nil
`
