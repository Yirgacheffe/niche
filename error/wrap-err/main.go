package main

import (
	"fmt"

	"./utils"
)

func main() {
	utils.Wrap()
	fmt.Println()

	utils.Unwrap()
	fmt.Println()

	utils.StackTrace()
}
