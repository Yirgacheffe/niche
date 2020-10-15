package main

import (
	"encoding/json"
	"fmt"
)

type profile struct {
	UserName  string `json:"user_name,omitempty"`
	Followers int    `json:"followers,omitempty,string"`
}

type student struct {
	FirstName string  `json:"f_name"`
	LastName  string  `json:"l_name,omitempty"`
	Email     string  `json:"-"`       // Always discard
	Age       int     `json:"-,"`      // '-' as field name
	IsMale    bool    `json:",string"` // to string
	Profile   profile `json:""`        // no effect to the filed
}

func main() {

	keyle := &student{
		FirstName: "Keyle",
		LastName:  "",
		Age:       33,
		Email:     "keyle@doe.com",
		IsMale:    false,
		Profile: profile{
			UserName:  "keyle909394",
			Followers: 873,
		},
	}

	keyleJSON, _ := json.MarshalIndent(keyle, "", "    ")
	fmt.Println(string(keyleJSON))

}
