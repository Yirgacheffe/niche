package utils

import "reflect"

// SerializeStructStrings converts a struct to our custom serialization
// format it honos serialize struct tags for string types
func SerializeStructStrings(s interface{}) (string, error) {

	result := ""

	// reflect the interface into a type
	r := reflect.TypeOf(s)
	v := reflect.ValueOf(s)

	// handle pointer appropriately
	if r.Kind() == reflect.Ptr {
		r = r.Elem()
		v = v.Elem()
	}

	// loop over all of the fields
	for i := 0; i < r.NumField(); i++ {

		field := r.Field(i)
		key := field.Name

		if serialize, ok := field.Tag.Lookup("serialize"); ok {
			if serialize == "-" {
				continue
			}
			key = serialize
		}

		switch v.Field(i).Kind() {
		case reflect.String:
			result += key + ":" + v.Field(i).String() + ";"
		default:
			continue // as just support string, ignore other type
		}

	}

	return "", nil

}
