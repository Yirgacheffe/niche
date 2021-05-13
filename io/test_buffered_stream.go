package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {

	buf := bytes.NewBufferString("Hello World!")

	fmt.Print("bytes written => ")
	fmt.Println(buf.WriteString("How are you???????????"))

	strReader := strings.NewReader("  Doing well?       ")
	fmt.Print("bytes read => ")
	fmt.Println(buf.ReadFrom(strReader))

	fmt.Print("bytes read => ")
	fmt.Println(buf.Read(make([]byte, 12)))

	fmt.Print("bytes read => ")
	fmt.Println(buf.WriteTo(os.Stdout))

	fmt.Print("bytes read => ")
	fmt.Println(buf.Read(make([]byte, 10)))

	fmt.Print("bytes read => ")
	fmt.Println(buf.WriteString("Hello!"))

	fmt.Print("bytes read => ")
	fmt.Println(buf.Read(make([]byte, 10)))

}
