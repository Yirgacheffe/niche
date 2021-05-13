package main

import "net/http"

type Client struct {
	*http.Client
	Resp chan *http.Response
	Errs chan error
}

func NewClient(client *http.Client, bufSize int) *Client {

	respCh := make(chan *http.Response, bufSize)
	errsCh := make(chan error, bufSize)

	return &Client{
		client,
		respCh,
		errsCh,
	}

}

func (c *Client) GetInAsync(url string) {
	resp, err := c.Get(url)

	if err != nil {
		c.Errs <- err
		return
	}
	c.Resp <- resp // put response into channel
}
