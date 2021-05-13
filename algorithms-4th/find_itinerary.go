package main

import "fmt"

func findItinerary(tickets map[string]string) {

	// Reverse the map
	// To create a new map
	toFrom := make(map[string]string)
	for k, v := range tickets {
		toFrom[v] = k
	}

	var starting string

	// Find started point
	for k, _ := range tickets {
		if _, ok := toFrom[k]; !ok {
			starting = k
			break
		}
	}

	to, ok := tickets[starting]

	for ok {
		fmt.Printf("%v -> %v, ", starting, to)
		starting = to
		to, ok = tickets[starting]
	}

	fmt.Println("")

}

func printIt(m map[string]string) {
	for k, v := range m {
		fmt.Printf("%v -> %v\n", k, v)
	}
}

func main() {

	tickets := map[string]string{
		"Goa":     "Chennai",
		"Delhi":   "Goa",
		"Bombay":  "Delhi",
		"Chennai": "Banglore",
	}
	findItinerary(tickets)

}
