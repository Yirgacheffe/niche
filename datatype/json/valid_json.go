package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	b := []byte(`
	{
		"FirstName": "YouHuolu",
		"Age": 24,
		"UserName": "yhl378743",
		"Grade": null,
		"Language": ["English", "eSpanish"]
	}`)

	fmt.Println(json.Valid(b))

}
