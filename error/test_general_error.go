package main

import (
    "database/sql"
    "fmt"
)

func foo() error {
	// return sql.ErrNoRows
	return fmt.Errorf("foo err, %v", sql.ErrNoRows)
}

func bar() error {
	return foo()
}

func main() {
	err := bar()

	if err == sql.ErrNoRows {
		fmt.Printf("Data not found, %+v\n", err)
		return
	} else {
		fmt.Printf("Not found with no type, %+v\n", err)
		return
	}

}
