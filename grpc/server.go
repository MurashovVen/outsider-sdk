package grpc

import (
	"context"
	"net"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"

	"github.com/MurashovVen/outsider-sdk/app/logger"
)

type (
	Server struct {
		*grpc.Server

		address string

		services []ServerRegisterer
	}

	ServerRegisterer interface {
		Register(*Server)
	}
)

func NewServer(
	address string, services []ServerRegisterer, options ...grpc.ServerOption,
) *Server {
	return &Server{
		address:  address,
		services: services,
		Server:   grpc.NewServer(options...),
	}
}

func (s *Server) ListenAndServe(ctx context.Context) error {
	lis, err := net.Listen("tcp", s.address)
	if err != nil {
		return err
	}

	for _, service := range s.services {
		service.Register(s)
	}

	eg, egCtx := errgroup.WithContext(ctx)

	eg.Go(
		func() error {
			return s.Serve(lis)
		},
	)

	eg.Go(
		func() error {
			<-egCtx.Done()

			s.GracefulStop()

			return egCtx.Err()
		},
	)

	return eg.Wait()
}

// todo:
// https://github.com/grpc-ecosystem/go-grpc-middleware/tree/main/interceptors/recovery
// https://github.com/grpc-ecosystem/go-grpc-middleware/tree/main/interceptors/timeout
// https://github.com/grpc-ecosystem/go-grpc-middleware/blob/main/interceptors/retry/retry.go
func DefaultServerOptions(log *logger.Logger) []grpc.ServerOption {
	return []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(
			logging.UnaryServerInterceptor(interceptorLogger(log.Logger)),
		),
		grpc.ChainStreamInterceptor(
			logging.StreamServerInterceptor(interceptorLogger(log.Logger)),
		),
	}
}
