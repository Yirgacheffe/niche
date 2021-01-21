package main

import (
	"fmt"
	"net/http"
)

func main() {

	urls := []string{
		"http://httpbin.org",
		"http://www.github.com",
		"http://golang.org",
	}

	c := NewClient(http.DefaultClient, len(urls))
	FetchAll(urls, c)

	for i := 0; i < len(urls); i++ {
		select {
		case resp := <-c.Resp:
			fmt.Printf("Status received for %s: %d\n", resp.Request.URL, resp.StatusCode)
		case err := <-c.Errs:
			fmt.Printf("Error received: %s\n", err)
		}
	}

}
