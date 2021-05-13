package main

import (
	"fmt"
	"sync"
)

type rect struct {
	length int
	width  int
}

func (r rect) area(wg *sync.WaitGroup) {
	defer wg.Done()

	if r.length < 0 {
		fmt.Printf("rect %v's length should be greater than zero\n", r)
		return
	}
	if r.width < 0 {
		fmt.Printf("rect %v's witdh should be greater than zero\n", r)
		return
	}

	area := r.length * r.width
	fmt.Printf("rect %v's area is: %d\n", r, area)
}

func main() {
	var wg sync.WaitGroup

	r1 := rect{-68, 38}
	r2 := rect{50, -10}
	r3 := rect{20, 100}

	rects := []rect{r1, r2, r3}
	for _, v := range rects {
		wg.Add(1)
		go v.area(&wg)
	}
	wg.Wait()
	fmt.Println("All goroutine finished...")
}
