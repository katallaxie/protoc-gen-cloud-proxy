package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	pb "github.com/katallaxie/protoc-gen-cloud-proxy/example"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	if err := pb.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
