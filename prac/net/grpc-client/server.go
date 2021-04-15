package main

import (
	"fmt"
	"net"

	"./greeter"
	"google.golang.org/grpc"
)

func main() {

	grpcServer := grpc.NewServer()
	greeter.RegisterGreeterServiceServer(grpcServer, &Greeter{Exclaim: true})

	listener, err := net.Listen("tcp", ":4410")
	if err != nil {
		panic(err)
	}

	fmt.Println("Listening on port :4410")
	grpcServer.Serve(listener)

}
