package main

import (
	"context"

	pb "github.com/katallaxie/protoc-gen-cloud-proxy/example"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/google/uuid"
)

func HandleRequest(ctx context.Context, event *pb.Insert_Request) (*pb.Insert_Response, error) {
	res := &pb.Insert_Response{Uuid: uuid.New().String()}

	return res, nil
}

func main() {
	lambda.Start(HandleRequest)
}
