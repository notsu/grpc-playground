package grpc

import (
	"context"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/notsu/grpc-playground/03-with-authentication/pong-service/proto"
)

var (
	errInvalidArguments = status.Errorf(codes.InvalidArgument, "Insufficient arguments")
)

// Server represents the gRPC server
type Server struct{}

// SayHello returns a pong message back to the client
func (s *Server) SayHello(ctx context.Context, r *proto.HelloRequest) (*proto.HelloResponse, error) {
	log.Printf("received: %v", r.Name)

	return &proto.HelloResponse{
		Message: "pong",
	}, nil
}
