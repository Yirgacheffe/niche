package main

import "fmt"

// Sliding window
func findSubarrayM1(a []int, sum int) {

	csum, s := a[0], 0
	n := len(a)

	for i := 1; i <= n; i++ {
		for (csum > sum) && (s < i-1) {
			csum -= a[s]
			s++
		}
		if csum == sum {
			fmt.Printf("Found: %d to %d\n", s, i-1)
			return
		}
		if i < n {
			csum += a[i]
		}
	}

	fmt.Println("Not found!!!")

}

func findSubarrayM2(a []int, sum int) {

	csum, s := 0, 0
	n := len(a)

	for i := 0; i < n; i++ {
		csum += a[i]
		for (csum > sum) && (s <= i) {
			csum -= a[s]
			s++
		}
		if csum == sum {
			fmt.Printf("Found: %d to %d\n", s, i)
			return
		}
	}

	fmt.Println("Not found!!!")

}

// Hashmap version?
func findSubarrayWithHash(a []int, sum int) {

	hmap := make(map[int]int)
	csum := 0

	for i := 0; i < len(a); i++ {
		csum += a[i]
		// this is the important part... :-(
		e := csum - sum

		if v, ok := hmap[e]; ok {
			fmt.Printf("Found: %d to %d\n", v+1, i)
			return
		} else {
			hmap[csum] = i
		}
	}

	fmt.Println("Not found!!!")

}

func main() {
	a := []int{1, 4, 20, 3, 10, 5}

	findSubarrayM1(a, 13)
	findSubarrayM2(a, 33)
	findSubarrayWithHash(a, 33)
}
