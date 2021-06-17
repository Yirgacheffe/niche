package main

type T struct {
	dummy        [0]func()
	AnotherField int
}

var x map[T]int // copmile error

func main() {
	var a, b T
	_ = a == b // compile error
}
