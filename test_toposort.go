package main

import (
	"fmt"
	"sort"
)

var prereqs = map[string][]string{
	"algorithms":            {"data structures"},
	"calculus":              {"linear algebra"},
	"compilers":             {"computer organzation", "data structures", "formal lauguages"},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operation system":      {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) []string {

	seen := make(map[string]bool)

	var order []string
	var visit func(items []string)

	visit = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visit(m[item])
				order = append(order, item)
			}
		}
	}

	// -----------------------
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	visit(keys)
	return order // --------------------##----

}
