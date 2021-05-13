package main

import (
	"fmt"
	"net/http"
	"sync"
)

func createServer(name string, port int) *http.Server {

	mux := http.NewServeMux()
	handleFunc := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello: "+name)
	}

	mux.HandleFunc("/", handleFunc)
	server := http.Server{
		Addr:    fmt.Sprintf(":%v", port),
		Handler: mux,
	}

	return &server // Return new server (pointer here)

}

func main() {

	// Create and add 2 go routine
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		serv := createServer("server 1", 9091)
		fmt.Println(serv.ListenAndServe())
		wg.Done()
	}()

	go func() {
		serv := createServer("server 2", 9092)
		fmt.Println(serv.ListenAndServe())
		wg.Done()
	}()

	wg.Wait() // Wait until go routing is done
}
