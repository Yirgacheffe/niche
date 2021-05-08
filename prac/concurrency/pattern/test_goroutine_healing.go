package main

import (
	"log"
	"os"
	"time"
)

type startGoroutineFn func(done <-chan bool, pulseInterval time.Duration) (heartbeat <-chan interface{})

func main() {

	newSteward := func(
		timeout time.Duration, startGoroutine startGoroutineFn,
	) startGoroutineFn {
		return func(done <-chan bool, pulseInterval time.Duration) <-chan interface{} {
			heartbeat := make(chan interface{})
			go func() {
				defer close(heartbeat)

				var wardDone chan bool
				var wardHeartbeat <-chan interface{}
				startWard := func() {
					wardDone = make(chan bool)
					wardHeartbeat = startGoroutine(or(wardDone, done), timeout/2) // or is pipeline function, ignore it right now
				}

				startWard()
				pulse := time.Tick(pulseInterval)

			monitorLoop:
				for {
					timeoutSignal := time.After(timeout)
					for {
						select {
						case <-pulse:
							select {
							case heartbeat <- struct{}{}:
							default:
							}
						case <-wardHeartbeat:
							continue monitorLoop
						case <-timeoutSignal:
							log.Println("steward: ward unhealthy; restarting")
							close(wardDone)
							startWard()
							continue monitorLoop
						case <-done:
							return
						}
					}

				}
			}()

			return heartbeat
		}
	}

	// start the goroutine healing testing...
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ltime | log.LUTC)

	// go work is type of startGoroutineFn ...
	doWork := func(done <-chan bool, _ time.Duration) <-chan interface{} {
		log.Println("ward: Hello, I am irresponsible!")
		go func() {
			<-done
			log.Println("ward: I am halting.")
		}()
		return nil
	}

	doWorkWithSteward := newSteward(4*time.Second, doWork)
	done := make(chan bool)

	time.AfterFunc(
		9*time.Second,
		func() {
			log.Println("main: halting steward and ward.")
			close(done)
		},
	)

	for range doWorkWithSteward(done, 4*time.Second) {
	}
	log.Println("Done")
}
