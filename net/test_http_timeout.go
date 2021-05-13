package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
)

func main() {

	client := &http.Client{Timeout: 100 * time.Microsecond}
	res, err := client.Get("https://jsonplaceholder.typicode.com/users/1")

	if err != nil {
		// get `url.Error` struct pointer from `err` interface
		urlErr := err.(*url.Error)

		// check if error occured due to timeout
		if urlErr.Timeout() {
			fmt.Println("Error occured due to a timeout.")
		}

		log.Fatal("Error:", err)
	} else {
		fmt.Println("Success:", res.StatusCode)
	}

}
