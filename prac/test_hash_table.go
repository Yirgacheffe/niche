package main

import "fmt"

const SIZE = 15

type Node struct {
	Value int
	Next  *Node
}

type HashTable struct {
	Size  int
	Table map[int]*Node
}

func insert(hash *HashTable, value int) int {
	idx := hashIt(value, hash.Size)
	element := Node{
		value,
		hash.Table[idx],
	}
	hash.Table[idx] = &element
	return idx
}

func traverse(hash *HashTable) {

	for k := range hash.Table {
		if hash.Table[k] != nil {
			t := hash.Table[k]
			for t != nil {
				fmt.Printf("%d -> ", t.Value)
				t = t.Next
			}
			fmt.Println()
		}
	}

}

func hashIt(v, size int) int {
	return (v % size)
}

func main() {

	/*
		table := make(map[int]*Node, SIZE)
		hash := &HashTable{
			Table: table,
			Size:  SIZE,
		}
	*/

	hash := &HashTable{
		SIZE,
		make(map[int]*Node, SIZE),
	}

	fmt.Println("Number of spaces:", hash.Size)
	for i := 0; i < 120; i++ {
		insert(hash, i)
	}

	traverse(hash) // Println element in hashed table

}
