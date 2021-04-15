package main

import (
	"fmt"
	"log"
	"net"

	"./chat"
	"google.golang.org/grpc"
)

func main() {

	fmt.Println("Go gRPC beginner turorial!")

	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := chat.Server{}
	server := grpc.NewServer()

	chat.RegisterChatServiceServer(server, &s)

	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to server: %s", err)
	}

}
