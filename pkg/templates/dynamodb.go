package templates

const dynamoDBTpl = `
svc := dynamodb.New(session.New())

{{ if .Methods.GetDynamodb.GetUpdateItem }}
input := &dynamodb.UpdateItemInput{
  ReturnValues:     aws.String("{{ .Methods.GetDynamodb.GetUpdateItem.ReturnValues }}"),
	TableName:        aws.String("{{ .Methods.GetDynamodb.GetUpdateItem.TableName }}"),
  UpdateExpression: aws.String("{{ .Methods.GetDynamodb.GetUpdateItem.UpdateExpression }}"),
}

_, err := svc.UpdateItem(input)
if err != nil {
	return nil, err
}
{{ end }}

// DynamoDB
return nil, nil
`
