package main

import (
	"fmt"
	"time"
)

func process(ch chan string) {
	time.Sleep(10500 * time.Millisecond)
	ch <- "processed"
}

func main() {

	// Case 1
	start := time.Now()
	var c1, c2 <-chan int

	select {
	case <-c1: // blocked because c1 and c2 are 'nil'
	case <-c2:
	default:
		fmt.Printf("In default after %v\n\n", time.Since(start))
	}

	// Case 2
	ch := make(chan string)
	go process(ch)

	for {
		time.Sleep(1000 * time.Millisecond)
		select {
		case v := <-ch:
			fmt.Println("received value: ", v)
			return
		default:
			// default make select non-blocking
			fmt.Println("no value received")
		}
	}

	// default case also works for nil channel

}
