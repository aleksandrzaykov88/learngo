package adder

import (
	"context"
	"grpcadder/pkg/api"
)

// GRPCServer ...
type GRPCServer struct {
	api.UnimplementedAdderServer
}

// Add ...
func (s *GRPCServer) Add(ctx context.Context, req *api.AddRequest) (*api.AddResponse, error) {
	return &api.AddResponse{Result: req.X + req.Y}, nil
}
