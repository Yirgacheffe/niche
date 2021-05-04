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

	done := make(chan bool)
	defer close(done)

	for num := range takeN(done, repeat(done, 2), 10) {
		fmt.Printf("%#v ", num)
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

	fmt.Println()
	fmt.Println("I am a break line.......")

	terminal := make(chan bool)
	defer close(terminal)

	rand.Seed(time.Now().Unix())

	rand := func() interface{} { return rand.Int() }
	for num := range takeN(terminal, repeatFn(terminal, rand), 10) {
		fmt.Println(num)
	}

}
