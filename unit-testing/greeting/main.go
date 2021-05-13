package main

import "fmt"

func main() {

	greetingEmpty := hello("")
	fmt.Println(greetingEmpty)

	greetingJohnson := hello("Johnson")
	fmt.Println(greetingJohnson)
}
