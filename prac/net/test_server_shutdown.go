package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(2)

	server := &http.Server{
		Addr:    ":9092",
		Handler: nil,
	}

	// Add shutdown function ......
	server.RegisterOnShutdown(
		func() {
			fmt.Println("RegisterOnShutdown(): Complete!")
			wg.Done()
		},
	)

	// Time After 3 seconds
	time.AfterFunc(
		3*time.Second,
		func() {
			err := server.Shutdown(context.Background())
			fmt.Println("server.Shutdown(): Complete!", err)
			wg.Done()
		},
	)

	fmt.Println("ListenAndServe():", server.ListenAndServe())

	wg.Wait()
	fmt.Println("Main(): Exit complete!")

}
