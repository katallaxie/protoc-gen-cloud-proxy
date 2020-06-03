package gen

import (
	g "github.com/aws-labs/aws-grpc-service-proxy/pkg/gen"

	google_protobuf "github.com/golang/protobuf/protoc-gen-go/descriptor"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
)

type generator struct {
	*g.Generator
}

type Generator interface {
	Generate()
}

func New() Generator {
	return &generator{Generator: g.New()}
}

func (g *generator) Make(protoFile *google_protobuf.FileDescriptorProto) (*plugin.CodeGeneratorResponse_File, error) {
	return nil, nil
}

func (g *generator) Generate() {
	g.Generator.Generate(g)
}
