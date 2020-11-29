package opts

import (
	"go.uber.org/zap"
)

const (
	// DefaultVerbose ...
	DefaultVerbose = false
)

// Opts ...
type Opts struct {
	// Verbose toggles verbosity
	Verbose bool
	// Path to th sources
	Path string
	// Logger ...
	Logger *zap.Logger
}

// Opt ...
type Opt func(*Opts)

// New ...
func New(opts ...Opt) *Opts {
	o := NewDefaultOpts()
	o.Configure(opts...)

	return o
}

// DefaultOpts ...
func NewDefaultOpts() *Opts {
	return &Opts{
		Verbose: DefaultVerbose,
	}
}

// WithLogger ...
func WithLogger(logger *zap.Logger) Opt {
	return func(opts *Opts) {
		opts.Logger = logger
	}
}

// Configure ...
func (s *Opts) Configure(opts ...Opt) error {
	for _, o := range opts {
		o(s)
	}

	return nil
}
