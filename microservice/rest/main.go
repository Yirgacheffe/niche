package main

import (
	"fmt"
	"net/http"
)

func main() {

	storage := MemStorage{}
	c := NewController(&storage)

	http.HandleFunc("/get", c.GetValue(false))
	http.HandleFunc("/get/default", c.GetValue(true))
	http.HandleFunc("/set", c.SetValue)

	fmt.Println("Listening on port :3333")
	if err := http.ListenAndServe(":3333", nil); err != nil {
		panic(err)
	}

}
