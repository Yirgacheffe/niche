package main

import (
	"fmt"
	"time"
)

func main() {
	d, _ := time.ParseDuration("3m32s")
	fmt.Println(d)
}
