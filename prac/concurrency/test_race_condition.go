package main

import (
	"fmt"
	"sync"
)

var x = 0

func increament(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	x += 1
	m.Unlock()
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	var m sync.Mutex

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go increament(&wg, &m)
	}
	wg.Wait()
	fmt.Println("final value of x", x)
}
