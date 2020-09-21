package main

import "fmt"

func replaceZero(a []int) []int {
	tmp := []int{}
	// Initial the temp array
	for _, v := range a {
		if v == 0 {
			tmp = append(tmp, -1)
		} else {
			tmp = append(tmp, v)
		}
	}
	return tmp
}

func bruteForce(a []int) {

	startIdx := 0

	n := len(a)
	maxlen := 0

	for i := 0; i < n; i++ {
		sum := a[i]

		for j := i + 1; j < n; j++ {
			sum += a[j]

			if sum == 0 {
				tmpLen := j - i + 1
				if maxlen <= tmpLen {
					startIdx = i
					maxlen = tmpLen
				}
			}
		}
	}

	fmt.Printf("Max lenth is: %d\n", maxlen)
	fmt.Printf("Index is: %d -> %d\n", startIdx, maxlen)

}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// // Using map store the sum - index pair
func betterForce(a []int) {

	hmap := make(map[int]int)
	sum := 0
	maxlen := 0
	endIdx := 0

	for i, v := range a {
		sum += v

		// Handle sum == 0 at last index
		if sum == 0 {
			maxlen = i + 1
			endIdx = i
		}

		if _, ok := hmap[sum]; ok {

			if i-hmap[sum] > maxlen {
				maxlen = i - hmap[sum]
				endIdx = i
			}

		} else {
			hmap[sum] = i
		}
	}

	startIdx := endIdx - maxlen + 1

	fmt.Println(startIdx)
	fmt.Println(endIdx)

}

func main() {
	a := []int{1, 0, 1, 1, 1, 0, 0}
	tmp := replaceZero(a)

	bruteForce(tmp)
	betterForce(tmp)
}
