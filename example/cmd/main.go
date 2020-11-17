package main

import (
	"context"
	"math/rand"
	"time"

	pb "github.com/katallaxie/protoc-gen-cloud-proxy/example"
	o "github.com/katallaxie/protoc-gen-cloud-proxy/pkg/opts"
	"github.com/katallaxie/protoc-gen-cloud-proxy/pkg/proxy"

	"go.uber.org/zap"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	defer logger.Sync()

	p := proxy.New()
	srv := pb.NewService(
		o.WithVerbose(),
		o.WithLogger(logger),
	)

	if err := p.Start(ctx, srv); err != nil {
		panic(err)
	}
}
