package main

import "fmt"

type Pair struct {
	fir int
	sec int
}

// Target is O(n * n), get value (a + b) = (c + d)
func find(a []int) bool {

	l := len(a)
	m := make(map[int]Pair)

	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			sum := a[i] + a[j]
			if _, ok := m[sum]; ok {
				p := m[sum]
				fmt.Printf("%d+%d = %d+%d\n", a[i], a[j], p.fir, p.sec)
				return true
			} else {
				m[sum] = Pair{a[i], a[j]}
			}
		}
	}

	return false

}

func main() {
	a := []int{14, 3, 4, 7, 1, 2, 9, 8, 10, 11}
	find(a)
}
