package main

import (
	"encoding/json"
	"fmt"
)

type student struct {
	FirstName      string
	LastName       string
	Email          string
	Age            int
	HeightInMeters float64
	IsMale         bool
}

func main() {

	mike := student{
		FirstName:      "Mike",
		LastName:       "Pompeo",
		Age:            56,
		HeightInMeters: 1.75,
		IsMale:         true,
	}

	mikeJSON, _ := json.Marshal(mike)

	fmt.Println(mikeJSON) // Slices of byte
	fmt.Println(string(mikeJSON))

}
