package main

import (
	"fmt"
	"net/http"
)

func main() {
	c := NewController()
	http.HandleFunc("/", c.Process)

	fmt.Println("Listening on port :3310")
	if err := http.ListenAndServe(":3310", nil); err != nil {
		panic(err)
	}

}
