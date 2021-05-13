package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type student struct {
	FirstName string `json:"f_name"`
	LastName  string `json:"l_name"`
	Age       int    `json:"age"`
	Email     string `json:"email,omitempty"`
}

func main() {

	b := []byte(`
	{
		"f_name": "Mike",
		"l_name": "Pence",
		"age": 63,
		"email": "mike.p@doe.com"
	}`)

	if ok := json.Valid(b); !ok {
		log.Fatal("Invalid json format found.")
	}

	var mike student
	err := json.Unmarshal(b, &mike)

	fmt.Printf("%v\n", err)
	fmt.Printf("%#v\n", mike)

}
