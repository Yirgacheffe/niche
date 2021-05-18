package main

import "fmt"

func countDistinctElementInWindow(a []int, k int) {

	hm := make(map[int]int)
	l := len(a)

	// First window
	for i := 0; i < k; i++ {
		key := a[i]

		if v, ok := hm[key]; !ok {
			hm[key] = 1
		} else {
			hm[key] = v + 1
		}
	}

	fmt.Printf("Distinct count: %d\n", len(hm))

	for i := k; i < l; i++ {

		// This is the import part, get prev index
		prev := a[i-k]

		// Remove prev window
		if v, ok := hm[prev]; ok {
			if v == 1 {
				delete(hm, prev)
			} else {
				hm[prev] = v - 1
			}
		}

		// Current window
		curr := a[i]
		if v, ok := hm[curr]; !ok {
			hm[curr] = 1
		} else {
			hm[curr] = v + 1
		}

		fmt.Printf("Distinct count: %d\n", len(hm))
	}

}

func main() {
	a := []int{1, 2, 1, 3, 4, 2, 3}
	countDistinctElementInWindow(a, 3)
}
