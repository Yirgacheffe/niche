package main

import "fmt"

func main() {
	ch := make(chan string, 2)
	ch <- "Hwool"
	ch <- "moon"
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
