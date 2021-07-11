package main

import (
	"encoding/json"
	"fmt"
)

type student map[string]interface{}

func main() {

	mike := student{
		"FirstName":      "Mike",
		"LastName":       "Pompeo",
		"Age":            56,
		"HeightInMeters": 1.75,
		"IsMale":         true,
	}

	mikeJSON, _ := json.MarshalIndent(mike, "", "  ")

	fmt.Println(mikeJSON) // Slices of byte
	fmt.Println(string(mikeJSON))

}
