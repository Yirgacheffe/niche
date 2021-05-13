package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	dst := bufio.NewWriter(os.Stdout)

	fmt.Println(dst.WriteString("Hello World!!!"))
	fmt.Println(dst.Write([]byte(" How are you?     \n")))

	fmt.Println(dst.Flush(), "Flushed!")

	fmt.Println(dst.WriteString("Goo000ooo00d?\n"))
	fmt.Println(dst.Flush(), "Flushed Again!")

}
