package main

import (
	"context"

	"golang.org/x/time/rate"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/tap"
)

type Tap struct {
	lim *rate.Limiter
}

func NewTap() *Tap {
	return &Tap{rate.NewLimiter(150, 5)}
}

func (t *Tap) Handler(ctx context.Context, info *tap.Info) (context.Context, error) {
	if !t.lim.Allow() {
		return nil,
			status.Errorf(codes.ResourceExhausted, "service is over rate limit")
	}
	return ctx, nil
}
