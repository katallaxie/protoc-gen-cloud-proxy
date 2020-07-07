package main

import (
  "time"
  "os"
  "math/rand"
  "fmt"

  pb "github.com/awslabs/protoc-gen-aws-service-proxy/example"
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
