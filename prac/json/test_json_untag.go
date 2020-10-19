package main

import (
	"encoding/json"
	"fmt"
)

type profile struct {
	UserName  string `json:"u_name"`
	Followers int    `json:"f_count"`
}

type student struct {
	FirstName      string   `json:"f_name"`
	LastName       string   `json:"-"`
	HeightInMeters float64  `json:"height"`
	IsMale         bool     `json:"male"`
	Languages      []string `json:,lang,omitempty`
	Profile        profile  `json:"profile"`
}

func main() {

	data := []byte(`
	{
		"f_name": "Shu",
		"l_name": "Pu",
		"height": 1.78,
		"male": true,
		"lang": null,
		"profile": {
			"u_name": "shupu9834", "f_count": 1324
		}
	}`)

	var shupu student

	fmt.Printf("%v\n", json.Unmarshal(data, &shupu))
	fmt.Printf("%#v\n", shupu)

}
