package main

import "fmt"

func CheckType(s interface{}) {

	switch s.(type) {
	case string:
		fmt.Println("It's string.")
	case int:
		fmt.Println("It's int.")
	default:
		fmt.Println("Not sure what it is...")
	}

}

func main() {

	CheckType("test")
	CheckType(64)
	CheckType(false)

	var ix interface{}
	ix = "test again"

	if s, ok := ix.(string); ok {
		fmt.Println("value is:", s)
	}

	if _, ok := ix.(int); !ok {
		fmt.Println("I am happy to handle this.")
	}

}
