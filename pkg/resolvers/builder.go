package resolvers

import (
	"context"
	"fmt"
	"path"
	"reflect"
	"strings"

	"github.com/cgentron/api/iface"

	"github.com/mitchellh/mapstructure"
	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
)

// Symbols variable stores the map of stdlib symbols per package.
var Symbols = map[string]map[string]reflect.Value{}

func init() {
	Symbols["github.com/cgentron/api/iface"] = map[string]reflect.Value{
		"ResolverHandler":  reflect.ValueOf((*iface.ResolverHandler)(nil)),
		"_ResolverHandler": reflect.ValueOf((*_iface_resolver_Handler)(nil)),
	}
}

// _iface_resolver_Handler is an interface wrapper for Handler type
type _iface_resolver_Handler struct {
	WResolve func(a0 context.Context, a1 []byte) ([]byte, error)
}

func (W _iface_resolver_Handler) Resolve(a0 context.Context, a1 []byte) ([]byte, error) {
	return W.WResolve(a0, a1)
}

// Descriptor The static part of a plugin configuration (prod).
type Descriptor struct {
	// ModuleName (required)
	ModuleName string `description:"plugin's module name." json:"moduleName,omitempty" toml:"moduleName,omitempty" yaml:"moduleName,omitempty"`

	// Version (required)
	Version string `description:"plugin's version." json:"version,omitempty" toml:"version,omitempty" yaml:"version,omitempty"`
}

// DevPlugin The static part of a plugin configuration (only for dev).
type DevPlugin struct {
	// GoPath plugin's GOPATH. (required)
	GoPath string `description:"plugin's GOPATH." json:"goPath,omitempty" toml:"goPath,omitempty" yaml:"goPath,omitempty"`

	// ModuleName (required)
	ModuleName string `description:"plugin's module name."  json:"moduleName,omitempty" toml:"moduleName,omitempty" yaml:"moduleName,omitempty"`
}

type resolverContext struct {
	// GoPath plugin's GOPATH
	GoPath string `json:"goPath,omitempty" toml:"goPath,omitempty" yaml:"goPath,omitempty"`

	// Import plugin's import/package
	Import string `json:"import,omitempty" toml:"import,omitempty" yaml:"import,omitempty"`

	// BasePkg plugin's base package name (optional)
	BasePkg string `json:"basePkg,omitempty" toml:"basePkg,omitempty" yaml:"basePkg,omitempty"`

	interpreter *interp.Interpreter
}

// Builder is a plugin builder.
type Builder struct {
	descriptors map[string]resolverContext
}

// NewBuilder ...
func NewBuilder(client *Client, resolvers map[string]Descriptor) (*Builder, error) {
	builder := &Builder{descriptors: map[string]resolverContext{}}

	for name, desc := range resolvers {
		manifest, err := ReadManifest(client.GoPath(), desc.ModuleName)
		if err != nil {
			return nil, fmt.Errorf("%s: failed to read manifest: %w", desc.ModuleName, err)
		}

		i := interp.New(interp.Options{GoPath: client.GoPath()})
		i.Use(stdlib.Symbols)
		i.Use(Symbols)

		_, err = i.Eval(fmt.Sprintf(`import "%s"`, manifest.Import))
		if err != nil {
			return nil, fmt.Errorf("%s: failed to import plugin code %q: %w", desc.ModuleName, manifest.Import, err)
		}

		builder.descriptors[name] = resolverContext{
			interpreter: i,
			GoPath:      client.GoPath(),
			Import:      manifest.Import,
		}
	}

	return builder, nil
}

// Build ...
func (b *Builder) Build(name string, config map[string]interface{}, resolverName string) (iface.Constructor, error) {
	if b.descriptors == nil {
		return nil, fmt.Errorf("plugin: no plugin definition in the static configuration: %s", name)
	}

	descriptor, ok := b.descriptors[name]
	if !ok {
		return nil, fmt.Errorf("plugin: unknown plugin type: %s", name)
	}

	r, err := newResolver(descriptor, config, resolverName)
	if err != nil {
		return nil, err
	}

	return r.NewResolver, err
}

// Resolver ...
type Resolver struct {
	resolverName string
	fnNew        reflect.Value
	config       reflect.Value
}

func newResolver(descriptor resolverContext, config map[string]interface{}, resolverName string) (*Resolver, error) {
	basePkg := descriptor.BasePkg
	if basePkg == "" {
		basePkg = strings.ReplaceAll(path.Base(descriptor.Import), "-", "_")
	}

	vConfig, err := descriptor.interpreter.Eval(basePkg + `.CreateConfig()`)
	if err != nil {
		return nil, fmt.Errorf("plugin: failed to eval CreateConfig: %w", err)
	}

	cfg := &mapstructure.DecoderConfig{
		DecodeHook:       mapstructure.StringToSliceHookFunc(","),
		WeaklyTypedInput: true,
		Result:           vConfig.Interface(),
	}

	decoder, err := mapstructure.NewDecoder(cfg)
	if err != nil {
		return nil, fmt.Errorf("plugin: failed to create configuration decoder: %w", err)
	}

	err = decoder.Decode(config)
	if err != nil {
		return nil, fmt.Errorf("plugin: failed to decode configuration: %w", err)
	}

	fnNew, err := descriptor.interpreter.Eval(basePkg + `.New`)
	if err != nil {
		return nil, fmt.Errorf("plugin: failed to eval New: %w", err)
	}

	return &Resolver{
		resolverName: resolverName,
		fnNew:        fnNew,
		config:       vConfig,
	}, nil
}

// NewResolver ...
func (m *Resolver) NewResolver() (iface.ResolverHandler, error) {
	args := []reflect.Value{m.config, reflect.ValueOf(m.resolverName)}
	results := m.fnNew.Call(args)

	if len(results) > 1 && results[1].Interface() != nil {
		return nil, results[1].Interface().(error)
	}

	resolver, ok := results[0].Interface().(iface.ResolverHandler)
	if !ok {
		return nil, fmt.Errorf("plugin: invalid handler type: %T", results[0].Interface())
	}

	return resolver, nil
}
