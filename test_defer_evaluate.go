package main

import (
	"fmt"
)

func printA(a int) {
    fmt.Println("value of a in deferred function", a)
}

func main() {
    a := 10
    defer printA(a)
    a = 20

    fmt.Println("value of a before deferred function call", a)
}