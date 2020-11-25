package templates

const sqsClientStreamingTpl = `svc := sqs.New(s.session)
for {
  msg, err := stream.Recv()
  if err == io.EOF {
    return stream.SendAndClose(&api.Empty{})
  }

  if err != nil {
    return err
  }

  input := &sqs.SendMessageInput{
    QueueUrl:               aws.String(msg.GetQueueUrl()),
    DelaySeconds:           aws.Int64(msg.GetDelaySeconds()),
    MessageBody:            aws.String(msg.MessageBody),
  }

  _, err = svc.SendMessage(input)
  if err != nil {
    return err
  }
}`
