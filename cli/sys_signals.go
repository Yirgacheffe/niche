package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// CatchSignal sets up a listener for
// SIGINT interrupts
func CatchSignal(ch chan os.Signal, done chan bool) {

	sig := <-ch
	fmt.Println("signal received:", sig)

	switch sig {
	case syscall.SIGINT:
		fmt.Println("handling a SIGINT now!")
	case syscall.SIGTERM:
		fmt.Println("handling a SIGTERM in a entirely different way!")
	default:
		fmt.Println("unexpected signal received!")
	}

	done <- true

}

func main() {

	signals := make(chan os.Signal)
	done := make(chan bool)

	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	go CatchSignal(signals, done)

	fmt.Println("press ctrl-c to terminated...")
	<-done
	fmt.Println("Done!")

}
