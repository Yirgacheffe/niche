package main

import (
	"fmt"
)

func encryptWithIdx(s string) string {

	n := len(s)
	if n > 128 {
		fmt.Println("Max size is 128, return back.")
		return s
	}

	chars := []byte(s)
	j := 0

	for i := (n - 1); i >= 0; i-- {
		chars[j] = (s[i] + byte(i)) % 128
		j++
	}

	return string(chars)

}

func main() {
	s0 := "geeks"
	s1 := "java"

	fmt.Println(encryptWithIdx(s0))
	fmt.Println(encryptWithIdx(s1))
}
