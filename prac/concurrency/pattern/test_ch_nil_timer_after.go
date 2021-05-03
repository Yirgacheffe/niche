package main

import (
	"fmt"
	"math/rand"
	"time"
)

func add(c chan int) {
	sum := 0
	t := time.NewTimer(time.Second)

	for {
		select {
		case num := <-c:
			sum = sum + num
		case <-t.C:
			c = nil
			fmt.Println(sum)
		}
	}
}

func emit(c chan int) {
	for {
		c <- rand.Intn(10)
	}
}

func main() {
	c := make(chan int)
	go add(c)
	go emit(c)

	time.Sleep(3 * time.Second)
}
