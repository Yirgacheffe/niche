package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {

	// `Buffer` implement `Write()` from `Writer` interface
	buf := new(bytes.Buffer)

	bufEncoder := json.NewEncoder(buf)

	bufEncoder.Encode(person{"Rossel Galler", 28})
	bufEncoder.Encode(person{"Monica Galler", 27})
	bufEncoder.Encode(person{"Jack Galler", 56})

	fmt.Println(buf) // Call `buf.String()` method to show

}
