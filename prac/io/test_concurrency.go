package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
)

func produce(data chan int, wg *sync.WaitGroup) {
	n := rand.Intn(999)
	data <- n

	wg.Done()
}

func consume(data chan int, done chan bool) {

	f, err := os.Create("test_concurrency_write.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()

	for d := range data {
		_, err = fmt.Fprintln(f, d)
		if err != nil {
			fmt.Println(err)
			done <- false
		}
	}

	fmt.Println("Write file concurrency successfully!")
	done <- true
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
