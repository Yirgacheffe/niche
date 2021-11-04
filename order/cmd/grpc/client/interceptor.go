package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
)

func orderClientUnaryInterceptor()  {}
func orderClientStreamInterceptor() {}

//
func WithClientInterceptor() grpc.DialOption {
	return grpc.WithUnaryInterceptor(clientInterceptor)
}

func clientInterceptor(ctx context.Context, method string, req interface{}, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	start := time.Now()
	err := invoker(ctx, method, req, reply, cc, opts...)

	log.Printf("invoke remote method=%s duration=%s error=%v", method, time.Since(start), err)
	return err
}
