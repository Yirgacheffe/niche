package main

import "fmt"

const (
	n = 41
	k = 3
)

func array_solution() int {

	var a [n]int
	for i := range a {
		a[i] = 1
	}

	var (
		i       = 0
		counter = 0
		remain  = n
	)

	for remain > 1 {
		if a[i] != 0 {
			counter++
			if counter == k {
				counter = 0
				a[i] = 0
				remain--
			}
		}

		i++
		if i == n {
			i = 0
		}
	}

	// Person p is the winner
	var p int
	for i := range a {
		if a[i] != 0 {
			p = i + 1
		}
	}
	return p
}

func main() {
	fmt.Println("the winner is: ", array_solution())
}
