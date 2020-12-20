// +build func_test

package resolvers

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBuilder(t *testing.T) {
	tests := []struct {
		desc    string
		name    string
		version string
		archive string
	}{
		{desc: "should build a new builder from demo", name: "github.com/cgentron/plugindemo", version: "master", archive: "https://github.com/cgentron/plugindemo/archive/main.zip"},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			c, err := NewClient()
			assert.NoError(t, err)

			ctx := context.Background()

			str, err := c.Fetch(ctx, tt.name, tt.version, tt.archive)
			assert.NoError(t, err)
			assert.Empty(t, "", str)

			err = c.Unzip(tt.name, tt.version)
			assert.NoError(t, err)

			resolvers := map[string]Descriptor{
				"demo": {ModuleName: tt.name, Version: tt.version},
			}

			b, err := NewBuilder(c, resolvers)
			assert.NoError(t, err)
			assert.NotNil(t, b)
		})
	}
}

func TestNewBuilder_Build(t *testing.T) {
	tests := []struct {
		desc    string
		name    string
		version string
		archive string
	}{
		{desc: "should build a new builder from demo", name: "github.com/cgentron/plugindemo", version: "master", archive: "https://github.com/cgentron/plugindemo/archive/main.zip"},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			c, err := NewClient()
			assert.NoError(t, err)

			ctx := context.Background()

			str, err := c.Fetch(ctx, tt.name, tt.version, tt.archive)
			assert.NoError(t, err)
			assert.Empty(t, "", str)

			err = c.Unzip(tt.name, tt.version)
			assert.NoError(t, err)

			resolvers := map[string]Descriptor{
				"demo": {ModuleName: tt.name, Version: tt.version},
			}

			b, err := NewBuilder(c, resolvers)
			assert.NoError(t, err)
			assert.NotNil(t, b)

			config := map[string]interface{}{}

			r, err := b.Build("demo", config, tt.name)
			assert.NoError(t, err)
			assert.NotNil(t, r)
		})
	}
}

func TestNewResolver(t *testing.T) {
	tests := []struct {
		desc    string
		name    string
		version string
		archive string
	}{
		{desc: "should build a new builder from demo", name: "github.com/cgentron/plugindemo", version: "master", archive: "https://github.com/cgentron/plugindemo/archive/main.zip"},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			c, err := NewClient()
			assert.NoError(t, err)

			ctx := context.Background()

			str, err := c.Fetch(ctx, tt.name, tt.version, tt.archive)
			assert.NoError(t, err)
			assert.Empty(t, "", str)

			err = c.Unzip(tt.name, tt.version)
			assert.NoError(t, err)

			resolvers := map[string]Descriptor{
				"demo": {ModuleName: tt.name, Version: tt.version},
			}

			b, err := NewBuilder(c, resolvers)
			assert.NoError(t, err)
			assert.NotNil(t, b)

			config := map[string]interface{}{}

			r, err := b.Build("demo", config, tt.name)
			assert.NoError(t, err)
			assert.NotNil(t, r)
		})
	}
}

func TestNewResolver_Resolve(t *testing.T) {
	tests := []struct {
		desc    string
		name    string
		version string
		archive string
		in      []byte
		out     []byte
	}{
		{desc: "should build a new builder from demo", name: "github.com/cgentron/plugindemo", version: "master", archive: "https://github.com/cgentron/plugindemo/archive/main.zip", in: []byte("foo"), out: []byte("demo")},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			c, err := NewClient()
			assert.NoError(t, err)

			ctx := context.Background()

			str, err := c.Fetch(ctx, tt.name, tt.version, tt.archive)
			assert.NoError(t, err)
			assert.Empty(t, "", str)

			err = c.Unzip(tt.name, tt.version)
			assert.NoError(t, err)

			resolvers := map[string]Descriptor{
				"demo": {ModuleName: tt.name, Version: tt.version},
			}

			b, err := NewBuilder(c, resolvers)
			assert.NoError(t, err)

			config := map[string]interface{}{}

			r, err := b.Build("demo", config, tt.name)
			assert.NoError(t, err)

			rr, err := r()
			assert.NoError(t, err)

			out, err := rr.Resolve(ctx, tt.in)
			assert.NoError(t, err)
			assert.Equal(t, tt.out, out)
		})
	}
}
