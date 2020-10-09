package main

import "fmt"

func hi(i int) {
	fmt.Println(i)
}

func main() {

	i := 10
	defer hi(i)

	i = i + 10
	fmt.Println(i)

}
