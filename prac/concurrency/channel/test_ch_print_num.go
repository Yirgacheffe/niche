package main

import (
	"fmt"
	"runtime"
	"time"
)

var signal = make(chan bool)

func printNbrs() {

	counter := 1

	for {
		select {
		case <-signal:
			return
		default:
			time.Sleep(100 * time.Millisecond)
			counter++
		}
	}

}

func main() {

	go printNbrs()

	fmt.Println("Before: goroutines", runtime.NumGoroutine())
	time.Sleep(1 * time.Second)

	signal <- true
	fmt.Println("After:  goroutines", runtime.NumGoroutine())

	fmt.Println("Program exited!!!")

}
