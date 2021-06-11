package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Func func(s string) (interface{}, error)

type Memo struct {
	f     Func
	cache map[string]result
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]result)}
}

// NOTE: not concurrency safe.
func (m *Memo) Get(key string) (interface{}, error) {
	res, ok := m.cache[key]
	if !ok {
		res.value, res.err = m.f(key)
		m.cache[key] = res
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
