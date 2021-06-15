package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func onemain() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Hello, QCon!")
	})

	go http.ListenAndServe("127.0.0.1:8001", http.DefaultServeMux) // debug
	http.ListenAndServe("0.0.0.0:8080", mux)                       // app traffic
}

//
func serveApp1() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Hello, QCon!")
	})
	http.ListenAndServe("0.0.0.0:8080", mux)
}

func serveDebug1() {
	http.ListenAndServe("127.0.0.1:8001", http.DefaultServeMux)
}

func twomain() {
	go serveDebug1()
	serveApp1()
}

//
func serveApp2() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Hello, QCon!")
	})
	if err := http.ListenAndServe("0.0.0.0:8080", mux); err != nil {
		log.Fatal(err)
	}
}

func serveDebug2() {
	if err := http.ListenAndServe("127.0.0.1:8001", http.DefaultServeMux); err != nil {
		log.Fatal(err)
	}
}

func threemain() {
	go serveDebug2()
	go serveApp2()
	select {}
}

// Only use log.Fatal from main.main or init functions.

func serveApp3() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Hello, QCon!")
	})
	return http.ListenAndServe("0.0.0.0:8080", mux)
}

func serveDebug3() error {
	return http.ListenAndServe("127.0.0.1:8001", http.DefaultServeMux)
}

func fourmain() {
	done := make(chan error, 2)
	go func() {
		done <- serveDebug3()
	}()
	go func() {
		done <- serveApp3()
	}()

	for i := 0; i < cap(done); i++ {
		if err := <-done; err != nil {
			fmt.Printf("error: %v\n", err)
		}
	}
}

//
func serve(addr string, handler http.Handler, stop <-chan struct{}) error {
	s := http.Server{
		Addr:    addr,
		Handler: handler,
	}

	go func() {
		<-stop // wait for stop signal
		s.Shutdown(context.Background())
	}()

	return s.ListenAndServe()
}

func serveApp(stop <-chan struct{}) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Hello, QCon!")
	})
	return serve("0.0.0.0:8080", mux, stop)
}

func serveDebug(stop <-chan struct{}) error {
	return serve("127.0.0.1:8001", http.DefaultServeMux, stop)
}

func main() {
	done := make(chan error, 2)
	stop := make(chan struct{})
	go func() {
		done <- serveDebug(stop)
	}()
	go func() {
		done <- serveApp(stop)
	}()

	var stopped bool
	for i := 0; i < cap(done); i++ {
		if err := <-done; err != nil {
			fmt.Printf("error: %v\n", err)
		}
		if !stopped {
			stopped = true
			close(stop)
		}
	}
}
