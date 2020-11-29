// +build codegen

package main

import (
	"log"
	"math/rand"
	"time"

	o "github.com/katallaxie/protoc-gen-cloud-proxy/internal/gen/opts"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var root = &cobra.Command{
	Use: "proto-gen-cli",
	RunE: func(cmd *cobra.Command, args []string) error {
		opts.Logger.Info("running code generation")

		return nil
	},
}

var opts *o.Opts

func init() {
	rand.Seed(time.Now().UnixNano())

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("cannot initialize zap logger: %v", err)
	}

	opts = o.New(o.WithLogger(logger))

	cobra.OnInitialize(initConfig)
	root.PersistentFlags().StringVarP(&opts.Path, "path", "p", opts.Path, "source files")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
}

func execute() {
	if err := root.Execute(); err != nil {
		opts.Logger.Fatal("%+v", zap.Error(err))
	}
}

func main() {
	execute()
}
