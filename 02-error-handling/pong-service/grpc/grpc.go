package grpc

import (
	"context"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/notsu/grpc-playground/02-error-handling/pong-service/proto"
)

var (
	errInvalidArguments = status.Errorf(codes.InvalidArgument, "Insufficient arguments")
)

// Server represents the gRPC server
type Server struct{}

// SayHello returns a pong message back to the client
func (s *Server) SayHello(ctx context.Context, r *proto.HelloRequest) (*proto.HelloResponse, error) {
	log.Printf("received: %v", r.Name)

	st := status.New(codes.InvalidArgument, "Insufficient arguments")
	dt, err := st.WithDetails(&proto.Error{
		Code:    100,
		Message: "Something wrong na ja",
	})
	if err != nil {
		return nil, st.Err()
	}

	return nil, dt.Err()
}
