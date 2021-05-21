package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links // visit appends to links each link found in n and returns
}

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		// push tag
		stack = append(stack, n.Data)
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = n.NextSibling {
		outline(stack, c)
	}
}

func main() {
	// call with './fetch https://golang.org | ./find_links_01'
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		log.Printf("find link: %v\n", err)
		os.Exit(1)
	}

	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}

	outline(nil, doc)
	// call with './fetch https://golang.org | ./find_links_01'

	add := func(r rune) rune {
		return r + 1
	}
	fmt.Println(strings.Map(add, "HAL-9000"))
	fmt.Println(strings.Map(add, "VMS"))

}
