package main

import (
	"fmt"
	"time"
)

func main() {

	doWork := func(done <-chan bool, pluseInterval time.Duration) (<-chan interface{}, <-chan time.Time) {
		heartbeat := make(chan interface{})
		results := make(chan time.Time)

		go func() {
			// defer close(heartbeat)
			// defer close(results)
			workGen, pluse := time.Tick(2*pluseInterval), time.Tick(pluseInterval)
			sendPluse := func() {
				select {
				case heartbeat <- struct{}{}:
				default:
					// ignore one heartbeat signal is acceptable ...
				}
			}
			sendResult := func(r time.Time) {
				for {
					select {
					case <-done:
						return
					case <-pluse: // this will not ignore pluse while work send
						sendPluse()
					case results <- r:
						return
					}
				}
			}

			for i := 0; i < 2; i++ {
				select {
				case <-done:
					return
				case r := <-workGen:
					sendResult(r)
				case <-pluse:
					sendPluse()
				}
			}
		}()

		return heartbeat, results
	}

	// Start testing
	done := make(chan bool)
	time.AfterFunc(10*time.Second, func() { close(done) })

	const timeout = 2 * time.Second
	heartbeat, results := doWork(done, timeout/2)

	for {
		select {
		case _, ok := <-heartbeat:
			if !ok {
				return
			}
			fmt.Println("pulse")
		case r, ok := <-results:
			if !ok {
				return
			}
			fmt.Printf("result %v\n", r.Second())
		case <-time.After(timeout):
			fmt.Println("worker goroutine is not healthy!")
			return
		}
	}

}
