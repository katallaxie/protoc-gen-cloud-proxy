package cmd

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"math"
	"net"
	"time"

	pb "github.com/awslabs/protoc-gen-aws-service-proxy/example"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	health "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/status"
)

type srv struct {
}

func (s *srv) Start(ctx context.Context, ready func()) func() error {
	return func() error {
		lis, err := net.Listen("tcp", viper.GetString("addr"))
		if err != nil {
			return err
		}

		srv := &service{}

		tlsConfig := &tls.Config{}
		tlsConfig.InsecureSkipVerify = true
		srv.tlsCfg = tlsConfig

		var kaep = keepalive.EnforcementPolicy{
			MinTime:             5 * time.Second, // If a client pings more than once every 5 seconds, terminate the connection
			PermitWithoutStream: true,            // Allow pings even when there are no active streams
		}

		var kasp = keepalive.ServerParameters{
			MaxConnectionIdle:     time.Duration(math.MaxInt64), // If a client is idle for 15 seconds, send a GOAWAY
			MaxConnectionAge:      time.Duration(math.MaxInt64), // If any connection is alive for more than 30 seconds, send a GOAWAY
			MaxConnectionAgeGrace: 5 * time.Second,              // Allow 5 seconds for pending RPCs to complete before forcibly closing connections
			Time:                  5 * time.Second,              // Ping the client if it is idle for 5 seconds to ensure the connection is still active
			Timeout:               1 * time.Second,              // Wait 1 second for the ping ack before assuming the connection is dead
		}

		ss := grpc.NewServer(grpc.KeepaliveEnforcementPolicy(kaep), grpc.KeepaliveParams(kasp))
		pb.RegisterExampleServer(ss, srv)
		health.RegisterHealthServer(ss, srv)

		ready()

		if err := ss.Serve(lis); err != nil {
			return err
		}

		return nil
	}
}

type service struct {
	tlsCfg *tls.Config
	pb.UnimplementedExampleServer
}

func (s *service) Insert(ctx context.Context, req *pb.Insert_Request) (*pb.Insert_Response, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	svc := lambda.New(session.New())
	input := &lambda.InvokeInput{
		FunctionName: aws.String("dartTest"),
		Payload:      b,
		Qualifier:    aws.String("$LATEST"),
	}

	result, err := svc.Invoke(input)
	if err != nil {
		return nil, err
	}

	var payload pb.Insert_Response
	if err := json.Unmarshal(result.Payload, &payload); err != nil {
		return nil, err
	}

	return &payload, nil
}

func (s *service) Check(ctx context.Context, in *health.HealthCheckRequest) (*health.HealthCheckResponse, error) {
	log.Infof("Received Check request: %v", in)

	return &health.HealthCheckResponse{Status: health.HealthCheckResponse_SERVING}, nil
}

func (s *service) Watch(in *health.HealthCheckRequest, _ health.Health_WatchServer) error {
	log.Infof("Received Watch request: %v", in)

	return status.Error(codes.Unimplemented, "unimplemented")
}
