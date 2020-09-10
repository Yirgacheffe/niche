package main

import (
	"fmt"
)

// find max length which sum is 0
func find(a []int) {

	var maxlen int
	n := len(a)

	for i := 0; i < n; i++ {
		sum := 0

		for j := i; j < n; j++ {
			sum += a[j]
			if sum == 0 {
				maxlen = max(maxlen, j-i+1)
			}
		}

	}

	fmt.Println(maxlen)

}

// Using map store the sum - index pair
func findWithHashMap(a []int) {

	hmap := make(map[int]int)
	sum := 0
	maxlen := 0

	for i, v := range a {
		sum += v

		// ??? for all zero cases ???
		if v == 0 && maxlen == 0 {
			maxlen = 1
		}

		if sum == 0 {
			maxlen = i + 1
		}

		if _, ok := hmap[sum]; ok {
			maxlen = max(maxlen, i-hmap[sum])
		} else {
			hmap[sum] = i
		}

	}

	fmt.Println(hmap)
	fmt.Println(maxlen)

}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	a := []int{15, -2, 2, -8, 1, 7, 10, 23, 3}
	find(a)
	findWithHashMap(a)

	b := []int{15, -2, 2, -8, 1, 7, 0, 10, 23, 3}
	find(b)
	findWithHashMap(b)

	allZero := []int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	findWithHashMap(allZero)

}
