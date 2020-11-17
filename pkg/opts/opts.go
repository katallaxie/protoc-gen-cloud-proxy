package opts

// Opts ...
type Opts struct {
	// Verbose ...
	Verbose bool
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
func Configure(opt *Opts, opts ...Opt) error {
	for _, o := range opts {
		o(opt)
	}

	return nil
}

// ConfigureLogger ...
func ConfigureLogger() error {
	return nil
}

// WithVerbose ...
func WithVerbose() Opt {
	return func(opts *Opts) {
		opts.Verbose = true
	}
}
