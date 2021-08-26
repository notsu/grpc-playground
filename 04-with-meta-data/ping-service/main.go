package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	client "github.com/notsu/grpc-playground/04-with-meta-data/ping-service/proto"
)

const (
	timestampFormat = time.StampNano
)

func main() {
	fmt.Println("Run ping-service")

	// Create metadata and context.
	md := metadata.Pairs("timestamp", time.Now().Format(timestampFormat))
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	conn, err := grpc.Dial("pong:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := client.NewGreeterClient(conn)

	sayHello(ctx, c)
}

func sayHello(ctx context.Context, c client.GreeterClient) {
	var (
		header, trailer metadata.MD
	)

	response, err := c.SayHello(ctx, &client.HelloRequest{Name: "pong"}, grpc.Header(&header), grpc.Trailer(&trailer))
	if err != nil {
		log.Printf("Error from the server: %s", err)

		return
	}

	// Header
	if t, ok := header["timestamp"]; ok {
		log.Println("timestamp from server header:")
		for i, e := range t {
			log.Printf(" %d. %s\n", i, e)
		}
	} else {
		log.Fatal("timestamp expected but doesn't exist in header")
	}
	if l, ok := header["location"]; ok {
		log.Println("location from server header:")
		for i, e := range l {
			log.Printf(" %d. %s\n", i, e)
		}
	} else {
		log.Fatal("location expected but doesn't exist in header")
	}

	log.Printf("Response from server: %s", response.Message)

	// Trailer
	if t, ok := trailer["timestamp"]; ok {
		log.Println("timestamp from server trailer:")
		for i, e := range t {
			log.Printf(" %d. %s\n", i, e)
		}
	} else {
		log.Fatal("timestamp expected but doesn't exist in trailer")
	}
}
