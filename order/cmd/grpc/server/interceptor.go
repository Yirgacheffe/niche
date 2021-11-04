package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
)

// demo of:: Unary Interceptor
func orderUnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Println("[Server Interceptor]", info.FullMethod)
	log.Printf("Pre  Proc Message : %s", req)

	m, err := handler(ctx, req)

	log.Printf("Post Proc Message : %s", m)
	return m, err
}

// demo of:: Server Stream interceptor
func orderServerStreamInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	log.Println("[Server Stream Interceptor] ", info.FullMethod)

	err := handler(srv, newWrappedStream(ss))
	if err != nil {
		log.Printf("RPC failed with error %v", err)
	}

	return err
}

//
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
