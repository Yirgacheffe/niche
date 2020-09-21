package main

import (
	"fmt"
	"os/user"
)

func main() {

	currUser, err := user.Current()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("user.Current => %v\n", currUser)

	_, err = user.Lookup("Yoruichi")
	if err != nil {
		fmt.Println(err)
	}

}
