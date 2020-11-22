package main

import (
	"context"
	"math/rand"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
	pb "github.com/katallaxie/protoc-gen-cloud-proxy/example"
	o "github.com/katallaxie/protoc-gen-cloud-proxy/pkg/opts"

	"go.uber.org/zap"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	defer logger.Sync()

	s, err := session.NewSession()
	if err != nil {
		panic(err)
	}

	options := o.New(o.WithVerbose(), o.WithLogger(logger), o.WithSession(s))
	p := pb.NewProxy(options)

	if err := p.Start(ctx); err != nil {
		panic(err)
	}
}
