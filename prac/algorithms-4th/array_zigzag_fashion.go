package main

import (
	"fmt"
	"strconv"
)

func zigzagSlow(a []int) {

	n := len(a)

	for i := 1; i < n-1; i += 2 {
		j := i + 1
		swap(a, i, j)
	}

	printArrayWithZigZagStyle(a)

}

func zigzagFast(a []int) {

	// 'true' we want '<', 'false' we want '>'
	flag := true
	n := len(a)

	for i := 0; i < n-1; i++ {
		j := i + 1
		if flag {
			if !(a[i] < a[j]) {
				swap(a, i, j)
			}
		} else {
			if !(a[i] > a[j]) {
				swap(a, i, j)
			}
		}
		flag = !flag
	}

	printArrayWithZigZagStyle(a)

}

func swap(a []int, i, j int) {
	t := a[i]
	a[i] = a[j]
	a[j] = t
}

func printArrayWithZigZagStyle(a []int) {

	o := []string{}
	o = append(o, strconv.Itoa(a[0]))

	for i := 1; i < len(a); i++ {
		var c string

		if a[i] > a[i-1] {
			c = "<"
		} else {
			c = ">"
		}

		o = append(o, c)
		o = append(o, strconv.Itoa(a[i]))
	}

	fmt.Printf("Output ZigZag style: %v\n", o)

}

func main() {

	sorttedArray := func() []int {
		return []int{
			1, 2, 3, 4, 6, 7, 8,
		}
	}

	zigzagSlow(sorttedArray())
	zigzagFast(
		[]int{4, 3, 7, 8, 6, 2, 1, 9, 11, 12},
	)
}
