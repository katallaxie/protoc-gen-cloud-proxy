// +build codegen

package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var root = &cobra.Command{
	Use: "proto-gen-cli",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

var logger *zap.Logger

func init() {
	rand.Seed(time.Now().UnixNano())
}

func execute() {
	fmt.Println("test")
	if err := root.Execute(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}

func main() {
	execute()
}
