package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func main() {

	var LENGTH int64 = 8
	MIN := 0
	MAX := 94

	SEED := time.Now().Unix()
	arguments := os.Args

	switch len(arguments) {
	case 2:
		LENGTH, _ = strconv.ParseInt(arguments[1], 10, 64)
	default:
		fmt.Println("Using default values!")
	}

	rand.Seed(SEED)
	startChar := "!"

	var i int64 = 1
	var passwdString string

	for {
		myRand := random(MIN, MAX)
		passwdString += string(startChar[0] + byte(myRand))

		if i == LENGTH {
			break
		}
		i++
	}

	fmt.Println(passwdString)
	fmt.Println() // #---------------------------------- #

}
