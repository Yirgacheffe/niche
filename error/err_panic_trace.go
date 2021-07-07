package main

import (
	"fmt"
	"log"
	"os"
	"runtime/debug"

	"golang.org/x/net/html"
)

func forEachNode(n *html.Node, pre, post func(*html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

func soleTitle(doc *html.Node) (title string, err error) {
	type bailout struct{}

	defer func() {
		switch p := recover(); p {
		case nil:
			// no panic found
		case bailout{}:
			err = fmt.Errorf("multiple title elements")
		default:
			panic(p) // unexpected panic
		}
	}()

	visitTitle := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			if title != "" {
				panic(bailout{}) // multiple title
			}
			title = n.FirstChild.Data
		}
	}

	forEachNode(doc, visitTitle, nil)

	if title == "" {
		return "", fmt.Errorf("no title element")
	}

	return title, nil
}

func recoverFullName() {
	if r := recover(); r != nil {
		fmt.Println("recovered from ", r)
		debug.PrintStack()
	}
}

func fullName(firstName *string, lastName *string) {
	defer recoverFullName()

	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
}

func main() {
	defer fmt.Println("deferred call in main")

	firstName := "Elon"
	fullName(&firstName, nil)
	fmt.Println("returned normally from main")

	doc, err := html.Parse(os.Stdin)
	if err != nil {
		log.Printf("find link: %v\n", err)
		os.Exit(1)
	}

	t, err := soleTitle(doc)
	if err != nil {
		panic(err)
	}

	fmt.Println("Got title element:", t) // ........
}
