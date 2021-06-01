package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

var c = make(chan int)

func square(ctx context.Context) {

	i := 0

	for {
		select {
		case <-ctx.Done():
			fmt.Println("terminated square")
			return
		case c <- i * i:
			i++
		}
	}

}

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	go square(ctx)

	for i := 0; i < 5; i++ {
		fmt.Println("The square is", <-c)
	}

	// Cancel the context
	// Can do it with 'defer context()'??
	cancel()

	time.Sleep(3 * time.Second)
	fmt.Println("Num of goroutines", runtime.NumGoroutine())

}
