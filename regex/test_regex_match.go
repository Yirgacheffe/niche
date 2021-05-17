package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {

	if len(os.Args) <= 1 {
		fmt.Println("Usage: regexp [string]")
		os.Exit(1)
	}

	if m, _ := regexp.MatchString("^[0-9]+$", os.Args[1]); m {
		fmt.Println("Number!")
	} else {
		fmt.Println("Not Number!")
	}

}
