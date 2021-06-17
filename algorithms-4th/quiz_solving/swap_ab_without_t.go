package main

import "fmt"

func main() {

	a, b := 1, 3
	fmt.Println("0: ", a, b)

	b = a + b
	a = b - a
	b = b - a

	fmt.Println("1: ", a, b)

	a = a ^ b
	b = a ^ b
	a = a ^ b

	fmt.Println("2: ", a, b)

}
