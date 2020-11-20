package main

import (
	"fmt"

	"./utils"
)

func main() {

	ws := []utils.WorkWith{
		utils.WorkWith{Data: "Example 1", Version: 1},
		utils.WorkWith{Data: "Example 2", Version: 2},
	}

	fmt.Printf("Initial list: %#v\n", ws)

	// map string to lower
	ws = utils.Map(ws, utils.LowerCaseData)
	fmt.Printf("After LowerCaseData Map: %#v\n", ws)

	// map increase the number
	ws = utils.Map(ws, utils.IncrementVersion)
	fmt.Printf("After IncrementVersion Map: %#v\n", ws)

	// remove all versions less than 3
	ws = utils.Filter(ws, utils.OldVersion(3))
	fmt.Printf("After OldVersion Filter: %#v\n", ws)

}
