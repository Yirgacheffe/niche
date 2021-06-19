package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {

	doWork := func(done <-chan bool, id int, wg *sync.WaitGroup, result chan<- int) {
		start := time.Now()
		defer wg.Done()

		// simulate the work load
		loadedTime := time.Duration(1+rand.Intn(10)) * time.Second

		select {
		case <-done:
			return
		case <-time.After(loadedTime): // simulate work
		}
		select {
		case <-done:
			return
		case result <- id:
		}

		took := time.Since(start)

		if took < loadedTime {
			took = loadedTime
		}

		fmt.Printf("%v took %v\n", id, took)
	}

	// testing start
	done := make(chan bool)
	result := make(chan int)

	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go doWork(done, i, &wg, result)
	}

	firstReturned := <-result

	close(done)
	wg.Wait()
	fmt.Printf("Received an answer from #%v\n", firstReturned)

}
