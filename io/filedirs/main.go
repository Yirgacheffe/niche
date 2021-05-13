package main

import "./utils"

func main() {
	if err := utils.Operate(); err != nil {
		panic(err)
	}
	if err := utils.CapitalizerExample(); err != nil {
		panic(err)
	}
}
