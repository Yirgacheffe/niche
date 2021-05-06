package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now()
	var c1, c2 <-chan int

	select {
	case <-c1: // blocked because c1 and c2 are 'nil'
	case <-c2:
	default:
		fmt.Printf("In default after %v\n\n", time.Since(start))
	}

}
