package main

//go:generate mkdir -p api
//go:generate protoc --go_out=plugins=grpc:./api --go_opt=paths=source_relative --proto_path=proto proto/annotations.proto
//go:generate protoc --go_out=plugins=grpc:./api --go_opt=paths=source_relative --proto_path=proto proto/empty.proto
//go:generate protoc --go_out=plugins=grpc:./api --go_opt=paths=source_relative --proto_path=proto proto/fields.proto
//go:generate protoc --go_out=plugins=grpc:./api --go_opt=paths=source_relative --proto_path=proto proto/methods.proto
//go:generate protoc --go_out=plugins=grpc:./api --go_opt=paths=source_relative --proto_path=proto proto/sqs.proto
//go:generate go run -tags codegen internal/gen/proto/main.go -path=pkg/aws aws/apis/*/*/api-2.json
