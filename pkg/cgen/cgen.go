package cgen

import (
	"github.com/golang/protobuf/proto"
	"github.com/spf13/cobra"
)

const (
	pluginPrefix    = "cgen-"
	extensionPrefix = "cgen-x-"
)

type pluginCall struct {
	Name       string
	Invocation string
}

// Invoking a plugin.
func (p *pluginCall) call(document proto.Message) {}

var root = &cobra.Command{
	Use: "cgen",
}

func init() {}

func New() *cobra.Command {
	return root
}
