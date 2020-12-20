package iface

import "context"

// Manifest The plugin manifest.
type Manifest struct {
	Name        string `yaml:"name"`
	Type        string `yaml:"type"`
	Import      string `yaml:"import"`
	Description string `yaml:"summary"`
}

// Constructor ...
type Constructor func() (ResolverHandler, error)

// ResolverHandler ...
type ResolverHandler interface {
	Resolve(context.Context, []byte) ([]byte, error)
}
