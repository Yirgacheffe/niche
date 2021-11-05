package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"

	"golang.org/x/net/netutil"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	ec "order/internal/grpc/ecommerce"
	pb "order/internal/grpc/impl"
	sv "order/internal/grpc/service"
)

var orderMap = make(map[string]ec.Order)

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

func anotherRun() {

	certs, err := tls.LoadX509KeyPair("out/example.com.crt", "out/example.com.key")
	if err != nil {
		log.Fatalf("failed to read client ca cert: %v", err)
	}

	certPool := x509.NewCertPool()
	bs, err := ioutil.ReadFile("out/my_root_ca.crt")
	if err != nil {
		log.Fatalf("failed to read client ca cert: %v", err)
	}

	ok := certPool.AppendCertsFromPEM(bs)
	if !ok {
		log.Fatalf("failed to append client certs")
	}

	tlsConfig := &tls.Config{
		ClientAuth:   tls.RequireAndVerifyClientCert,
		Certificates: []tls.Certificate{certs},
		ClientCAs:    certPool,
	}

	tlsCreds := credentials.NewTLS(tlsConfig)
	srv := grpc.NewServer(grpc.Creds(tlsCreds), grpc.MaxConcurrentStreams(64), grpc.InTapHandle(NewTap().Handler))

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", 28702))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	if err := srv.Serve(netutil.LimitListener(listener, 1024)); err != nil {
		log.Fatalf("failed to serve gRPC: %v", err) // ....
	}
}
