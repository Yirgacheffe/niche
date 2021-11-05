package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Unary
func orderClientUnaryInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption) error {

	log.Println("Method: " + method)
	err := invoker(ctx, method, req, reply, cc, opts...)
	log.Println(reply)
	return err
}

type retryBackoff struct{}

func (r retryBackoff) Next() chan time.Time {

	retryCh := make(chan time.Time)
	maxRetries := 3

	go func() {
		pulse := time.Tick(time.Second * 3)

		for i := 0; i < maxRetries; i++ {

			select {
			case t, ok := <-pulse:
				if ok {
					retryCh <- t
				}
			}
		}
	}()

	return retryCh
}

func retryClientInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption) error {

	var (
		start    = time.Now()
		attempts = 0
		err      error
		backoff  retryBackoff // is something to set retry logic
	)

	for {
		attempts += 1
		select {
		case <-ctx.Done():
			err = status.Errorf(codes.DeadlineExceeded, "timeout reached before next retry attempt")
		case <-backoff.Next():
			startAttempt := time.Now()
			err = invoker(ctx, method, req, reply, cc, opts...)
			if err != nil {
				log.Printf("not able to invoker : %v, after %f's, retry\n", err, time.Since(startAttempt).Seconds())
				continue
			}
		}
		break
	}

	log.Printf("finished in %f's\n", time.Since(start).Seconds())
	return err
}

func isIdempotent(ctx context.Context) bool {
	val, ok := ctx.Value("idempotent").(bool)
	if !ok {
		return true
	}

	return val
}

func NotIdempotent(ctx context.Context) context.Context {
	return context.WithValue(ctx, "idempotent", false)
}

// Stream
func orderClientStreamInterceptor(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string,
	streamer grpc.Streamer,
	opts ...grpc.CallOption) (grpc.ClientStream, error) {

	log.Println("[Client Interceptor] ", method)
	s, err := streamer(ctx, desc, cc, method, opts...)
	if err != nil {
		return nil, err
	}

	return newWrappedStream(s), nil
}

type wrappedStream struct {
	grpc.ClientStream
}

func newWrappedStream(s grpc.ClientStream) grpc.ClientStream {
	return &wrappedStream{s}
}

func (w *wrappedStream) SendMsg(m interface{}) error {
	log.Println("[Client Stream Interceptor]"+"Send a message (Type: %T) at %v", m, time.Now().Format(time.RFC3339))
	return w.ClientStream.SendMsg(m)
}
func (w *wrappedStream) RecvMsg(m interface{}) error {
	log.Println("[Client Stream Interceptor]"+"Receive a message (Type: %T) at %v", m, time.Now().Format(time.RFC3339))
	return w.ClientStream.RecvMsg(m)
}

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
