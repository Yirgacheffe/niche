package main

import (
	"fmt"
	"strconv"
)

func main() {

	i := 2
	s := "1000"

	if len(s) > 1 {
		i, _ := strconv.Atoi(s)
		i = i + 5
	}

	fmt.Println(i)

}
