package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

type Func func(s string) (interface{}, error)

type Memo struct {
	f     Func
	cache map[string]result
	mu    sync.Mutex
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]result)}
}

func (m *Memo) Get(key string) (interface{}, error) {
	m.mu.Lock()
	res, ok := m.cache[key]
	if !ok {
		res.value, res.err = m.f(key)
		m.cache[key] = res
	}
	m.mu.Unlock()

	return res.value, res.err
}

// Better than 'Get' version
func (m *Memo) Get1(key string) (interface{}, error) {
	m.mu.Lock()
	res, ok := m.cache[key]
	m.mu.Unlock()

	if !ok {
		res.value, res.err = m.f(key)

		m.mu.Lock()
		m.cache[key] = res
		m.mu.Unlock()
	}

	return res.value, res.err
}

type result struct {
	value interface{}
	err   error
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
