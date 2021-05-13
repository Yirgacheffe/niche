package main

import "fmt"

// 2 pointer method
// array must be sorted
func findPair(a []int, sum int) {

	lo := 0
	hi := len(a) - 1

	for lo < hi {
		if a[lo]+a[hi] == sum {
			fmt.Printf("Found: %d %d\n", a[lo], a[hi])
			return
		} else if a[lo]+a[hi] < sum {
			lo++
		} else {
			hi--
		}
	}

	fmt.Println("Not found!")

}

// Map method, Hash Table, Haseset, bitset
func findPairWithHashing(a []int, sum int) {

	n := len(a)
	m := map[int]bool{}

	for i := 0; i < n; i++ {
		t := sum - a[i]
		_, ok := m[t]
		if ok {
			fmt.Println("Found!")
			return
		}
		m[a[i]] = true
	}

	fmt.Println("Not Found!")

}

func main() {

	// Please sort the array before we find the pair
	// Using any algrithm the sort the array
	a := []int{-8, 1, 4, 6, 10, 45}
	findPair(a, 17)
	findPairWithHashing(a, 18)
}
