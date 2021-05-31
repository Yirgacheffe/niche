package main

import (
	"net/http"
)

func main() {
	p := &Proxy{
		Client: http.DefaultClient, BaseURL: "https://httpbin.org/",
	}

	http.Handle("/", p)

	if err := http.ListenAndServe(":9910", nil); err != nil {
		panic(err)
	}

}
