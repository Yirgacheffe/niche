package internal

import (
	"context"
	"sync"

	"grpc-json/keyvalue"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	err_invalid_token = status.Errorf(codes.Unauthenticated, "invalid token")
	err_key_not_found = status.Errorf(codes.NotFound, "key not found")
)

type KeyValue struct {
	mutex sync.RWMutex
	m     map[string]string
}

func NewKeyValue() *KeyValue {
	return &KeyValue{
		m: make(map[string]string),
	}
}

func (k *KeyValue) Set(ctx context.Context, r *keyvalue.SetKeyValueRequest) (*keyvalue.KeyValueResponse, error) {
	k.mutex.Lock()
	k.m[r.GetKey()] = r.GetValue()
	k.mutex.Unlock()

	return &keyvalue.KeyValueResponse{Value: r.GetValue()}, nil
}

func (k *KeyValue) Get(ctx context.Context, r *keyvalue.GetKeyValueRequest) (*keyvalue.KeyValueResponse, error) {
	k.mutex.Lock()
	defer k.mutex.RUnlock()

	v, ok := k.m[r.GetKey()]
	if !ok {
		return nil, err_key_not_found
	}

	return &keyvalue.KeyValueResponse{Value: v}, nil
}
