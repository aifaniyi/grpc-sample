package main

import (
	"log"
	"net"

	"github.com/aifaniyi/grpc-sample/pkg/dto"
	"github.com/aifaniyi/grpc-sample/pkg/server"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("error binding to port %s", port)
	}

	s := grpc.NewServer()

	dto.RegisterGreetingServiceServer(s, &server.Server{})

	log.Printf("starting server on port %s", port)
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
