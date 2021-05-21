package main

import (
	"fmt"
	"os"
	"runtime"
)

func printA(a int) {
	fmt.Println("value of a in deferred function", a)
}

func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x)
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}

func printStack() {
	var buf [4069]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}

func main() {
	a := 10
	defer printA(a)
	a = 20

	fmt.Println("value of a before deferred function call", a)

	defer printStack()
	f(3)

}
