package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"

	ec "order/internal/grpc/ecommerce"
	pb "order/internal/grpc/impl"
	sv "order/internal/grpc/service"
)

var orderMap = make(map[string]ec.Order)

// demo of:: Unary Interceptor
func orderUnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

	log.Println("[Server Interceptor]", info.FullMethod)
	log.Printf("Pre  Proc Message : %s", req)

	m, err := handler(ctx, req)

	log.Printf("Post Proc Message : %s", m)
	return m, err
}

// demo of:: Server Stream interceptor
type wrappedStream struct {
	grpc.ServerStream
}

func newWrappedStream(s grpc.ServerStream) grpc.ServerStream {
	return &wrappedStream{s}
}

func (w *wrappedStream) RecvMsg(m interface{}) error {
	log.Printf("[Server Stream Interceptor Wrapper]"+"Receive a message (Type: %T) at %s", m, time.Now().Format(time.RFC3339))
	return w.ServerStream.RecvMsg(m)
}

func (w *wrappedStream) SendMsg(m interface{}) error {
	log.Printf("[Server Stream Interceptor Wrapper]"+"Sending a message (Type: %T) at %v", m, time.Now().Format(time.RFC3339))
	return w.ServerStream.SendMsg(m)
}

func orderServerStreamInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	log.Println("[Server Stream Interceptor] ", info.FullMethod)

	err := handler(srv, newWrappedStream(ss))
	if err != nil {
		log.Printf("RPC failed with error %v", err)
	}

	return err
}

func main() {

	var port int

	flag.IntVar(&port, "p", 50051, "Port No. default 50051")
	flag.Parse()

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	initSampleData()

	// srv := grpc.NewServer()
	srv := grpc.NewServer(
		grpc.UnaryInterceptor(orderUnaryServerInterceptor), grpc.StreamInterceptor(orderServerStreamInterceptor))

	sv.RegisterProductServiceServer(srv, pb.NewProductServer())
	ec.RegisterOrderManagementServer(srv, pb.NewOrderManagementServer(orderMap))

	log.Printf("gRPC listener on: %d", port)
	if err := srv.Serve(listener); err != nil {
		log.Fatalf("failed to serve gRPC: %v", err) // ....
	}

}

func initSampleData() {
	orderMap["101"] = ec.Order{Id: "101", Items: []string{"iPhone13", "Yamazki"}, Destination: "Hua Yuan No.350", Price: 2550.00}
	orderMap["102"] = ec.Order{Id: "102", Items: []string{"Apple Watch S7"}, Destination: "San Jose, CA", Price: 1850.90}
	orderMap["103"] = ec.Order{Id: "103", Items: []string{"Google Home Mini", "Google Nest Hub"}, Destination: "HeXi District", Price: 2000.00}
	orderMap["104"] = ec.Order{Id: "104", Items: []string{"Amazon Kindle"}, Destination: "San Jose, UK", Price: 199.00}
	orderMap["105"] = ec.Order{Id: "105", Items: []string{"Apple Mac Mini", "Apple iPhone 12"}, Destination: "Mountain View, CA", Price: 300.00}
}
