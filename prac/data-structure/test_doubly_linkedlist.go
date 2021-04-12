package main

import "fmt"

type Node struct {
	Previous *Node
	Value    int
	Next     *Node
}

func addNode(t *Node, v int) int {

	if t == nil {
		t = &Node{nil, v, nil}
		root = t
		return 0
	}

	if v == t.Value {
		fmt.Println("Node already exists:", v)
		return -1
	}

	if t.Next == nil {
		temp := t
		t.Next = &Node{temp, v, nil}
		return -2
	}

	return addNode(t.Next, v)

}

func traverse(t *Node) {

	if t == nil {
		fmt.Println("-> Empty list!")
		return
	}

	for t != nil {
		fmt.Println("-> ", t.Value)
		t = t.next
	}

	fmt.Println()

}

func reverse(t *Node) {

	if t == nil {
		fmt.Println("-> Empty list!")
		return
	}

	var tmp Node

	// Find last node in the link
	for t != nil {
		tmp = t
		t = t.Next
	}

	for tmp.Previous != nil {
		fmt.Printf("%d -> ", tmp.Value)
		tmp = tmp.Previous
	}

	fmt.Printf("%d -> ", tmp.Value)
	fmt.Println()

}

func size(t *Node) int {

}

func lookupNode(t *Node, v int) int {

}

var root = new(Node)

func main() {
	fmt.Println("Hello world!")
}
