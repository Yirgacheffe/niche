package main

import "fmt"

func sender(ch chan int) {

	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4

	close(ch)
}

func main() {

	c := make(chan int, 3)
	go sender(c)

	fmt.Printf("len of channel c is: %v cap is: %v\n", len(c), cap(c))

	for v := range c {
		fmt.Printf("read value is: %v len is: %v\n", v, len(c))
	}

}
