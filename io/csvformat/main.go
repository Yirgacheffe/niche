package main

import (
	"fmt"

	"./utils"
)

func main() {

	var err error

	err = utils.AddMoviesFromText()
	if err != nil {
		panic(err)
	}

	err = utils.AddMoviesFromFile()
	if err != nil {
		panic(err)
	}

	err = utils.WriteCSVOutput()
	if err != nil {
		panic(err)
	}

	buffer, err := utils.WriteCSVBuffer()
	if err != nil {
		panic(err)
	}

	fmt.Println("Buffer =", buffer.String())

}
