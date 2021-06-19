package main

import (
	"fmt"
	"sync"
	"time"
)

func sqrWorker(wg *sync.WaitGroup, tasks <-chan int, results chan<- int, instance int) {

	for nbr := range tasks {
		// simulate block task, so that each go routine can run the tasks
		time.Sleep(time.Millisecond)
		fmt.Printf("Sending result by worker %v\n", instance)
		results <- nbr * nbr
	}

	wg.Done()
}

func main() {

	fmt.Println("[mian] main started")

	var wg sync.WaitGroup

	tasks := make(chan int, 10)
	results := make(chan int, 10)

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go sqrWorker(&wg, tasks, results, i)
	}

	for i := 0; i < 5; i++ {
		tasks <- i * 2
	}

	fmt.Println("[main] Wrote 5 tasks")
	close(tasks)

	// wait until all tasks get done
	wg.Wait()

	for i := 0; i < 5; i++ {
		result := <-results
		fmt.Println("[main] Result", i, ":", result)
	}

	fmt.Println("[main] main() stopped")

}
