package main

import "fmt"

const MaxUint = ^uint(0)
const MaxInt = int(^uint(0) >> 1)

const Is64bitArch = ^uint(0)>>63 == 1
const Is32bitArch = ^uint(0)>>63 == 0
const WordBits = 32 << (^uint(0) >> 63) // 64或32

func main() {
	fmt.Println(MaxUint)
	fmt.Println(MaxInt)

	fmt.Println(Is64bitArch)
	fmt.Println(Is32bitArch)
	fmt.Println(WordBits)
}
