package main

import (
	"fmt"
	"runtime"
)

func main() {

	curVersion := runtime.Version()

	fmt.Println(curVersion)

}
