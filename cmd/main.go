package main

import (
	"os"

	cmd "github.com/katallaxie/protoc-gen-cloud-proxy/pkg/cgen"
)

func init() {}

func Execute() {
	c := cmd.New()

	if err := c.Execute(); err != nil {
		os.Exit(1)
	}
}
