package main

import (
	"fmt"
	"time"
)

func write(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("Successfully wrote", i, "to ch")
	}
	close(ch)
}

func main() {
	ch := make(chan int, 2)
	go write(ch)

	time.Sleep(2 * time.Second)
	for v := range ch {
		fmt.Println("Read value", v, "from ch")
		time.Sleep(2 * time.Second)
	}

	ch1 := make(chan string, 2)
	ch1 <- "Hwool"
	ch1 <- "moon"
	fmt.Println(<-ch1)
	fmt.Println(<-ch1)

	ch2 := make(chan string, 3)
	ch2 <- "Howen"
	ch2 <- "West wood"

	fmt.Println("capacity is", cap(ch2))
	fmt.Println("length is", len(ch2))
	fmt.Println("read value", <-ch2)
	fmt.Println("new length is", len(ch2))
}
