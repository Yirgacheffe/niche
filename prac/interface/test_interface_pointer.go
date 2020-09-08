package main

import "fmt"

type Describer interface {
	Describe()
}

type Person struct {
	name string
	age  int
}

func (p Person) Describe() {
	fmt.Printf("%s is %d years old\n", p.name, p.age)
}

type Address struct {
	state   string
	country string
}

func (a *Address) Describe() {
	fmt.Printf("address is %s - %s\n", a.state, a.country)
}

func main() {

	p1 := Person{"Nirawana", 10}
	p2 := Person{"RadioHead", 13}

	var d1 Describer
	d1 = p1
	d1.Describe()

	d1 = &p2
	d1.Describe()

	// Address testing
	a1 := Address{"Akense", "USA"}

	var d2 Describer
	// d2 = a1
	d2 = &a1
	d2.Describe()

}
