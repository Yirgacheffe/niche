package main

import "./utils"

func main() {

	if err := utils.WorkWithTemp(); err != nil {
		panic(err)
	}

}
