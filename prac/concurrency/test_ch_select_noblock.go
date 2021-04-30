package main

import (
	"fmt"
	"time"
)

var start time.Time

func init() {
	start = time.Now()
}

func main() {

	fmt.Println("main started:", time.Since(start))

	ch1 := make(chan string, 2)
	ch2 := make(chan string, 2)

	ch1 <- "value 1"
	ch1 <- "value 2"
	ch2 <- "value 3"
	ch2 <- "value 4"

	select {
	case v := <-ch1:
		fmt.Println("response time from ch1:", v, time.Since(start))
	case v := <-ch2:
		fmt.Println("response time from ch2:", v, time.Since(start))
	}

	fmt.Println("main finished:", time.Since(start))

}
