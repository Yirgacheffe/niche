package main

import "fmt"

func main() {

	a, b := 1, 3
	fmt.Println(a, b)

	b = a + b
	a = b - a
	b = b - a

	fmt.Println(a, b)

}
