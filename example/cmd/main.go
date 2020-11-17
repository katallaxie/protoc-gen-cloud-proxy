package main

import (
	"context"
	"math/rand"
	"time"

	o "github.com/katallaxie/protoc-gen-cloud-proxy/pkg/opts"
	"github.com/katallaxie/protoc-gen-cloud-proxy/pkg/proxy"

	pb "github.com/katallaxie/protoc-gen-cloud-proxy/example"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	p := proxy.New()
	srv := pb.NewService(o.WithVerbose())

	if err := p.Start(ctx, srv); err != nil {
		panic(err)
	}
}
