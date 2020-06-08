package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"context"
	"crypto/tls"
	"encoding/json"
	"math"
	"net"

	"github.com/andersnormal/pkg/server"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	pb "github.com/awslabs/protoc-gen-aws-service-proxy/example"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	health "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/status"
)

// RootCmd ...
var RootCmd = &cobra.Command{
	Use:   "server",
	Short: "Runs the gRPC service",
	Long:  `Not yet`,
	RunE:  runE,
}

func addFlags(cmd *cobra.Command) {
	cmd.Flags().String("log-format", "text", "log format")
	cmd.Flags().String("log-level", "info", "log-level")
	cmd.Flags().String("addr", ":9090", "address")
	cmd.Flags().String("topic", "monolog", "kafka, topic")
	cmd.Flags().StringSlice("brokers", []string{"localhost:9092"}, "kafka brokers")
	cmd.Flags().String("version", "2.3.2", "kafka version")

	// set the link between flags
	viper.BindPFlag("log-format", cmd.Flags().Lookup("log-format"))
	viper.BindPFlag("log-level", cmd.Flags().Lookup("log-level"))
	viper.BindPFlag("addr", cmd.Flags().Lookup("addr"))
	viper.BindPFlag("topic", cmd.Flags().Lookup("topic"))
	viper.BindPFlag("brokers", cmd.Flags().Lookup("brokers"))
	viper.BindPFlag("version", cmd.Flags().Lookup("version"))
}

func init() {
	// initialize cobra
	cobra.OnInitialize(initConfig)

	// adding flags
	addFlags(RootCmd)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// set the default format, which is basically text
	log.SetFormatter(&log.TextFormatter{})

	viper.AutomaticEnv() // read in environment variables that match

	// config logger
	logConfig()
}

func logConfig() {
	// reset log format
	if viper.GetString("log-format") == "json" {
		log.SetFormatter(&log.JSONFormatter{})
	}

	// set the configured log level
	if level, err := log.ParseLevel(viper.GetString("log-level")); err == nil {
		log.SetLevel(level)
	}
}

type root struct {
	logger *log.Entry
}

func runE(cmd *cobra.Command, args []string) error {
	// create a new root
	root := new(root)

	// init logger
	root.logger = log.WithFields(log.Fields{
		"verbose": viper.GetBool("verbose"),
		"brokers": viper.GetStringSlice("brokers"),
		"topic":   viper.GetString("topic"),
	})

	// create root context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// create server
	s, _ := server.WithContext(ctx)

	// log ...
	root.logger.Info("Starting server ...")

	// debug listener
	debug := server.NewDebugListener(
		server.WithPprof(),
		server.WithStatusAddr(":8443"),
	)
	s.Listen(debug, true)

	// listen for grpc
	s.Listen(&srv{}, true)

	// listen for the server and wait for it to fail,
	// or for sys interrupts
	if err := s.Wait(); err != nil {
		root.logger.Error(err)
	}

	// noop
	return nil
}

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

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
