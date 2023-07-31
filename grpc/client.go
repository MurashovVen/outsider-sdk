package grpc

import (
	"time"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"google.golang.org/grpc"
	"google.golang.org/grpc/backoff"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/MurashovVen/outsider-sdk/app/logger"
)

type ClientConn struct {
	*grpc.ClientConn
}

func MustConnect(address string, options ...grpc.DialOption) *grpc.ClientConn {
	conn, err := Connect(address, options...)
	if err != nil {
		panic("dialing grpc: " + err.Error())
	}

	return conn
}

func Connect(address string, options ...grpc.DialOption) (*grpc.ClientConn, error) {
	return grpc.Dial(address, options...)
}

// todo:
// https://github.com/grpc-ecosystem/go-grpc-middleware/tree/main/interceptors/recovery
// https://github.com/grpc-ecosystem/go-grpc-middleware/tree/main/interceptors/timeout
// https://github.com/grpc-ecosystem/go-grpc-middleware/blob/main/interceptors/retry/retry.go
func DefaultDialOptions(log *logger.Logger) []grpc.DialOption {
	return []grpc.DialOption{
		grpc.WithConnectParams(
			grpc.ConnectParams{
				Backoff:           backoff.DefaultConfig,
				MinConnectTimeout: 5 * time.Second,
			},
		),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainUnaryInterceptor(
			logging.UnaryClientInterceptor(interceptorLogger(log.Logger)),
		),
		grpc.WithStreamInterceptor(
			logging.StreamClientInterceptor(interceptorLogger(log.Logger)),
		),
	}
}
