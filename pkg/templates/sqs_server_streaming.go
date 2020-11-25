package templates

const sqsServerStreamingTpl = `
queueUrl := req.GetQueueUrl()

svc := sqs.New(s.session)
ll := s.logger.With(zap.String("queue_url", queueUrl))

input := &sqs.ReceiveMessageInput{
  QueueUrl:              aws.String(queueUrl),
  MessageAttributeNames: aws.StringSlice(req.MessageAttributeNames),
  MaxNumberOfMessages:   aws.Int64(req.MaxNumberOfMessages),
  WaitTimeSeconds:       aws.Int64(req.WaitTimeSeconds),
}

output, err := svc.ReceiveMessageWithContext(stream.Context(), input)
if err != nil {
  return err
}

ll.Info("receiving messages", zap.Int("messages", len(output.Messages)))

receiptHandles := make([]*sqs.DeleteMessageBatchRequestEntry, 0, len(output.Messages))

for _, msg := range output.Messages {
  ll.Info("streaming message", zap.String("message_id", aws.StringValue(msg.MessageId)))

  m := &api.SQS_Message{
    Body:      aws.StringValue(msg.Body),
    MessageId: aws.StringValue(msg.MessageId),
  }

  if err := stream.Send(m); err != nil {
    ll.Error("sending message", zap.Error(err))

    return err
  }

  receiptHandles = append(receiptHandles, &sqs.DeleteMessageBatchRequestEntry{Id: msg.MessageId, ReceiptHandle: msg.ReceiptHandle})
}

if len(receiptHandles) == 0 {
  return nil
}

deleteMessageInput := sqs.DeleteMessageBatchInput{
  QueueUrl: aws.String(queueUrl),
  Entries:  receiptHandles,
}

ll.Info("deleting messages", zap.Int("messages", len(receiptHandles)))

_, err = svc.DeleteMessageBatch(&deleteMessageInput)
if err != nil {
  ll.Error("deleting messages", zap.Error(err))

  return err
}

return nil`
