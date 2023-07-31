package grpc

import (
	"context"

	"github.com/MurashovVen/outsider-sdk/app"
)

var (
	_ app.Work = (*Server)(nil)
)

func (s *Server) Runner(ctx context.Context) func() error {
	return func() error {
		return s.ListenAndServe(ctx)
	}
}

func (s *Server) Name() string {
	return "GRPCServer"
}
