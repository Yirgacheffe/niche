package main

import (
	"log"
	"net/http"
)

func main() {
	Routes()

	log.Println("listener: Started: Listenning on :9901")
	http.ListenAndServe(":9901", nil)
}
