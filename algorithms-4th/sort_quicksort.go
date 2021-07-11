package main

import "fmt"

func quickSort(a []int, low, high int) {

	if low < high {
		pivot := partition(a, low, high)

		quickSort(a, low, pivot-1)
		quickSort(a, pivot+1, high)
	}

}

func partition(a []int, low, high int) int {

	pi := a[high]
	i := low - 1

	for j := low; j <= high-1; j++ {
		if a[j] < pi {
			i++
			swap(a, i, j)
		}
	}

	swap(a, i+1, high)
	return i + 1

}

func swap(a []int, i, j int) {

	t := a[i]
	a[i] = a[j]
	a[j] = t

}

func main() {

	a := []int{10, 80, 30, 90, 40, 50, 70, 100, 25, 60, 75}
	n := len(a)

	fmt.Println(a)
	quickSort(a, 0, n-1)
	fmt.Println(a)

}
