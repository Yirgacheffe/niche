package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	strRder := strings.NewReader("Hello world! How are you?\n")

	// io.Stdout is 'Writer', strRder is string reader
	io.Copy(os.Stdout, strRder)
	fmt.Println()

	// CopyN is only  transfer specific bytes to the destination
	io.CopyN(os.Stdout, strRder, 12)

}
