package main

import (
	"fmt"

	"./utils"
)

func main() {

	if err := utils.EmptyStruct(); err != nil {
		panic(err)
	}

	fmt.Println()

	if err := utils.FullStruct(); err != nil {
		panic(err)
	}

}
