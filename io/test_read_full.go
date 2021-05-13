package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {

	src := strings.NewReader("Hello amazing world!") // 20 characters
	buf := make([]byte, 14)

	r1, err1 := io.ReadFull(src, buf)
	fmt.Printf("bytes read: %d, value: %s, error: %v\n", r1, buf[:r1], err1)

	r2, err2 := io.ReadFull(src, buf)
	fmt.Printf("bytes read: %d, value: %s, error: %v\n", r2, buf[:r2], err2)

	r3, err3 := io.ReadFull(src, buf)
	fmt.Printf("bytes read: %d, value: %s, error: %v\n", r3, buf[:r3], err3)

}
