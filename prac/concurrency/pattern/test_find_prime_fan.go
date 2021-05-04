package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

// Not finished, also need takeN function
func main() {

	fanIn := func(done <-chan bool, channels ...<-chan int) <-chan int {

		var wg sync.WaitGroup
		multiplexedStream := make(chan int)

		multiplex := func(c <-chan int) {
			defer wg.Done()

			for v := range c {
				select {
				case <-done:
					return
				case multiplexedStream <- v:
				}
			}
		}

		wg.Add(len(channels))

		// call several goroutine to merge the channel value
		for _, c := range channels {
			go multiplex(c)
		}

		// ....... shall we just wait and close the channel?
		go func() {
			wg.Wait()
			close(multiplexedStream)
		}()

		return multiplexedStream
	}

	primeFinder := func(done <-chan bool, valueStream <-chan int) <-chan int {
		primeStream := make(chan int)
		go func() {
			defer close(primeStream)
			for v := range valueStream {
				select {
				case <-done:
					return
				case primeStream <- v:
					// find primes algorithem put here ...
				}
			}
		}()
		return primeStream
	}

	repeatFn := func(done <-chan bool, fn func() int) <-chan int {
		valueStream := make(chan int)
		go func() {
			defer close(valueStream)
			for {
				select {
				case <-done:
					return
				case valueStream <- fn():
				}
			}
		}()
		return valueStream
	}

	// #1 General finders
	randFn := func() int { return rand.Intn(50000000) }

	done := make(chan bool)
	defer close(done)

	start := time.Now()

	randIntStream := repeatFn(done, randFn) // no 'toInt' function as we already have it
	fmt.Println("Primes:")
	for prime := range primeFinder(done, randIntStream) {
		fmt.Printf("\t%v\n", prime)
	}

	fmt.Printf("Search took: %v\n", time.Since(start))

	// #2 fanin-out finders
	randFns := func() int { return rand.Intn(50000000) }

	dones := make(chan bool)
	defer close(dones)

	starts := time.Now()
	randIntStreams := repeatFn(dones, randFns) // no 'toInt' function as we already have it

	// find cpu numbers
	numFinders := runtime.NumCPU()
	fmt.Printf("Spinning up %d prime finders.\n", numFinders)

	primeChs := make([]<-chan int, numFinders)

	fmt.Println("Primes:")
	for i := 0; i < numFinders; i++ {
		primeChs[i] = primeFinder(done, randIntStreams)
	}

	// take the prime out
	for prime := range fanIn(done, primeChs...) {
		fmt.Printf("\t%d\n", prime)
	}

	fmt.Printf("Search took: %v\n", time.Since(starts))

}
