package proxy

import (
	"context"

	"github.com/andersnormal/pkg/server"
	"github.com/spf13/cobra"

	o "github.com/katallaxie/protoc-gen-cloud-proxy/pkg/opts"
)

type proxy struct {
	opts     *o.Opts
	cmd      *cobra.Command
	listener Listener
}

// Proxy ...
type Proxy interface {
	// Start ...
	Start(context.Context) error
}

// Listener
type Listener server.Listener

// New ..
func New(l Listener, opts *o.Opts) Proxy {
	p := new(proxy)
	p.opts = opts
	p.listener = l

	return p
}

// WithContext ...
func (p *proxy) Start(ctx context.Context) error {
	// create root context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// create server
	s, _ := server.WithContext(ctx)

	if p.opts.Debug {
		// debug listener
		debug := server.NewDebugListener(
			server.WithPprof(),
			server.WithStatusAddr(p.opts.StatusAddr),
		)
		s.Listen(debug, true)
	}

	// listen for grpc
	s.Listen(p.listener, true)

	// listen for the server and wait for it to fail,
	// or for sys interrupts
	if err := s.Wait(); err != nil {
		return err
	}

	// noop
	return nil
}

// +private

func configure(p *proxy, opts ...o.Opt) error {
	for _, o := range opts {
		o(p.opts)
	}

	return nil
}
