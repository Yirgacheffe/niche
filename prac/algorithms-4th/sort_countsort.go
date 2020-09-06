package main

import "fmt"

func sort(a []int) []int {

	// make the counter from max range
	counter := make([]int, 52)
	for i := range counter {
		counter[i] = 0
	}

	for _, e := range a {
		counter[e]++
	}

	// Add pre valud to current!!!
	for i := 1; i < len(counter); i++ {
		counter[i] += counter[i-1]
	}

	// move the correct position
	res := make([]int, len(a))
	for i := 0; i < len(a); i++ {
		e := a[i]
		t := counter[e] - 1

		res[t] = e
		counter[e] = counter[e] - 1
	}

	return res

}

func main() {

	a := []int{5, 1, 9, 51, 9, 2, 0}

	fmt.Println(a)
	b := sort(a)
	fmt.Println(b)

}
