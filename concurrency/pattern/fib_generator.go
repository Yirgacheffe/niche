package main

import "fmt"

func fib(length int) <-chan int {

	ch := make(chan int, length)

	go func() {
		for i, j := 0, 1; i < length; i, j = i+j, i {
			ch <- i
		}
		close(ch)
	}()

	// closed but buffered, read safe
	return ch

}

func main() {
	for v := range fib(10) {
		fmt.Println("current fib value is:", v)
	}
}
