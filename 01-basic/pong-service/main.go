package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	server "github.com/notsu/grpc-playground/01-basic/pong-service/grpc"
	"github.com/notsu/grpc-playground/01-basic/pong-service/proto"
)

func main() {
	fmt.Println("Run pong-service")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	proto.RegisterGreeterServer(grpcServer, &server.Server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
