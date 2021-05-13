package main

import (
	"log"

	"github.com/pkg/errors"
)

// OriginalError returns the error original error
func OriginalError() error {
	return errors.New("error occured")
}

// PassThroughError calls OriginalError
// and forwards
// the error along after wrapping.
func PassThroughError() error {
	err := OriginalError()
	return errors.Wrap(err, "in PassThroughError")
}

// FinalDestination deals with the error and
// doesn't forward it
func FinalDestination() {
	err := PassThroughError()

	if err != nil {
		// log it when unexpected error happened
		log.Printf("an error occured: %s\n", err.Error())
		return
	}
}
