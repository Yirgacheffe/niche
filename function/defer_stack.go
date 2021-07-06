package main

import "fmt"

func main() {
	name := "Lamztick"
	fmt.Printf("Original name is: %s\n", name)
	fmt.Printf("Reverse name is: ")

	for _, v := range []rune(name) {
		defer fmt.Printf("%c", v)
	}
}
