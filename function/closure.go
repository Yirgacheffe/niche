package main

import "fmt"

func simpleHello() func(string) string {
	t := "Hello"
	c := func(s string) string {
		t := t + " " + s
		return t
	}
	return c
}

func main() {
	x := 10
	func() {
		y := 20
		fmt.Println(x + y)
	}()

	t1 := simpleHello()
	t2 := simpleHello()

	fmt.Println(t1("World"))
	fmt.Println(t2("Everyone"))
	fmt.Println(t1("Gopher"))
	fmt.Println(t2("!"))

}
