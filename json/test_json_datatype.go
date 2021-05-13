package main

import (
	"encoding/json"
	"fmt"
)

type Profile struct {
	UserName  string
	Followers int
	Grades    map[string]string
}

type student struct {
	FirstName, LastName string
	Age                 int
	Profile             Profile
	Languages           []string
}

// Profile can be anonymous and promoted to correct struct 'profile'
// In that case
// The field in profile become field of 'Student'
type student2 struct {
	FirstName string
	LastName  string
	Age       int
	Profile
	Languages []string
}

func main() {

	terasa := student{
		FirstName: "Terasa",
		LastName:  "May",
		Age:       63,
		Profile: Profile{
			UserName:  "terasa17994",
			Followers: 463,
			Grades:    map[string]string{"Mathemetics": "A", "French": "A+"},
		},
		Languages: []string{"English", "French"},
	}

	terasaJSON, _ := json.MarshalIndent(terasa, "", "    ")
	fmt.Println(string(terasaJSON))

	// Fields in 'Profile' will be promoted into 'student'
	// 'Grade' will be null as we are not setting value to it
	john := student2{
		FirstName: "Boris",
		LastName:  "Johnson",
		Age:       56,
		Profile: Profile{
			UserName:  "boris8733244",
			Followers: 500,
		},
		Languages: []string{"English", "French"},
	}

	johnJSON, _ := json.MarshalIndent(john, "", "    ")
	fmt.Println(string(johnJSON))

}
