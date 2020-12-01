package main

import (
	"fmt"
	"runtime"
)

func square(ch chan int) {
	for i := 0; i < 4; i++ {
		num := <-ch
		fmt.Println(num * num)
	}
}

func main() {

	fmt.Println("---- main started ----")

	ch := make(chan int, 3)

	go square(ch)
	fmt.Println("active goroutines:", runtime.NumGoroutine())

	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4 // block here

	fmt.Println("active goroutines:", runtime.NumGoroutine())

	go square(ch)
	fmt.Println("active goroutines:", runtime.NumGoroutine())

	ch <- 5
	ch <- 6
	ch <- 7
	ch <- 8 // block here

	fmt.Println("active goroutines:", runtime.NumGoroutine())

	fmt.Println("---- main finished ---")

}
