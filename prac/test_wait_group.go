package main

import (
	"fmt"
	"sync"
	"time"
)

func process(i int, wg *sync.WaitGroup) {
	fmt.Println("Start routine: ", i)
	time.Sleep(2 * time.Second)
	fmt.Println("Go routine end: ", i)
	wg.Done()
}

func main() {

	nbr := 3
	var wg sync.WaitGroup
	for i := 0; i < nbr; i++ {
		wg.Add(1)
		go process(i, &wg)
	}
	wg.Wait()
	fmt.Println("All go routines finished")
}
