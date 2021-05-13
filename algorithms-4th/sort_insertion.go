package main

import "fmt"

func sort(a []int) {

	n := len(a)

	for i := 0; i < n; i++ {
		for j := i; j > 0 && (a[j] < a[j-1]); j-- {
			exch(a, j, j-1)
		}
	}

}

func exch(a []int, i, min int) {
	var t int
	t = a[i]
	a[i] = a[min]
	a[min] = t
}

func main() {

	a := []int{4, 8, 1, 29, 39, 5, 23, 100, 34, 2, 84, 9, 14}
	fmt.Println(a)
	sort(a)
	fmt.Println(a)

}
