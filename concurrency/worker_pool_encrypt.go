package main

import (
	"context"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type op string

const (
	Hash    op = "encrypt"
	Compare    = "decrypt"
)

type WorkRequest struct {
	Op      op
	Text    []byte
	Compare []byte
}

type WorkResponse struct {
	Wr      WorkRequest
	Result  []byte
	Matched bool
	Err     error
}

func Dispatch(nbrWorker int) (context.CancelFunc, chan WorkRequest, chan WorkResponse) {
	ctx, cancel := context.WithCancel(context.Background())

	in := make(chan WorkRequest, 10)
	out := make(chan WorkResponse, 10)

	for i := 0; i < nbrWorker; i++ {
		go Worker(ctx, i, in, out)
	}

	return cancel, in, out
}

func Worker(ctx context.Context, id int, in chan WorkRequest, out chan WorkResponse) {
	for {
		select {
		case <-ctx.Done():
			return
		case r := <-in:
			fmt.Printf("worker id: %d, performing %s workn", id, r.Op)
			out <- Process(r)
		}
	}
}

func Process(r WorkRequest) WorkResponse {
	switch r.Op {
	case Hash:
		return hashWork(r)
	case Compare:
		return compareWork(r)
	default:
		return WorkResponse{Err: errors.New("unsupported operation!")}
	}
}

func hashWork(r WorkRequest) WorkResponse {
	v, err := bcrypt.GenerateFromPassword(
		r.Text,
		bcrypt.DefaultCost,
	)

	return WorkResponse{Result: v, Err: err, Wr: r}
}

func compareWork(r WorkRequest) WorkResponse {
	var matched bool
	err := bcrypt.CompareHashAndPassword(r.Compare, r.Text)
	if err == nil {
		matched = true
	}

	return WorkResponse{Matched: matched, Err: err, Wr: r}
}

func main() {

	cancel, in, out := Dispatch(10)
	defer cancel()

	for i := 0; i < 10; i++ {
		in <- WorkRequest{Op: Hash, Text: []byte(fmt.Sprintf("messages %d", i))}
	}

	for i := 0; i < 10; i++ {
		res := <-out
		if res.Err != nil {
			panic(res.Err)
		}

		in <- WorkRequest{Op: Compare, Text: res.Wr.Text, Compare: res.Result}
	}

	for i := 0; i < 10; i++ {
		res := <-out
		if res.Err != nil {
			panic(res.Err)
		}

		fmt.Printf("string: %s; matched: %v\n", string(res.Wr.Text), res.Matched)
	}

}
