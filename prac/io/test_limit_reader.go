package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {

	src := strings.NewReader("Hello amazing world!")

	// create 10 capacity length reader
	limited := io.LimitReader(src, 10)
	buf := make([]byte, 3)

	for {
		n, err := limited.Read(buf)
		fmt.Printf("%d bytes read, data: %s\n", n, buf[:n])

		if err == io.EOF {
			fmt.Println("----end-of-file----")
			break
		} else if err != nil {
			fmt.Println("----unknow-error-happened----")
			break
		}
	}

}
