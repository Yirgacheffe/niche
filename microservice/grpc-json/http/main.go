package main

import (
	"grpc-json/internal"
	"net/http"
)

func main() {
	c := Controller{KeyValue: internal.NewKeyValue()}

	http.HandleFunc("/set", c.SetHandler)
	http.HandleFunc("/get", c.GetHandler)

	if err := http.ListenAndServe(":9910", nil); err != nil {
		panic(err)
	}
}
