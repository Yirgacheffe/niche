package main

import "fmt"

func main() {
	ch := make(chan string, 3)
	ch <- "Howen"
	ch <- "West wood"

	fmt.Println("capacity is", cap(ch))
	fmt.Println("length is", len(ch))
	fmt.Println("read value", <-ch)
	fmt.Println("new length is", len(ch))
}
