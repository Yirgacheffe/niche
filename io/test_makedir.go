package main

import (
	"fmt"
	"os"
)

func main() {
	os.Mkdir("antropology", 0777)
	os.MkdirAll("antropology/xyz/013", 0777)

	err := os.Remove("antropology")
	if err != nil {
		fmt.Println(err)
	}

	os.RemoveAll("antropology")

}
