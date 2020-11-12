package main

import "./utils"

func main() {

	if err := utils.Base64Example(); err != nil {
		panic(err)
	}

	if err := utils.Base64ExampleEncoder(); err != nil {
		panic(err)
	}

	if err := utils.GobExample(); err != nil {
		panic(err)
	}

}
