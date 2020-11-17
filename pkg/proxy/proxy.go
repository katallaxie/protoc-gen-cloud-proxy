package proxy

import (
	"context"

	"github.com/andersnormal/pkg/server"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

type proxy struct {
	logger *zap.Logger
	opts   *Opts
	cmd    *cobra.Command
}

// Proxy ...
type Proxy interface {
	// Start ...
	Start(context.Context, Listener) error
}

// Listener
type Listener server.Listener

// Opts ...
type Opts struct{}

// Opt ...
type Opt func(*Opts)

// New ..
func New(opts ...Opt) Proxy {
	options := new(Opts)

	p := new(proxy)
	p.opts = options

	configure(p, opts...)

	return p
}

// WithContext ...
func (p *proxy) Start(ctx context.Context, l Listener) error {
	// create root context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// create server
	s, _ := server.WithContext(ctx)

	// debug listener
	debug := server.NewDebugListener(
		server.WithPprof(),
		server.WithStatusAddr(":8443"),
	)
	s.Listen(debug, true)

	// listen for grpc
	s.Listen(l, true)

	// listen for the server and wait for it to fail,
	// or for sys interrupts
	if err := s.Wait(); err != nil {
		return err
	}

	// noop
	return nil
}

// +private

func configure(p *proxy, opts ...Opt) error {
	for _, o := range opts {
		o(p.opts)
	}

	return nil
}
