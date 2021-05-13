package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func panicHandler(w http.ResponseWriter, r *http.Request) {
	panic("oops")
}

func main() {
	h := ApplyMiddleware(Handler, Logger(log.New(os.Stdout, "", 0)), SetID(200))

	http.HandleFunc("/", h)
	http.HandleFunc("/panic", panicHandler)
	fmt.Println("Listening on port :3310")

	if err := http.ListenAndServe(":3310", nil); err != nil {
		panic(err)
	}

}
