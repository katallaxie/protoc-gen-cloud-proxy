package plugins

import (
	"io/ioutil"
	"os"

	"github.com/golang/protobuf/proto"
	pb "github.com/katallaxie/protoc-gen-cloud-proxy/api"
	"github.com/spf13/cobra"
)

type RuntimeHandler func() (string, error)

type Runtime struct {
	Request  *pb.Request
	Response *pb.Response
	Opts     *Opts

	cmd *cobra.Command
}

func New(opts ...Opt) *Runtime {
	options := NewOpts(opts...)

	r := new(Runtime)
	r.Opts = options
	r.Response = &pb.Response{}

	cmd := &cobra.Command{}
	cmd.Flags().StringVar(&r.Opts.Input, "input", r.Opts.Input, "input")
	cmd.Flags().BoolVar(&r.Opts.Verbose, "verbose", r.Opts.Verbose, "verbose")

	r.cmd = cmd

	return r
}

func (r *Runtime) Execute(handler RuntimeHandler) error {
	cmd := r.cmd
	cmd.RunE = r.runE(handler)

	return cmd.Execute()
}

func (r *Runtime) runE(handler RuntimeHandler) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		data, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			return r.handleError(err)
		}

		r.Request = &pb.Request{}
		if err := proto.Unmarshal(data, r.Request); err != nil {
			return r.handleError(err)
		}

		return nil
	}
}

func (r *Runtime) handleError(err error) error {
	r.Response.Errors = append(r.Response.Errors, err.Error())
	resp, err := proto.Marshal(r.Response)
	if err != nil {
		return err
	}

	os.Stdout.Write(resp)

	return err
}
