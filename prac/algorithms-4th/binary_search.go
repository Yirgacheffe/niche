package main

import "fmt"

func binarySearch(key int, a []int) int {

	// Slices must be sorted
	lo := 0
	hi := len(a) - 1

	for lo <= hi {
		mid := lo + (hi-lo)/2
		if key < a[mid] {
			hi = mid - 1
		} else if key > a[mid] {
			lo = mid + 1
		} else {
			return mid
		}
	}

	return -1

}

func main() {

	a := []int{1, 5, 7, 14, 20, 25, 32, 35, 40, 44, 45, 48, 50, 56}
	idx := binarySearch(48, a)

	fmt.Printf("The index is: %d\n", idx)

}
