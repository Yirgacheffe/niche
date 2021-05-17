package main

import (
	"fmt"
	"io"
	"net/http"
)

var helloIndexHTML = "<DOCTYPE-TYPE html><html><head><title>Hello World</title></head><body>Hello, World!</body></html>"

func helloHTTPHandler(res http.ResponseWriter, r *http.Request) {

	res.Header().Set("Content-type", "text/html")
	io.WriteString(res, helloIndexHTML)

}

func main() {
	fmt.Println("---------------------------------------------")

	http.HandleFunc("/hello", helloHTTPHandler)
	http.ListenAndServe(":9000", nil)
	http.Handle("/assets/", http.FileServer(http.Dir("assets")))
}
