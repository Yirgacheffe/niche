package main

import (
	"fmt"
	"sort"
)

type Person struct {
	name   string
	height int
	weight int
}

func main() {

	persons := make([]Person, 5)

	persons[0] = Person{"Mihalis", 180, 90}
	persons[1] = Person{"Bill", 134, 45}
	persons[2] = Person{"Marietta", 155, 45}
	persons[3] = Person{"Epifanios", 144, 50}
	persons[4] = Person{"Athina", 134, 45}

	/*
		persons := make([]Person, 0)

		persons = append(persons, Person{"Mihalis", 180, 90})
		persons = append(persons, Person{"Bill", 134, 45})
		persons = append(persons, Person{"Marietta", 155, 45})
		persons = append(persons, Person{"Epifanios", 144, 50})
		persons = append(persons, Person{"Athina", 134, 45})
	*/

	sort.Slice(persons, func(i, j int) bool {
		return persons[i].height < persons[j].height
	})
	fmt.Println("<:", persons)

	sort.Slice(persons, func(i, j int) bool {
		return persons[i].height > persons[j].height
	})

	fmt.Println(">:", persons)

}
