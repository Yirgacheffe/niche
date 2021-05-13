package main

import (
	"fmt"
	"net/http"
)

// Controller embeds http.Client
type Controller struct {
	*http.Client
}

// DoOps with controller object
func (c *Controller) DoOps() error {

	resp, err := c.Client.Get("http://www.google.com")
	if err != nil {
		return err
	}

	fmt.Println("results of ClientOps:", resp.StatusCode)
	return nil

}
