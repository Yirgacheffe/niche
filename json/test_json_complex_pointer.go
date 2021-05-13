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
	Profile        *profile
}

func main() {

	d := []byte(`
	{
		"FirstName": "john", 
		"HeightInMeters": 1.75, 
		"IsMale": null,
		"Languages": ["English"],
		"Subjects": ["Math", "Science"],
		"Grades": null,
		"Profile": {
			"Followers": 1989
		}
	}`)

	// Profile will be merged
	var johnson student = student{
		IsMale:    true,
		Languages: [2]string{"Korea", "Chinese"},
		Subjects:  nil,
		Grades:    map[string]string{"Science": "A+"},
		Profile: &profile{
			UserName: "johnsondoe89",
		},
	}

	fmt.Printf("Error: %v\n", json.Unmarshal(d, &johnson))
	fmt.Printf("%#v\n", johnson)
	fmt.Printf("%#v\n", johnson.Profile)

}
