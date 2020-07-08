// Code generated by aws-grpc-service-proxy. DO NOT EDIT.

package proto

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"net"
	"os"
	"time"

	"github.com/andersnormal/pkg/server"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/golang/protobuf/jsonpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

// RootCmd ...
var RootCmd = &cobra.Command{
	Use:   "server",
	Short: "Runs the gRPC service",
	RunE:  runE,
}

func addFlags(cmd *cobra.Command) {
	cmd.Flags().String("log-format", "text", "log format")
	cmd.Flags().String("log-level", "info", "log-level")
	cmd.Flags().String("addr", ":9090", "address")

	// set the link between flags
	viper.BindPFlag("log-format", cmd.Flags().Lookup("log-format"))
	viper.BindPFlag("log-level", cmd.Flags().Lookup("log-level"))
	viper.BindPFlag("addr", cmd.Flags().Lookup("addr"))
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

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// SongJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of SongJSONMarshaler This struct is safe to replace or modify but
// should not be done so concurrently.
var SongJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *Song) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}
	buf := &bytes.Buffer{}
	if err := SongJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

var _ json.Marshaler = (*Song)(nil)

// SongJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of Song. This struct is safe to replace or modify but
// should not be done so concurrently.
var SongJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *Song) UnmarshalJSON(b []byte) error {
	return SongJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*Song)(nil)

// InsertJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of InsertJSONMarshaler This struct is safe to replace or modify but
// should not be done so concurrently.
var InsertJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *Insert) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}
	buf := &bytes.Buffer{}
	if err := InsertJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

var _ json.Marshaler = (*Insert)(nil)

// InsertJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of Insert. This struct is safe to replace or modify but
// should not be done so concurrently.
var InsertJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *Insert) UnmarshalJSON(b []byte) error {
	return InsertJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*Insert)(nil)

// Insert_RequestJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of Insert_RequestJSONMarshaler This struct is safe to replace or modify but
// should not be done so concurrently.
var Insert_RequestJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *Insert_Request) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}
	buf := &bytes.Buffer{}
	if err := Insert_RequestJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

var _ json.Marshaler = (*Insert_Request)(nil)

// Insert_RequestJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of Insert_Request. This struct is safe to replace or modify but
// should not be done so concurrently.
var Insert_RequestJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *Insert_Request) UnmarshalJSON(b []byte) error {
	return Insert_RequestJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*Insert_Request)(nil)

// Insert_ResponseJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of Insert_ResponseJSONMarshaler This struct is safe to replace or modify but
// should not be done so concurrently.
var Insert_ResponseJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *Insert_Response) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}
	buf := &bytes.Buffer{}
	if err := Insert_ResponseJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

var _ json.Marshaler = (*Insert_Response)(nil)

// Insert_ResponseJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of Insert_Response. This struct is safe to replace or modify but
// should not be done so concurrently.
var Insert_ResponseJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *Insert_Response) UnmarshalJSON(b []byte) error {
	return Insert_ResponseJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*Insert_Response)(nil)

// UpdateJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of UpdateJSONMarshaler This struct is safe to replace or modify but
// should not be done so concurrently.
var UpdateJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *Update) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}
	buf := &bytes.Buffer{}
	if err := UpdateJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

var _ json.Marshaler = (*Update)(nil)

// UpdateJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of Update. This struct is safe to replace or modify but
// should not be done so concurrently.
var UpdateJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *Update) UnmarshalJSON(b []byte) error {
	return UpdateJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*Update)(nil)

// Update_RequestJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of Update_RequestJSONMarshaler This struct is safe to replace or modify but
// should not be done so concurrently.
var Update_RequestJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *Update_Request) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}
	buf := &bytes.Buffer{}
	if err := Update_RequestJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

var _ json.Marshaler = (*Update_Request)(nil)

// Update_RequestJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of Update_Request. This struct is safe to replace or modify but
// should not be done so concurrently.
var Update_RequestJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *Update_Request) UnmarshalJSON(b []byte) error {
	return Update_RequestJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*Update_Request)(nil)

// Update_ResponseJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of Update_ResponseJSONMarshaler This struct is safe to replace or modify but
// should not be done so concurrently.
var Update_ResponseJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *Update_Response) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}
	buf := &bytes.Buffer{}
	if err := Update_ResponseJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

var _ json.Marshaler = (*Update_Response)(nil)

// Update_ResponseJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of Update_Response. This struct is safe to replace or modify but
// should not be done so concurrently.
var Update_ResponseJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *Update_Response) UnmarshalJSON(b []byte) error {
	return Update_ResponseJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*Update_Response)(nil)

// ListSongsJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of ListSongsJSONMarshaler This struct is safe to replace or modify but
// should not be done so concurrently.
var ListSongsJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *ListSongs) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}
	buf := &bytes.Buffer{}
	if err := ListSongsJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

var _ json.Marshaler = (*ListSongs)(nil)

// ListSongsJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of ListSongs. This struct is safe to replace or modify but
// should not be done so concurrently.
var ListSongsJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *ListSongs) UnmarshalJSON(b []byte) error {
	return ListSongsJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*ListSongs)(nil)

// ListSongs_RequestJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of ListSongs_RequestJSONMarshaler This struct is safe to replace or modify but
// should not be done so concurrently.
var ListSongs_RequestJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *ListSongs_Request) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}
	buf := &bytes.Buffer{}
	if err := ListSongs_RequestJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

var _ json.Marshaler = (*ListSongs_Request)(nil)

// ListSongs_RequestJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of ListSongs_Request. This struct is safe to replace or modify but
// should not be done so concurrently.
var ListSongs_RequestJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *ListSongs_Request) UnmarshalJSON(b []byte) error {
	return ListSongs_RequestJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*ListSongs_Request)(nil)

// ListSongs_ResponseJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of ListSongs_ResponseJSONMarshaler This struct is safe to replace or modify but
// should not be done so concurrently.
var ListSongs_ResponseJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *ListSongs_Response) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}
	buf := &bytes.Buffer{}
	if err := ListSongs_ResponseJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

var _ json.Marshaler = (*ListSongs_Response)(nil)

// ListSongs_ResponseJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of ListSongs_Response. This struct is safe to replace or modify but
// should not be done so concurrently.
var ListSongs_ResponseJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *ListSongs_Response) UnmarshalJSON(b []byte) error {
	return ListSongs_ResponseJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*ListSongs_Response)(nil)

// EmptyJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of EmptyJSONMarshaler This struct is safe to replace or modify but
// should not be done so concurrently.
var EmptyJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *Empty) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}
	buf := &bytes.Buffer{}
	if err := EmptyJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

var _ json.Marshaler = (*Empty)(nil)

// EmptyJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of Empty. This struct is safe to replace or modify but
// should not be done so concurrently.
var EmptyJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *Empty) UnmarshalJSON(b []byte) error {
	return EmptyJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*Empty)(nil)

type srv struct {
}

type service struct {
	tlsCfg *tls.Config
	UnimplementedExampleServer
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
		RegisterExampleServer(ss, srv)
		// health.RegisterHealthServer(ss, srv)

		ready()

		if err := ss.Serve(lis); err != nil {
			return err
		}

		return nil
	}
}

// Here goes a message Insert

func (s *service) Insert(ctx context.Context, req *Insert_Request) (*Insert_Response, error) {
	// Encode the Message into JSON representation
	b, err := req.MarshalJSON()
	if err != nil {
		return nil, err
	}

	svc := lambda.New(session.New())
	input := &lambda.InvokeInput{
		FunctionName: aws.String("arn:aws:lambda:us-east-2:123456789012:function:my-function"),
		Payload:      b,
		Qualifier:    "$LATEST",
	}

	result, err := svc.Invoke(input)
	if err != nil {
		return nil, err
	}

	var payload Insert_Response
	if err := payload.UnmarshalJSON(result.Payload); err != nil {
		return nil, err
	}

	return nil, nil

}

// Here goes a message Update

func (s *service) Update(ctx context.Context, req *Update_Request) (*Update_Response, error) {

	// DynamoDB
	return nil, nil

}
