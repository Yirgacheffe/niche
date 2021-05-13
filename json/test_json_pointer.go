package main

import (
	"encoding/json"
	"fmt"
)

type profileI interface {
	Follow()
}

type profile struct {
	UserName  string
	Followers int
}

func (p *profile) Follow() {
	p.Followers++
}

type student struct {
	FirstName string
	LastName  string
	Primary   profileI
	Secondry  profileI
}

func main() {

	barret := &student{
		FirstName: "Amy",
		LastName:  "Barret",
		Primary: &profile{
			UserName: "barret7634", Followers: 170,
		},
	}

	barret.Primary.Follow()

	barretJSON, _ := json.MarshalIndent(barret, "", "\t")
	fmt.Println(string(barretJSON))

}
