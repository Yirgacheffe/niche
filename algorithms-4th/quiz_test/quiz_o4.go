package main

import "fmt"

func main() {

	a := [2]uint8{5, 6}
	b := [2]uint8{5, 6}

	if a == b {
		fmt.Println("Equal.")
	} else {
		fmt.Println("Not equal.")
	}

}
