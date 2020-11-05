package utils

import (
	"encoding/json"
	"fmt"
)

const (
	jsonBlob = `{"name": "De Guo"}`
	fullJson = `{"name": "De Guo", "age": 40}`
)

type Example struct {
	Name string `json:"name"`
	Age  int    `json:"age,omitempty"`
}

// BaseEncoding shows encoding and decoding with normal types
func BaseEncoding() error {

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
	err = json.Unmarshal([]byte(fullJson), &e)
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
