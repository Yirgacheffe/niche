package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func produce(data chan int, wg *sync.WaitGroup) {
	n := rand.Intn(999)
	data <- n

	wg.Done()
}

func consume(data chan int, done chan bool) {
	// WIP ...
}

func main() {

	data := make(chan int)
	done := make(chan bool)

	wg := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go produce(data, &wg)
	}

	go consume(data, done)

	go func() {
		wg.Wait()
		close(data)
	}()

	d := <-done
	if d == true {
		fmt.Println("File written ok.")
	} else {
		fmt.Println("File written failed.")
	}

}
