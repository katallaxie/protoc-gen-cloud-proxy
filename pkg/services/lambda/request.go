package lambda

import (
	"github.com/valyala/fasthttp"

	"github.com/katallaxie/protoc-gen-cloud-proxy/pkg/services"
)

type InvocationType string

const (
	InvocationTypeSync   InvocationType = "RequestResponse"
	InvocationTypeAsync  InvocationType = "Event"
	InvocationTypeDryRun InvocationType = "DryRun"
)

type LogType string

const (
	LogTypeNone LogType = "None"
	LogTypeTail LogType = "Tail"
)

type lambda struct {
	cfg Config
}

func New() services.Request {
	l := &lambda{}

	return l
}

func (l *lambda) Do() error {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(l.cfg.Endpoint())

	req.Header.Add("X-Amz-Target", "invoke")
	req.Header.Add("X-Amz-Invocation-Type", string(l.cfg.invocationType))
	req.Header.Add("X-Amz-Log-Type", string(l.cfg.logType))
	req.Header.Add("Content-Type", "application/x-amz-json-1.1")
	// Content-Length ...

	// Post the request ...
	req.Header.SetMethod("POST")

	// get security credentials here ... this should be a general thing

	resp := fasthttp.AcquireResponse()
	client := &fasthttp.Client{}
	if err := client.Do(req, resp); err != nil {
		return err
	}

	return nil
}
