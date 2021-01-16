package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"

	client "github.com/notsu/grpc-playground/01-basic/ping-service/proto"
)

var (
	method = "bidiHello"
)

func main() {
	fmt.Println("Run ping-service")
	ctx := context.Background()

	conn, err := grpc.Dial("pong:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := client.NewGreeterClient(conn)

	switch method {
	case "sayHello":
		sayHello(ctx, c)
	case "lotsOfReplies":
		lotsOfReplies(ctx, c)
	case "lotsOfGreetings":
		lotsOfGreetings(ctx, c)
	case "bidiHello":
		bidiHello(ctx, c)
	}
}

func sayHello(ctx context.Context, c client.GreeterClient) {
	response, err := c.SayHello(ctx, &client.HelloRequest{Name: "pong"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}

	log.Printf("Response from server: %s", response.Message)
}

func lotsOfReplies(ctx context.Context, c client.GreeterClient) {
	done := make(chan bool)

	stream, err := c.LotsOfReplies(ctx, &client.HelloRequest{Name: "pong"})
	if err != nil {
		log.Fatalf("Error when connect to the server: %s", err)
	}

	go func(stream client.Greeter_LotsOfRepliesClient) {
		for {
			msg, err := stream.Recv()
			if err == io.EOF {
				done <- true
				return
			}

			if err != nil {
				log.Fatalf("Error when receive a message: %s", err)
			}

			log.Printf("Receive a message: %s", msg.Message)
		}
	}(stream)

	<-done

	log.Println("Done!")
}

func lotsOfGreetings(ctx context.Context, c client.GreeterClient) {
	stream, err := c.LotsOfGreetings(ctx)
	if err != nil {
		log.Fatalf("Error when connect to the server: %s", err)
	}

	for i := 1; i < 10; i++ {
		stream.Send(&client.HelloRequest{
			Name: fmt.Sprintf("Request number %d", i),
		})
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Printf("Failed to close and receive: %s", err)
	}

	log.Printf("Receive a response: %s", resp.Message)

	log.Println("Done!")
}

func bidiHello(ctx context.Context, c client.GreeterClient) {
	done := make(chan bool)

	stream, err := c.BidiHello(ctx)
	if err != nil {
		log.Fatalf("Error when connect to the server: %s", err)
	}

	go func(stream client.Greeter_BidiHelloClient) {
		for {
			msg, err := stream.Recv()
			if err == io.EOF {
				done <- true
				return
			}

			if err != nil {
				log.Fatalf("Error when receive a message: %s", err)
			}

			log.Printf("Receive a message: %s", msg.Message)
		}
	}(stream)

	go func(stream client.Greeter_BidiHelloClient) {
		i := 1

		for {
			msg := fmt.Sprintf("Request number %d", i)

			err := stream.Send(&client.HelloRequest{
				Name: msg,
			})
			if err == io.EOF {
				done <- true
				return
			}

			if err != nil {
				log.Fatalf("Error when receive a message: %s", err)
			}

			log.Printf("Send a message: %s", msg)

			i++

			time.Sleep(3 * time.Second)
		}
	}(stream)

	<-done

	log.Println("Done!")
}
