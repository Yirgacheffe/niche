package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "order/internal/grpc/impl"
	sv "order/internal/grpc/service"
)

func main() {

	var port int

	flag.IntVar(&port, "p", 50051, "Port No. default 50051")
	flag.Parse()

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srv := grpc.NewServer()
	sv.RegisterProductServiceServer(srv, pb.NewProductServer())

	log.Printf("gRPC listener on: %d", port)
	if err := srv.Serve(listener); err != nil {
		log.Fatalf("failed to serve gRPC: %v", err) // ....
	}

}
