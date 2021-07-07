package main

import (
	"database/sql"
	"fmt"
)

func foo() error {
	return fmt.Errorf("foo err, %v", sql.ErrNoRows) // != sql.ErrNoRows
}

func bar() error {
	return foo()
}

func main() {
	err := bar()

	if err == sql.ErrNoRows {
		fmt.Printf("Data not found, %+v\n", err)
	} else {
		fmt.Printf("Not found with no type, %+v\n", err)
	}
}
