package templates

const serviceTpl = `

type srv struct {
	opts *o.Opts
}

type service struct {
  tlsCfg *tls.Config
  logger *zap.Logger
  session *session.Session
	UnimplementedExampleServer
}

func New(opts *o.Opts) proxy.Listener {
  s := new(srv)
  s.opts = opts

  return s
}

func NewProxy(opts *o.Opts) proxy.Proxy {
  s := New(opts)
  p := proxy.New(s, opts)

	return p
}

func (s *srv) Start(ctx context.Context, ready func()) func() error {
	return func() error {
		lis, err := net.Listen("tcp", s.opts.Addr)
		if err != nil {
			return err
		}

		ll := s.opts.Logger.With(zap.String("addr", s.opts.Addr))
		srv := &service{
      session: s.opts.Session,
      logger: s.opts.Logger,
    }

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

		grpc_zap.ReplaceGrpcLogger(ll)

		ss := grpc.NewServer(
			grpc.KeepaliveEnforcementPolicy(kaep),
			grpc.KeepaliveParams(kasp),
			grpc_middleware.WithUnaryServerChain(grpc_zap.UnaryServerInterceptor(ll)))

		RegisterExampleServer(ss, srv)
		grpc_health_v1.RegisterHealthServer(ss, health.NewServer())
		// grpc_health_v1.RegisterHealthServer(ss, srv)

		ready()

		ll.Info("start listening")

		if err := ss.Serve(lis); err != nil {
			return err
		}

		return nil
	}
}
`
