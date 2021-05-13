package main

import "fmt"

// p is target set
// s is the subset candidate
func isSubsetInBruteforce(p, s []int) bool {

	pl := len(p)
	sl := len(s)

	j := 0
	i := 0

	// loop array subset
	for j = 0; j < sl; j++ {

		// loop array p
		for i = 0; i < pl; i++ {
			if s[j] == p[i] {
				break
			}
		}

		if i == pl {
			return false
		}
	}

	return true

}

// Version 1: Brute force
// Version 2: Binery Search
func main() {

	p := []int{11, 1, 13, 36, 25, 9, 14, 20}
	s := []int{25, 20, 30}
	x := []int{11, 25, 9, 14}

	isSubset1 := isSubsetInBruteforce(p, s)
	fmt.Println(isSubset1)

	isSubset2 := isSubsetInBruteforce(p, x)
	fmt.Println(isSubset2)

}
