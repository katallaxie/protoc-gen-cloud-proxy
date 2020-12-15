package cgen

import (
	"bytes"
	"os"
	"os/exec"

	"github.com/golang/protobuf/proto"
	"github.com/spf13/cobra"

	pb "github.com/katallaxie/protoc-gen-cloud-proxy/api"
)

const (
	pluginPrefix    = "cgen-"
	extensionPrefix = "cgen-x-"
)

type root struct {
	opts        *Opts
	pluginCalls []*pluginCall
}

func newRoot(opts ...Opt) *root {
	options := NewOpts(opts...)

	r := new(root)
	r.pluginCalls = make([]*pluginCall, 0)
	r.opts = options

	return r
}

func (r *root) preRunE(cmd *cobra.Command, args []string) error {
	for _, plugin := range r.opts.Plugins {
		p := &pluginCall{Name: pluginPrefix + plugin}
		r.pluginCalls = append(r.pluginCalls, p)
	}

	return nil
}

func (r *root) runE(cmd *cobra.Command, args []string) error {
	outputs := make([]*pb.Message, 0)

	for _, plugin := range r.pluginCalls {
		out, err := plugin.call(nil)
		if err != nil {
			return err
		}

		outputs = append(outputs, out...)
	}

	return nil
}

type pluginCall struct {
	Name       string
	Invocation string
}

// Invoking a plugin.
func (p *pluginCall) call(document proto.Message) ([]*pb.Message, error) {
	req := &pb.Request{}

	// save for later invocation parameters
	// invocation := p.Invocation

	version := &pb.Version{}
	version.Major = 0
	version.Minor = 1
	version.Patch = 0
	req.CompilerVersion = version

	msg, _ := proto.Marshal(req)

	cmd := exec.Command(p.Name, "-plugin")
	cmd.Stdin = bytes.NewReader(msg)
	cmd.Stderr = os.Stderr

	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	resp := &pb.Response{}
	if err := proto.Unmarshal(out, resp); err != nil {
		return nil, err
	}

	return nil, nil
}

func init() {}

func New(opts ...Opt) *cobra.Command {
	r := newRoot(opts...)

	cmd := &cobra.Command{
		Use:     "cgen",
		PreRunE: r.preRunE,
		RunE:    r.runE,
	}

	cmd.Flags().StringSliceVar(&r.opts.Plugins, "plugin", r.opts.Plugins, "plugin name")

	return cmd
}
