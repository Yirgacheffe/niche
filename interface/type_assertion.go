package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

type Describer interface {
	Describe()
}

type People struct {
	name string
	age  int
}

func (p People) Describe() {
	fmt.Printf("%s is %d years old\n", p.name, p.age)
}

func assert(i interface{}) {
	v, ok := i.(int)
	fmt.Println(v, ok)
}

func findType(i interface{}) {
	switch v := i.(type) {
	case string:
		fmt.Printf("%s\n", v)
	case int:
		fmt.Printf("%d\n", v)
	case Describer:
		v.Describe()
	default:
		fmt.Printf("Unknown Type\n")
	}
}

func main() {
	i := 10
	assert(i)
	var s interface{} = "xyzzz"
	assert(s)

	findType("Yz983jjdf")
	findType(873)
	findType(83.33)

	p := People{
		name: "Neview",
		age:  24,
	}

	findType(p)

	var w io.Writer = os.Stdout
	if f, ok := w.(*os.File); ok {
		fmt.Println(f.Name())
	}

	if b, ok := w.(*bytes.Buffer); ok {
		fmt.Println(b.WriteString("{}"))
	}
}
