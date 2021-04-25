package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	h := ApplyMiddleware(Handler, Logger(log.New(os.Stdout, "", 0)), SetID(200))

	http.HandleFunc("/", h)
	fmt.Println("Listening on port :3310")

	if err := http.ListenAndServe(":3310", nil); err != nil {
		panic(err)
	}

}
