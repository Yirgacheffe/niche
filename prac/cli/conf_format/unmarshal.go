package main

import "fmt"

const (
	exampleJSON = `{"name": "example1", "age": 64}`
	exampleYAML = `name: example2
age: 68
`
)

// UnmarshalAll takes data in various formats
// and converts them into structs
func UnmarshalAll() error {

	j := JSONData{}
	y := YAMLData{}

	if err := j.Decode([]byte(exampleJSON)); err != nil {
		return err
	}

	fmt.Println("json unmarshal=", j)

	if err := y.Decode([]byte(exampleYAML)); err != nil {
		return err
	}

	fmt.Println("yaml unmarshal=", y)

	return nil // Re, return as no any error happened

}
