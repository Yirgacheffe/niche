package main

import "./utils"

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

}
