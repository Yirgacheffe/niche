package main

import (
	"errors"
	"fmt"
)

// ErrorTyped is a way to make a package level
// error to check against.
// I.e. if err == TypedError
var ErrorTyped = errors.New("this is a typed error")

// ErrorValue is a way to make package level
// error to check against.
// I.e. if err == ErrorValue
var ErrorValue = errors.New("this is a typed error")

// TypedError is a way to make an error type
// you can do err.(type) == ErrorValue
type TypedError struct {
	error
}

// BasicErrors demonstrates some way to create errors
func BasicErrors() {

	err := errors.New("this is a quick and easy way to create to error")
	fmt.Println("errors.New: ", err)

	err = fmt.Errorf("an error occured: %s", "someting happened")
	fmt.Println("fmt.Errorf: ", err)

	err = ErrorTyped
	fmt.Println("typed error: ", err)

	err = ErrorValue
	fmt.Println("value error: ", err)

	err = TypedError{errors.New("typed error")}
	fmt.Println("typed error: ", err)

}
