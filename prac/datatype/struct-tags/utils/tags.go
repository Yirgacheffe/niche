package utils

// Person represent struct a person
type Person struct {
	Name  string `serialize:"name"`
	City  string `serialize:"city"`
	State string
	Misc  string `serialize:"-"`
	Year  int    `serialize:"year"`
}

// EmptyStruct demonstrate serialize and deserialize for an Empty struct with tags
func EmptyStruct() error {
	return nil
}
