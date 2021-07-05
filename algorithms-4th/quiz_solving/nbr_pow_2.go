package main

import (
	"fmt"
	"math"
)

const ten = "10 = 8 + 2     => 2^3 + 2 "
const sev = " 7 = 4 + 2 + 1 => 2^2 + 2^1 + 2^0"

func main() {
	n, i, sum := 15, 0, 0

	for n >= 1 {

		// divide by 2, check if gcd is 1
		g := n % 2
		if g == 1 {
			i++
		}
		n /= 2
		sum += int(math.Pow(2, float64(i-1)))

	}

	fmt.Printf("%d can be composed by %d (power of 2)\n", sum, i)
}
