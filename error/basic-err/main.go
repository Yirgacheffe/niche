package main

import "fmt"

func someDemoFunc() error {
	return ErrorTyped
}

func main() {

	BasicErrors()

	err := someDemoFunc()
	fmt.Println("custom demo error: ", err)

}
