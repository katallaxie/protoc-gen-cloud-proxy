package cgen

import (
	"go.uber.org/zap"
)

const (
	// DefaultLogFormat ...
	DefaultLogFormat = "text"
	// DefaultLogLevel ...
	DefaultLogLevel = "warn"
	// DefaultVerbose ...
	DefaultVerbose = false
	// DefaultDebug ...
	DefaultDebug = false
)

// Opts ...
type Opts struct {
	// Verbose ...
	Verbose bool
	// Debug ...
	Debug bool
	// LogFormat ...
	LogFormat string
	// LogLevel ...
	LogLevel string
	// Logger ...
	Logger *zap.Logger
	// Plugins ...
	Plugins []string
}

// Opt ...
type Opt func(*Opts)

// DefaultOpts ...
func NewDefaultOpts() *Opts {
	return &Opts{
		LogFormat: DefaultLogFormat,
		LogLevel:  DefaultLogLevel,
		Debug:     DefaultDebug,
		Verbose:   DefaultVerbose,
		Plugins:   make([]string, 0),
	}
}

// New ...
func NewOpts(opts ...Opt) *Opts {
	o := NewDefaultOpts()
	o.Configure(opts...)

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

// WithDebug ...
func WithDebug() Opt {
	return func(opts *Opts) {
		opts.Debug = true
	}
}
