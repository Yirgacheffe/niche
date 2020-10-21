package main

import "./utils"

func main() {
	err := utils.WorkWithBuffer()
	if err != nil {
		panic(err)
	}

	utils.SearchString()
	utils.ModifyString()
	utils.StringReader()
}
