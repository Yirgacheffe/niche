package main

import (
	"fmt"

	"github.com/Yirgacheffe/hello-mod/calc"
	"github.com/Yirgacheffe/hello-mod/hello"
)

func main() {

	sum := calc.Add(1, 2)
	fmt.Println("result comes from calc module:", sum)

	hello.Hello() // just call hello for testing

}
