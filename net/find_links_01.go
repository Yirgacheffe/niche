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

// outline2
var depth int

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
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

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		depth++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
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

	// outline
	outline(nil, doc)

	// call with './fetch https://golang.org | ./find_links_01'
	add := func(r rune) rune {
		return r + 1
	}
	fmt.Println(strings.Map(add, "HAL-9000"))
	fmt.Println(strings.Map(add, "VMS"))

	// another outline
	forEachNode(doc, startElement, endElement)
}
