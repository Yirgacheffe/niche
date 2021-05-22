package main

import (
	"encoding/json"
	"fmt"
)

type profile struct {
	UserName  string
	Followers int
}

type student struct {
	FirstName      string
	LastName       string
	HeightInMeters float64
	IsMale         bool
	Languages      [2]string
	Subjects       []string
	Grades         map[string]string
	Profile        profile
}

func main() {

	data := []byte(`
	{
		"FirstName": "john", 
		"HeightInMeters": 1.75, 
		"IsMale": null,
		"Languages": ["English", "Spanish", "German"],
		"Subjects": ["Math", "Science"],
		"Grades": {"Math": "A"},
		"Profile": {
			"UserName": "jondoe83",
			"Followers": 1989
		}
	}`)

	// Language 'German' will be dropped
	// Subject 'Art' will be dropped
	// Slice 'Grades' will be merged
	var john student = student{
		IsMale:   true,
		Subjects: []string{"Art"},
		Grades:   map[string]string{"Science": "A+"},
	}

	fmt.Printf("Error: %v\n", json.Unmarshal(data, &john))
	fmt.Printf("%#v\n", john)

}
