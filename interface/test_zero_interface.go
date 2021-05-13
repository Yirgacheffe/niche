package main

import "fmt"

type Describer interface {
	Describe()
}

func main() {
	var d Describer
	if d == nil {
		fmt.Printf("d1 is nil and has type %T value %v\n", d, d)
	}

	d.Describe()
}
