package gen

import (
	"io"
	"os"

	google_protobuf "github.com/golang/protobuf/protoc-gen-go/descriptor"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
)

type Generator struct {
	reader io.Reader
	writer io.Writer
}

type fileMaker interface {
	Make(*google_protobuf.FileDescriptorProto) (*plugin.CodeGeneratorResponse_File, error)
}

// New ...
func New() *Generator {
	return &Generator{
		reader: os.Stdin,
		writer: os.Stdout,
	}
}

// Generate ...
func (g *Generator) Generate(make fileMaker) {

}
