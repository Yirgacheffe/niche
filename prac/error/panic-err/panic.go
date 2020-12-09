package main

import (
	"fmt"
	"strconv"
)

// Panic divide by zero cause panic
func Panic() {
	zero, err := strconv.ParseInt("0", 10, 64)
	if err != nil {
		panic(err)
	}

	a := 1 / zero // get a panic
	fmt.Println("will never get here", a)
}

// Catcher catch all panic
func Catcher() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic occured:", r)
		}
	}()

	Panic() // Call 'Panic' divide by zero
}
