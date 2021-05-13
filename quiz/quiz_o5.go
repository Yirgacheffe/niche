package main

import "fmt"

type rect struct {
	len, wid int
}

func (r rect) area() int {
	return r.len * r.wid
}

func main() {
	r := &rect{len: 5, wid: 6}
	fmt.Println(r.area())
}
