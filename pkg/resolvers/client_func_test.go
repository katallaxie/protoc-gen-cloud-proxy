// +build func_test

package resolvers

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFetch(t *testing.T) {
	tests := []struct {
		desc    string
		name    string
		version string
		archive string
	}{
		{desc: "should unzip demo plugin", name: "demo", version: "master", archive: "https://github.com/cgentron/plugindemo/archive/main.zip"},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			c, err := NewClient()
			assert.NoError(t, err)

			ctx := context.Background()

			str, err := c.Fetch(ctx, tt.name, tt.version, tt.archive)
			assert.NoError(t, err)
			assert.Empty(t, "", str)
		})
	}
}

func TestUnzip(t *testing.T) {
	tests := []struct {
		desc    string
		name    string
		version string
		archive string
	}{
		{desc: "should unzip demo", name: "demo", version: "master", archive: "https://github.com/cgentron/plugindemo/archive/main.zip"},
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
		})
	}
}

func TestReadManifest(t *testing.T) {
	tests := []struct {
		desc       string
		name       string
		version    string
		archive    string
		moduleName string
	}{
		{desc: "should read demo manifest", name: "demo", version: "master", archive: "https://github.com/cgentron/plugindemo/archive/main.zip"},
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

			manifest, err := c.ReadManifest(tt.name)
			assert.NoError(t, err)
			assert.Equal(t, "Demo Plugin", manifest.Name)
			assert.Equal(t, "github.com/cgentron/plugindemo", manifest.Import)
		})
	}
}
