package main

import (
	"fmt"
	"sync"
)

func main() {

	var numCalsCreated int

	calcPool := &sync.Pool{
		New: func() interface{} {
			numCalsCreated += 1
			mem := make([]byte, 1024)
			return &mem
		},
	}

	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())

	const numWorkers = 1024 * 1024
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	for i := numWorkers; i > 0; i-- {
		go func() {
			defer wg.Done()
			mem := calcPool.Get().(*[]byte)
			defer calcPool.Put(mem)
			// do something with the 'mem'
		}()
	}

	wg.Wait()
	fmt.Printf("%d calculators were created.", numCalsCreated)

}
