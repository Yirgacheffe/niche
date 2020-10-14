package main

import (
	"encoding/json"
	"fmt"
)

type profile struct {
	UserName  string
	Followers int
	Grades    map[string]string
}

// Profile can be anonymous and promoted to correct struct 'profile'
type student struct {
	FirstName, LastName string
	Age                 int
	Profile             profile
	Languages           []string
}

func main() {

	terasa := student{
		FirstName: "Terasa",
		LastName:  "May",
		Age:       63,
		Profile: profile{
			UserName:  "terasa17994",
			Followers: 463,
			Grades:    map[string]string{"Mathemetics": "A", "French": "A+"},
		},
		Languages: []string{"English", "French"},
	}

	terasaJSON, _ := json.MarshalIndent(terasa, "", "    ")
	fmt.Println(string(terasaJSON))

}
