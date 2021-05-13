package main

import (
	"fmt"
	"net/http"
)

// DoOps depends a client then fetch 'google'
func DoOps(c *http.Client) error {

	resp, err := c.Get("http://www.google.com")
	if err != nil {
		return err
	}

	fmt.Println("results of DoOps:", resp.StatusCode)
	return nil

}

// DefaultGetGolang use default client
// to get golang.org
func DefaultGetGolang() error {

	resp, err := http.Get("https://golang.org")
	if err != nil {
		return err
	}

	fmt.Println("results of DefaultGetGolang:", resp.StatusCode)
	return nil

}
