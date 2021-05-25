package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := make([]byte, 0, 100)

	str = strconv.AppendInt(str, 4567, 10)
	str = strconv.AppendBool(str, false)
	str = strconv.AppendQuote(str, "ajdkfhd")
	str = strconv.AppendQuoteRune(str, '包')

	fmt.Println(string(str))
	fmt.Println(strconv.Itoa(1023))
}
