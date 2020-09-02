package main

import "fmt"

func describe(i interface{}) {
	fmt.Printf("Type = %T, Value = %v\n", i, i)
}

func main() {

	n := 10
	describe(n)

	s := "Hello World"
	describe(s)

	strt := struct {
		name string
	}{
		name: "Kurosaki",
	}
	describe(strt)

}
