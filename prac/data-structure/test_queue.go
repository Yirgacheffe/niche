package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

var size = 0
var queue = new(Node)

func Push(t *Node, v int) bool {

	if t == nil {
		queue = &Node{
			v,
			nil,
		}

		size++
		return true
	}

	temp := &Node{v, t}
	queue = temp

	size++
	return true

}

func Pop(t *Node) (int, bool) {

	if size == 0 {
		return 0, false
	}

	if size == 1 {
		queue = nil
		size--
		return t.Value, true
	}

	var temp *Node

	for t.Next != nil {
		temp = t
		t = t.Next
	}

	v := (temp.Next).Value
	temp.Next = nil

	size--
	return v, true

}

func traverse(t *Node) {

	if size == 0 {
		fmt.Println("-> Empty Queue!")
		return
	}

	for t != nil {
		fmt.Printf("%d -> ", t.Value)
		t = t.Next
	}

	fmt.Println("")

}

func main() {

	// Test the function of a queue data structure
	queue = nil
	Push(queue, 10)

	fmt.Println("Size:", size)
	traverse(queue)

	v, b := Pop(queue)
	if b {
		fmt.Println("Pop:", v)
	}
	fmt.Println("Size:", size)

	for i := 0; i < 5; i++ {
		Push(queue, i)
	}

	traverse(queue)
	fmt.Println("Size:", size)

	v, b = Pop(queue)
	if b {
		fmt.Println("Pop:", v)
	}
	fmt.Println("Size:", size)

	v, b = Pop(queue)
	if b {
		fmt.Println("Pop:", v)
	}

	traverse(queue)
	fmt.Println("Size:", size)

}
