package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", NegHandler)

	fmt.Println("Listening on port :3310")
	err := http.ListenAndServe(":3310", nil)
	if err != nil {
		panic(err)
	}
}
