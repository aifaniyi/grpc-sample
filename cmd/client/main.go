package main

import (
	"context"
	"io"
	"log"

	"github.com/aifaniyi/grpc-sample/pkg/dto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address = "localhost:8080"
)

func main() {
	opts := grpc.WithTransportCredentials(insecure.NewCredentials())

	conn, err := grpc.Dial(address, opts)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Unary: client greeting request
	client := dto.NewGreetingServiceClient(conn)
	request := &dto.GreetingServiceRequest{Count: 10}

	resp, err := client.Greeting(context.Background(), request)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("response: %v", resp)

	// Streaming server: client lots of greeting replies request
	stream, err := client.LotsOfGreetingReplies(context.Background(), request)
	if err != nil {
		log.Fatal(err)
	}

	for {
		resp, err := stream.Recv()
		if err == nil {
			log.Print(resp)
			continue
		}

		if err == io.EOF {
			return
		}

		if err != nil {
			log.Fatalf("receive error : %v", err)
		}
	}
}
