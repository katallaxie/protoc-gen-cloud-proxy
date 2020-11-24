package templates

const sqsClientStreamingTpl = `
svc := sqs.New(s.session)

for {
  msg, err := stream.Recv()
  if err != nil {
    return err
  }

  bb, err := msg.MarshalJSON()
  if err != nil {
    return err
  }

  input := &sqs.SendMessageInput{
    QueueUrl:    aws.String(""),
    MessageBody: aws.String(string(bb)),
  }

  _, err := svc.SendMessage(input)
  if err != nil {
    return err
  }

  output := &api.Sqs_Output{}

  if err := stream.SendMsg(output); err != nil {
    return err
  }
}
`
