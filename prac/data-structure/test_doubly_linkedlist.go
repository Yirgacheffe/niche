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

}

func reverse(t *Node) {

}

func size(t *Node) int {

}

func lookupNode(t *Node, v int) int {

}

var root = new(Node)

func main() {
	fmt.Println("Hello world!")
}
