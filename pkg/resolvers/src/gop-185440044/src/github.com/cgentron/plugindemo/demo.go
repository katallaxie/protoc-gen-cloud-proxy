package plugindemo

import (
	"context"

	"github.com/cgentron/api/iface"
)

// Config ...
type Config struct{}

// CreateConfig ...
func CreateConfig() *Config {
	return &Config{}
}

// Demo ...
type Demo struct {
	name string
}

// New ...
func New(config *Config, name string) (iface.ResolverHandler, error) {
	return &Demo{
		name: name,
	}, nil
}

func (a *Demo) Resolve(context context.Context, msg []byte) ([]byte, error) {
	return []byte("demo"), nil
}
