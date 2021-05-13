package main

import (
	"encoding/json"
	"fmt"
)

const (
	a = 1 + iota
	b
	c
)

// Type `profile` implement MarshalJSON
type profile struct {
	UserName  string
	Followers int
}

func (p profile) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{"f_cnt": %d}`, p.Followers)), nil
}

//
type age int

// MarshalText -- "" --> \"\"
func (a age) MarshalText() ([]byte, error) {
	return []byte(fmt.Sprintf(`{"age": %d}`, int(a))), nil
}

type student struct {
	FirstName string
	LastName  string
	Age       age
	Profile   profile
}

func main() {

	kakashi := &student{
		FirstName: "Kakashi",
		LastName:  "Hatake",
		Age:       28,
		Profile: profile{
			UserName:  "kakashi34984",
			Followers: 129,
		},
	}

	kakashiJSON, _ := json.MarshalIndent(kakashi, "", "    ")
	fmt.Println(string(kakashiJSON))

}
