package main

import (
	"errors"
	"fmt"
)

func main() {
	msg := "not found"
	a, b := errors.New(msg), errors.New(msg)

	fmt.Println(a == b)
}
