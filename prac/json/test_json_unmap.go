package main

import (
	"encoding/json"
	"fmt"
)

type student map[string]interface{}

func main() {

	data := []byte(`
	{
		"id": 763848, 
		"f_name": "boloke", 
		"l_name": "russel", 
		"height": 1.75,
		"male": true,
		"lang": null,
		"subjects": ["Math", "Science"],
		"profile": {
			"u_name": "realrussel8313", "followers": 7383
		}
	}`)

	var russel student

	fmt.Printf("Error: %v\n", json.Unmarshal(data, &russel))
	fmt.Printf("%#v\n\n", russel)

	i := 1
	for k, v := range russel {
		fmt.Printf("%d: Key (`%T`)`%v`, value (`%T`)`%#v`\n", i, k, k, v, v)
		i++
	}

}
