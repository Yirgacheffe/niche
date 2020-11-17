package utils

import "fmt"

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

	p := Person{}
	res, err := SerializeStructStrings(&p)
	if err != nil {
		return err
	}

	fmt.Printf("Empty struct: %#v\n", p)
	fmt.Println("Serialize Results:", res)

	newP := Person{}
	if err := DeserializeStructStrings(res, &newP); err != nil {
		return err
	}

	fmt.Printf("Deserialize results: %#v\n", newP)
	return nil

}

// FullStruct demonstrates serialize and deserialize for an Full struct
// with tags
func FullStruct() error {
	p := Person{
		Name:  "Aaron",
		City:  "Seattle",
		State: "WA",
		Misc:  "some awesome fact in the fieldwards that will happen",
		Year:  2017,
	}

	res, err := SerializeStructStrings(&p)
	if err != nil {
		return err
	}

	fmt.Printf("Full struct: %#v\n", p)
	fmt.Println("Serialize Results:", res)

	newP := Person{}
	if err := DeserializeStructStrings(res, &newP); err != nil {
		return err
	}

	fmt.Printf("Deserialize result: %#v\n", newP)
	return nil

}
