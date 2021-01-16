package grpc

import (
	"context"
	"fmt"
	"io"
	"log"
	"sync"
	"time"

	"github.com/notsu/grpc-playground/01-basic/pong-service/proto"
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

// LotsOfReplies returns pong messages streaming back to the client
func (s *Server) LotsOfReplies(r *proto.HelloRequest, srv proto.Greeter_LotsOfRepliesServer) error {
	log.Printf("received: %v", r.Name)

	// use wait group to allow process to be concurrent
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func(count int64) {
			defer wg.Done()

			//time sleep to simulate server process time
			time.Sleep(time.Duration(count) * time.Second)

			resp := proto.HelloResponse{
				Message: fmt.Sprintf("pong #%d for name: %s", count, r.Name),
			}

			if err := srv.Send(&resp); err != nil {
				log.Printf("send error %v", err)
			}

			log.Printf("finishing request number : %d", count)
		}(int64(i))
	}

	wg.Wait()

	return nil
}

// LotsOfGreetings receives streaming greetings and send back a pong message
func (s *Server) LotsOfGreetings(srv proto.Greeter_LotsOfGreetingsServer) error {
	i := 1

	for {
		in, err := srv.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}

		log.Printf("receive %s", in.Name)

		if i > 5 {
			out := &proto.HelloResponse{
				Message: "pong",
			}

			if err := srv.SendAndClose(out); err != nil {
				return err
			}

			return nil
		}

		i++
	}
}

// BidiHello steams greeting for both client and server
func (s *Server) BidiHello(srv proto.Greeter_BidiHelloServer) error {
	i := 1

	for {
		in, err := srv.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}

		log.Printf("receive %s", in.Name)

		out := &proto.HelloResponse{
			Message: "pong",
		}

		if err := srv.Send(out); err != nil {
			return err
		}

		i++
	}
}
