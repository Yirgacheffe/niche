package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type profile struct {
	UserName string
	Follower string // format float64
}

func (p *profile) UnmarshalJSON(data []byte) error {

	var c map[string]interface{}

	_ = json.Unmarshal(data, &c)
	fmt.Printf("container: %T / %#v \n\n", c, c)

	// Extract interface value ... ...
	iUserName, _ := c["UserName"]
	iFollower, _ := c["Follower"]

	fmt.Printf("iUserName: %T/%#v \n", iUserName, iUserName)
	fmt.Printf("iFollower: %T/%#v \n", iFollower, iFollower)

	fmt.Println()

	// Extract concreate value ... ...
	userName, _ := iUserName.(string)
	follower, _ := iFollower.(float64)

	fmt.Printf("username: %T/%#v \n", userName, userName)
	fmt.Printf("follower: %T/%#v \n", follower, follower)

	// Assign vlaue to profile struct
	p.UserName = strings.ToUpper(userName)
	p.Follower = fmt.Sprintf("%.2fk", follower/1000)

	return nil

}

type student struct {
	FirstName string
	Profile   profile
}

func main() {

	data := []byte(`
	{
		"FirstName": "shikamaru",
		"Profile": {
			"UserName": "shikamaru983", "Follower": 1783
		}
	}`)

	var maru student

	fmt.Printf("Error: %v\n", json.Unmarshal(data, &maru))
	fmt.Printf("%#v\n", maru)

}
