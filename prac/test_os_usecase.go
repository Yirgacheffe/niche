package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(os.Args[1:])

	fmt.Println("os.Args => ", os.Args)

	fmt.Printf("os.PathSeparator => %v\n", os.PathSeparator)
	fmt.Printf("os.PathSeparator => %c\n", os.PathSeparator)

	fmt.Printf("os.DevNull => %v\n", os.DevNull)

}
