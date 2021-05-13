package main

import "fmt"

func main() {
	fmt.Println("before panic")
	Catcher()
	fmt.Println("after panic")
}
