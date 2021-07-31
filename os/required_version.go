package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

func main() {

	curVersion := runtime.Version()

	major := strings.Split(curVersion, ".")[0][2]
	minor := strings.Split(curVersion, ".")[1]

	m1, _ := strconv.Atoi(string(major))
	m2, _ := strconv.Atoi(minor)

	if m1 == 1 && m2 < 8 {
		fmt.Println("Need Go version 1.8 or higher!")
		return
	}

	fmt.Println("You are using Go version 1.8 or higher!")

}
