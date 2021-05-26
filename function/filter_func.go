package main

import "fmt"

type student struct {
	firstname string
	lastname  string
	age       int
}

func filter(s []student, f func(student) bool) []student {
	var r []student

	for _, v := range s {
		if f(v) == true {
			r = append(r, v)
		}
	}

	return r
}

func main() {
	s := []student{{"Jake", "Black", 35}, {"Alyson", "Felix", 28}}

	f := filter(
		s,
		func(x student) bool {
			return x.age < 30
		})

	c := filter(
		s,
		func(b student) bool {
			return b.firstname == "Jake"
		})

	fmt.Println(f)
	fmt.Println(c)
}
