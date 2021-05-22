package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type person struct {
	Name string
	Age  int
}

func main() {

	jsonStream := strings.NewReader(`
		{"Name": "Rossel Galler", "Age": 28}
		{"Name": "Monica Galler", "Age": 27}
		{"Name": "Jack Galler", "Age": 56}
	`)

	decoder := json.NewDecoder(jsonStream)
	var rossel, monica person

	decoder.Decode(&rossel)
	decoder.Decode(&monica)

	fmt.Printf("%#v\n", rossel)
	fmt.Printf("%#v\n", monica)

}
