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

	s1 := student{
		"Jake",
		"Blake",
		35,
	}

	s2 := student{
		"Alyson",
		"Felix",
		28,
	}

	s := []student{s1, s2}
	f := filter(s, func(x student) bool {
		return x.age < 30
	})

	c := filter(s, func(b student) bool {
		return b.firstname == "Jake"
	})

	fmt.Println(f)
	fmt.Println(c)

}
