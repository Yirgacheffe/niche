package main

import "fmt"

type author struct {
	firstName string
	lastName  string
	bio       string
}

func (a *author) fullName() string {
	return fmt.Sprintf("%s %s", a.firstName, a.lastName)
}

type post struct {
	title   string
	content string
	author
}

func (p *post) details() {
	fmt.Printf("Title: %s, Content: %s, Author: %s, Bio: %s\n", p.title, p.content, p.fullName(), p.bio)
}

type website struct {
	posts []post
}

func (w *website) contents() {
	fmt.Println("Content of Website")
	for _, v := range w.posts {
		v.details()
		fmt.Println()
	}
}

func main() {
	author1 := author{
		"Naveen",
		"Ramanathan",
		"Golang Enthusiast",
	}
	post1 := post{
		"Inheritance in Go",
		"Go supports composition instead of inheritance",
		author1,
	}
	post2 := post{
		"Struct instead of Class in Go",
		"Go does not support classes but method be added",
		author1,
	}

	w := website{
		posts: []post{post1, post2},
	}

	// post1.details()
	w.contents()
}
