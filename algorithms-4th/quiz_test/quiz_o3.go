package main

import "fmt"

func hello(n ...int) {
	n[0] = 8
}

func main() {

	a := []int{0, 1, 3, 5, 7}
	hello(a...)

	fmt.Println(a[0])

}
