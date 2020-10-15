package main

import (
	"bytes"
	"fmt"
)

func main() {

	in := bytes.NewReader([]byte("example"))
	out := &bytes.Buffer{}

	fmt.Print("Stdout on copy = ")
	if err := Copy(in, out); err != nil {
		panic(err)
	}

	fmt.Println("Out byte buffer = ", out.String())

	fmt.Print("Stdout on Pipe = ")
	if err := PipeExample(); err != nil {
		panic(err)
	}

}
