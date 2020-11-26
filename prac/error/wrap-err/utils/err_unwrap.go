package utils

import (
	"fmt"

	"github.com/pkg/errors"
)

// Unwrap will unwrap an error and do type assertion
func Unwrap() {

	err := error(TypedError{errors.New("an error occured")})
	err = errors.Wrap(err, "wrapped")

	fmt.Println("wrapped error: ", err)

	switch errors.Cause(err).(type) {
	case TypedError:
		fmt.Println("a typed error occured: ", err)
	default:
		fmt.Println("an unknown error occured")
	}

}

// StackTrace will print all  the stack for the error
func StackTrace() {

	err := error(TypedError{errors.New("an error occured")})
	err = errors.Wrap(err, "wrapped")

	fmt.Printf("%+v\n", err)

}
