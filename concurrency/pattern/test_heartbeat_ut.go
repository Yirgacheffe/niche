package main

import (
	"fmt"
	"time"
)

func doWork(done <-chan bool, nums ...int) (<-chan interface{}, <-chan int) {

	heartbeat := make(chan interface{}, 1)
	results := make(chan int)

	go func() {
		defer close(heartbeat)
		defer close(results)

		time.Sleep(2 * time.Second) // simulate time cost work

		for _, n := range nums {
			select {
			case heartbeat <- struct{}{}:
			default:
				// ......
			}

			select {
			case <-done:
				return
			case results <- n:
			}
		}
	}()

	return heartbeat, results // ..........
}

func main() {

	done := make(chan bool)
	defer close(done)

	intSlice := []int{0, 1, 2, 3, 5, 6}
	heartbeat, results := doWork(done, intSlice...)

	<-heartbeat // Get a heartbeat avoid timeout setting

	i := 0
	for r := range results {
		if expected := intSlice[i]; r != expected {
			fmt.Printf("index %v: expected %v, but received %v\n", i, expected, r)
		}
		i++
	}

	/*
		for {
			select {
			case v, ok := <-results:
				if !ok {
					return
				}
				fmt.Printf("%v\n", v)
			case <-heartbeat:
				fmt.Println("pulse!")
			}
		}
	*/

}
