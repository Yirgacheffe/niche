package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	_, err := os.Stat("some-file-notexist.kkk")

	fmt.Println(os.IsNotExist(err))
	fmt.Println(err == os.ErrNotExist)
	fmt.Println(errors.Is(err, os.ErrNotExist))
}
