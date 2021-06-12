package main

import "fmt"

func reverseBits(n int) int {
	var rev int

	for n > 0 {
		rev <<= 1
		if n&1 == 1 {
			rev ^= 1
		}
		n >>= 1
	}

	return rev
}

func main() {
	fmt.Println(reverseBits(11))
}
