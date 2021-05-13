package main

import "./utils"

func main() {

	if err := utils.NullEncoding(); err != nil {
		panic(err)
	}

	if err := utils.RegularEncoding(); err != nil {
		panic(err)
	}

	if err := utils.PointerEncoding(); err != nil {
		panic(err)
	}

}
