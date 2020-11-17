package opts

import (
	"go.uber.org/zap"
)

// Opts ...
type Opts struct {
	// Verbose ...
	Verbose bool
	// Logger ...
	Logger *zap.Logger
}

// Opt ...
type Opt func(*Opts)

// New ...
func New() *Opts {
	o := new(Opts)
	o.Verbose = false

	return o
}

// Configure ...
func (s *Opts) Configure(opts ...Opt) error {
	for _, o := range opts {
		o(s)
	}

	return nil
}

// ConfigureLogger ...
func ConfigureLogger() error {
	return nil
}

// WithLogger ...
func WithLogger(logger *zap.Logger) Opt {
	return func(opts *Opts) {
		opts.Logger = logger
	}
}

// WithVerbose ...
func WithVerbose() Opt {
	return func(opts *Opts) {
		opts.Verbose = true
	}
}
