package utils

import (
	"fmt"

	"github.com/pkg/errors"
)

// WrappedError demonstrates error wrapping and
// annotating an error
func WrappedError(e error) error {
	return errors.Wrap(e, "An Error happened in WrappedError")
}

// TypedError is a error we can check against
type TypedError struct {
	error
}

// Wrap shows what happens when we wrap an error
func Wrap() {
	e := errors.New("Standard error")
	fmt.Println("Regular error - ", WrappedError(e))

	t := TypedError{errors.New("Typed error")}
	fmt.Println("Typed error - ", WrappedError(t))

	fmt.Println("Nil error - ", WrappedError(nil))
}
