package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
)

func badMain() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello Gopher!")
	})

	go func() {
		if err := http.ListenAndServe(":9010", nil); err != nil {
			log.Fatal(err)
		}
	}()

	for {
		// block the main goroutine, doesn't do any IO ...............
	}
}

func alsobadMain() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello Gopher!")
	})

	go func() {
		if err := http.ListenAndServe(":9010", nil); err != nil {
			log.Fatal(err)
		}
	}()

	for {
		runtime.Gosched() // Yields the processor, allowing other goroutine to run
	}
}

func alsobaddMain() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello Gopher!")
	})

	go func() {
		if err := http.ListenAndServe(":9010", nil); err != nil {
			log.Fatal(err)
		}
	}()

	select {
	// Empty select will block forever, not spinning a whole CPU .....
	}
}

// main - Normal
func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(rw, "Hello, Gopher!")
	})

	if err := http.ListenAndServe(":9010", nil); err != nil {
		log.Fatal(err)
	}
}

// if your goroutine cannot make progress until it gets the result from another,
// oftentimes it is simpler to just do the work yourself rather than to delegate it.
