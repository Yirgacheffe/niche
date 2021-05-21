package main

import "fmt"

func hello() []string {
	return nil
}

func main() {

	a := hello

	if a == nil {
		fmt.Println("nil")
	} else {
		fmt.Println("not nil")
	}

}
