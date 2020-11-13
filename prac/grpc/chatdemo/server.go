package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {

	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to server: %s", err)
	}

}
