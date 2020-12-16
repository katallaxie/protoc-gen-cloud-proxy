package plugins

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
	// Invocation ...
	Invocation string
	// Input ...
	Input string
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
