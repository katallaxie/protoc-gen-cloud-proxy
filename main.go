package main

import (
	"github.com/katallaxie/protoc-gen-cloud-proxy/pkg/module"

	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
)

func main() {
	pgs.
		Init(pgs.DebugEnv("DEBUG_PGV")).
		RegisterModule(module.Proxy()).
		RegisterPostProcessor(pgsgo.GoFmt()).
		Render()
}
