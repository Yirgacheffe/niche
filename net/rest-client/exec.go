package main

import "fmt"

func Exec() error {
	c := NewAPIClient("username", "password")

	statusCode, err := c.GetGoogle()
	if err != nil {
		return err
	}

	fmt.Println("Result of GetHttpbin: ", statusCode)

	return nil
}
