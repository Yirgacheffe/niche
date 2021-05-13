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
		fmt.Printf("%d -> ", t.Value)
		t = t.Next
	}

	fmt.Println()

}

func reverse(t *Node) {

	if t == nil {
		fmt.Println("-> Empty list!")
		return
	}

	var tmp *Node

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
	if t == nil {
		fmt.Println("-> Empty list!")
		return 0
	}

	n := 0

	for t != nil {
		n++
		t = t.Next
	}

	return n
}

func lookupNode(t *Node, v int) bool {
	if t == nil {
		return false
	}

	if t.Value == v {
		return true
	}

	if t.Next == nil {
		return false
	}

	return lookupNode(t.Next, v)
}

var root = new(Node)

func main() {
	fmt.Println(root)
	root = nil
	traverse(root)

	addNode(root, 1)
	addNode(root, 1)
	traverse(root)

	addNode(root, 10)
	addNode(root, 5)
	addNode(root, 0)
	addNode(root, 0)
	traverse(root)

	addNode(root, 100)
	reverse(root)
	fmt.Println("Size:", size(root))

	traverse(root)
}
