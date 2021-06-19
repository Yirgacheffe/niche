package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"

	"log"
)

// double 'd' make the alignment of variables ... :-)
var readdStream = make(chan int)
var writeStream = make(chan int)

func set(v int) {
	writeStream <- v
}
func read() int {
	return <-readdStream
}

func monitor() {
	var value int // is this the best way or...?
	for {
		select {
		case v := <-writeStream:
			value = v
			fmt.Printf("%d ", value)
		case readdStream <- value:
		}
	}
}

func main() {

	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ltime | log.LUTC)

	if len(os.Args) != 2 {
		log.Println("Please input an integer!")
		return
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Println(err)
		return
	}

	go monitor()

	log.Printf("Going to create %d random numbers.\n", n)
	rand.Seed(time.Now().Unix())

	var wg sync.WaitGroup

	for i := n; i > 0; i-- {
		wg.Add(1)
		go func() {
			defer wg.Done()
			set(rand.Intn(10 * n))
		}()
	}

	wg.Wait()
	fmt.Println()
	log.Printf("Last value: %d\n", read())

}
