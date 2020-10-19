package lambda

import (
	"fmt"

	"github.com/katallaxie/protoc-gen-cloud-proxy/pkg/config"
)

type Config struct {
	functionName   string
	invocationType InvocationType
	logType        LogType
	config.Config
}

func (c *Config) Endpoint() string {
	return fmt.Sprintf("https://lambda.%s.amazonaws.com/2015-03-31/functions/%s/invocations", c.AWSRegion, c.functionName)
}
