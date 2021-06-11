package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Func func(s string) (interface{}, error)

type Memo struct{ requests chan request }

func New(f Func) *Memo {
	memo := &Memo{requests: make(chan request)}
	go memo.server(f)
	return memo
}

func (m *Memo) Get(key string) (interface{}, error) {
	response := make(chan result)
	m.requests <- request{key, response}
	res := <-response
	return res.value, res.err
}

func (m *Memo) Close() { close(m.requests) }

func (m *Memo) server(f Func) {
	cache := make(map[string]*entry)

	for req := range m.requests {
		e := cache[req.key]
		if e == nil {
			e = &entry{ready: make(chan struct{})}
			cache[req.key] = e
			go e.call(f, req.key)
		}

		go e.deliver(req.response)
	}
}

type result struct {
	value interface{}
	err   error
}

type entry struct {
	res   result
	ready chan struct{}
}

func (e *entry) call(f Func, key string) {
	e.res.value, e.res.err = f(key)
	close(e.ready)
}

func (e *entry) deliver(response chan<- result) {
	<-e.ready
	response <- e.res
}

type request struct {
	key      string
	response chan<- result
}

func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func incomingUrls() []string {
	return []string{"https://golang.org", "https://godoc.org", "https://play.golang.org"}
}

func main() {
	m := New(httpGetBody)
	for _, url := range incomingUrls() {
		start := time.Now()
		v, err := m.Get(url)
		if err != nil {
			log.Println(err)
		}

		fmt.Printf("%s, %s, %d bytes\n", url, time.Since(start), len(v.([]byte)))
	}
}
