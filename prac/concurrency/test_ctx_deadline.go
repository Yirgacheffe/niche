package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

var startTime = time.Now()

func worker(ctx context.Context, seconds int) {

	select {
	case <-ctx.Done():
		fmt.Printf("%0.2fs - worker(%d) killed !!!\n", time.Since(startTime).Seconds(), seconds)
		return
	case <-time.After(time.Duration(seconds) * time.Second):
		// This case is weird, default case to run ???
		fmt.Printf("%0.2fs - worker(%d) completed.\n", time.Since(startTime).Seconds(), seconds)
	}

}

func main() {

	deadline := time.Now().Add(3 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)

	// mannually cancel if main returned before deadline
	// let's say if return in 1 second
	defer cancel()

	go worker(ctx, 2)

	// These 2 will be killed as deadline is 3 seconds
	go worker(ctx, 6)
	go worker(ctx, 8)

	time.Sleep(5 * time.Second)
	fmt.Println("Number of goroutine", runtime.NumGoroutine())

}
