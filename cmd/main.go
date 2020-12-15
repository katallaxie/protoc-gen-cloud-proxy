package main

import (
	"os"

	cmd "github.com/katallaxie/protoc-gen-cloud-proxy/internal/cgen"
)

func init() {}

func main() {
	c := cmd.New()

	if err := c.Execute(); err != nil {
		os.Exit(1)
	}
}
