package utils

import (
	"encoding/json"
	"fmt"
)

const (
	jsonBlob = `{"name": "De Guo"}`
	fullJSON = `{"name": "De Guo", "age": 40}`
)

// Example is a basic struct with age and name fields
type Example struct {
	Name string `json:"name"`
	Age  int    `json:"age,omitempty"`
}

// RegularEncoding shows encoding and decoding with
// normal types
func RegularEncoding() error {

	e := Example{}

	// no age example
	err := json.Unmarshal([]byte(jsonBlob), &e)
	if err != nil {
		return err
	}
	fmt.Printf("Regular Unmarshal, no age: %+v\n", e)

	value, err := json.Marshal(&e)
	if err != nil {
		return err
	}
	fmt.Println("Regular Marshal, no age:", string(value))

	// age = 0 example
	err = json.Unmarshal([]byte(fullJSON), &e)
	if err != nil {
		return err
	}
	fmt.Printf("Regular Unmarshal, age = 0: %+v\n", e)

	value, err = json.Marshal(&e)
	if err != nil {
		return err
	}
	fmt.Println("Regular Marshal, age = 0:", string(value))

	return nil

}
