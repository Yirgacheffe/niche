package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// JSONData is our common data struct with JSON struct tags
type JSONData struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// ToJSON dumps the JSONData struct to
// a JSON format bytes.Buffer
func (j *JSONData) ToJSON() (*bytes.Buffer, error) {

	b, err := json.Marshal(j)
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(b)
	return buf, nil

}

// Decode will decode into JSONData
func (j *JSONData) Decode(data []byte) error {
	return json.Unmarshal(data, j)
}

// JSONExampleWithMap shows ways to use types beyond structs and other
// useful functions
func JSONExampleWithMap() error {

	res := make(map[string]string)
	err := json.Unmarshal([]byte(`{"key1": "value1"}`), &res)
	if err != nil {
		return err
	}

	fmt.Println("unmarshal into a map instead of a struct:", res)

	b := bytes.NewReader([]byte(`{"key2": "value2"}`))
	decoder := json.NewDecoder(b)

	if err := decoder.Decode(&res); err != nil {
		return err
	}

	fmt.Println("use decoders/encoders to work with streams:", res)
	return nil

}
