package main

import "fmt"

func sendData(sendch chan<- int) {
	sendch <- 10
}

func main() {
	chnl := make(chan int)
	go sendData(chnl)

	fmt.Println(<-chnl)
	// chnl := make(chan<- int)
}
