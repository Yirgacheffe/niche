package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	repeat := func(done <-chan bool, values ...interface{}) <-chan interface{} {
		valueStream := make(chan interface{})

		go func() {
			defer close(valueStream)
			for {
				for _, v := range values {
					select {
					case <-done:
						return
					case valueStream <- v:
					}
				}
			}
		}()

		return valueStream
	}

	takeN := func(done <-chan bool, valueStream <-chan interface{}, num int) <-chan interface{} {
		takeStream := make(chan interface{})

		go func() {
			defer close(takeStream)
			for i := 0; i < num; i++ {
				select {
				case <-done:
					return
				case takeStream <- valueStream:
				}
			}
		}()

		return takeStream
	}

	repeatFn := func(done <-chan bool, fn func() interface{}) <-chan interface{} {
		valueStream := make(chan interface{})

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

	toString := func(done <-chan bool, valueStream <-chan interface{}) <-chan string {
		stringStream := make(chan string)

		go func() {
			defer close(stringStream)
			for {
				for v := range valueStream {
					select {
					case <-done:
						return
					case stringStream <- v.(string): // panic, not convert
					}
				}
			}
		}()

		return stringStream
	}

	// # No 1.
	done := make(chan bool)
	defer close(done)

	for num := range takeN(done, repeat(done, 2), 10) {
		fmt.Printf("%#v ", num)
	}

	fmt.Println()
	fmt.Println("I am a break line.......")

	// # No 2.
	terminal := make(chan bool)
	defer close(terminal)

	rand.Seed(time.Now().Unix())

	rand := func() interface{} { return rand.Int() }
	for num := range takeN(terminal, repeatFn(terminal, rand), 10) {
		fmt.Println(num)
	}

	// # No 3. failed due to 'toString' function
	doneCh := make(chan bool)
	defer close(doneCh)

	var message string
	for token := range toString(doneCh, takeN(doneCh, repeat(doneCh, "I", "am."), 5)) {
		message += token
	}

	fmt.Printf("message: %s...", message)

}
