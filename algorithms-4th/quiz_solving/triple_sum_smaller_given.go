package main

import "fmt"

func tripleSetSmallerThanGiven(a []int, sum int) {

	counter := 0
	n := len(a)

	for i := 0; i < n-2; i++ {

		t := sum - a[i]
		lo := i + 1
		hi := n - 1

		for lo < hi {

			if a[lo]+a[hi] < t {
				counter += hi - lo
				lo++
			} else {
				hi--
			}
		}
		// increment i++ to n-2
	}

	fmt.Println("Output:", counter)
}

func main() {

	// If the array is not sorted, use any sort method you want
	a := []int{1, 3, 4, 5, 7, 8}
	tripleSetSmallerThanGiven(a, 12)

	b := []int{-2, 0, 1, 3}
	tripleSetSmallerThanGiven(b, 2)

}
