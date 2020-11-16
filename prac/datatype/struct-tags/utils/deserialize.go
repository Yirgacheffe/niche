package utils

import (
	"errors"
	"reflect"
	"strings"
)

// DeserializeStructStrings converts a serialized string using our custom
// serialization format to a struct
func DeserializeStructStrings(s string, res interface{}) error {

	r := reflect.TypeOf(res)

	// using a pointer, check if it is correct
	if r.Kind() != reflect.Ptr {
		return errors.New("argument res must be a pointer")
	}

	// deref the pointer
	r = r.Elem()
	value := reflect.ValueOf(res).Elem()

	// split our serialization string into a map
	vals := strings.Split(s, ";")
	valMap := make(map[string]string)

	for _, v := range vals {
		keyv := strings.Split(v, ":")
		if len(keyv) != 2 {
			continue
		}
		valMap[keyv[0]] = keyv[1]
	}

	// loop over the fields
	for i := 0; i < r.NumField(); i++ {
		field := r.Field(i)

		if serialize, ok := field.Tag.Lookup("serialize"); ok {
			// "-" will be ignored
			if serialize == "-" {
				continue
			}

			if val, ok := valMap[serialize]; ok {
				value.Field(i).SetString(val)
			}
		} else if val, ok := valMap[field.Name]; ok {
			value.Field(i).SetString(val)
		}
	}

	return nil

}
