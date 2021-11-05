package main

import (
	"context"

	trace "cloud.google.com/go/trace/apiv1"
	"google.golang.org/grpc"
)

var TraceClient *trace.Client

func init() {
	var err error
	TraceClient, err := trace.NewClient(context.Background(), "order-trace-id")
	if err != nil {
		panic(err)
	}
}

func WithTracing() grpc.DialOption {
	return grpc.WithUnaryInterceptor(TraceClient.GRPCClientInterceptor())
}

func Tracing() grpc.ServerOption {
	return grpc.WithUnaryInterceptor(TraceClient.GRPCServerInterceptor())
}

// server.go
func runServer() error {
	var tlsCreds
	conn, err := grpc.Dial("localhost:5052", grpc.WithTransportCredentials(tlsCreds),
		WithClientInterceptor(),
		WithTracing())

	srv := grpc.NewServer(grpc.Creds(tlsCreds), ServerInterceptor(), Tracing())

}

/*

message Error {
	int64  code          = 1;
	string message       = 2;
	bool   temporary     = 3;
	int64  userErrorCode = 4;
}

service Cache {
	rpc Dump(DumpReq) returns (DumpResp) {}
}

message DumpReq {

}

message DumpResp {
	repeated DumpItem items = 1;
}

message DumpItem {
	string key = 1;
	bytes  val = 2;
}

*/