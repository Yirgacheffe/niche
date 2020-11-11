package utils

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type pos struct {
	X      int
	Y      int
	Object string
}

// GobExample demostrates usting the gob package
func GobExample() error {

	buffer := bytes.Buffer()
	p := pos{
		X: 10, Y: 15, Object: "wrench",
	}

	// If p was an interface call
	// gob.Register first
	e := gob.NewEncoder(&buffer)
	if err := e.Encode(&p); err != nil {
		return err
	}

	fmt.Println("Gob encoded valued length: ", len(buffer.Bytes()))

	p2 := pos{}
	d := gob.NewDecoder(&buffer)
	if err := d.Decode(&p2); err != nil {
		return err
	}

	fmt.Println("Gob Decode value: ", p2)

	return nil

}
