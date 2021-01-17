package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"

	client "github.com/notsu/grpc-playground/02-error-handling/ping-service/proto"
)

func main() {
	fmt.Println("Run ping-service")

	// Timeout in 10 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn, err := grpc.Dial("pong:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := client.NewGreeterClient(conn)

	sayHello(ctx, c)
}

func sayHello(ctx context.Context, c client.GreeterClient) {
	response, err := c.SayHello(ctx, &client.HelloRequest{Name: "pong"})
	if err != nil {
		s := status.Convert(err)

		for _, d := range s.Details() {
			switch info := d.(type) {
			case *client.Error:
				log.Printf("Custom failure: %s", info)
			default:
				log.Fatalf("Error when calling SayHello: %s", err)
			}
		}

		return
	}

	log.Printf("Response from server: %s", response.Message)
}
