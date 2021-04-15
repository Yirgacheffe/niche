package main

import (
	"bitbucket-repos/internal/grpc/impl"
	"bitbucket-repos/internal/grpc/service"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {

	netListener := getNetListener(7000)
	grpcServ := grpc.NewServer()

	repoServiceImpl := impl.NewRepositoryServiceGrpcImpl()
	service.RegisterRepositoryServiceServer(grpcServ, repoServiceImpl)

	if err := grpcServ.Serve(netListener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}

func getNetListener(port uint) net.Listener {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		panic(fmt.Sprintf("failed to listen: %v", err))
	}

	return listener
}
