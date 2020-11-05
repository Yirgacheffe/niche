package utils

import (
	"encoding/json"
	"fmt"
)

// ExamplePointer is the same, but uses a *Int
type ExamplePointer struct {
	Name string `json:"name"`
	Age  *int   `json:"age,omitempty"`
}

// PointerEncoding shows methods for dealing with nil/omitted values
func PointerEncoding() error {

	e := ExamplePointer{}

	// no age example
	err := json.Unmarshal([]byte(jsonBlob), &e)
	if err != nil {
		return err
	}
	fmt.Printf("Pointer Unmarshal, no age: %+v\n", e)

	value, err := json.Marshal(&e)
	if err != nil {
		return err
	}
	fmt.Println("Pointer Marshal, no age:", string(value))

	// age = 0 example
	err = json.Unmarshal([]byte(fullJSON), &e)
	if err != nil {
		return err
	}
	fmt.Printf("Pointer Unmarshal, age = 0: %+v\n", e)

	value, err = json.Marshal(&e)
	if err != nil {
		return err
	}
	fmt.Println("Pointer Marshal, age = 0:", string(value))

	return nil

}
