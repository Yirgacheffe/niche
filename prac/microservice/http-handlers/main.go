package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/name", HelloHandler)
	http.HandleFunc("/greeting", GreetingHandler)

	fmt.Println("Listening on port :3333")

	err := http.ListenAndServe(":3333", nil)
	panic(err)

}
