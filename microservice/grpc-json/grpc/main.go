package main

import (
	"grpc-json/internal"
	kv "grpc-json/keyvalue"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	port = "50051"
)

func main() {
	l, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	kv.RegisterKeyValueServer(s, internal.NewKeyValue())

	log.Fatalf("Listener on port ", port)

	if err := s.Serve(l); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
