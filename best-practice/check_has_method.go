package main

import "fmt"

type A int
type B int

func (b B) M(x int) string {
	return fmt.Sprint(b, ": ", x)
}

func check(v interface{}) bool {
	_, has := v.(interface{ M(int) string })
	return has
}

func main() {
	var a A = 123
	var b B = 789

	fmt.Println(check(a))
	fmt.Println(check(b))
}
