package utils

import (
	"errors"
	"reflect"
)

// DeserializeStructStrings converts a serialized string using our custom
// serialization format to a struct
func DeserializeStructStrings(s string, res interface{}) error {

	r := reflect.TypeOf(res)

	// using a pointer, check if it is correct
	if r.Kind() != reflect.Ptr {
		return errors.New("argument res must be a pointer")
	}

	r = r.Elem()
	v := reflect.ValueOf(res).Elem()

	return nil
}
